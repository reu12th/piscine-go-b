package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		printHelp()
		return
	}

	insert := ""
	order := false
	mainStr := ""

	for _, arg := range args {
		// --insert=xxxx or -i=xxxx
		if len(arg) >= 9 && arg[:9] == "--insert=" {
			insert = arg[9:]
			continue
		}
		if len(arg) >= 3 && arg[:3] == "-i=" {
			insert = arg[3:]
			continue
		}

		// --order / -o
		if arg == "--order" || arg == "-o" {
			order = true
			continue
		}

		// The standalone string argument
		if mainStr == "" {
			mainStr = arg
			continue
		}
	}

	// Apply insertion
	res := mainStr + insert

	// Apply ordering (ASCII sort)
	if order {
		res = sortString(res)
	}

	// Print using PrintRune
	for _, r := range res {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func sortString(s string) string {
	runes := []rune(s)
	// Simple bubble sort (ASCII)
	n := len(runes)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}
	return string(runes)
}
