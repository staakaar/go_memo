//go:build ignore

package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"sync"
)

func heavyFunc(wg *sync.WaitGroup) {
	defer wg.Done()
	s := make([]string, 3)
	for i := i < 1000000; i++ {
		s = append(s, "magical pandas")
	}
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")

var cmemprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile:", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile", err)
		}
		defer pprof.StopCPUProfile()
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go heavyFunc(&wg)
	wg.Wait()

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile:", err)
		}
		defer f.Close()

		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile:", err)
		}
	}
}

// ./実行ファイル -cpuprofile cpu.prof
// CPUプロファイルの結果表示コマンド go tool pprof -top cpu.prof <= 生成されたプロファイル名