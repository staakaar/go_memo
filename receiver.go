//go:build ignore

package main

import "fmt"

type I int

func (i I) Add(n int) I {
	return i + I(n)
}

func main() {
	var n I = 0

	n = n.Add(1).Add(2)
	fmt.Println(n)

	// レシーバなしでの呼び出し可能
	add := n.Add
	fmt.Println(add(3))

	// n.Add I型のnに対して定義されているメソッド
	fmt.Printf("%T\n", n.Add) // func(int) main.I
	fmt.Printf("%T\n", I.Add) // func(main.I, int) main.I レシーバが第一引数い定義されている
}
