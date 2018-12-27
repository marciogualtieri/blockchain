/*
Package block implements Zilliqa's blocks, e.g., TX-Block, Micro Block,
Final Block.
*/
package block

import (
	"blockchain/utils/collection"
	"bytes"
	"encoding/gob"
	"github.com/golang-collections/go-datastructures/bitarray"
)

/*
MicroBlockMarker marks a micro block proposed by the committee.
*/
const MicroBlockMarker = 0

/*
FinalBlockMarker marks a final block.
*/
const FinalBlockMarker = 1

/*
TXBlock struct represents a TX-BLOCK.
*/
type TXBlock struct {
	ConsensusSignatureMap1stRound bitarray.BitArray
	ConsensusSignatureMap2ndRound bitarray.BitArray
	ConsensusSignature1stRound    []byte
	ConsensusSignature2ndRound    []byte
	Header                        TXBlockHeader
}

/*
TXBlockHashSet struct contains the Merkle tree root hash, the state delta
hash, and the hash of all its micro blocks info.
*/
type TXBlockHashSet struct {
	MerkleRootHash      []byte
	StateDeltaHash      []byte
	MicroBlocksInfoHash []byte
}

/*
TXBlockHeader struct contains header info for a TX-Block.
*/
type TXBlockHeader struct {
	BlockType      uint8
	Version        uint32
	GasLimit       uint64
	RewardsHigh    uint64
	RewardsLow     uint64
	PreviousHash   []byte
	BlockNum       uint64
	HashSet        TXBlockHashSet
	NumTXs         uint32
	MinerPublicKey []byte
	DSBlockNum     uint64
}

/*
Serialize TXBlockHeader
*/
func (object TXBlockHeader) Serialize() ([]byte, error) {
	bytesBuffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&bytesBuffer)
	err := encoder.Encode(&object)
	if err == nil {
		return bytesBuffer.Bytes(), nil
	}
	return nil, err
}

/*
Message (header + signature map size in bits + signature map size in bytes +
signature map) as a byte array
*/
func (txBlock TXBlock) Message() ([]byte, error) {
	message := txBlock.ConsensusSignature1stRound
	headerBytes, err := txBlock.Header.Serialize()
	if err != nil {
		return nil, err
	}
	signaturesBytes := collection.BitArrayToByteArray(txBlock.ConsensusSignatureMap1stRound)
	message = append(message, headerBytes...)
	message = append(message, byte(txBlock.ConsensusSignatureMap1stRound.Capacity()))
	message = append(message, byte(cap(signaturesBytes)))
	message = append(message, signaturesBytes...)
	return message, nil
}

/*
DeserializeTXBlockHeader is the deserializer for TXBlockHeader...
*/
func DeserializeTXBlockHeader(object []byte) (*TXBlockHeader, error) {
	header := TXBlockHeader{}
	bytesBuffer := bytes.Buffer{}
	_, err := bytesBuffer.Write(object)
	if err == nil {
		decoder := gob.NewDecoder(&bytesBuffer)
		err = decoder.Decode(&header)
		if err == nil {
			return &header, nil
		}
	}
	return nil, err
}
