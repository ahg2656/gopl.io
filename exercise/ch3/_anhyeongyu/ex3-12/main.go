package main

import "fmt"

func reverseAnagrams(s1, s2 string) bool {
	var result bool
	if s1 == reverse(s2) {
		result = true
	} else {
		result = false
	}
	return result
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	fmt.Println(reverseAnagrams("abc", "cbc"))
}
