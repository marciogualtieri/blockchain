package block

import (
	"blockchain/utils/crypto"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/golang-collections/go-datastructures/bitarray"
)

/*
TestTxBlockHashSet is the TX-Block hash set used in tests.
*/
var TestTxBlockHashSet = TXBlockHashSet{
	MerkleRootHash:      chainhash.HashB([]byte("Test Merkle Root")),
	StateDeltaHash:      chainhash.HashB([]byte("Test State Delta")),
	MicroBlocksInfoHash: chainhash.HashB([]byte("Test Micro Blocks Info")),
}

/*
TestTXBlockHeader is the TX-Block header used in tests.
*/
var TestTXBlockHeader = TXBlockHeader{
	BlockType:      MicroBlockMarker,
	Version:        123,
	GasLimit:       456,
	RewardsHigh:    789,
	RewardsLow:     012,
	PreviousHash:   chainhash.HashB([]byte("Test Previous Block Hash")),
	BlockNum:       345,
	HashSet:        TestTxBlockHashSet,
	NumTXs:         789,
	MinerPublicKey: crypto.TestPublicKey1,
	DSBlockNum:     901,
}

var templateTestBlock = TXBlock{
	ConsensusSignatureMap1stRound: nil,
	ConsensusSignatureMap2ndRound: nil,
	ConsensusSignature1stRound:    nil,
	ConsensusSignature2ndRound:    nil,
	Header:                        TestTXBlockHeader,
}

/*
TestTXBlock is the TX-Block used in tests for the happy path.
*/
var TestTXBlock = createTestTXBlock()

/*
TestTXBlockNotEnoughSignatures is the TX-Block used in tests for when not enough
votes.
*/
var TestTXBlockNotEnoughSignatures = createTestTXBlockNotEnoughSignatures()

/*
TestTXBlockCorruptSignature is the TX-Block used in tests for when the
multisignature is invalid.
*/
var TestTXBlockCorruptSignature = createTestTXBlockCorruptSignature()

/*
TestTXBlockCorruptSignatureMap is the TX-Block used in tests for when the
multisignature is invalid.
*/
var TestTXBlockCorruptSignatureMap = createTestTXBlockCorruptSignatureMap()

func createTestTXBlock() TXBlock {
	block := templateTestBlock
	block.ConsensusSignatureMap1stRound = createConsensusSignatureMap(3, 2)
	block.ConsensusSignatureMap2ndRound = createConsensusSignatureMap(3, 2)
	message, _ := block.Message()
	block.ConsensusSignature2ndRound = createCombinedConsensusSignature(message)
	return block
}

func createTestTXBlockNotEnoughSignatures() TXBlock {
	block := templateTestBlock
	block.ConsensusSignatureMap1stRound = createConsensusSignatureMap(3, 1)
	block.ConsensusSignatureMap2ndRound = createConsensusSignatureMap(3, 1)
	message, _ := block.Message()
	block.ConsensusSignature2ndRound = createCombinedConsensusSignature(message)
	return block
}

func createTestTXBlockCorruptSignature() TXBlock {
	block := templateTestBlock
	block.ConsensusSignatureMap1stRound = createConsensusSignatureMap(3, 2)
	block.ConsensusSignatureMap2ndRound = createConsensusSignatureMap(3, 2)
	message, _ := block.Message()
	block.ConsensusSignature2ndRound = createCombinedConsensusSignature(message)
	block.ConsensusSignature2ndRound[0] = 0xff
	return block
}

func createTestTXBlockCorruptSignatureMap() TXBlock {
	block := templateTestBlock
	block.ConsensusSignatureMap1stRound = createConsensusSignatureMap(3, 1)
	block.ConsensusSignatureMap2ndRound = createConsensusSignatureMap(3, 2)
	message, _ := block.Message()
	block.ConsensusSignature2ndRound = createCombinedConsensusSignature(message)
	block.ConsensusSignatureMap1stRound = createConsensusSignatureMap(3, 2)
	return block
}

func createConsensusSignatureMap(committeeSize uint64, numSignatures uint64) bitarray.BitArray {
	consensusSignatures := bitarray.NewBitArray(committeeSize)
	for i := 0; i < int(numSignatures); i++ {
		consensusSignatures.SetBit(uint64(i))
	}
	return consensusSignatures
}

func createCombinedConsensusSignature(message []byte) []byte {
	signature1, _ := crypto.PartialSignMessage(message, crypto.TestPrivateKey1,
		crypto.TestPrivateNonce1, crypto.CombinedPublicNonces1)
	signature2, _ := crypto.PartialSignMessage(message, crypto.TestPrivateKey2,
		crypto.TestPrivateNonce2, crypto.CombinedPublicNonces2)
	signature3, _ := crypto.PartialSignMessage(message, crypto.TestPrivateKey3,
		crypto.TestPrivateNonce3, crypto.CombinedPublicNonces3)
	combinedSignatures, _ := crypto.CombineSignatures([][]byte{signature1,
		signature2, signature3})
	return combinedSignatures
}
