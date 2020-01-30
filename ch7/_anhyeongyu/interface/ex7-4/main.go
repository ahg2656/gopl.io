// strings.NewReader 함수는 인자인 문자열을 읽어 io.Reader(및 기타) 인터페이스를 충족하는 값을 반환한다.
// NewReader의 간단한 버전을 직접 구현하고 이를 사용해 문자열에서 입력을 받는 HTML 파서(5.2절)를 만들어라.
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
)

type MyReader string

func (r *MyReader) Read(p []byte) (n int, err error) {
	n = len(*r)
	copy(p, []byte(*r))
	err = io.EOF
	return
}

func NewReader(s string) io.Reader {
	var r MyReader
	r = MyReader(s)
	return &r
}

func main() {
	// func html.Parse(r io.Reader) (*html.Node, error)
	/*
		type Reader interface {
			Read(p []byte) (n int, err error)
		}
	*/
	doc, _ := html.Parse(NewReader("<html><body><h1>hello</h1></body></html>"))
	fmt.Println(doc.FirstChild.LastChild.FirstChild.FirstChild.Data)
}
