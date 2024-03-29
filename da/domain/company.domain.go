// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package domain

import (
	"time"
)

const TableNameCompany = "Company"

// Company mapped from table <Company>
type Company struct {
	ID               int64     `gorm:"column:Id;primaryKey" json:"Id"`                      // Id
	CompanyName      string    `gorm:"column:Company_Name" json:"Company_Name"`             // 公司名称
	CompanyType      int32     `gorm:"column:Company_Type" json:"Company_Type"`             // 公司类型(0:政府 1: 消防 3 :企业 )
	CompanyDesc      string    `gorm:"column:Company_Desc" json:"Company_Desc"`             // 描述
	RegisterDateTime time.Time `gorm:"column:Register_Date_Time" json:"Register_Date_Time"` // 注册时间
	Status           int32     `gorm:"column:Status" json:"Status"`                         // 状态(0: 正常 1:停用)
	InsertDateTime   time.Time `gorm:"column:Insert_Date_Time" json:"Insert_Date_Time"`     // 插入时间
	InsertUser       int64     `gorm:"column:Insert_User" json:"Insert_User"`               // 插入用户
	UpdateDateTime   time.Time `gorm:"column:Update_Date_Time" json:"Update_Date_Time"`     // 更新时间
	UpdateUser       int64     `gorm:"column:Update_User" json:"Update_User"`               // 更新用户
	Version          int32     `gorm:"column:Version;default:1" json:"Version"`             // 版本
	Deleted          string    `gorm:"column:Deleted;not null;default:0" json:"Deleted"`    // 删除标志(0:正常 1 删除)
}

// TableName Company's table name
func (*Company) TableName() string {
	return TableNameCompany
}
