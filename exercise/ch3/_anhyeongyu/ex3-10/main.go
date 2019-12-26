package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer

	for len(s) > 0 {
		var p int
		if len(s)%3 != 0 {
			p = len(s) % 3
		} else {
			p = 3
		}

		buf.WriteString(s[:p] + ",")
		s = s[p:]
	}

	sc := buf.String()

	return sc[:len(sc)-1]
}

func main() {
	fmt.Println(comma("1234"))
}
