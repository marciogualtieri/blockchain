package crypto

import (
	"github.com/decred/dcrd/dcrec/secp256k1/schnorr"
	"math/big"
)

/*
TestMessage is the byte array message used in tests.
*/
var TestMessage = []byte("Test Message")

/*
TestSignature1 is a partial signature used in tests.
*/
var TestSignature1, _ = PartialSignMessage(TestMessage, TestPrivateKey1, TestPrivateNonce1, CombinedPublicNonces1)

/*
TestSignature2 is a partial signature used in tests.
*/
var TestSignature2, _ = PartialSignMessage(TestMessage, testPrivateKey2, TestPrivateNonce2, CombinedPublicNonces2)

/*
TestSignature3 is a partial signature used in tests.
*/
var TestSignature3, _ = PartialSignMessage(TestMessage, testPrivateKey3, TestPrivateNonce3, CombinedPublicNonces3)

var testPrivateNonce1, testPublicNonce1, _ = GenerateNounce(TestMessage, TestPrivateKey1)
var testPrivateNonce2, testPublicNonce2, _ = GenerateNounce(TestMessage, testPrivateKey2)
var testPrivateNonce3, testPublicNonce3, _ = GenerateNounce(TestMessage, testPrivateKey3)

/*
TestPrivateNonce1 is a private nonce used in tests.
*/
var TestPrivateNonce1 = testPrivateNonce1

/*
TestPrivateNonce2 is a private nonce used in tests.
*/
var TestPrivateNonce2 = testPrivateNonce2

/*
TestPrivateNonce3 is a private nonce used in tests.
*/
var TestPrivateNonce3 = testPrivateNonce3

/*
TestPublicNonce1 is a public nonce used in tests.
*/
var TestPublicNonce1 = testPublicNonce1

/*
TestPublicNonce2 is a public nonce used in tests.
*/
var TestPublicNonce2 = testPublicNonce2

/*
TestPublicNonce3 is a public nonce used in tests.
*/
var TestPublicNonce3 = testPublicNonce3

/*
TestSignature is the signature used in tests.
*/
var TestSignature = schnorr.NewSignature(big.NewInt(123), big.NewInt(456)).Serialize()

var testPrivateKey1, testPublicKey1, _ = GenerateKeys()

/*
TestPrivateKey1 is a private key used in tests.
*/
var TestPrivateKey1 = testPrivateKey1

/*
TestPublicKey1 is a public key used in tests.
*/
var TestPublicKey1 = testPublicKey1

var testPrivateKey2, testPublicKey2, _ = GenerateKeys()

/*
TestPrivateKey2 is a private key used in tests.
*/
var TestPrivateKey2 = testPrivateKey2

/*
TestPublicKey2 is a public key used in tests.
*/
var TestPublicKey2 = testPublicKey2

var testPrivateKey3, testPublicKey3, _ = GenerateKeys()

/*
TestPrivateKey3 is a private key used in tests.
*/
var TestPrivateKey3 = testPrivateKey3

/*
TestPublicKey3 is a public key used in tests.
*/
var TestPublicKey3 = testPublicKey3

/*
CombinedPublicKey is the aggregated public key used in tests.
*/
var CombinedPublicKey, _ = CombinePublicKeys([][]byte{TestPublicKey1, TestPublicKey2, TestPublicKey3})

/*
CombinedPublicNonces1 is a aggregated public nounce used in tests.
*/
var CombinedPublicNonces1, _ = CombinePublicKeys([][]byte{TestPublicNonce2, TestPublicNonce3})

/*
CombinedPublicNonces2 is a aggregated public nounce used in tests.
*/
var CombinedPublicNonces2, _ = CombinePublicKeys([][]byte{TestPublicNonce1, TestPublicNonce3})

/*
CombinedPublicNonces3 is a aggregated public nounce used in tests.
*/
var CombinedPublicNonces3, _ = CombinePublicKeys([][]byte{TestPublicNonce1, TestPublicNonce2})
