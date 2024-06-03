//go:build ignore

package main

import (
	"bytes"
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"sync"
)

func downloadCSV(wg *sync.WaitGroup, urls []string, ch chan []byte) {
	defer wg.Done()
	defer close(ch)

	for _, u := range urls {
		resp, err := http.Get(u)
		if err != nil {
			log.Println("cannot download CSV:", err)
			continue
		}

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			resp.Body.Close()
			log.Println("cannot read content:", err)
			continue
		}
		resp.Body.Close()
		ch <- b
	}
}

func insertRecords(records []string) {
	panic("unimplemented")
}

func main() {
	urls := []string{
		"http://my-server.com/data01.csv",
		"http://my-server.com/data02.csv",
		"http://my-server.com/data03.csv",
	}

	// バイト列を転送するためのchannelを作成
	ch := make(chan []byte)

	var wg sync.WaitGroup
	wg.Add(1)
	go downloadCSV(&wg, urls, ch)

	for _, b := range ch {
		r := csv.NewReader(bytes.NewReader(r))
		for {
			records, err := r.Read()
			if err != nil {
				log.Fatal(err)
			}
			insertRecords(records)
		}
	}
	wg.Wait()
}
