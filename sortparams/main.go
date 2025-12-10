package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	paramsList := os.Args[1:]
	paramsLen := len(paramsList)

	for i := 0; i < paramsLen; i++ {
		for j := 0; j < paramsLen-i-1; j++ {
			if paramsList[j] > paramsList[j+1] {
				tmp := paramsList[j]
				paramsList[j] = paramsList[j+1]
				paramsList[j+1] = tmp
			}
		}
	}

	for i := range paramsList {
		for j := range paramsList[i] {
			z01.PrintRune(rune(paramsList[i][j]))
		}
		z01.PrintRune('\n')
	}
}
