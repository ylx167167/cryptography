package cpabe

import (
	"fmt"
	"os"
)

type MSP struct {
	Matrix [][]int
	Label  []byte
	Rows   int
	Cols   int
}

func mspInit(msp *MSP, rows int, cols int) {
	msp.Matrix = make([][]int, rows)
	for i := range msp.Matrix {
		msp.Matrix[i] = make([]int, cols)
	}
	msp.Label = make([]byte, rows)
	msp.Rows = rows
	msp.Cols = cols
}

func MspClear(msp MSP) {
	//不用实现
}

func MspSetup(msp *MSP, fileName string) error {
	// fAttr, err := os.Open(fileName)
	fAttr, err := os.OpenFile(fileName, os.O_RDONLY, 0)
	if err != nil {
		fmt.Printf("Open file failure!")
		return err
	}
	defer fAttr.Close()
	var rows int
	var cols int
	fmt.Fscanf(fAttr, "%d %d\n", &rows, &cols)
	mspInit(msp, rows, cols)
	for i := 0; i < rows; i++ {
		fmt.Fscanf(fAttr, "%c\n", &msp.Label[i])
		fmt.Print(string(msp.Label[i]))
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Fscanf(fAttr, "%d ", &msp.Matrix[i][j])
			fmt.Print(msp.Matrix[i][j])
		}
		fmt.Printf("\n")
	}
	return nil
}
