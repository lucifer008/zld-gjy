// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSysRolePermissionGroup = "Sys_Role_Permission_Group"

// SysRolePermissionGroup mapped from table <Sys_Role_Permission_Group>
type SysRolePermissionGroup struct {
	ID     int64 `gorm:"column:Id;primaryKey" json:"Id"`
	SysID  int64 `gorm:"column:Sys_Id" json:"Sys_Id"`   // Id
	SysId2 int64 `gorm:"column:Sys_Id2" json:"Sys_Id2"` // 角色Id
}

// TableName SysRolePermissionGroup's table name
func (*SysRolePermissionGroup) TableName() string {
	return TableNameSysRolePermissionGroup
}