package models

import "zld-jy/da/model"

type Customers struct {
	Base
	CustomerName string `form:"customerName" json:"customerName" example:"客户名称"`
}
type CustomerModel struct {
	model.Customer
}
