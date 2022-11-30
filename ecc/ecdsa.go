package ecc

import (
	"bytes"
	"compress/gzip"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"math/big"
	"os"
	"strings"

	"github.com/obscuren/ecies"
)

func GenerateKeyECDSA() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
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

func Encrypt(message string, ecdsaPublicKey *ecdsa.PublicKey) string {
	eciesPublicKey := GenerateciesPub(ecdsaPublicKey)
	cipherBytes, _ := ecies.Encrypt(rand.Reader, eciesPublicKey, []byte(message), nil, nil)
	cipherString := hex.EncodeToString(cipherBytes)
	return cipherString
}

func Decrypt(message string, ecdsaPrivateKey *ecdsa.PrivateKey) string {
	eciesPrivateKey, _ := GenerateeciesPri(ecdsaPrivateKey)
	bytes, _ := hex.DecodeString(message)
	decrypeMessageBytes, _ := eciesPrivateKey.Decrypt(rand.Reader, bytes, nil, nil)
	decrypeMessageString := string(decrypeMessageBytes[:])
	return decrypeMessageString
}

func SavekeyPair(pri *ecdsa.PrivateKey, pub *ecdsa.PublicKey) (*pem.Block, *pem.Block) {
	priderStream, _ := x509.MarshalECPrivateKey(pri)
	priblock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: priderStream,
	}

	pubderStream, _ := x509.MarshalPKIXPublicKey(pub)
	pubblock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubderStream,
	}
	return priblock, pubblock
}
func SavekeyPairFile(pri *ecdsa.PrivateKey, pub *ecdsa.PublicKey, ID string, path string) error {
	priblock, pubblock := SavekeyPair(pri, pub)
	pripath := ID + "_prikey.pem"
	pubpath := ID + "_pubkey.pem"
	prifile, err1 := os.Create(pripath)
	pubfile, err2 := os.Create(pubpath)
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	var err error
	err = pem.Encode(prifile, priblock)
	if err != nil {
		return err
	}
	err = pem.Encode(pubfile, pubblock)
	if err != nil {
		return err
	}
	return nil
}

func ReadKeyPair(pripem []byte, pubpem []byte) (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	priblock, _ := pem.Decode(pripem)
	privateKey, _ := x509.ParseECPrivateKey(priblock.Bytes)
	pubblock, _ := pem.Decode(pubpem)
	x509EncodedPub := pubblock.Bytes
	genericpublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
	publickey := genericpublicKey.(*ecdsa.PublicKey)
	return privateKey, publickey, nil

}

func ReadKeyPairFile(pripath string, pubpath string) (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	//打开文件
	prifile, err := os.Open(pripath)
	if err != nil {
		return nil, nil, err
	}
	pubfile, err := os.Open(pubpath)
	if err != nil {
		return nil, nil, err
	}
	pristat, err := prifile.Stat()
	if err != nil {
		return nil, nil, err
	}
	pubstat, err := pubfile.Stat()
	if err != nil {
		return nil, nil, err
	}
	pribuf := make([]byte, pristat.Size())
	pubbuf := make([]byte, pubstat.Size())
	prifile.Read(pribuf)
	pubfile.Read(pubbuf)
	defer prifile.Close()
	defer pubfile.Close()
	privateKey, publicKey, err := ReadKeyPair(pribuf, pubbuf)
	return privateKey, publicKey, err
}
