package main;

import (
	"github.com/kataras/iris"
)

func home(ctx iris.Context) {
	ctx.Gzip(true);
	ctx.View("home.html")
}

func dashboard(ctx iris.Context) {
	ctx.Gzip(true);
	ctx.View("dashboard.html")	
}

func d3(ctx iris.Context) {
	ctx.Gzip(true);
	ctx.View("d3.html")
}
