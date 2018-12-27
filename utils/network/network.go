/*
Package network contains Zilliqa's network entities, such as peers, nodes, etc
*/
package network

import "math"

/*
Peer struct represents a machine
*/
type Peer struct {
	Address string
	Port    uint32
}

/*
DSEntrusted structs represents a DS committee member
*/
type DSEntrusted struct {
	PublicKey []byte
	Peer      Peer
}

/*
Mediator struct stores global variables
*/
type Mediator struct {
	DSCommittee       []DSEntrusted
	ToleranceFraction float64
}

/*
NumberForConsensus is the mininum number of signatures necessary for consensus...
*/
func (mediator Mediator) NumberForConsensus(shardSize int) int {
	return int(math.Ceil(float64(shardSize) * mediator.ToleranceFraction))
}
