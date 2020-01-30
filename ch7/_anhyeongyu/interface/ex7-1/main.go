// ByteCounter의 아이디어를 사용해 단어 단위 카운터와 라인 단위 카운터를 구현하라.
// bufio.ScanWords가 유용할 것이다.

package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int
type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	buf := bytes.NewBuffer(p)
	scanner := bufio.NewScanner(buf)

	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)

	// Count the words.
	for scanner.Scan() {
		*c++
	}
	return len(p), nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	buf := bytes.NewBuffer(p)
	scanner := bufio.NewScanner(buf)

	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanLines)

	// Count the lines.
	for scanner.Scan() {
		*c++
	}
	return len(p), nil
}

func main() {
	var w WordCounter
	var l LineCounter

	w.Write([]byte("Thie	is a sample."))
	fmt.Println(w)

	w = 0

	// WordCounter
	fmt.Fprint(&w, "This is a sentence containing lots of words.")
	fmt.Println(w)

	// LineCounter
	fmt.Fprint(&l, "asdf asdf\n asdf asdf")
	fmt.Println(l)
}
