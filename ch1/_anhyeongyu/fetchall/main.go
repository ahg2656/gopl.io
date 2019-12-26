// Fetchall은 URL을 병렬로 반입하고 시간과 크기를 보고한다.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// _ "net/http/pprof"
// 사용되지 않으면 import에서 지워지는 것을 방지하며 init 함수를 실행함

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // 고루틴 시작
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // ch 채널로 송신
		return
	}

	// iotuil.ReadAll() - 한번에 모두 읽고 가지고 있기 때문에 메모리를 모두 차지한다.
	// io.Copy() - 일정 크기의 버퍼를 두고 읽고 쓰기를 반복하기 때문에 메모리를 적게 차지

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // 리소스 누출 방지
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	sec := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", sec, nbytes, url)
}
