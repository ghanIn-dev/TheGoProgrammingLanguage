package main

import (
	"os"
	"strings"
)

func main() {
	echo1()
	echo3()
}

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
}

func echo3() {
	strings.Join(os.Args[1:], " ")
}
