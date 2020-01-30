package main

import "fmt"

type ByteCounter int

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println((c))

	c = 0 // 카운터 초기화
	var name = "Dolly"

	/*
		func fmt.Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
	*/
	/*
		type Writer interface {
			Write(p []byte) (n int, err error)
		}
	*/
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	// ByteCounter가 int 타입임에도
	// int와 ByteCounter의 타입은 다르다.
	*c += ByteCounter(len(p))
	return len(p), nil
}
