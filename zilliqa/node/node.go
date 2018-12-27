/*
Package node implements processing routines for Zilliqa's Final Block...
*/
package node

import (
	"blockchain/utils/crypto"
	"blockchain/utils/network"
	"blockchain/zilliqa/block"
	"errors"
	"github.com/golang-collections/go-datastructures/bitarray"
)

/*
Node struture implements a blockchain node.
*/
type Node struct {
	mediator network.Mediator
}

/*
VerifyFinalBlockConsensusSignature verifies final block's consensus signature.
*/
func (node Node) VerifyFinalBlockConsensusSignature(txBlock block.TXBlock) (bool, error) {
	keys, err := collectCommitteePublicKeys(node.mediator.DSCommittee)
	if err != nil {
		return false, nil
	}

	signaturesTally := committeeSignatureTally(txBlock.ConsensusSignatureMap1stRound,
		node.mediator.DSCommittee)
	signaturesTallyForConsensus :=
		node.mediator.NumberForConsensus(len(node.mediator.DSCommittee))

	if signaturesTally < signaturesTallyForConsensus {
		return false, errors.New("not enough signatures for consensus")
	}

	combinedPublicKeys, err := crypto.CombinePublicKeys(keys)
	if err != nil {
		return false, err
	}

	message, err := txBlock.Message()
	if err != nil {
		return false, err
	}

	return crypto.VerifySignature(message, txBlock.ConsensusSignature2ndRound, combinedPublicKeys)
}

func collectCommitteePublicKeys(committee []network.DSEntrusted) ([][]byte, error) {
	keys := make([][]byte, cap(committee))
	for i, member := range committee {
		keys[i] = member.PublicKey
	}
	return keys, nil
}

func committeeSignatureTally(signatures bitarray.BitArray,
	committee []network.DSEntrusted) int {
	tally := 0
	for i := range committee {
		hasSigned, _ := signatures.GetBit(uint64(i))
		if hasSigned {
			tally++
		}
	}
	return tally
}
