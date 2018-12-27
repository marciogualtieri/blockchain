package network

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumForConsensus(t *testing.T) {
	assert.Equal(t, 2, TestMediator.NumberForConsensus(3))
}
