package repository

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"nginx-gray-go/models"
)

func RegisterDB() {

	beego.LoadAppConfig("ini", "..\\conf\\app.conf")
	//if errConfig != nil {
	//	panic("获取配置文件出错")
	//}
	sqlConn := beego.AppConfig.String("sqlconn")

	fmt.Print(sqlConn)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", sqlConn, 30)

	orm.RegisterModel(new(models.ServiceNode))
	orm.Debug = true
}
