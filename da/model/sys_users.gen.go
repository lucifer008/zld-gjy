// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysUser = "Sys_Users"

// SysUser mapped from table <Sys_Users>
type SysUser struct {
	Id2             int64     `gorm:"column:Id2;primaryKey" json:"Id2"`                  // Id
	Id13            int64     `gorm:"column:Id13" json:"Id13"`                           // Id
	Id5             int64     `gorm:"column:Id5" json:"Id5"`                             // 雇员Id
	UserName        string    `gorm:"column:User_Name" json:"User_Name"`                 // 用户名称
	UserEmail       string    `gorm:"column:User_Email" json:"User_Email"`               // 用户邮箱
	UserPassword    string    `gorm:"column:User_Password" json:"User_Password"`         // 登录密码
	UserStatus      int32     `gorm:"column:User_Status" json:"User_Status"`             // 用户状态(0:正常 1 :停用 )
	UserType        int32     `gorm:"column:User_Type" json:"User_Type"`                 // 用户类型
	InsertDateTime3 time.Time `gorm:"column:Insert_Date_Time3" json:"Insert_Date_Time3"` // 插入时间
	InsertUser2     int64     `gorm:"column:Insert_User2" json:"Insert_User2"`           // 插入用户
	UpdateDateTime4 time.Time `gorm:"column:Update_Date_Time4" json:"Update_Date_Time4"` // 更新时间
	UpdateUser2     int64     `gorm:"column:Update_User2" json:"Update_User2"`           // 更新用户
	Version3        int32     `gorm:"column:Version3" json:"Version3"`                   // 版本
	Deleted2        string    `gorm:"column:Deleted2;not null" json:"Deleted2"`          // 删除标志(0:正常 1 删除)
	CompayID        int64     `gorm:"column:Compay_Id" json:"Compay_Id"`                 // 公司
}

// TableName SysUser's table name
func (*SysUser) TableName() string {
	return TableNameSysUser
}
