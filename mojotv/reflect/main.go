package main

import "reflect"

type User struct {
	Email  string `mcl:"email"`
	Name   string `mcl:"name"`
	Age    int    `mcl:"age"`
	Github string `mcl:"github" default:"a8m"`
}

func main() {
	var user interface{} = User{}

	t := reflect.TypeOf(user)

	if t.Kind() != reflect.Struct {
		return
	}
}
