package main

import "fmt"

// pc[i]은 i의 인구수다.
var pc [256]byte

// 패키지 수준의 변수 선언 후 init() 함수
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func popCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main() {
	fmt.Println(popCount(253))
	fmt.Println(byte(255 >> (5 * 8)))
}
