package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)

	value := []rune{
		'x', ' ', '=', ' ', '0' + 4, '0' + 2, ',', ' ', 'y', ' ', '=', ' ', '0' + 2, '0' + 1,
	}

	for _, ch := range value {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
