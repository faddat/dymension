package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dymensionxyz/dymension/v3/x/rollapp/types"
)

func (k msgServer) SubmitInvalidDataBatch(goCtx context.Context, msg *types.MsgInvalidDataBatch) (*types.MsgInvalidDataBatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !k.RollappsEnabled(ctx) {
		return nil, types.ErrRollappsDisabled
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// load rollapp object for stateful validations
	_, isFound := k.GetRollapp(ctx, msg.RollappId)
	if !isFound {
		return nil, types.ErrUnknownRollappID
	}

	err := k.VerifyInvalidDataBatch(ctx, msg)

	if err == nil {
		//FIXME: handle deposit burn on wrong FP
		k.Logger(ctx).Info("unable to verif non-available proof ", "rollappID", msg.RollappId)

	}

	return &types.MsgInvalidDataBatchResponse{}, nil
}
