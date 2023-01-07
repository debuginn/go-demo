package main

import (
	"fmt"
	"github.com/kataras/iris"
)

type DataDemo1 struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	app := iris.New()

	app.Post("/data1", Handler1)

	err := app.Run(
		iris.Addr(":8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
	fmt.Printf("iris start app is failed, err:%+v\n", err)
}

func Handler1(ctx iris.Context) {
	user := DataDemo1{}

	err := ctx.ReadForm(&user)
	if err != nil {
		fmt.Printf("read form data is failed, err:%+v\n", err)
		return
	}

	fmt.Printf("user data :%+v", user)
}
