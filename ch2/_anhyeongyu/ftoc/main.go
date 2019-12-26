// Ftoc는 화씨-섭씨 변환을 두 번 출력한다.
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))   // "212°F = 100°C"
}

// 캡슐화
// 첫글자가 대문자가 아니기 때문에 export 할 수 없다.
func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
