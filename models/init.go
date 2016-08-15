package company

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"time"
)

func init(){
	orm.RegisterModel(new(News), new(Message), new(MessageFlag))
}

func InitDB(flag bool){
	dbhost := beego.AppConfig.DefaultString("dbhost", "localhost")
	dbport := beego.AppConfig.DefaultString("dbport", "3306")
	dbusername := beego.AppConfig.DefaultString("dbusername", "root")
	dbpassword := beego.AppConfig.DefaultString("dbpassword", "root")
	dbname := beego.AppConfig.DefaultString("dbname", "frontsurf")

	dsn := dbusername+":"+dbpassword+"@tcp("+dbhost+":"+dbport+")/"+dbname+"?charset=utf8"
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.DefaultTimeLoc, _ = time.LoadLocation("Asia/Shanghai")    // 默认 time.Local
	orm.RunSyncdb("default", false, true)   //自动建表
	orm.Debug = flag
}