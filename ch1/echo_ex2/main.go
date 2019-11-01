//echo test2
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", ""
	for index, arg := range os.Args[1:] {
		s += sep + strconv.Itoa(index+1) + " " + arg
		sep = "\n"
	}
	fmt.Println(s)
}
