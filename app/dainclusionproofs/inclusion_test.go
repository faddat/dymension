package inclusion_test

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	openrpcns "github.com/rollkit/celestia-openrpc/types/namespace"
	"github.com/stretchr/testify/require"

	inclusion "github.com/dymensionxyz/dymension/v3/app/dainclusionproofs"
	rollapptypes "github.com/dymensionxyz/dymension/v3/x/rollapp/types"
)

func TestInclusionProof(t *testing.T) {

	require := require.New(t)

	nameidstr := "e06c57a64b049d6463ef"
	namespaceBytes, err := hex.DecodeString(nameidstr)
	require.NoError(err)
	ns, err := openrpcns.New(openrpcns.NamespaceVersionZero, append(openrpcns.NamespaceVersionZeroPrefix, namespaceBytes...))
	require.NoError(err)

	file, err := os.Open("./blob_inclusion_proof.json")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	// Decode the JSON-encoded data into your struct
	jsonDecoder := json.NewDecoder(file)
	proof := rollapptypes.BlobInclusionProof{}
	err = jsonDecoder.Decode(&proof)
	require.NoError(err)

	fmt.Println("Namespace", hex.EncodeToString(ns.Bytes()))

	inclusionProof := &inclusion.InclusionProof{}

	inclusionProof.Blob = proof.GetBlob()
	inclusionProof.DataRoot = proof.GetDataroot()
	inclusionProof.Nmtproofs = proof.GetNmtproofs()
	inclusionProof.Nmtroots = proof.GetNmtroots()
	inclusionProof.RowProofs = proof.GetRproofs()

	err = inclusionProof.VerifyBlobInclusion(ns.Bytes(), proof.GetDataroot())
	require.NoError(err)

}