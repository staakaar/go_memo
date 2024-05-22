//go:build ignore

package main

import (
	"fmt"
)

type F struct {
	Name string
	Age  int
}

func (f *F) String() string {
	return fmt.Sprintf("Name=%q, AGE=%d", f.Name, f.Age)
}

func init() {
	f := &F{
		Name: "Taro",
		Age:  24,
	}

	fmt.Printf("%v\n", f)

	/** フィールド名と内容 */
	fmt.Printf("%#v+\n", f)

	/** struct名も出力 */
	fmt.Printf("%#v+\n", f)

	/** 型名を出力 */
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", *f)
}
