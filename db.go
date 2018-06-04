package main;

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"	
	"github.com/go-pg/pg"
//	"github.com/go-pg/pg/orm"
);

type Config struct {
	User string
	Password string
	Database string
}


// dbconnect
func db() *pg.DB {
	c := Config{};
	y, _ := ioutil.ReadFile("./config.yaml");
	yaml.Unmarshal(y, &c);
	db := pg.Connect(&pg.Options{
		User: c.User,
		Password: c.Password,
		Database: c.Database,
	});
	return db;	
}



//func create () bool {	
//	for _, model := range []interface{}{&Advices{}} {
//		err := db.CreateTable(model, &orm.CreateTableOptions{
//			Temp:          false, 
//			FKConstraints: true,
//		})
//		if err != nil {
//			panic(err)
//		}
//	}
//	println("create table");
//	return true;
//}

func insert (advice *Advices) {
	fmt.Println(advice);
	db := db();
	err := db.Insert(advice)
	if err != nil { panic(err) };
}
