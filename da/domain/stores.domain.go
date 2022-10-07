// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package domain

const TableNameStore = "Stores"

// Store mapped from table <Stores>
type Store struct {
	ID           int64  `gorm:"column:Id;primaryKey" json:"Id"`            // 仓储Id
	StoreName    string `gorm:"column:Store_Name" json:"Store_Name"`       // 仓储名称
	StoreTel     string `gorm:"column:Store_Tel" json:"Store_Tel"`         // 联系电话
	StoreAddress string `gorm:"column:Store_Address" json:"Store_Address"` // 地址
	Memo         string `gorm:"column:Memo" json:"Memo"`                   // 备注
}

// TableName Store's table name
func (*Store) TableName() string {
	return TableNameStore
}