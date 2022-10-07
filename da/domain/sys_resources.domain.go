// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package domain

import (
	"time"
)

const TableNameSysResource = "Sys_Resources"

// SysResource mapped from table <Sys_Resources>
type SysResource struct {
	ID                     int64     `gorm:"column:Id;primaryKey" json:"Id"`                                  // 资源Id
	SysID                  int64     `gorm:"column:Sys_Id" json:"Sys_Id"`                                     // 资源Id
	ResourceIdentity       string    `gorm:"column:Resource_Identity;not null" json:"Resource_Identity"`      // 资源标示
	ResourceType           int32     `gorm:"column:Resource_Type;not null" json:"Resource_Type"`              // 资源类型(0 app   1 模块  2  菜单 3  页面 4  按钮 5  扩展按钮)
	ResourceStatus         int32     `gorm:"column:Resource_Status" json:"Resource_Status"`                   // 资源状态(0:正常 1 :停用)
	ResourceActopnType     int32     `gorm:"column:Resource_Actopn_Type" json:"Resource_Actopn_Type"`         // 资源动作类型 (-1 无 1.增加;2 删 3 改 4 上传 5导出 6 查询 .)
	ResourceActionIdentity string    `gorm:"column:Resource_Action_Identity" json:"Resource_Action_Identity"` // 资源动作标示(针对扩展按钮，备用字段)
	ResourceDesc           string    `gorm:"column:Resource_Desc" json:"Resource_Desc"`                       // 资源描述
	ResourceOrders         int32     `gorm:"column:Resource_Orders" json:"Resource_Orders"`                   // 资源序序号
	InsertDateTime         time.Time `gorm:"column:Insert_Date_Time;not null" json:"Insert_Date_Time"`        // 插入时间
	InsertUser             int64     `gorm:"column:Insert_User;not null" json:"Insert_User"`                  // 插入用户
	UpdateDateTime         time.Time `gorm:"column:Update_Date_Time" json:"Update_Date_Time"`                 // 更新时间
	UpdateUser             int64     `gorm:"column:Update_User" json:"Update_User"`                           // 更新用户
	Version                int32     `gorm:"column:Version;not null" json:"Version"`                          // 版本
	Deleted                string    `gorm:"column:Deleted;not null" json:"Deleted"`                          // 删除标志
	CompayId3              int64     `gorm:"column:Compay_Id3;not null" json:"Compay_Id3"`                    // 公司
}

// TableName SysResource's table name
func (*SysResource) TableName() string {
	return TableNameSysResource
}