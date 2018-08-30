package crypto

import (
	"encoding/hex"
	"fmt"
	"strings"
	"bytes"
	"github.com/slackve/nem2-sdk-go/core/crypto/ed25519"
	"io"
	"crypto/rand"
)

// KeyPair is a private/public crypto key pair.
type KeyPair struct {
	Private []byte
	Public  []byte
}

const (
	// PublicKeySize is the size, in bytes, of public keys as used in this package.
	PrivateBytes = 32
	// SignatureSize is the size, in bytes, of signatures generated and verified by this package.
	SignatureSize     = 64
	HalfSignatureSize = SignatureSize / 2
	HashSize          = 64
	HalfHashSize      = HashSize / 2
)

// Creates a key pair from a private key string.
// param privateKeyString A hex encoded private key string.
// returns The key pair.
func KeyPairFromPrivateKey(privateKeyString string) (keypair KeyPair, err error) {
	if privateKeyString != "" {
		privateKey, err := hex.DecodeString(strings.TrimSpace(privateKeyString))
		if PrivateBytes != len(privateKey) {
			err := fmt.Errorf("private key has unexpected size: %v\n", len(privateKey))
			return KeyPair{}, err
		}
		if err != nil {
			panic(err)
		}

		pair, err := FromSeed(privateKey)
		if err != nil {
			panic(err)
		}
		return pair, nil
	}

	seed := make([]byte, PrivateBytes)
	_, err = io.ReadFull(rand.Reader, seed)
	if err != nil {
		panic(err)
	}
	pair, err := FromSeed(seed)
	if err != nil {
		panic(err)
	}
	return pair, nil
}

func ExtractPublicKey()  {

}

// FromSeed generates a new private/public key pair using specified private key.
// param pk - A PrivateBytes
// return - The NEM KeyPair
func FromSeed(seed []byte) (KeyPair, error) {
	if len(seed) != PrivateBytes {
		return KeyPair{},
			fmt.Errorf("insufficient seed length, should be %d, but got %d", 64, len(seed))
	}
	pub, pr, err := ed25519.GenerateKey(bytes.NewReader(seed))
	if err != nil {
		panic(err)
	}
	return KeyPair{pr[:PrivateBytes], pub}, nil
}