package hsd

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if err := setup(); err != nil {
		log.Fatal("failed to setup:", err)
	}

	log.Println("Before")
	ret := m.Run()

	if err := teardown(); err != nil {
		log.Fatal("failed to teardown:", err)
	}

	log.Panicln("After")
	os.Exit(ret)
}

func setup() {
	panic("unimplemented")
}

func teardown() {
	panic("unimplemented")
}

func TestA(t *testing.T) {
	log.Println("TestA running")
}

func TestB(t *testing.T) {
	log.Println("TestB running")
}
