package action

type BaseAction struct {
}

// Query 基类方法
func (action BaseAction) Query(params interface{}) (total int64, data interface{}) {
	return total, data
}
