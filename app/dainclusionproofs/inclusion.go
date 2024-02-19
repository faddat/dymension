package inclusion

import (
	"bytes"
	"crypto/sha256"
	"errors"

	"github.com/celestiaorg/nmt"
	"github.com/cometbft/cometbft/crypto/merkle"
	cmtcrypto "github.com/cometbft/cometbft/proto/tendermint/crypto"
	"github.com/rollkit/celestia-openrpc/types/blob"
)

type InclusionProof struct {
	Blob      []byte
	Nmtproofs [][]byte
	Nmtroots  [][]byte
	RowProofs [][]byte
	DataRoot  []byte
}

func (ip *InclusionProof) VerifyBlobInclusion(namespace []byte, dataRoot []byte) error {

	if !bytes.Equal(ip.DataRoot, dataRoot) {
		return errors.New("data root not matching")
	}

	var nmtProofs []*nmt.Proof
	for _, codedNMTProof := range ip.Nmtproofs {
		var unmarshalledProof nmt.Proof
		err := unmarshalledProof.UnmarshalJSON(codedNMTProof)
		if err != nil {
			return err
		}
		nmtProofs = append(nmtProofs, &unmarshalledProof)
	}

	b, _, err := ip.blobsAndCommitments(namespace, ip.Blob)
	if err != nil {
		return err
	}

	shares, err := blob.SplitBlobs(*b)
	if err != nil {
		return err
	}
	index := 0

	for i, nmtProof := range nmtProofs {
		sharesNum := nmtProof.End() - nmtProof.Start()
		var leafs [][]byte

		for j := index; j < index+sharesNum; j++ {
			leaf := shares[j].ToBytes()
			leafs = append(leafs, leaf)
		}
		if !nmtProof.VerifyInclusion(sha256.New(), namespace, leafs, ip.Nmtroots[i]) {
			return errors.New("blob not included")
		}

		index += sharesNum
	}

	for j, rowProof := range ip.RowProofs {

		var proof cmtcrypto.Proof
		err := proof.Unmarshal(rowProof)
		if err != nil {
			return err
		}
		rProof, err := merkle.ProofFromProto(&proof)
		if err != nil {
			return err
		}
		err = rProof.Verify(ip.DataRoot, ip.Nmtroots[j])
		if err != nil {
			return err
		}
	}
	return nil
}

func (ip *InclusionProof) blobsAndCommitments(namespace []byte, daBlob []byte) (*blob.Blob, []byte, error) {
	b, err := blob.NewBlobV0(namespace, daBlob)
	if err != nil {
		return nil, nil, err
	}

	commitment, err := blob.CreateCommitment(b)
	if err != nil {
		return nil, nil, err
	}

	return b, commitment, nil
}