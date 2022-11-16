package cpabe

import (
	"fmt"
	"os"
	"time"

	"github.com/Nik-U/pbc"
)

func GeneratePrime() {
	// params := pbc(160, 512)
	//test
}

func SetupSingularPairing() *pbc.Pairing {
	var rbits uint32 = 256
	var qbits uint32 = 1624
	params := pbc.GenerateA(rbits, qbits)
	pairing := params.NewPairing()
	pairing := params.NewPairing()1
	return pairing
}

func SetupOrdinaryPairing() *pbc.Pairing {
	var rbits uint32 = 256
	var qbits uint32 = 3248
	params := pbc.GenerateE(rbits, qbits)
	pairing := params.NewPairing()
	return pairing
}

// int attrNo,pairing_t *pairing, MSP *msp
func Setup(attrNo int, pairing pbc.Pairing, msp *MSP, fileName string) error {
	fSetup, err := os.OpenFile(fileName, os.O_WRONLY, 0)
	if err != nil {
		fmt.Printf("Open file failure!")
		return err
	}
	defer fSetup.Close()
	//golang的时间计算

	g := pairing.NewG1()
	g.Rand()
	h := pairing.NewG1()

	//initial the alpha and a in Z_p
	alpha := pairing.NewZr()
	a := pairing.NewZr()
	alpha.Rand()
	a.Rand()
	//public key e(g,g)^alpha
	pubKey := pairing.NewGT() //initial the publicKey
	gAlpha := pairing.NewG2() //initial the gAlpha
	gA := pairing.NewG2()     //initial the gA
	gAlpha.PowZn(g, alpha)    //gAlpha=g^alpha
	gA.PowZn(g, a)            //gA=g^a
	// start := time.Now().UnixNano() // 获取当前时间
	// pubKey.Pair(g, gAlpha)         //publicKey = e(g,g^alpha) = e(g,g)^alpha
	// setupEnd := time.Now().UnixNano()
	// setupTime := float64((start - setupEnd) / 1e6)

	start := time.Now()    // 获取当前时间
	pubKey.Pair(g, gAlpha) //publicKey = e(g,g^alpha) = e(g,g)^alpha
	setupTime := time.Since(start)

	//Master secret key
	msk := pairing.NewG2()
	msk.Set(gAlpha) //msk = g^alpha
	//write the master key and public key to file
	fG, err1 := os.OpenFile("publicKey/g.key", os.O_RDWR, 0)
	fGA, err2 := os.OpenFile("publicKey/gA.key", os.O_RDWR, 0)
	fPub, err3 := os.OpenFile("publicKey/eGG.key", os.O_RDWR, 0)
	fMsk, err := os.OpenFile("MSK/msk.key", os.O_RDWR, 0)
	if err != nil || err1 != nil || err2 != nil || err3 != nil {
		fmt.Printf("Open file failure!")
		return err
	}
	defer fG.Close()
	defer fGA.Close()
	defer fPub.Close()
	defer fMsk.Close()
	defer fSetup.Close()

	var fH *os.File
	fmt.Fprintf(fG, "%s\n", g)
	fmt.Fprintf(fPub, "%s\n", pubKey)
	fmt.Fprintf(fGA, "%s\n", gA)
	fmt.Fprintf(fSetup, "%s\r\n", setupTime)
	var count int = 0
	hCmd := make([]byte, 0, 100) //the command line for the pointer of FILE* fH
	var attrName string          //the name of attribute
	hCmd = append(hCmd, "publicKey/h"...)
	for {
		if count != attrNo {
			break //如果count!=attrNo则退出
		}
		attrName = fmt.Sprintf("%c", msp.Label[count])
		hCmd = append(hCmd, attrName...)
		hCmd = append(hCmd, ".key"...)
		fH, err = os.OpenFile(string(hCmd), os.O_WRONLY, 0)
		if err != nil {
			fH.Close()
			fmt.Printf("Open file failure!")
			return err
		}
		h.Rand()
		fmt.Fprintf(fH, "%s", h)
		hCmd = append(hCmd[:0], hCmd[(len(hCmd)):]...) //清空buffer
		hCmd = append(hCmd, "publicKey/h"...)
		fH.Close()
		count++
	}

	return nil
}
