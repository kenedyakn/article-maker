package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	a := App{}
	// You need to set your Username and Password here
	a.Initialize(DB_USER, DB_PASSWORD, DB_NAME)
	a.Run(":9000")
}
