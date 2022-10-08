package action

import (
	"context"
	"zld-jy/da/base"
	"zld-jy/da/dao"
	"zld-jy/da/domain"
	"zld-jy/models"
)

//订单相关
type OrdersAction struct {
	BaseAction
}

var OrderActionInstance OrdersAction

func init() {
	OrderActionInstance = OrdersAction{}
}

// Query
//@Summary 查询订单信息
//@Schemes
//@Description 查询订单信息
//@Tags 订单
//@Param q query models.Orders true "查询参数"
//@Accept json
//@Produce json
//@Success 200
//@Router /orders/query [get]
func (oa OrdersAction) Query(orders models.Orders) (total int64, data []models.OrdersModel, err error) {
	qur := dao.Use(base.DB)
	result, total, err := qur.Order.WithContext(context.Background()).FindByPage(orders.CurrentIndex, orders.PageSize)
	if err != nil {
		panic(err)
	}
	for _, v := range result {
		var res models.OrdersModel
		res = models.OrdersModel{Order: domain.Order{ID: v.ID, OrderNo: v.OrderNo, OutOrderNo: v.OutOrderNo, ProjectName: v.ProjectName,
			ShipmentDate: v.ShipmentDate, Status: v.Status, DeliverProvinceCode: v.DeliverProvinceCode, DeliverCityCode: v.DeliverCityCode,
			DeliverAreaCode: v.DeliverAreaCode, DeliverAddress: v.DeliverAddress, ArriveAddress: v.ArriveAddress, ArriveProvinceCode: v.ArriveProvinceCode,
			ArriveCityCode: v.ArriveCityCode, ArriveAreaCode: v.ArriveAreaCode, ArriveDate: v.ArriveDate, ReceiptCustomer: v.ReceiptCustomer, ReceiptCustomerTel: v.ReceiptCustomerTel, PricingMethod: v.PricingMethod,
		}}
		data = append(data, res)
	}
	return total, data, err
}
