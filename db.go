package main;

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"	
	"github.com/go-pg/pg"
);

type Config struct {
	User string
	Password string
	Database string
}

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

func insert(advice *Advices) {
	fmt.Println(advice);
	db := db();
	err := db.Insert(advice)
	if err != nil { panic(err) };
}

func find() ([]Advices, error){
	db := db();
	var advices []Advices
	_, err := db.Query(&advices, `SELECT * FROM advices`);
	return advices, err
}
