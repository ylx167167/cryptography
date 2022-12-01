package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"math/rand"

	"github.com/ylx167167/cryptography/aes"
	cpabe "github.com/ylx167167/cryptography/cp-abe"
	"github.com/ylx167167/cryptography/ecc"
)

// import cpabe "github.com/ylx167167/cryptography/cp-abe"

// import
func main() {

	// test7()
	// test8()
	// test9("test")
	// test10()
	// Test11()
	// Test12()
	// test13()

	test14()
}

func Test5() {
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
func Test6() {
	pri, pub, err := ecc.ReadKeyPairFile("A_prikey.pem", "A_pubkey.pem")
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Print(pri)
	fmt.Print(pub)
}

func Test7() {
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

func Test8() {
	apri, apub, _ := ecc.ReadKeyPairFile("A_prikey.pem", "A_pubkey.pem")
	// a, b := ecc.GenerateeciesPri(apri)
	var message string = "this is a message, 这是需要加密的数据"
	e := ecc.Encrypt(message, apub)
	d := ecc.Decrypt(e, apri)
	fmt.Print(d)
}

func Test9(s string) {
	// s := "Test"
	// hash := sha256.New()
	// hash.Write([]byte(s))
	// bs := hash.Sum(nil)
	// fmt.Print(string(bs))
	h := sha256.New()
	h.Write([]byte(s))
	fmt.Printf("%x", h.Sum(nil))
}

// 生成既能签名和加密的算法
func Test10() {
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
	message := "Test"
	signature, _ := ecc.Sign(&dsapri, message)
	T, _ := ecc.Verify(message, signature, &dsapub)
	fmt.Print(T)
}

type Marble struct {
	ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	Name       string `json:"name"`    //the fieldtags are needed to keep case from bouncing around
	Color      string `json:"color"`
	Size       int    `json:"size"`
	Owner      string `json:"owner"`
}

func Test11() {
	str := "{\"color\":\"blue\",\"docType\":\"marble\",\"name\":\"marble1\",\"owner\":\"tom\",\"size\":35}"
	var marbleJSON Marble
	err := json.Unmarshal([]byte(str), &marbleJSON)
	if err != nil {
		fmt.Print("解析失败")
		return
	}
	fmt.Print(marbleJSON.Name)

}

type Systemtype struct {
	Org        string `json:"org"`
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"` //放置公钥结构体
}

func Test12() {
	pri, pub, _ := ecc.GenerateKeyECDSA()
	priblock, pubblock := ecc.SavekeyPair(pri, pub)
	typepri := bytes.NewBuffer(pem.EncodeToMemory(priblock))
	typepub := bytes.NewBuffer(pem.EncodeToMemory(pubblock))

	sys := &Systemtype{"org1test", typepub.String(), typepri.String()}
	fmt.Print(typepri.String())
	// sys := &Systemtype{"org1test", string(priblock.Bytes), string(pubblock.Bytes)}
	// 	Org:        "org1test",
	// 	PublicKey:  strpri,
	// 	PrivateKey: strpub,
	// }
	systemBytes, err1 := json.Marshal(sys)
	if err1 != nil {
		fmt.Print(err1)
		return
	}
	// fmt.Print(string(systemBytes))
	var sys2 Systemtype
	err := json.Unmarshal(systemBytes, &sys2)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Print("\n")

	fmt.Print(sys2.PrivateKey)
	var out bytes.Buffer
	b64 := base64.NewEncoder(base64.StdEncoding, &out)
	if _, err := b64.Write(priblock.Bytes); err != nil {
		return
	}
	fmt.Print("\n")
	fmt.Print(out.String())
	var out2 bytes.Buffer
	var testbyte []byte
	b65 := base64.NewDecoder(base64.StdEncoding, &out2)
	if _, err := b65.Read(testbyte); err != nil {
		return
	}
	priblock1 := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: testbyte,
	}
	fmt.Print("\n")
	typepri1 := bytes.NewBuffer(pem.EncodeToMemory(priblock1))
	fmt.Print(string(typepri1.String()))

	// out.Close()
}

// 长度为62
var bt []byte = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")

func RandStr1(n int) string {
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = bt[rand.Int31()%62]
	}
	return string(result)
}

func test13() {
	result := RandStr1(12)
	fmt.Print(result)
}

func test14() {
	text := "123"                                        // 你要加密的数据
	AesKey := []byte("#HvL%$o0oNNoOZnk#o2qbqCeQB1iXeIR") // 对称秘钥长度必须是16的倍数

	fmt.Printf("明文: %s\n秘钥: %s\n", text, string(AesKey))
	encrypted, err := aes.AesEncrypt([]byte(text), AesKey)
	if err != nil {
		panic(err)
	}
	fmt.Printf("加密后: %s\n", base64.StdEncoding.EncodeToString(encrypted))
	//encrypteds, _ := base64.StdEncoding.DecodeString("xvhqp8bT0mkEcAsNK+L4fw==")
	origin, err := aes.AesDecrypt(encrypted, AesKey)
	if err != nil {
		panic(err)
	}
	fmt.Printf("解密后明文: %s\n", string(origin))
}
