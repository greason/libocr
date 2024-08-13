package ocr3confighelper

import (
	"crypto/ed25519"
	"math/big"
)

type KeyV2 struct {
	privateKey *ed25519.PrivateKey
	PublicKey  ed25519.PublicKey
	Version    int
}

func MustNewV2XXXTestingOnly(k *big.Int) KeyV2 {
	seed := make([]byte, ed25519.SeedSize)
	copy(seed, k.Bytes())
	privKey := ed25519.NewKeyFromSeed(seed)
	return KeyV2{
		privateKey: &privKey,
		PublicKey:  ed25519PubKeyFromPrivKey(privKey),
		Version:    2,
	}
}

func ed25519PubKeyFromPrivKey(privKey ed25519.PrivateKey) ed25519.PublicKey {
	publicKey := make([]byte, ed25519.PublicKeySize)
	copy(publicKey, privKey[32:])
	return publicKey
}
