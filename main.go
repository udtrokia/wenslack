package main;

import (
	"github.com/kataras/iris"
);

type Advices struct {
	Uint string
	Position string
	Advice string
}

func post(ctx iris.Context) {
	advice :=  Advices{};
	err := ctx.ReadJSON(&advice);
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	};
	insert(&advice);
	ctx.Writef("Post: %#v", &advice);
}


func main() {
	app := iris.New();
	app.Post("/", post);
	app.Run(iris.Addr(":8080"));
}
