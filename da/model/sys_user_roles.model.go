// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSysUserRole = "Sys_User_Roles"

// SysUserRole mapped from table <Sys_User_Roles>
type SysUserRole struct {
	ID     int64 `gorm:"column:Id;primaryKey" json:"Id"`
	SysID  int64 `gorm:"column:Sys_Id" json:"Sys_Id"`   // 角色Id
	SysId2 int64 `gorm:"column:Sys_Id2" json:"Sys_Id2"` // Id
}

// TableName SysUserRole's table name
func (*SysUserRole) TableName() string {
	return TableNameSysUserRole
}
