package action

import (
	"zld-jy/models"
)

var CustomerActions CustomerAction

func init() {
	CustomerActions = CustomerAction{}
}

type CustomerAction struct {
}

// Query
//@Summary 查询客户信息
// @Schemes
// @Description 查询客户信息
// @Tags 客户
// @Param q query models.Customers true "查询参数"
// @Accept json
// @Produce json
// @Success 200
// @Router /customers/query [get]
func (customerAction CustomerAction) Query(customers models.Customers) (total int64, data []models.Customers) {
	total = 10
	return total, data
}
