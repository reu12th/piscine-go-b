package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 4 || args[1] != "-c" {
		os.Exit(1)
	}

	buf := 0
	count := args[2]

	for i := 0; i < len(count); i++ {
		if count[i] < '0' || count[i] > '9' {
			os.Exit(1)
		}
		digit := int(count[i] - '0')
		buf = buf*10 + digit
	}

	// files
	multiplefiles := len(args[3:]) > 1
	errors := false
	for idx, filename := range args[3:] {
		file, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			errors = true
			continue
		}
		if multiplefiles {
			if idx > 0 {
				fmt.Print("\n")
			}
			fmt.Printf("==> %v <==\n", filename)
		}
		startInd := len(file) - buf

		if startInd < 0 {
			startInd = 0
		}
		fmt.Print(string(file[startInd:]))
	}
	if errors {
		os.Exit(1)
	}
}
