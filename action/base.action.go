package action

import "time"

type BaseAction struct {
}

var currentDatetime = time.Now()
var insertDatetime = currentDatetime
var updateDatetime = currentDatetime
var insertUser int64 = 1
var updateUser int64 = 1
var deleted = "0"
var version int32 = 1
var companyId int64 = 1

// Query 基类方法
func (action BaseAction) Query(params interface{}) (total int64, data interface{}) {
	return total, data
}
