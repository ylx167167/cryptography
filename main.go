package main

import "github.com/ylx167167/cryptography/cpabe"

func main() {
	pairing := cpabe.SetupSingularPairing()
	var msp cpabe.MSP
	cpabe.MspSetup(&msp, "attribute.dat")
	cpabe.Setup(msp.Rows, *pairing, &msp, "setupPairingTime.txt")
}
