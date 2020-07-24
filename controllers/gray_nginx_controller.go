package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	orm2 "github.com/astaxie/beego/orm"
	"nginx-gray-go/models"
	"nginx-gray-go/rediss"
	"nginx-gray-go/utils"
)

type NginxGrayController struct {
	beego.Controller
}

func (this *NginxGrayController) AddOneToNginx() {
	var id, _ = this.GetInt("id")
	writeToRedis(id, models.ADD)
}

func (this *NginxGrayController) DeleteOneFromNginx() {
	var id, _ = this.GetInt("id")
	writeToRedis(id, models.REMOVE)
}

func (this *NginxGrayController) UpdateOneToNginx() {
	var id, _ = this.GetInt("id")
	writeToRedis(id, models.UPDATE)
}

func writeToRedis(id int, operateType string) {

	if id == 0 {
		panic("id 不能为空")
	}
	serviceNode := models.ServiceNode{Id: id}

	orm := orm2.NewOrm()
	err := orm.Read(&serviceNode)
	if err != nil {
		panic("数据不存在")
	}

	serviceNodeRedis := models.ServiceNodeRedis{}
	serviceNodeRedis.Operate = operateType
	utils.CopyFields(&serviceNode, &serviceNodeRedis)

	jsonString, _ := json.Marshal(serviceNodeRedis)

	redisClient := rediss.GetRedisInstance()
	setErr := redisClient.Set(rediss.LocalCtx, serviceNodeRedis.Domain, string(jsonString), 0).Err()
	if setErr != nil {
		panic("redis set 出错")
	}

	//url := getNginxUrl()
	//
	//url = url + "/update-ups"
	//resp, err := http.Get(url)
	//
	//if resp.StatusCode != 200 {
	//	panic("内部转发错误" + err.Error())
	//}

}

var nginxRunUrl string

func getNginxUrl() string {
	if nginxRunUrl != "" {
		return nginxRunUrl
	}
	beego.LoadAppConfig("ini", "..\\conf\\app.conf")
	nginxRunUrl := beego.AppConfig.String("nginx_run_url")
	return nginxRunUrl
}
