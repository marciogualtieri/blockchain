package block

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTXBlockHeaderSerializationWithValidObject(t *testing.T) {
	serializedObject, err := TestTXBlockHeader.Serialize()
	assert.Nil(t, err)
	assert.NotNil(t, serializedObject)
	deserializedObject, err := DeserializeTXBlockHeader(serializedObject)
	assert.Nil(t, err)
	assert.Equal(t, TestTXBlockHeader, *deserializedObject)
}

func TestTXBlockHeaderSerializationWithCorruptObject(t *testing.T) {
	serializedObject, _ := TestTXBlockHeader.Serialize()
	serializedObject[0] = 0x77
	_, err := DeserializeTXBlockHeader(serializedObject)
	assert.NotNil(t, err)
}
