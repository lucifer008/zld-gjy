package action

import (
	"context"
	"zld-jy/da/base"
	"zld-jy/da/dao"
	"zld-jy/models"
	"zld-jy/utils"
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
// @Schemes
// @Description 查询客户信息
// @Tags 客户
// @Param q dao models.Customers true "查询参数"
// @Accept json
// @Produce json
// @Success 200
// @Router /customers/dao [get]
func (customerAction CustomerAction) Query(customers models.Customers) (total int64, data []models.CustomerModel, err error) {
	qur := dao.Use(base.DB)
	if customers.CustomerName != "" {
		result, total, err := qur.Customer.WithContext(context.Background()).Where(qur.Customer.CustomerName.Like(customers.CustomerName)).FindByPage(customers.CurrentIndex, customers.PageSize)
		if err != nil {
			panic(err)
		}
		utils.SimpleCopyProperties(&data, result)
		return total, data, err
	}
	result, total, err := qur.Customer.WithContext(context.Background()).FindByPage(customers.CurrentIndex, customers.PageSize)
	if err != nil {
		panic(err)
	}
	utils.SimpleCopyProperties(&data, result)
	return total, data, err
}
