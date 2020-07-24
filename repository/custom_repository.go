package repository

import (
	orm2 "github.com/astaxie/beego/orm"
	"nginx-gray-go/utils"
	"nginx-gray-go/vo"
)

func PageQuery(pageNo int, pageSize int, sqlParam map[string]interface{}, table string, resultSaved interface{}) vo.PageResponse {

	log := utils.GetLogger()

	if pageNo == 0 {
		pageNo = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}

	orm := orm2.NewOrm()
	querySetter := orm.QueryTable(table)

	for key, value := range sqlParam {
		querySetter = querySetter.Filter(key, value)
	}

	count, _ := querySetter.Count()
	log.Debug(" 条件查询行数: ", count)

	limitStart := utils.ProcessPageStart(pageNo, pageSize)
	querySetter.OrderBy("-id").Limit(pageSize, limitStart).All(resultSaved)

	pageResponse := utils.HandlePage(pageNo, pageSize, utils.Int64toint(count))
	pageResponse.Content = resultSaved

	return pageResponse
}
