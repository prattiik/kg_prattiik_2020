package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	result := isMappingExists(args[0], args[1])
	fmt.Println(result)
}

func isMappingExists(s1 string, s2 string) bool {
	buffer1 := make([]int, 128)
	buffer2 := make([]int, 128)
	for i := 0; i < len(s1); i++ {
		if buffer1[s1[i]] != buffer2[s2[i]] {
			if buffer2[s2[i]] == 0 {
				return false
			}
		}
		buffer1[s1[i]] = i + 1
		buffer2[s2[i]] = i + 1
	}
	return true
}
