package fraudproof

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/dymensionxyz/dymension/x/rollapp/types"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/store/rootmulti"

	"github.com/cosmos/cosmos-sdk/baseapp"

	abci "github.com/tendermint/tendermint/abci/types"

	fraudtypes "github.com/cosmos/cosmos-sdk/baseapp"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

var (
	ErrInvalidPreStateAppHash = errors.New("invalid pre state app hash")
	ErrInvalidAppHash         = errors.New("invalid app hash")
)

type FraudProofVerifier interface {
	InitFromFraudProof(fraudProof *fraudtypes.FraudProof) error
	VerifyFraudProof(fraudProof *fraudtypes.FraudProof) error
}

type RollappFPV struct {
	host *baseapp.BaseApp
	app  *baseapp.BaseApp
}

var _ FraudProofVerifier = (*RollappFPV)(nil)

// New creates a new FraudProofVerifier
func New(host *baseapp.BaseApp, appName string, logger log.Logger) *RollappFPV {
	//TODO: use logger?
	//TODO: default home directory?

	//TODO: test with dymension app for working reference
	// encCdc := rollappparams.MakeEncodingConfig()
	// encCdc := rollappparams.MakeEncodingConfig()
	// rollapp := rollappevm.NewRollapp(log.NewNopLogger(), db.NewMemDB(), nil, true, map[int64]bool{}, "/tmp", 0, encCdc, nil)

	// encCdc := app.MakeEncodingConfig()
	// rollapp := app.New(log.NewNopLogger(), db.NewMemDB(), nil, true, map[int64]bool{}, "/tmp", 0, encCdc, nil)

	newApp := baseapp.NewBaseApp(appName, log.NewNopLogger(), db.NewMemDB(), nil)
	//FIXME: remove this
	if host != nil {
		newApp.SetMsgServiceRouter(host.MsgServiceRouter())
		newApp.SetBeginBlocker(host.GetBeginBlocker())
		newApp.SetEndBlocker(host.GetEndBlocker())
	}
	// newApp.msgServiceRouter = app.msgServiceRouter
	// newApp.beginBlocker = app.beginBlocker
	// newApp.endBlocker = app.endBlocker

	return &RollappFPV{
		host: host,
		app:  newApp,
	}
}

// InitFromFraudProof initializes the FraudProofVerifier from a fraud proof
func (fpv *RollappFPV) InitFromFraudProof(fraudProof *fraudtypes.FraudProof) error {
	// check app is initialized
	if fpv.app == nil {
		// return types.ErrAppNotInitialized
		return fmt.Errorf("app not initialized")
	}

	_, err := fraudProof.ValidateBasic()
	if err != nil {
		return err
	}

	cmsHost := fpv.host.CommitMultiStore().(*rootmulti.Store)
	storeKeys := cmsHost.StoreKeysByName()
	// modules := fraudProof.GetModules()
	// iavlStoreKeys := make([]storetypes.StoreKey, 0, len(modules))
	// for _, module := range modules {
	// iavlStoreKeys = append(iavlStoreKeys, storeKeys[module])
	// }

	iavlStoreKeys := make([]storetypes.StoreKey, 0, len(storeKeys))
	for _, storeKey := range storeKeys {
		iavlStoreKeys = append(iavlStoreKeys, storeKey)
	}
	fpv.app.MountStores(iavlStoreKeys...)

	storeKeyToIAVLTree, err := fraudProof.GetDeepIAVLTrees()
	if err != nil {
		return err
	}
	cmsStore := fpv.app.CommitMultiStore().(*rootmulti.Store)
	for storeKey, iavlTree := range storeKeyToIAVLTree {
		cmsStore.SetDeepIAVLTree(storeKey, iavlTree)
	}

	err = fpv.app.LoadLatestVersion()
	if err != nil {
		return err
	}

	//is it enough? rollkit uses:
	//	// This initial height is used in `BeginBlock` in `validateHeight`
	// options = append(options, SetInitialHeight(blockHeight))
	fpv.app.InitChain(abci.RequestInitChain{InitialHeight: fraudProof.BlockHeight})

	return nil

}

// VerifyFraudProof implements the ABCI application interface. It loads a fresh BaseApp using
// the given Fraud Proof, runs the given fraudulent state transition within the Fraud Proof,
// and gets the app hash representing state of the resulting BaseApp. It returns a boolean
// representing whether this app hash is equivalent to the expected app hash given.
func (fpv *RollappFPV) VerifyFraudProof(fraudProof *fraudtypes.FraudProof) error {

	//TODO: check app is initialized

	//TODO: pass rollapp name as well
	appHash := fpv.app.GetAppHashInternal()
	fmt.Println("appHash - prestate", hex.EncodeToString(appHash))

	if !bytes.Equal(fraudProof.PreStateAppHash, appHash) {
		return ErrInvalidPreStateAppHash
	}

	//TODO: verifyy all data exists in fraud proof

	// Execute fraudulent state transition
	fpv.app.BeginBlock(*fraudProof.FraudulentBeginBlock)
	fmt.Println("appHash - beginblock", hex.EncodeToString(fpv.app.GetAppHashInternal()))

	for _, tx := range fraudProof.FraudulentDeliverTx {
		resp := fpv.app.DeliverTx(*tx)
		if !resp.IsOK() {
			panic(resp.Log)
		}
	}
	fmt.Println("appHash - posttx", hex.EncodeToString(fpv.app.GetAppHashInternal()))

	fpv.app.EndBlock(*fraudProof.FraudulentEndBlock)
	appHash = fpv.app.GetAppHashInternal()
	fmt.Println("appHash - endblock", hex.EncodeToString(appHash))
	if !bytes.Equal(appHash, fraudProof.ExpectedValidAppHash) {
		return types.ErrInvalidAppHash
	}
	return nil
}