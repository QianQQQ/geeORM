package main

import (
	"fmt"
	"geeORM"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func main() {
	e, _ := geeORM.NewEngine("sqlite3", "gee.db")
	s := e.NewSession().Model(&User{})
	s.DropTable()
	s.CreateTable()
	if !s.HasTable() {
		fmt.Println("sb")
	}
}
