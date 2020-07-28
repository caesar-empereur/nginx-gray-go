package controllers

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	orm2 "github.com/astaxie/beego/orm"
	"io/ioutil"
	"net/http"
	"nginx-gray-go/models"
	"nginx-gray-go/rediss"
	"nginx-gray-go/utils"
	"nginx-gray-go/vo"
	"strconv"
	"strings"
)

var log = utils.GetLogger()

type NginxGrayController struct {
	beego.Controller
}

func (this * NginxGrayController) GetListFromNginx() {
	var backends= this.GetString("backends")
	apiResponse := vo.ApiResponse{}
	if backends == ""{
		apiResponse.Success=false
		apiResponse.Message="参数为空"
	} else {
		apiResponse.Success=true
		nginxRunUrl := beego.AppConfig.String("nginx_run_url")
		nginxRunUrl = nginxRunUrl + "?upstream=" + backends
		resp, err := http.Get(nginxRunUrl)

		if err != nil {
			apiResponse.Success=false
			apiResponse.Message="内部转发错误" + err.Error()
		} else {
			stringResp, err :=ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err.Error())
			}
			apiResponse.Message=string(stringResp)
			log.Info("nginx 6000 返回的结果 " + string(stringResp))
		}
	}
	this.Data["json"] = apiResponse
	this.ServeJSON()
}

func (this *NginxGrayController) AddOneToNginx() {
	var id, _ = this.GetInt("id")
	doOperate(id, models.ADD, this)
}

func (this *NginxGrayController) DeleteOneFromNginx() {
	var id, _ = this.GetInt("id")
	doOperate(id, models.REMOVE, this)
}

func (this *NginxGrayController) UpdateOneToNginx() {
	var id, _ = this.GetInt("id")
	doOperate(id, models.UPDATE, this)
}

func doOperate(id int, operateType string, this *NginxGrayController) {
	apiResponse := vo.ApiResponse{}
	if id == 0 {
		apiResponse.Success=false
		apiResponse.Message="id 不能为空"
	} else {
		msg, err := writeToNginx(id, operateType)
		if err != nil || msg == ""{
			apiResponse.Success=false
			apiResponse.Message=err.Error()
		} else {
			apiResponse.Success=true
			apiResponse.Message = msg
		}
	}

	this.Data["json"] = apiResponse
	this.ServeJSON()
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

func writeToNginx(id int, operateType string) (string, error) {
	if id == 0 {
		panic("id 不能为空")
		return "", errors.New("id 不能为空")
	}
	serviceNode := models.ServiceNode{Id: id}

	orm := orm2.NewOrm()
	err := orm.Read(&serviceNode)
	if err != nil {
		return "", errors.New("数据不存在")
	}

	nginxRunUrl := beego.AppConfig.String("nginx_run_url")
	nginxRunUrl = nginxRunUrl + "?upstream=" + serviceNode.Zone + "&server=" + serviceNode.Ip + ":"+serviceNode.Port
	if operateType == models.UPDATE{
		nginxRunUrl = nginxRunUrl + "&weight="+strconv.Itoa(serviceNode.Weight)
	} else {
		nginxRunUrl = nginxRunUrl + "&" + operateType + "="
	}

	log.Info("请求的 nginx 6000 的地址为 " + nginxRunUrl)

	resp, err := http.Get(nginxRunUrl)

	if err != nil {
		return "", errors.New("内部转发错误, " + err.Error())
	}

	stringResp, err :=ioutil.ReadAll(resp.Body)

	log.Info("nginx 6000 返回的结果 " + string(stringResp))

	if err != nil {
		return "", errors.New("内部转发错误, " + err.Error())
	}
	if strings.Contains(string(stringResp), "server") {
		if operateType == models.ADD{
			serviceNode.Join=true
		}
		if operateType == models.REMOVE {
			serviceNode.Join=false
		}
		id, err := orm.Update(&serviceNode, "join")
		if err != nil {
			return "", errors.New(strconv.Itoa(utils.Int64toint(id)) + "同时更新到 mysql 错误, " + err.Error())
		}
		log.Info("更新到 mysql 成功, " + strconv.Itoa(utils.Int64toint(id)))
	}
	return string(stringResp), nil
}
