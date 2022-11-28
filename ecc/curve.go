package ecc

import (
	"bytes"
	"compress/gzip"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"strings"
)

func GenerateKey() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

// ECDSA Sign
func Sign(privateKey *ecdsa.PrivateKey, messageHash string) (string, error) {
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, []byte(messageHash))
	if err != nil {
		return "", err
	}
	rStr, _ := r.MarshalText()
	sStr, _ := s.MarshalText()
	var result bytes.Buffer
	w := gzip.NewWriter(&result)
	defer w.Close()
	_, err = w.Write([]byte(string(rStr) + "+" + string(sStr)))
	if err != nil {
		return "", err
	}
	w.Flush()
	return hex.EncodeToString(result.Bytes()), nil
}

// ECDSA Verify
func Verify(messageHash, signature string, publicKey *ecdsa.PublicKey) (bool, error) {
	sigBytes, err := hex.DecodeString(signature)
	if err != nil {
		return false, err
	}
	reader, err := gzip.NewReader(bytes.NewBuffer(sigBytes))
	if err != nil {
		return false, err
	}
	defer reader.Close()
	buf := make([]byte, 1024)
	count, err := reader.Read(buf)
	if err != nil {
		return false, err
	}
	rsArr := strings.Split(string(buf[:count]), "+")
	if len(rsArr) != 2 {
		return false, err
	}
	var r, s big.Int
	err = r.UnmarshalText([]byte(rsArr[0]))
	if err != nil {
		return false, err
	}
	err = s.UnmarshalText([]byte(rsArr[1]))
	if err != nil {
		return false, err
	}
	result := ecdsa.Verify(publicKey, []byte(messageHash), &r, &s)
	return result, nil
}
