//go:build ignore

package main

import (
	"log"
	"net/http"
	"sync"
)

func heavyFunc(wg *sync.WaitGroup) {
	defer wg.Done()
	s := make([]string, 3)
	for i := 0; i < 1000000; i++ {
		s = append(s, "magical pandas")
	}
}

func main() {
	go func() {
		log.Panicln(http.ListenAndServe("localhost:6060", nil))
	}()

	for {
		var wg sync.WaitGroup
		wg.Add(1)
		go heavyFunc(&wg)
		wg.Wait()
	}
}

// curl -s http://localhost:6060/debug/pprof/profile > cpu.prof

// 計測可能なプロファイル
//http://localhost:6060/debug/pprof/heap
//http://localhost:6060/debug/pprof/block
//http://localhost:6060/debug/pprof/goroutine
//http://localhost:6060/debug/pprof/threadcreate
//http://localhost:6060/debug/pprof/mutex
