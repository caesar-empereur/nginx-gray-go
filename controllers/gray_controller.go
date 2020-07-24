package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	orm2 "github.com/astaxie/beego/orm"
	"nginx-gray-go/models"
	"nginx-gray-go/repository"
	"nginx-gray-go/utils"
	"nginx-gray-go/vo"
)

type GrayController struct {
	beego.Controller
}

func (this *GrayController) List() {
	var pageNo, _ = this.GetInt("pageNo")
	var pageSize, _ = this.GetInt("pageSize")

	var nodes []models.ServiceNode
	sqlParam := make(map[string]interface{})

	pageResponse := repository.PageQuery(pageNo, pageSize, sqlParam, "service_node", &nodes)

	this.Data["json"] = pageResponse
	this.ServeJSON()
}

func (this *GrayController) AddOne() {
	log := utils.GetLogger()

	var serviceNode models.ServiceNode
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &serviceNode)

	apiResponse := vo.ApiResponse{}
	if err != nil {
		apiResponse.Success = false
		apiResponse.Message = "参数缺失"
	} else {
		apiResponse.Success = true
	}

	orm := orm2.NewOrm()
	id, err := orm.Insert(serviceNode)

	if err != nil {
		panic("插入错误")
	}
	log.Debug(string(id))
}

func (this *GrayController) UpdateOne() {
	log := utils.GetLogger()

	var serviceNode models.ServiceNode
	json.Unmarshal(this.Ctx.Input.RequestBody, &serviceNode)

	apiResponse := vo.ApiResponse{}
	if serviceNode.Id == 0 {
		apiResponse.Success = false
		apiResponse.Message = "参数缺失"
	} else {
		apiResponse.Success = true
	}

	orm := orm2.NewOrm()
	id, err := orm.Update(&serviceNode, "ip", "weight", "port", "inner_access")

	if err != nil {
		panic("更新错误")
	}
	log.Debug(string(id))
}

func (this *GrayController) DeleteOne() {
	apiResponse := vo.ApiResponse{}

	var id, _ = this.GetInt("id")
	if id == 0 {
		apiResponse.Success = false
		apiResponse.Message = "参数缺失"
		this.Data["json"] = apiResponse
		this.ServeJSON()
		return
	}

	orm := orm2.NewOrm()
	_, err := orm.Delete(&models.ServiceNode{Id: id})

	if err != nil {
		panic("删除错误")
	}
}
