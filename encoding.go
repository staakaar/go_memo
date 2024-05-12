package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

var content = `
{
	"species": "taro",
	"description": "説明",
	"dimensions": {
		"height": 24,
		"width": 10
	}
}
`

type Dimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Data struct {
	Species     string     `json:"species"`
	Description string     `json:"description"`
	Dimensions  Dimensions `json:"dimensinos"`
}

func init() {
	var data Data
	err := json.Unmarshal([]byte(content), &data)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open("input.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err := json.NewDecoder(f).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	dec := json.NewDecoder(f)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		err := dec.Decode(&data)
		if err != nil {
			break
		}
		// doSomething(&data)
	}
}
