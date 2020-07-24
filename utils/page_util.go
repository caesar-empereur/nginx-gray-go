package utils

import "nginx-gray-go/vo"

func HandlePage(pageNo int, pageSize int, resCount int) vo.PageResponse {

	pageResponse := vo.PageResponse{}
	pageResponse.PageNo = pageNo
	pageResponse.PageSize = pageSize
	pageResponse.TotalElements = resCount
	var rest int
	if resCount%pageSize > 0 {
		rest = 1
	} else {
		rest = 0
	}
	pageResponse.TotalPages = resCount/pageSize + rest
	pageResponse.First = pageNo == 1
	pageResponse.Last = pageResponse.TotalPages == pageNo

	return pageResponse
}

func ProcessPageStart(pageNo int, pageSize int) int {
	return (pageNo - 1) * pageSize
}
