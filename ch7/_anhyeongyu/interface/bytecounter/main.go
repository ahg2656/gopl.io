package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

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


