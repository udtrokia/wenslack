// utils.go

package main;

import (
	"github.com/go-pg/pg/orm"
)

func create () bool {
	
	for _, model := range []interface{}{&Advices{}} {
		err := db().CreateTable(model, &orm.CreateTableOptions{
			Temp:          false, 
			FKConstraints: true,
		})
		if err != nil {
			panic(err)
		}
	}
	println("create table;");
	return true;
}
