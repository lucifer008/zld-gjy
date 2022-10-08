package action

import (
	"context"
	"zld-jy/da/base"
	"zld-jy/da/dao"
	"zld-jy/da/domain"
	"zld-jy/models"
)

var CustomerActions CustomerAction

func init() {
	CustomerActions = CustomerAction{}
}

type CustomerAction struct {
	BaseAction
}

// Query
//@Summary 查询客户信息
//@Schemes
//@Description 查询客户信息
//@Tags 客户
//@Param q query models.Customers true "查询参数"
//@Accept json
//@Produce json
//@Success 200
//@Router /customers/query [get]
func (customerAction CustomerAction) Query(customers models.Customers) (total int64, data []models.CustomerModel, err error) {
	qur := dao.Use(base.DB)
	if customers.CustomerName != "" {
		result, total, err := qur.Customer.WithContext(context.Background()).Where(qur.Customer.CustomerName.Like(customers.CustomerName)).FindByPage(customers.CurrentIndex, customers.PageSize)
		if err != nil {
			panic(err)
		}

		//for key, v := range result {
		//	var res models.CustomerModel
		//	utils.SimpleCopyProperties(&v, res)
		//	data[key] = res
		//}
		for _, v := range result {
			var res models.CustomerModel
			res = models.CustomerModel{Customer: domain.Customer{ID: v.ID, CustomerName: v.CustomerName, CustomerAddress: v.CustomerAddress,
				ContactTel: v.ContactTel, ContactMan: v.ContactMan, EnterTel: v.EnterTel, ProvinceCode: v.ProvinceCode, CityCode: v.CityCode, AreaCode: v.AreaCode}}
			data = append(data, res)
		}
		return total, data, err
	}
	result, total, err := qur.Customer.WithContext(context.Background()).FindByPage(customers.CurrentIndex, customers.PageSize)
	if err != nil {
		panic(err)
	}
	for _, v := range result {
		var res models.CustomerModel
		res = models.CustomerModel{Customer: domain.Customer{ID: v.ID, CustomerName: v.CustomerName, CustomerAddress: v.CustomerAddress,
			ContactTel: v.ContactTel, ContactMan: v.ContactMan, EnterTel: v.EnterTel, ProvinceCode: v.ProvinceCode, CityCode: v.CityCode, AreaCode: v.AreaCode}}
		data = append(data, res)
	}
	//utils.SimpleCopyProperties(&data, result)
	return total, data, err
}
