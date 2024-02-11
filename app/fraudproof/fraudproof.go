package fraudproof

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/dymensionxyz/dymension/v3/x/rollapp/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store/rootmulti"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"

	abci "github.com/tendermint/tendermint/abci/types"

	fraudtypes "github.com/cosmos/cosmos-sdk/baseapp"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	rollappevm "github.com/dymensionxyz/rollapp-evm/app"

	_ "github.com/evmos/evmos/v12/crypto/codec"
	_ "github.com/evmos/evmos/v12/crypto/ethsecp256k1"
	_ "github.com/evmos/evmos/v12/types"
)

// TODO: Move to types package
var (
	ErrInvalidPreStateAppHash = errors.New("invalid pre state app hash")
	ErrInvalidAppHash         = errors.New("invalid app hash")
)

type FraudProofVerifier interface {
	InitFromFraudProof(fraudProof *fraudtypes.FraudProof) error
	VerifyFraudProof(fraudProof *fraudtypes.FraudProof) error
}

type RollappFPV struct {
	app  *baseapp.BaseApp
	keys map[string]storetypes.StoreKey
}

var _ FraudProofVerifier = (*RollappFPV)(nil)

// New creates a new FraudProofVerifier
func New(appName string) *RollappFPV {
	cfg := rollappevm.MakeEncodingConfig()

	//TODO: use logger? default home directory?
	rollappApp := rollappevm.NewRollapp(log.NewNopLogger(), dbm.NewMemDB(), nil, false, map[int64]bool{}, "/tmp", 0, cfg, simapp.EmptyAppOptions{})

	rollapp := baseapp.NewBaseApp(appName, log.NewNopLogger(), dbm.NewMemDB(), cfg.TxConfig.TxDecoder())
	rollapp.SetMsgServiceRouter(rollappApp.MsgServiceRouter())
	rollapp.SetBeginBlocker(rollappApp.GetBeginBlocker())
	rollapp.SetEndBlocker(rollappApp.GetEndBlocker())

	cms := rollappApp.CommitMultiStore().(*rootmulti.Store)
	storeKeys := cms.StoreKeysByName()

	return &RollappFPV{
		app:  rollapp,
		keys: storeKeys,
	}
}

// InitFromFraudProof initializes the FraudProofVerifier from a fraud proof
func (fpv *RollappFPV) InitFromFraudProof(fraudProof *fraudtypes.FraudProof) error {
	// check app is initialized
	if fpv.app == nil {
		return fmt.Errorf("app not initialized")
	}

	_, err := fraudProof.ValidateBasic()
	if err != nil {
		return err
	}

	fpv.app.SetInitialHeight(fraudProof.BlockHeight + 1) //FIXME: why +1?

	cms := fpv.app.CommitMultiStore().(*rootmulti.Store)
	storeKeys := fpv.keys
	modules := fraudProof.GetModules()
	iavlStoreKeys := make([]storetypes.StoreKey, 0, len(modules))
	for _, module := range modules {
		iavlStoreKeys = append(iavlStoreKeys, storeKeys[module])
	}

	fpv.app.MountStores(iavlStoreKeys...)

	storeKeyToIAVLTree, err := fraudProof.GetDeepIAVLTrees()
	if err != nil {
		return err
	}
	for storeKey, iavlTree := range storeKeyToIAVLTree {
		cms.SetDeepIAVLTree(storeKey, iavlTree)
	}

	err = fpv.app.LoadLatestVersion()
	if err != nil {
		return err
	}

	fpv.app.InitChain(abci.RequestInitChain{})

	return nil

}

// VerifyFraudProof checks the validity of a given fraud proof.
//
// The function takes a FraudProof object as an argument and returns an error if the fraud proof is invalid.
//
// The function performs the following checks:
// 1. It checks if the pre-state application hash of the fraud proof matches the current application hash.
// 2. It executes a fraudulent state transition.
// 3. Finally, it checks if the post-state application hash matches the expected valid application hash in the fraud proof.
//
// If any of these checks fail, the function returns an error. Otherwise, it returns nil.
//
// Note: This function modifies the state of the RollappFPV object it's called on.
func (fpv *RollappFPV) VerifyFraudProof(fraudProof *fraudtypes.FraudProof) error {
	appHash := fpv.app.GetAppHashInternal()
	fmt.Println("appHash - prestate", hex.EncodeToString(appHash))

	if !bytes.Equal(fraudProof.PreStateAppHash, appHash) {
		return ErrInvalidPreStateAppHash
	}

	// Execute fraudulent state transition
	if fraudProof.FraudulentBeginBlock != nil {
		panic("fraudulent begin block not supported")
		// fpv.app.BeginBlock(*fraudProof.FraudulentBeginBlock)
		// fmt.Println("appHash - beginblock", hex.EncodeToString(fpv.app.GetAppHashInternal()))
	} else {
		// Need to add some dummy begin block here since its a new app
		fpv.app.ResetDeliverState()
		fpv.app.SetBeginBlocker(nil)
		fpv.app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{Height: fraudProof.BlockHeight + 1}}) //FIXME: why +1?
		fmt.Println("appHash - dummy beginblock", hex.EncodeToString(fpv.app.GetAppHashInternal()))

		if fraudProof.FraudulentDeliverTx != nil {
			// skip IncrementSequenceDecorator check in AnteHandler
			fpv.app.SetAnteHandler(nil)
			SetRollappAddressPrefixes("ethm")

			resp := fpv.app.DeliverTx(*fraudProof.FraudulentDeliverTx)
			if !resp.IsOK() {
				panic(resp.Log)
			}
			fmt.Println("appHash - posttx", hex.EncodeToString(fpv.app.GetAppHashInternal()))
			SetRollappAddressPrefixes("dym")
		} else {
			panic("fraudulent end block not supported")
			// fpv.app.EndBlock(*fraudProof.FraudulentEndBlock)
			// fmt.Println("appHash - endblock", hex.EncodeToString(fpv.app.GetAppHashInternal()))
		}
	}

	appHash = fpv.app.GetAppHashInternal()
	fmt.Println("appHash - final", hex.EncodeToString(appHash))
	if !bytes.Equal(appHash, fraudProof.ExpectedValidAppHash) {
		return types.ErrInvalidAppHash
	}
	return nil
}

func SetRollappAddressPrefixes(prefix string) {
	// Set prefixes
	accountPubKeyPrefix := prefix + "pub"
	validatorAddressPrefix := prefix + "valoper"
	validatorPubKeyPrefix := prefix + "valoperpub"
	consNodeAddressPrefix := prefix + "valcons"
	consNodePubKeyPrefix := prefix + "valconspub"

	// Set config
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccountNoAssert(prefix, accountPubKeyPrefix)
	config.SetBech32PrefixForValidatorNoAssert(validatorAddressPrefix, validatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNodeNoAssert(consNodeAddressPrefix, consNodePubKeyPrefix)
}