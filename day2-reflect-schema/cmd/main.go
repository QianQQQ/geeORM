package main

import (
	"fmt"
	"geeORM"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e, _ := geeORM.NewEngine("sqlite3", "gee.db")
	defer e.Close()
	s := e.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}
