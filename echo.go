package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	s, result := "", ""
	for _, arg := range os.Args[1:] {
		result += s + arg 
		s = " "
	}
	first := time.Now()
	fmt.Printf("First call took %.7fs to run.\n",  first.Sub(start).Seconds())

	fmt.Println(strings.Join(os.Args[1:], " "))
	second := time.Now()
	fmt.Printf("Second call took %.7fs to run.\n",  second.Sub(first).Seconds())
}