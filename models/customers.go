package models

type Customers struct {
	Base
	CustomerName string `form:"customerName" json:"customerName" example:"客户名称"`
}
