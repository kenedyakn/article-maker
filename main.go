package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	a := App{}
	// You need to set your Username and Password here
	a.Initialize("root", "123!@#QWEasd", "article_maker_db")

	a.Run(":9000")
}
