// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameOrganization = "Organization"

// Organization mapped from table <Organization>
type Organization struct {
	ID             int64     `gorm:"column:Id;primaryKey" json:"Id"`                   // Id
	ComID          int64     `gorm:"column:Com_Id" json:"Com_Id"`                      // Id
	OrgID          int64     `gorm:"column:Org_Id" json:"Org_Id"`                      // Id
	OrgName        string    `gorm:"column:Org_Name" json:"Org_Name"`                  // 组织名称
	OrgCode        string    `gorm:"column:Org_Code" json:"Org_Code"`                  // 组织代码
	OrgLevel       int32     `gorm:"column:Org_Level" json:"Org_Level"`                // 组织级别
	BeginDate      time.Time `gorm:"column:Begin_Date" json:"Begin_Date"`              // 开始日期
	EndDate        time.Time `gorm:"column:End_Date" json:"End_Date"`                  // 结束日期
	InsertDateTime time.Time `gorm:"column:Insert_Date_Time" json:"Insert_Date_Time"`  // 插入时间
	InsertUser     int64     `gorm:"column:Insert_User" json:"Insert_User"`            // 插入用户
	UpdateDateTime time.Time `gorm:"column:Update_Date_Time" json:"Update_Date_Time"`  // 更新时间
	UpdateUser     int64     `gorm:"column:Update_User" json:"Update_User"`            // 更新用户
	Version        int32     `gorm:"column:Version;default:1" json:"Version"`          // 版本
	Deleted        string    `gorm:"column:Deleted;not null;default:0" json:"Deleted"` // 删除标志(0:正常 1 删除)
}

// TableName Organization's table name
func (*Organization) TableName() string {
	return TableNameOrganization
}