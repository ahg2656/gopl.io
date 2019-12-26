// Boiling은 물의 끓는점을 출력한다.
package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
	// 출력:
	// boiling point = 212°F or 100°C
	// 패키지 수준의 개체의 이름은 패키지 내의 모든 파일에서 참조 가능
}
