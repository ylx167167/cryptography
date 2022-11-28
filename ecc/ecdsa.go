package ecc

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

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
