/*
Package crypto implements cryptographic utilities
*/
package crypto

import (
	"crypto/rand"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/decred/dcrd/dcrec/secp256k1"
	"github.com/decred/dcrd/dcrec/secp256k1/schnorr"
)

/*
GenerateKeys generates the private and public keys.
*/
func GenerateKeys() ([]byte, []byte, error) {
	privateKeyBytes, _, _, err := secp256k1.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	privateKey, publicKey := secp256k1.PrivKeyFromBytes(privateKeyBytes)
	return privateKey.Serialize(), publicKey.Serialize(), nil
}

/*
GenerateNounce generates the private and public nonce for the private key.
*/
func GenerateNounce(message []byte, privateKeyBytes []byte) ([]byte, []byte, error) {
	privateKey, _ := secp256k1.PrivKeyFromBytes(privateKeyBytes)
	privateNonce, publicNonce, err := schnorr.GenerateNoncePair(nil, chainhash.HashB(message),
		privateKey, nil, nil)
	if err != nil {
		return nil, nil, err
	}
	return privateNonce.Serialize(), publicNonce.Serialize(), nil
}

/*
CombinePublicKeys combines the input public keys for multi-signature.
*/
func CombinePublicKeys(publicKeysBytes [][]byte) ([]byte, error) {
	publicKeys, err := parsePublicKeysBytes(publicKeysBytes)
	if err != nil {
		return nil, err
	}
	aggregatedKeys := schnorr.CombinePubkeys(publicKeys)
	return aggregatedKeys.Serialize(), nil
}

func parsePublicKeysBytes(publicKeysBytes [][]byte) ([]*secp256k1.PublicKey, error) {
	keys := make([]*secp256k1.PublicKey, cap(publicKeysBytes))
	for i, keyBytes := range publicKeysBytes {
		key, err := secp256k1.ParsePubKey(keyBytes)
		if err != nil {
			return nil, err
		}
		keys[i] = key
	}
	return keys, nil
}

/*
PartialSignMessage partially signs the message for multi-signature.
*/
func PartialSignMessage(message []byte, privateKeyBytes []byte, privateNonceBytes []byte,
	CombinedPublicNoncesBytes []byte) ([]byte, error) {
	privateKey, _ := secp256k1.PrivKeyFromBytes(privateKeyBytes)
	privateNonce, _ := secp256k1.PrivKeyFromBytes(privateNonceBytes)
	CombinedPublicNonces, err := secp256k1.ParsePubKey(CombinedPublicNoncesBytes)
	if err != nil {
		return nil, err
	}
	signature, err := schnorr.PartialSign(secp256k1.S256(), chainhash.HashB(message),
		privateKey, privateNonce, CombinedPublicNonces)

	if err != nil {
		return nil, err
	}

	return signature.Serialize(), nil
}

/*
SignMessage single signs the message.
*/
func SignMessage(message []byte, privateKeyBytes []byte) ([]byte, error) {
	signature, err := signMessage(message, privateKeyBytes)
	if err != nil {
		return nil, err
	}
	signatureBytes := signature.Serialize()
	return signatureBytes, nil
}

func signMessage(message []byte, privateKeyBytes []byte) (*schnorr.Signature, error) {
	privateKey, _ := secp256k1.PrivKeyFromBytes(privateKeyBytes)
	r, s, err := schnorr.Sign(privateKey, chainhash.HashB(message))
	if err != nil {
		return nil, err
	}
	signature := schnorr.NewSignature(r, s)
	return signature, nil
}

/*
CombineSignatures combines signatures for multi sign.
*/
func CombineSignatures(signaturesBytes [][]byte) ([]byte, error) {
	signatures, err := parseSignaturesBytes(signaturesBytes)
	if err != nil {
		return nil, err
	}
	combinedSignature, err := schnorr.CombineSigs(secp256k1.S256(), signatures)
	if err != nil {
		return nil, err
	}
	return combinedSignature.Serialize(), err
}

func parseSignaturesBytes(signaturesBytes [][]byte) ([]*schnorr.Signature, error) {
	signatures := make([]*schnorr.Signature, cap(signaturesBytes))
	for i, signatureBytes := range signaturesBytes {
		signature, err := schnorr.ParseSignature(signatureBytes)
		if err != nil {
			return nil, err
		}
		signatures[i] = signature
	}
	return signatures, nil
}

/*
VerifySignature verifies a signature...
*/
func VerifySignature(message []byte, signatureBytes []byte,
	publicKeyBytes []byte) (bool, error) {
	signature, err := schnorr.ParseSignature(signatureBytes)
	if err != nil {
		return false, err
	}
	publicKey, err := secp256k1.ParsePubKey(publicKeyBytes)
	if err != nil {
		return false, err
	}
	return schnorr.Verify(publicKey, chainhash.HashB(message), signature.R, signature.S), nil
}
