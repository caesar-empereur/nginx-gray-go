package routers

import (
	"github.com/astaxie/beego"
	"nginx-gray-go/controllers"
)

func init() {

	beego.Router("/node/page", &controllers.GrayController{}, "get:List")
	beego.Router("/node/add", &controllers.GrayController{}, "post:AddOne")
	beego.Router("/node/update", &controllers.GrayController{}, "post:UpdateOne")
	beego.Router("/node/delete", &controllers.GrayController{}, "post:DeleteOne")

	beego.Router("/nginx/list", &controllers.NginxGrayController{}, "get:GetListFromNginx")
	beego.Router("/nginx/add", &controllers.NginxGrayController{}, "get:AddOneToNginx")
	beego.Router("/nginx/delete", &controllers.NginxGrayController{}, "get:DeleteOneFromNginx")
	beego.Router("/nginx/update", &controllers.NginxGrayController{}, "get:UpdateOneToNginx")

}
