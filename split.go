package main

import (
	"fmt"
)

func main() {
	s := "Hello,World,This,Is,Go"
	sep := ","
	result := splitString(s, sep)
	for _, str := range result {
		fmt.Println(str)
	}
}

func splitString(s string, sep string) []string {
	var result []string
	start := 0
	sepLen := len(sep)
	sLen := len(s)

	for i := 0; i <= sLen-sepLen; {
		if s[i:i+sepLen] == sep {
			result = append(result, s[start:i])
			start = i + sepLen
			i += sepLen
		} else {
			i++
		}
	}
	result = append(result, s[start:])
	return result
}
