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
