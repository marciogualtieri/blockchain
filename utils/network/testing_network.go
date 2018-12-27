package network

import (
	"blockchain/utils/crypto"
)

/*
TestPeer1 is a network peer used in tests.
*/
var TestPeer1 = Peer{Address: "host1", Port: 111}

/*
TestPeer2 is a network peer used in tests.
*/
var TestPeer2 = Peer{Address: "host2", Port: 222}

/*
TestPeer3 is a network peer used in tests.
*/
var TestPeer3 = Peer{Address: "host3", Port: 333}

/*
TestAggregatedPublicKey is the DS committee aggregated public key used in tests.
*/
var TestAggregatedPublicKey, _ = crypto.CombinePublicKeys(
	[][]byte{crypto.TestPublicKey1, crypto.TestPublicKey2, crypto.TestPublicKey3})

/*
CommitteeMember1 is a member of the DS commmitte used in tests.
*/
var CommitteeMember1 = DSEntrusted{Peer: TestPeer1,
	PublicKey: crypto.TestPublicKey1}

/*
CommitteeMember2 is a member of the DS commmitte used in tests.
*/
var CommitteeMember2 = DSEntrusted{Peer: TestPeer2,
	PublicKey: crypto.TestPublicKey2}

/*
CommitteeMember3 is a member of the DS commmitte used in tests.
*/
var CommitteeMember3 = DSEntrusted{Peer: TestPeer3,
	PublicKey: crypto.TestPublicKey3}

func createTestMediator() Mediator {
	mediator := Mediator{
		DSCommittee:       make([]DSEntrusted, 3),
		ToleranceFraction: float64(2) / 3}
	mediator.DSCommittee[0] = CommitteeMember1
	mediator.DSCommittee[1] = CommitteeMember2
	mediator.DSCommittee[2] = CommitteeMember3
	return mediator
}

/*
TestMediator is the mediator used in tests.
*/
var TestMediator = createTestMediator()
