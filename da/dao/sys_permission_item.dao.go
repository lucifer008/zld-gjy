// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"zld-jy/da/domain"
)

func newSysPermissionItem(db *gorm.DB) sysPermissionItem {
	_sysPermissionItem := sysPermissionItem{}

	_sysPermissionItem.sysPermissionItemDo.UseDB(db)
	_sysPermissionItem.sysPermissionItemDo.UseModel(&domain.SysPermissionItem{})

	tableName := _sysPermissionItem.sysPermissionItemDo.TableName()
	_sysPermissionItem.ALL = field.NewAsterisk(tableName)
	_sysPermissionItem.ID = field.NewInt64(tableName, "Id")
	_sysPermissionItem.SysID = field.NewInt64(tableName, "Sys_Id")
	_sysPermissionItem.SysId2 = field.NewInt64(tableName, "Sys_Id2")
	_sysPermissionItem.PermissionItemVal = field.NewInt32(tableName, "Permission_Item_Val")
	_sysPermissionItem.PermissionType = field.NewInt32(tableName, "Permission_Type")
	_sysPermissionItem.PermissionDesc = field.NewString(tableName, "Permission_Desc")

	_sysPermissionItem.fillFieldMap()

	return _sysPermissionItem
}

type sysPermissionItem struct {
	sysPermissionItemDo sysPermissionItemDo

	ALL               field.Asterisk
	ID                field.Int64  // Id
	SysID             field.Int64  // Id
	SysId2            field.Int64  // 资源Id
	PermissionItemVal field.Int32  // 权限值
	PermissionType    field.Int32  // 权限类型
	PermissionDesc    field.String // 权限描述

	fieldMap map[string]field.Expr
}

func (s sysPermissionItem) Table(newTableName string) *sysPermissionItem {
	s.sysPermissionItemDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysPermissionItem) As(alias string) *sysPermissionItem {
	s.sysPermissionItemDo.DO = *(s.sysPermissionItemDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysPermissionItem) updateTableName(table string) *sysPermissionItem {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "Id")
	s.SysID = field.NewInt64(table, "Sys_Id")
	s.SysId2 = field.NewInt64(table, "Sys_Id2")
	s.PermissionItemVal = field.NewInt32(table, "Permission_Item_Val")
	s.PermissionType = field.NewInt32(table, "Permission_Type")
	s.PermissionDesc = field.NewString(table, "Permission_Desc")

	s.fillFieldMap()

	return s
}

func (s *sysPermissionItem) WithContext(ctx context.Context) *sysPermissionItemDo {
	return s.sysPermissionItemDo.WithContext(ctx)
}

func (s sysPermissionItem) TableName() string { return s.sysPermissionItemDo.TableName() }

func (s sysPermissionItem) Alias() string { return s.sysPermissionItemDo.Alias() }

func (s *sysPermissionItem) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysPermissionItem) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 6)
	s.fieldMap["Id"] = s.ID
	s.fieldMap["Sys_Id"] = s.SysID
	s.fieldMap["Sys_Id2"] = s.SysId2
	s.fieldMap["Permission_Item_Val"] = s.PermissionItemVal
	s.fieldMap["Permission_Type"] = s.PermissionType
	s.fieldMap["Permission_Desc"] = s.PermissionDesc
}

func (s sysPermissionItem) clone(db *gorm.DB) sysPermissionItem {
	s.sysPermissionItemDo.ReplaceDB(db)
	return s
}

type sysPermissionItemDo struct{ gen.DO }

func (s sysPermissionItemDo) Debug() *sysPermissionItemDo {
	return s.withDO(s.DO.Debug())
}

func (s sysPermissionItemDo) WithContext(ctx context.Context) *sysPermissionItemDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysPermissionItemDo) ReadDB() *sysPermissionItemDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysPermissionItemDo) WriteDB() *sysPermissionItemDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysPermissionItemDo) Clauses(conds ...clause.Expression) *sysPermissionItemDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysPermissionItemDo) Returning(value interface{}, columns ...string) *sysPermissionItemDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysPermissionItemDo) Not(conds ...gen.Condition) *sysPermissionItemDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysPermissionItemDo) Or(conds ...gen.Condition) *sysPermissionItemDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysPermissionItemDo) Select(conds ...field.Expr) *sysPermissionItemDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysPermissionItemDo) Where(conds ...gen.Condition) *sysPermissionItemDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysPermissionItemDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *sysPermissionItemDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysPermissionItemDo) Order(conds ...field.Expr) *sysPermissionItemDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysPermissionItemDo) Distinct(cols ...field.Expr) *sysPermissionItemDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysPermissionItemDo) Omit(cols ...field.Expr) *sysPermissionItemDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysPermissionItemDo) Join(table schema.Tabler, on ...field.Expr) *sysPermissionItemDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysPermissionItemDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sysPermissionItemDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysPermissionItemDo) RightJoin(table schema.Tabler, on ...field.Expr) *sysPermissionItemDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysPermissionItemDo) Group(cols ...field.Expr) *sysPermissionItemDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysPermissionItemDo) Having(conds ...gen.Condition) *sysPermissionItemDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysPermissionItemDo) Limit(limit int) *sysPermissionItemDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysPermissionItemDo) Offset(offset int) *sysPermissionItemDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysPermissionItemDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sysPermissionItemDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysPermissionItemDo) Unscoped() *sysPermissionItemDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysPermissionItemDo) Create(values ...*domain.SysPermissionItem) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysPermissionItemDo) CreateInBatches(values []*domain.SysPermissionItem, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysPermissionItemDo) Save(values ...*domain.SysPermissionItem) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysPermissionItemDo) First() (*domain.SysPermissionItem, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*domain.SysPermissionItem), nil
	}
}

func (s sysPermissionItemDo) Take() (*domain.SysPermissionItem, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*domain.SysPermissionItem), nil
	}
}

func (s sysPermissionItemDo) Last() (*domain.SysPermissionItem, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*domain.SysPermissionItem), nil
	}
}

func (s sysPermissionItemDo) Find() ([]*domain.SysPermissionItem, error) {
	result, err := s.DO.Find()
	return result.([]*domain.SysPermissionItem), err
}

func (s sysPermissionItemDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*domain.SysPermissionItem, err error) {
	buf := make([]*domain.SysPermissionItem, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysPermissionItemDo) FindInBatches(result *[]*domain.SysPermissionItem, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysPermissionItemDo) Attrs(attrs ...field.AssignExpr) *sysPermissionItemDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysPermissionItemDo) Assign(attrs ...field.AssignExpr) *sysPermissionItemDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysPermissionItemDo) Joins(fields ...field.RelationField) *sysPermissionItemDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysPermissionItemDo) Preload(fields ...field.RelationField) *sysPermissionItemDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysPermissionItemDo) FirstOrInit() (*domain.SysPermissionItem, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*domain.SysPermissionItem), nil
	}
}

func (s sysPermissionItemDo) FirstOrCreate() (*domain.SysPermissionItem, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*domain.SysPermissionItem), nil
	}
}

func (s sysPermissionItemDo) FindByPage(offset int, limit int) (result []*domain.SysPermissionItem, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysPermissionItemDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysPermissionItemDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysPermissionItemDo) Delete(models ...*domain.SysPermissionItem) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysPermissionItemDo) withDO(do gen.Dao) *sysPermissionItemDo {
	s.DO = *do.(*gen.DO)
	return s
}
