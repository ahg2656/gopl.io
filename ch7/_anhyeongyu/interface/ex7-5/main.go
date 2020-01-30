// io 패키지의 LimitReader 함수는 io.Reader인 r과 바이트 수 n을 받아 r에서 읽고
// n바이트 이후에 파일 끝 상태를 보고하는 새 Reader를 반환한다. 이를 구현하라.

// func LimitReader(r io.Reader, limit int) io.Reader

package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type limitReader struct {
	r        io.Reader
	n, limit int
}

func (lr *limitReader) Read(p []byte) (n int, err error) {
	n, err = lr.r.Read(p[:lr.limit])
	lr.n += n
	if lr.n >= lr.limit {
		err = io.EOF
	}
	return
}

func LimitReader(r io.Reader, limit int) io.Reader {
	return &limitReader{r: r, limit: limit}
}

func main() {
	// func strings.NewReader(s string) *strings.Reader
	/*
		package strings
		func (r *Reader) Read(b []byte) (n int, err error) {
			...
		}
	*/
	/*
		package io
		type Reader interface {
			Read(p []byte) (n int, err error)
		}
	*/
	// strings.Reader는 io.Reader 인터페이스에 충족한다.

	reader := LimitReader(strings.NewReader("안녕, 세계!"), 5)
	buffer := &bytes.Buffer{}

	/*
		func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error) {
			...
		}
	*/
	bytes, _ := buffer.ReadFrom(reader)
	fmt.Println(bytes)  // 5
	fmt.Println(buffer) // 안�
}
