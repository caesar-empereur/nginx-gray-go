package models

type ServiceNodeRedis struct {
	Id int `orm:"column(id);pk" json:"id"`

	Ip string `orm:"column(ip)" json:"ip"`

	Port string `orm:"column(port)" json:"port"`

	ServiceName string `orm:"column(service_name)" json:"service_name"`

	Domain string `orm:"column(domain)" json:"domain"`

	Weight int `orm:"column(weight)" json:"weight"`

	InnerAccess bool `orm:"column(inner_access)" json:"inner_access"`

	Operate string `orm:"column(operate)" json:"operate"`

	Zone string `orm:"column(zone)" json:"zone"`
}
