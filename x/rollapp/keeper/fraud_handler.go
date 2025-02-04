package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	tmtypes "github.com/cosmos/ibc-go/v6/modules/light-clients/07-tendermint/types"

	common "github.com/dymensionxyz/dymension/v3/x/common/types"
	"github.com/dymensionxyz/dymension/v3/x/rollapp/types"
)

// HandleFraud handles the fraud evidence submitted by the user.
func (k Keeper) HandleFraud(ctx sdk.Context, rollappID, clientId string, fraudHeight uint64, seqAddr string) error {
	// Get the rollapp from the store
	rollapp, found := k.GetRollapp(ctx, rollappID)
	if !found {
		return sdkerrors.Wrapf(types.ErrInvalidRollappID, "rollapp with ID %s not found", rollappID)
	}

	stateInfo, err := k.FindStateInfoByHeight(ctx, rollappID, fraudHeight)
	if err != nil {
		return err
	}

	//check height is not finalized
	if stateInfo.Status == common.Status_FINALIZED {
		return sdkerrors.Wrapf(types.ErrDisputeAlreadyFinalized, "state info for height %d is already finalized", fraudHeight)
	}

	//check height is not reverted
	if stateInfo.Status == common.Status_REVERTED {
		return sdkerrors.Wrapf(types.ErrDisputeAlreadyReverted, "state info for height %d is already reverted", fraudHeight)
	}

	//check the sequencer for this height is the same as the one in the fraud evidence
	if stateInfo.Sequencer != seqAddr {
		return sdkerrors.Wrapf(types.ErrWrongProposerAddr, "sequencer address %s does not match the one in the state info", seqAddr)
	}

	// slash the sequencer, clean delayed packets
	err = k.hooks.FraudSubmitted(ctx, rollappID, fraudHeight, seqAddr)
	if err != nil {
		return err
	}

	//mark the rollapp as frozen. revert all pending states to finalized
	rollapp.Frozen = true
	k.SetRollapp(ctx, rollapp)

	k.RevertPendingStates(ctx, rollappID)

	//TODO: get the clientId from rollapp object, instead of by proposal
	if clientId != "" {
		err = k.freezeClientState(ctx, clientId)
		if err != nil {
			return err
		}
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeFraud,
			sdk.NewAttribute(types.AttributeKeyRollappId, rollappID),
			sdk.NewAttribute(types.AttributeKeyFraudHeight, fmt.Sprint(fraudHeight)),
			sdk.NewAttribute(types.AttributeKeyFraudSequencer, seqAddr),
			sdk.NewAttribute(types.AttributeKeyClientID, clientId),
		),
	)

	return nil
}

// freeze IBC client state
func (k Keeper) freezeClientState(ctx sdk.Context, clientId string) error {
	clientState, ok := k.ibcclientKeeper.GetClientState(ctx, clientId)
	if !ok {
		return sdkerrors.Wrapf(types.ErrInvalidClientState, "client state for clientID %s not found", clientId)
	}

	tmClientState, ok := clientState.(*tmtypes.ClientState)
	if !ok {
		return sdkerrors.Wrapf(types.ErrInvalidClientState, "client state with ID %s is not a tendermint client state", clientId)
	}

	tmClientState.FrozenHeight = clienttypes.NewHeight(tmClientState.GetLatestHeight().GetRevisionHeight(), tmClientState.GetLatestHeight().GetRevisionNumber())
	k.ibcclientKeeper.SetClientState(ctx, clientId, tmClientState)

	return nil
}

// revert all pending states of a rollapp
func (k Keeper) RevertPendingStates(ctx sdk.Context, rollappID string) {
	// TODO (#631): Prefix store by rollappID for efficient querying
	queuePerHeight := k.GetAllBlockHeightToFinalizationQueue(ctx)
	for _, queue := range queuePerHeight {
		leftPendingStates := []types.StateInfoIndex{}
		for _, stateInfoIndex := range queue.FinalizationQueue {
			//keep pending packets not related to this rollapp in the queue
			if stateInfoIndex.RollappId != rollappID {
				leftPendingStates = append(leftPendingStates, stateInfoIndex)
				continue
			}

			stateInfo, _ := k.GetStateInfo(ctx, stateInfoIndex.RollappId, stateInfoIndex.Index)
			stateInfo.Status = common.Status_REVERTED
			k.SetStateInfo(ctx, stateInfo)
		}

		if len(leftPendingStates) == 0 {
			k.RemoveBlockHeightToFinalizationQueue(ctx, queue.CreationHeight)
		} else {
			k.SetBlockHeightToFinalizationQueue(ctx, types.BlockHeightToFinalizationQueue{
				CreationHeight:    queue.CreationHeight,
				FinalizationQueue: leftPendingStates,
			})
		}
	}
}
