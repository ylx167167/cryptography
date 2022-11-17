package cpabe

import (
	"fmt"
	"os"

	"github.com/Nik-U/pbc"
)

func Element_fread(fp *os.File, format string, e *pbc.Element, base int) {
	temp1 := make([]byte, 0, 1024)
	temp2 := make([]byte, 0, 1024)
	tempAll := make([]byte, 0, 2048)
	fmt.Fscanf(fp, "%s %s", &temp1, &temp2)
	tempAll = append(tempAll, temp1...)
	tempAll = append(tempAll, temp2...)
	// fmt.Printf("string(tempAll) = %s\n", string(tempAll))
	_, err := e.SetString(string(tempAll), base)
	if err != true {
		fmt.Print("e.SetString failure")
	}
}

func Element_fread_line(fp *os.File, e *pbc.Element, base int) {

	tempAll := make([]byte, 0, 2048)
	fp.ReadAt(tempAll, 2048) //fgets
	e.SetString(string(tempAll), base)
}

func User_fread(filename string) (int, []int, [][]byte, [][]byte) {
	fUser, err := os.OpenFile(filename, os.O_RDONLY, 0777)
	if err != nil {
		fmt.Printf("open file failure\n")
	}
	var userNo int
	fmt.Fscanf(fUser, "%d\n", &userNo)
	attrNo := make([]int, userNo)
	userName := make([][]byte, userNo)
	attribute := make([][]byte, userNo)
	for i := range userName {
		userName[i] = make([]byte, 100) //姓名不能超过100个字符
	}
	var j int = 0
	var k int = 0
	for i := 0; i < userNo; i++ {
		fmt.Fscanf(fUser, "%s", &userName[i])
		fmt.Print(userNo, " :", string(userName[i]))
		fmt.Fscanf(fUser, "%d\n", &attrNo[i])
		fmt.Print("", attrNo[i])
		attribute[i] = make([]byte, attrNo[i])
		j = 0
		k = attrNo[i]
		for {
			if k == 0 {
				break
			}
			fmt.Fscanf(fUser, "%c\n", &attribute[i][j])
			j++
			k--
		}
		fmt.Printf("\n")
	}
	return userNo, attrNo, userName, attribute
}
