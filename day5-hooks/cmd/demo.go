package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func p(i ...interface{}) {
	fmt.Println(i)
}

func main() {
	p(1)
}
