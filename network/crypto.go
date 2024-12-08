package network

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func SignOrder(message string) (signature []byte, publicKey *ecdsa.PublicKey, err error) {
	// Generate private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, nil, err
	}

	// Hash the message
	hash := sha256.Sum256([]byte(message))

	// Sign the hash
	signature, err = crypto.Sign(hash[:], privateKey)
	if err != nil {
		return nil, nil, err
	}

	return signature, &privateKey.PublicKey, nil
}

func PublicKeyToEthereumAddress(publicKey *ecdsa.PublicKey) string {
	// Convert public key to bytes
	publicKeyBytes := crypto.FromECDSAPub(publicKey)

	// Keccak-256 hash of public key
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:]) // Remove first byte (0x04 prefix)
	hashBytes := hash.Sum(nil)

	// Take last 20 bytes
	address := hashBytes[12:]

	// Convert to hexadecimal with 0x prefix
	return "0x" + hex.EncodeToString(address)
}