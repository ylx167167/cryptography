package ecc

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"github.com/obscuren/ecies"
)

func GenerateKeyECIES() (*ecies.PrivateKey, *ecies.PublicKey, error) {
	privateKey, err := ecies.GenerateKey(rand.Reader, elliptic.P256(), ecies.ParamsFromCurve(elliptic.P256()))
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

func GenerateeciesPri(pri *ecdsa.PrivateKey) (*ecies.PrivateKey, *ecies.PublicKey) {
	eciesPri := ecies.ImportECDSA(pri)
	eciesPub := &eciesPri.PublicKey
	return eciesPri, eciesPub
}

func GenerateciesPub(pub *ecdsa.PublicKey) *ecies.PublicKey {
	return &ecies.PublicKey{
		X:      pub.X,
		Y:      pub.Y,
		Curve:  pub.Curve,
		Params: ecies.ParamsFromCurve(pub.Curve),
	}
}
