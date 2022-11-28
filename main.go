package main

import (
	"fmt"

	cpabe "github.com/ylx167167/cryptography/cp-abe"
	"github.com/ylx167167/cryptography/ecc"
)

func main() {

	// test7()
	test8()
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

	a, b := ecc.GenerateeciesKey(apri)
	fmt.Print("\n")
	fmt.Print(a)
	fmt.Print("\n")
	fmt.Print(b)
}

func test8() {
	apri, _, _ := ecc.ReadKeyPairFile("A_prikey.pem", "A_pubkey.pem")
	a, b := ecc.GenerateeciesKey(apri)
	var message string = "this is a message, 这是需要加密的数据"
	e := ecc.Encrypt_ecies(message, b)

	d := ecc.Decrypt_ecies(e, a)
	fmt.Print(d)
}
