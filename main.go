package main

import (
	_ "news-module/routers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	m "news-module/models"
)

func main() {
	beego.SetLogger("file", `{"filename":"logs/frontsurf.log","maxlines":1000000}`)
	m.InitDB(true)
	beego.Run()
}

