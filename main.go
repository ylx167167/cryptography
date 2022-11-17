package main

import (
	cpabe "github.com/ylx167167/cryptography/cp-abe"
)

func main() {
	pairing := cpabe.SetupSingularPairing()
	var msp cpabe.MSP
	cpabe.MspSetup(&msp, "attribute.dat")
	cpabe.Setup(msp.Rows, *pairing, &msp, "setupPairingTime.txt")
	cpabe.User_fread("user.file")
	userNo, attrNo, userName, attribute := cpabe.User_fread("user.file")
	for i := 0; i < userNo; i++ {
		cpabe.KeyGen(*pairing, attrNo[i], attribute[i], string(userName[i]))
	}
}
