package main;

import (
	"time"
	"fmt"
	"encoding/json"
	"github.com/udtrokia/bowie"
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
		ctx.StatusCode(iris.StatusInternalServerError);
		ctx.WriteString(err.Error());
	};

	out, err := json.Marshal(advice);
	
	ziggy := Bowie.Ziggy("wenslack.db", 0666);
	ziggy.Star(
		[]byte(""),
		[]byte(string(out)),
		true);
	
	ctx.Writef("Post: %#v", &advice);
}

func get_list(ctx iris.Context) {
	ziggy := Bowie.Ziggy("wenslack.db", 0666);
	ziggy.Oddity(func (pairs []Bowie.Asher){
		//out, err := json.Marshal(pairs);		
		fmt.Printf("\n%#v\n", pairs);
		fmt.Printf("%+v\n", pairs);
		fmt.Printf("%v\n", pairs);

//		_pairs, err := find();
//		if err != nil {panic(err)};
//		fmt.Printf("\n%#v\n", _pairs);
//		fmt.Printf("%+v\n", _pairs);
//		fmt.Printf("%v\n", _pairs);
//		
		ctx.JSON(pairs);
	})
}

func main() {
	app := iris.New();
	app.Logger().SetLevel("debug");
	app.Use(recover.New());
	app.Use(logger.New());
	
	authConfig := basicauth.Config{
		Users:   map[string]string{"username": "password" },
		Realm:   "Authorization Required", 
		Expires: time.Duration(30) * time.Minute,
	}

	authentication := basicauth.New(authConfig);

	crs := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowCredentials: true,
	});

	app.StaticWeb("/static", "./assets");
	
	page := app.Party("/", crs).AllowMethods(iris.MethodOptions);
	page.Use(authentication);
	app.RegisterView(iris.HTML("./assets/html", ".html"));
	{
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
