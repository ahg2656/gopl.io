// Echo4는 커맨드라인 인수를 출력한다.
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep)) // 포인트 문법을 써야 값에 접근이 가능하다.
	if !*n {
		fmt.Println()
	}
}
