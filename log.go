package main

import (
	"log"
	"os"
)

func init() {
	f, err := os.Open("text.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)
	log.Println("app started")
}
