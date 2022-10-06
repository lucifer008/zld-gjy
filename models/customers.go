package models

import "zld-jy/da/domain"

type Customers struct {
	Base
	CustomerName string `form:"customerName" json:"customerName" example:"客户名称"`
}
type CustomerModel struct {
	domain.Customer
}
