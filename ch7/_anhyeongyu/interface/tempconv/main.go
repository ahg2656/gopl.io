package tempconv

import (
	"flag"
	"fmt"
)

// 2.5절
type Celsius float64
type Fahrenheit float64

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32.0) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

// *celsiusFlag는 flag.Value 인터페이스를 충족한다.
type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // 오류 확인이 필요 없다.
	switch unit {
	case "C", "ºC":
		f.Celsius = Celsius(value)
		return nil
	case "F", "ºF":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag는 Celsius플래그를 지정된 시간,
// 기본 값, 사용법으로 정의하고 플래그 변수의 주소를 반환한다.
// 인자 flag에는 양과 단위가 있어야 한다. 예) "100C"
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
