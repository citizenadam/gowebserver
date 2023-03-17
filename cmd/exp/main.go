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

	m := make(map[string]User)

	acct1 := User{
		Name:       "John Smith",
		Age:        22,
		Bio:        `Realtime Dog Dealer`,
		Authstatus: Authstatus{},
	}

	acct2 := User{
		Name:       "Jenny Smith",
		Age:        21,
		Bio:        `Doge Entertainment`,
		Authstatus: Authstatus{},
		Tags: struct {
			owner int
			cats  int
			dogs  int
		}{},
	}

	m["acct1"] = acct1
	m["acct2"] = acct2

	err = t.Execute(os.Stdout, m)
	if err != nil {
		panic(err)
	}
}
