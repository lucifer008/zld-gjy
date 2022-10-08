package models

import "zld-jy/da/domain"

type Orders struct {
	Base
}
type OrdersModel struct {
	domain.Order
}
type NewOrderModel struct {
	Orders      domain.Order
	OrdersGoods []domain.OrdersGood
}
