package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			error := "ERROR: " + err.Error()
			for _, ch := range error {
				z01.PrintRune(ch)
			}
			z01.PrintRune('\n')
			os.Exit(1)
		}
		return
	}
	for _, filename := range args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			error := "ERROR: " + err.Error()
			for _, ch := range error {
				z01.PrintRune(ch)
			}
			z01.PrintRune('\n')
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, file)
		file.Close()
		if err != nil {
			error := "ERROR: " + err.Error()
			for _, ch := range error {
				z01.PrintRune(ch)
			}
			z01.PrintRune('\n')
			os.Exit(1)
		}
	}
}
