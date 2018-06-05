package main;

import (
	"time"
	
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/basicauth"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
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

func get_list(ctx iris.Context) {
	list, err := find();
	if err != nil {
		ctx.Writef("Error", err);		
	}
	ctx.JSON(list);
}

func main() {
	app := iris.New();
	app.Logger().SetLevel("debug");
	app.Use(recover.New());
	app.Use(logger.New());
	
	authConfig := basicauth.Config{
		Users:   map[string]string{"username": "password" },
		Realm:   "Authorization Required", // defaults to "Authorization Required"
		Expires: time.Duration(30) * time.Minute,
	}

	authentication := basicauth.New(authConfig)

	crs := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowCredentials: true,
	});

	app.StaticWeb("/static", "./assets");
	
	page := app.Party("/", crs).AllowMethods(iris.MethodOptions);
	page.Use(authentication);
	app.RegisterView(iris.HTML("./assets/html", ".html"));
	{
		//page.Get("/", home);
		page.Get("/", dashboard);
		page.Get("dashboard", dashboard);
	}
	
	api := app.Party("/api/", crs).AllowMethods(iris.MethodOptions);
	{
		api.Post("advices", post);
		api.Post("get_list", get_list);
	}

	app.Run(iris.Addr(":8080"));
}
