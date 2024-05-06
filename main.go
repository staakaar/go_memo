package main

import "fmt"

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
			fmt.Printf("%[1]T: %[1]s\n", e) // panicの場合 string my error
			fmt.Printf("%[1]T: %[1]s\n", e) // error型 runtime.boundsError: runtime error: index out of range [2] with length 2
		}
	}()

	var a [2]int
	n := 2
	println(a[n])
}
