// Code generated by "zld-jy/da/field". DO NOT EDIT.
// Code generated by "zld-jy/da/field". DO NOT EDIT.
// Code generated by "zld-jy/da/field". DO NOT EDIT.

package model

const TableNameSysUserGroupRole = "sys_User_Group_Roles"

// SysUserGroupRole mapped from table <sys_User_Group_Roles>
type SysUserGroupRole struct {
	Id4 int64 `gorm:"column:Id4;primaryKey" json:"Id4"`
	Id3 int64 `gorm:"column:Id3" json:"Id3"` // 角色Id
	Id8 int64 `gorm:"column:Id8" json:"Id8"` // Id
}

// TableName SysUserGroupRole's table name
func (*SysUserGroupRole) TableName() string {
	return TableNameSysUserGroupRole
}
