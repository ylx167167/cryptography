package ecc

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"

	"github.com/obscuren/ecies"
)

func GenerateeciesKey(pri *ecdsa.PrivateKey) (*ecies.PrivateKey, *ecies.PublicKey) {
	eciesPri := ecies.ImportECDSA(pri)
	eciesPub := &eciesPri.PublicKey
	return eciesPri, eciesPub
}

func Encrypt_ecies(message string, eciesPublicKey *ecies.PublicKey) string {
	cipherBytes, _ := ecies.Encrypt(rand.Reader, eciesPublicKey, []byte(message), nil, nil)
	cipherString := hex.EncodeToString(cipherBytes)
	return cipherString
}

func Decrypt_ecies(message string, eciesPrivateKey *ecies.PrivateKey) string {
	bytes, _ := hex.DecodeString(message)
	decrypeMessageBytes, _ := eciesPrivateKey.Decrypt(rand.Reader, bytes, nil, nil)
	decrypeMessageString := string(decrypeMessageBytes[:])
	return decrypeMessageString
}
