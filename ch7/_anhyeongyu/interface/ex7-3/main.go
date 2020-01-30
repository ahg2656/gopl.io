package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

// String 메소드는 전달된 값을 출력하는데 사용되며
// 그 값은 문자열을 받는 임의의 포맷의 피연산자
// 또는 포메팅하지 않고 출력하는 Print 등이 있다.
/*
	type Stringer interface {
		String() string
	}
*/

func (t *tree) String() string {
	var sorted = make([]int, 0)
	sorted = appendValues(sorted, t)
	var buffer bytes.Buffer
	for i, v := range sorted {
		if i > 0 {
			buffer.WriteString(" ")
		}
		buffer.WriteString(strconv.Itoa(v))
	}
	return buffer.String()
}

// Sort는 값을 직접 정렬한다.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues는 t의 원소를 values에 순서대로 추가하고
// 결과 슬라이스를 반환한다.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// &tree{value: value} 반환과 같다.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}

	var root *tree
	for _, v := range data {
		root = add(root, v)
	}
	fmt.Println(root)
	// String 메소드가 없는 경우
	// 출력: &{10 0xc00007e060 0xc00007e080}
}
