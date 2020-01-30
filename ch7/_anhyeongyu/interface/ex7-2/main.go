// io.Writer를 받아서 이를 감싸는 새 Writer와 새 Writer로 써진 바이트 수를 저장하는
// int64 변수 포인터를 반환하는 CountingWriter 함수를 다음의 시그니처대로 작성하라.

// func CountingWriter(w io.Writer) (io.Writer, *int64)

package main

import (
	"fmt"
	"io"
	"os"
)

type ByteCounter struct {
	w   io.Writer
	cnt int64
}

func (bc *ByteCounter) Write(p []byte) (n int, err error) {
	n, err = bc.w.Write(p)
	bc.cnt += int64(n)
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	bc := ByteCounter{w, 0}
	return &bc, &bc.cnt
}

func main() {
	writer, count := CountingWriter(os.Stdout)
	fmt.Println(*count) // 0
	fmt.Fprintf(writer, "Hello, world!\n")
	fmt.Println(*count) // 14
	fmt.Fprintf(writer, "Bye, golang.\n")
	fmt.Println(*count) // 27
}
