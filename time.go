package main

import (
	"fmt"
	"log"
	"time"
)

func init() {
	now := time.Now()
	fmt.Println(now.Format(time.RFC3339))

	d, err := time.ParseDuration("3s")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(d)

	e, err := time.ParseDuration("4m")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e)

	w, err := time.ParseDuration("5h")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(w)

}
