package vo

type PageResponse struct {
	PageNo int `json:"pageNo"`

	PageSize int `json:"pageSize"`

	TotalPages int `json:"totalPages"`

	TotalElements int `json:"totalElements"`

	First bool `json:"first"`

	Last bool `json:"last"`

	Content interface{} `json:"content"`
}
