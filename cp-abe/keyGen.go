package cpabe

import (
	"fmt"
	"os"

	"github.com/Nik-U/pbc"
)

func KeyGen(pairing pbc.Pairing, attrNo int, attribute []byte, userName string) {
	fMsk, err1 := os.OpenFile("MSK/msk.key", os.O_RDWR|os.O_CREATE, 0777)    //fMsk to read the master key
	fG, err2 := os.OpenFile("publicKey/g.key", os.O_RDWR|os.O_CREATE, 0777)  //fG to read the public key -genterator g
	fGA, err := os.OpenFile("publicKey/gA.key", os.O_RDWR|os.O_CREATE, 0777) //fGA to read the public key -gA
	if err != nil || err2 != nil || err1 != nil {
		fmt.Print("os.OpenFile failure")
	}
	defer fMsk.Close()
	defer fG.Close()
	defer fGA.Close()            //close all file pointer
	var fH *os.File              //fH to read the public key-h attribute
	hCmd := make([]byte, 0, 100) //the command line for the pointer of FILE* fH
	var attrName string          //the name of attribute

	g := pairing.NewG2()
	Element_fread(fG, "%s %s", g, 10)
	msk := pairing.NewG2()
	Element_fread(fMsk, "%s %s", msk, 10)
	gA := pairing.NewG2()
	Element_fread(fGA, "%s %s", gA, 10)
	h := make([]pbc.Element, attrNo)
	for i := 0; i < attrNo; i++ {
		hCmd = append(hCmd, "publicKey/h"...)
		attrName = fmt.Sprintf("%c", attribute[i])
		hCmd = append(hCmd, []byte(attrName)...)
		hCmd = append(hCmd, ".key"...)
		fmt.Print(string(hCmd), "\n")
		fH, err = os.OpenFile(string(hCmd), os.O_RDWR|os.O_CREATE, 0777)
		if err != nil {
			fmt.Print("os.OpenFile failure\n")
		}
		h[i] = *pairing.NewG2()
		Element_fread(fH, "%s %s\n", &h[i], 10)

		hCmd = append(hCmd[:0], hCmd[(len(hCmd)):]...) //清空buffer
		fH.Close()
	}

	//start to calculate private key and write file
	t := pairing.NewZr()
	L := pairing.NewG2()
	K := pairing.NewG2()
	Kx := pairing.NewG2()
	temp := pairing.NewG2()
	fileL := make([]byte, 0, 100)
	fileK := make([]byte, 0, 100)
	fileKx := make([]byte, 0, 100)
	fileL = append(fileL, []byte(userName)...)
	fileL = append(fileL, "/L.key"...)
	fileK = append(fileK, []byte(userName)...)
	fileK = append(fileK, "/K.key"...)
	fileKx = append(fileKx, []byte(userName)...)
	fileKx = append(fileKx, "/Kx.key"...)
	fL, err := os.OpenFile(string(fileL), os.O_RDWR|os.O_CREATE, 0777)    //fL to write the privateKey L
	fK, err1 := os.OpenFile(string(fileK), os.O_RDWR|os.O_CREATE, 0777)   //fK to write the privateKey K
	fKx, err2 := os.OpenFile(string(fileKx), os.O_RDWR|os.O_CREATE, 0777) //fKx to write the privateKey Kx
	if err != nil || err2 != nil || err1 != nil {
		fmt.Print("os.OpenFile failure")
	}
	defer fL.Close()
	defer fK.Close()
	defer fKx.Close()

	t.Rand()
	L.PowZn(g, t)
	fmt.Fprintf(fL, "%s", L)
	temp.PowZn(gA, t) //first K = g^at
	K.Mul(temp, msk)  //second K = K*g^alpha
	fmt.Fprintf(fK, "%s", K)
	for i := 0; i < attrNo; i++ {
		Kx.Set0()
		Kx.PowZn(&h[i], t)           //Kx = hx^t
		fmt.Fprintf(fKx, "%s\n", Kx) //Kx = hx^t
	}

}
