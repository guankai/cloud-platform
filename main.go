package main

import (
	_ "cloud-platform/routers"
	"github.com/astaxie/beego"
	"cloud-platform/utils/db"
	_ "github.com/go-sql-driver/mysql"
	_ "cloud-platform/models"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	db.InitDatabase()
	// set default database
	//orm.RegisterDataBase("default", "mysql", "root:cloudos@54.223.92.8/service_find?charset=utf8", 30)
	// create table
	//orm.RunSyncdb("default", false, true)
	beego.Run()
}
