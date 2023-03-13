package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
	Bio  string
	Authstatus
	Tags struct {
		owner int
		cats  int
		dogs  int
	}
}

type Authstatus struct {
	Login int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name:       "John Smith",
		Age:        22,
		Bio:        `<script>alert("Haha, you have been hac0rd");</script>`,
		Authstatus: Authstatus{},
	}

	// user2 := User{
	// 	Name: "Jenny Smith",
	// 	Age:  21,
	// }

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
