package main

import (
	_ "MarmotCMS/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// Model Struct
type User struct {
	Id   int
	Name string `orm:"size(100)"`
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@/marmot_cms?charset=utf8")
}

func main() {
	beego.Run()
}
