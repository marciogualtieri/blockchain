package node

import (
	"blockchain/utils/network"
	"blockchain/zilliqa/block"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVerifyFinalBlockConsensusSignature(t *testing.T) {
	node := Node{mediator: network.TestMediator}
	verification, err :=
		node.VerifyFinalBlockConsensusSignature(block.TestTXBlock)
	assert.Nil(t, err)
	assert.True(t, verification)
}

func TestVerifyFinalBlockConsensusSignatureNotEnoughVotes(t *testing.T) {
	node := Node{mediator: network.TestMediator}
	verification, err :=
		node.VerifyFinalBlockConsensusSignature(block.TestTXBlockNotEnoughSignatures)
	assert.NotNil(t, err)
	assert.Equal(t, "not enough signatures for consensus", err.Error())
	assert.False(t, verification)
}

func TestVerifyFinalBlockConsensusSignatureCorruptSignature(t *testing.T) {
	node := Node{mediator: network.TestMediator}
	verification, err :=
		node.VerifyFinalBlockConsensusSignature(block.TestTXBlockCorruptSignature)
	assert.Nil(t, err)
	assert.False(t, verification)
}

func TestVerifyFinalBlockConsensusSignatureCorruptSignatureMap(t *testing.T) {
	node := Node{mediator: network.TestMediator}
	verification, err :=
		node.VerifyFinalBlockConsensusSignature(block.TestTXBlockCorruptSignatureMap)
	assert.Nil(t, err)
	assert.False(t, verification)
}
