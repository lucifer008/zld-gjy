// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSysPermissionItem = "Sys_Permission_Item"

// SysPermissionItem mapped from table <Sys_Permission_Item>
type SysPermissionItem struct {
	ID                int64  `gorm:"column:Id;primaryKey" json:"Id"`                        // Id
	SysID             int64  `gorm:"column:Sys_Id" json:"Sys_Id"`                           // Id
	SysId2            int64  `gorm:"column:Sys_Id2" json:"Sys_Id2"`                         // 资源Id
	PermissionItemVal int32  `gorm:"column:Permission_Item_Val" json:"Permission_Item_Val"` // 权限值
	PermissionType    int32  `gorm:"column:Permission_Type" json:"Permission_Type"`         // 权限类型
	PermissionDesc    string `gorm:"column:Permission_Desc" json:"Permission_Desc"`         // 权限描述
}

// TableName SysPermissionItem's table name
func (*SysPermissionItem) TableName() string {
	return TableNameSysPermissionItem
}
