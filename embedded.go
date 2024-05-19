//go:build ignore

package main

import "fmt"

type Attr struct {
	Name string
	Age  int
}

type AttrEx struct {
	Name string
}

type Teacher struct {
	Attr
	AttrEx
	Subject string
}

type Student struct {
	Attr
	Score int
}

func main() {
	teacher := Teacher{
		Attr: Attr{
			Name: "John Schwartz",
			Age:  43,
		},
		AttrEx: AttrEx{
			Name: "JS",
		},
		Subject: "Math",
	}

	student := Student{
		Attr: Attr{
			Name: "Robert Smith",
			Age:  17,
		},
		Score: 87,
	}

	fmt.Println(teacher.Attr.Name, teacher.Subject)
	fmt.Println(student.Name, student.Score)
}

func (a Attr) String() string {
	return fmt.Sprintf("%s(%d)", a.Name, a.Age)
}

func (a AttrEx) String() string {
	return fmt.Sprintf("(a.k.a. %s)", a.Name)
}
