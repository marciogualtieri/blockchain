# Blockchain Multisignature (Zilliqa's approach) in golang

## Overview

The purpose of this project is to showcase my knowledge of blockchain technologies through the implementation of a multisignature consensus scheme in the [Go language](https://golang.org/) (thus, it additionally serves the purpose of showcasing my OOP programming skills in this language).

Specifically, I'm implementing the multisignature consensus scheme used by [Zilliqa](https://zilliqa.com/)'s blockchain.

Zilliqa's blockchain implements sharding to solve the escalability issue found in other current blockchains, such as [Etherium](https://www.ethereum.org/) and [BitCoin](https://bitcoin.org/en/).

In this particular blockchain, an elected dynamic group of nodes (a shard of nodes, named the "DS Committee") should reach an agreement about if a block should or shouldn't be added to the blockchain. The committee members vote and produce a bit map of signatures, which is validated using a [Schnoor multisignature scheme](https://medium.com/@blairlmarshall/signature-verification-multi-signatures-19886fafe97b).

For a complete description of sharding and consensus in Zilliqa, refer to [the white paper](https://docs.zilliqa.com/whitepaper.pdf). For a general overview on Zilliqa's sharding, you may refer to the following blog articles:

- (The Zilliqa Design Story Piece by Piece: Part 1 (Network Sharding))[https://blog.zilliqa.com/https-blog-zilliqa-com-the-zilliqa-design-story-piece-by-piece-part1-d9cb32ea1e65]

- (The Zilliqa Design Story Piece by Piece: Part 2 (Consensus Protocol))[https://blog.zilliqa.com/the-zilliqa-design-story-piece-by-piece-part-2-consensus-protocol-e38f6bf566e3]

- (The Zilliqa Design Story Piece by Piece: Part 3 (Making Consensus Efficient))[https://blog.zilliqa.com/the-zilliqa-design-story-piece-by-piece-part-3-making-consensus-efficient-7a9c569a8f0e]


For a an overview of Zilliqa's blockchain current functionalities, you may refer to [this article](https://www.coinbureau.com/review/zilliqa-zil/).

## Design

I'm implementing specifically the 2nd round consensus verification, that is, the following Go function:

    func (node Node) VerifyFinalBlockConsensusSignature(txBlock block.TXBlock) (bool, error) { ... }

This function receives a final block and validates its signature map using a Schnorr multisignature scheme.

It's based on the following C++ function from Zalliqa's [source code](https://github.com/Zilliqa/Zilliqa/tree/master/src/libNode):

    bool Node::VerifyFinalBlockCoSignature(const TxBlock& txblock) { ... }

The implementation also includes all supporting code, including the supporting code necessary for testing.

The following UML diagrams shows the classes (Golang's receivers) implemented in the code:

![Blockchain Classes](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.githubusercontent.com/marciogualtieri/blockchain/master/uml/classes.plantuml)

Additionally the following packages with helper functions have been implemented:

![Blockchain Packages](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.githubusercontent.com/marciogualtieri/blockchain/master/uml/packages.plantuml)

## Installing Dependencies

The following command will install all project package dependencies:

    go get ./...

## Running Tests

To run all tests, execute the following command:

  $ go test -v ./...

You should get an output similar to the following:

  === RUN   TestBitArrayToByteArray
  --- PASS: TestBitArrayToByteArray (0.00s)
  PASS
  ok  	blockchain/utils/collection	0.005s
  === RUN   TestGenerateKeys
  --- PASS: TestGenerateKeys (0.00s)
  === RUN   TestAggregatePublicKeys
  --- PASS: TestAggregatePublicKeys (0.00s)
  === RUN   TestCombineSignatures
  --- PASS: TestCombineSignatures (0.00s)
  === RUN   TestVerifySingleSignature
  --- PASS: TestVerifySingleSignature (0.00s)
  === RUN   TestVerifyMultiplePartialSignature
  --- PASS: TestVerifyMultiplePartialSignature (0.00s)
  PASS
  ok  	blockchain/utils/crypto	0.124s
  === RUN   TestNumForConsensus
  --- PASS: TestNumForConsensus (0.00s)
  PASS
  ok  	blockchain/utils/network	0.100s
  === RUN   TestTXBlockHeaderSerializationWithValidObject
  --- PASS: TestTXBlockHeaderSerializationWithValidObject (0.00s)
  === RUN   TestTXBlockHeaderSerializationWithCorruptObject
  --- PASS: TestTXBlockHeaderSerializationWithCorruptObject (0.00s)
  PASS
  ok  	blockchain/zilliqa/block	0.109s
  === RUN   TestVerifyFinalBlockConsensusSignature
  --- PASS: TestVerifyFinalBlockConsensusSignature (0.00s)
  === RUN   TestVerifyFinalBlockConsensusSignatureNotEnoughVotes
  --- PASS: TestVerifyFinalBlockConsensusSignatureNotEnoughVotes (0.00s)
  === RUN   TestVerifyFinalBlockConsensusSignatureCorruptSignature
  --- PASS: TestVerifyFinalBlockConsensusSignatureCorruptSignature (0.00s)
  === RUN   TestVerifyFinalBlockConsensusSignatureCorruptSignatureMap
  --- PASS: TestVerifyFinalBlockConsensusSignatureCorruptSignatureMap (0.00s)
  PASS
  ok  	blockchain/zilliqa/node	0.082s