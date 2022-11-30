package main

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"

	cpabe "github.com/ylx167167/cryptography/cp-abe"
	"github.com/ylx167167/cryptography/ecc"
)

func main() {

	// test7()
	test8()
	// test9("test")
	// test10()
}
func test5() {
	pairing := cpabe.SetupSingularPairing()
	var msp cpabe.MSP
	cpabe.MspSetup(&msp, "attribute.dat")
	cpabe.Setup(msp.Rows, *pairing, &msp, "setupPairingTime.txt")
	cpabe.User_fread("user.file")
	userNo, attrNo, userName, attribute := cpabe.User_fread("user.file")
	for i := 0; i < userNo; i++ {
		cpabe.KeyGen(*pairing, attrNo[i], attribute[i], string(userName[i]))
	}
	// args:=
}
func test6() {
	pri, pub, err := ecc.ReadKeyPairFile("A_prikey.pem", "A_pubkey.pem")
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Print(pri)
	fmt.Print(pub)
}

func test7() {
	apri, apub, _ := ecc.ReadKeyPairFile("A_prikey.pem", "A_pubkey.pem")
	fmt.Print(apri)
	fmt.Print("\n")
	fmt.Print(apub)

	a, b := ecc.GenerateeciesPri(apri)
	fmt.Print("\n")
	fmt.Print(a)
	fmt.Print("\n")
	fmt.Print(b)
}

func test8() {
	apri, apub, _ := ecc.ReadKeyPairFile("A_prikey.pem", "A_pubkey.pem")
	// a, b := ecc.GenerateeciesPri(apri)
	var message string = "this is a message, 这是需要加密的数据"
	e := ecc.Encrypt(message, apub)
	d := ecc.Decrypt(e, apri)
	fmt.Print(d)
}

func test9(s string) {
	// s := "test"
	// hash := sha256.New()
	// hash.Write([]byte(s))
	// bs := hash.Sum(nil)
	// fmt.Print(string(bs))
	h := sha256.New()
	h.Write([]byte(s))
	fmt.Printf("%x", h.Sum(nil))
}

// 生成既能签名和加密的算法
func test10() {
	iespri, iespub, _ := ecc.GenerateKeyECIES()
	var dsapub ecdsa.PublicKey = ecdsa.PublicKey{
		X:     iespub.X,
		Y:     iespri.Y,
		Curve: iespub.Curve,
	}
	var dsapri ecdsa.PrivateKey = ecdsa.PrivateKey{
		PublicKey: dsapub,
		D:         iespri.D,
	}
	message := "test"
	signature, _ := ecc.Sign(&dsapri, message)
	T, _ := ecc.Verify(message, signature, &dsapub)
	fmt.Print(T)
}
