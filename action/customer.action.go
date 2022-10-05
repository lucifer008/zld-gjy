package action

import "zld-jy/models"

var CustomerActions CustomerAction

func init() {
	CustomerActions = CustomerAction{}
}

type CustomerAction struct {
}

func (customerAction CustomerAction) Query(customers models.Customers) {

}
