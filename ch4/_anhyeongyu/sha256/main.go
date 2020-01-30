package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
)

func main() {

	flag.Parse()

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// %x : 배열이나 바이트 슬라이스의 모든 원소를 16진수 출력
	// %t : 불리언 값 출력
	// %T : 값의 타입 출력

	fmt.Println(countDiff(&c1, &c2))
}

func countDiff(c1pt, c2pt *[32]byte) int {
	cnt := 0
	for i := 0; i < 32; i++ {
		if c1pt[i] == c2pt[i] {
			cnt++
			fmt.Println(i, string(c1pt[i]), c2pt[i])
		}
	}

	return cnt
}
