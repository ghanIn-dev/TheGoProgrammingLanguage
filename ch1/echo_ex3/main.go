//test3 비효율적인 버전과(echo1) echo3의 성능비교
package main

import (
	"fmt"
	"strings"
	"time"
)

var args = []string{"hi", "there", "buddy", "boy", "5", "6", "7", "8", "9"}

func main() {

	start := time.Now()
	echo1()
	end := time.Now()
	fmt.Printf("echo1 time : %v \n", end.Sub(start))

	start1 := time.Now()
	echo2()
	end1 := time.Since(start1)
	fmt.Printf("echo2 time : %s \n", end1)

	start3 := time.Now()
	echo3()
	end3 := time.Since(start3)
	fmt.Printf("echo3 time : %s \n", end3)
}

func echo1() {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(args, " "))
}
