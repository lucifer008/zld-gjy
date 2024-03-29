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

func newSysPermissionGroup(db *gorm.DB) sysPermissionGroup {
	_sysPermissionGroup := sysPermissionGroup{}

	_sysPermissionGroup.sysPermissionGroupDo.UseDB(db)
	_sysPermissionGroup.sysPermissionGroupDo.UseModel(&domain.SysPermissionGroup{})

	tableName := _sysPermissionGroup.sysPermissionGroupDo.TableName()
	_sysPermissionGroup.ALL = field.NewAsterisk(tableName)
	_sysPermissionGroup.ID = field.NewInt64(tableName, "Id")
	_sysPermissionGroup.SysID = field.NewInt64(tableName, "Sys_Id")
	_sysPermissionGroup.PermissionGroupType = field.NewInt32(tableName, "Permission_Group_Type")
	_sysPermissionGroup.PermissionGroupDesc = field.NewString(tableName, "Permission_Group_Desc")
	_sysPermissionGroup.InsertDateTime = field.NewTime(tableName, "Insert_Date_Time")
	_sysPermissionGroup.InsertUser = field.NewInt64(tableName, "Insert_User")
	_sysPermissionGroup.UpdateDateTime = field.NewTime(tableName, "Update_Date_Time")
	_sysPermissionGroup.UpdateUser = field.NewInt64(tableName, "Update_User")
	_sysPermissionGroup.Version = field.NewInt32(tableName, "Version")
	_sysPermissionGroup.Deleted = field.NewString(tableName, "Deleted")
	_sysPermissionGroup.CompayID = field.NewInt64(tableName, "Compay_Id")

	_sysPermissionGroup.fillFieldMap()

	return _sysPermissionGroup
}

type sysPermissionGroup struct {
	sysPermissionGroupDo sysPermissionGroupDo

	ALL                 field.Asterisk
	ID                  field.Int64  // Id
	SysID               field.Int64  // Id
	PermissionGroupType field.Int32  // 许可组类型
	PermissionGroupDesc field.String // 许可组描述
	InsertDateTime      field.Time   // 插入时间
	InsertUser          field.Int64  // 插入用户
	UpdateDateTime      field.Time   // 更新时间
	UpdateUser          field.Int64  // 更新用户
	Version             field.Int32  // 版本
	Deleted             field.String // 删除标志(0:正常 1 删除)
	CompayID            field.Int64  // 公司

	fieldMap map[string]field.Expr
}

func (s sysPermissionGroup) Table(newTableName string) *sysPermissionGroup {
	s.sysPermissionGroupDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysPermissionGroup) As(alias string) *sysPermissionGroup {
	s.sysPermissionGroupDo.DO = *(s.sysPermissionGroupDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysPermissionGroup) updateTableName(table string) *sysPermissionGroup {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "Id")
	s.SysID = field.NewInt64(table, "Sys_Id")
	s.PermissionGroupType = field.NewInt32(table, "Permission_Group_Type")
	s.PermissionGroupDesc = field.NewString(table, "Permission_Group_Desc")
	s.InsertDateTime = field.NewTime(table, "Insert_Date_Time")
	s.InsertUser = field.NewInt64(table, "Insert_User")
	s.UpdateDateTime = field.NewTime(table, "Update_Date_Time")
	s.UpdateUser = field.NewInt64(table, "Update_User")
	s.Version = field.NewInt32(table, "Version")
	s.Deleted = field.NewString(table, "Deleted")
	s.CompayID = field.NewInt64(table, "Compay_Id")

	s.fillFieldMap()

	return s
}

func (s *sysPermissionGroup) WithContext(ctx context.Context) *sysPermissionGroupDo {
	return s.sysPermissionGroupDo.WithContext(ctx)
}

func (s sysPermissionGroup) TableName() string { return s.sysPermissionGroupDo.TableName() }

func (s sysPermissionGroup) Alias() string { return s.sysPermissionGroupDo.Alias() }

func (s *sysPermissionGroup) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysPermissionGroup) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 11)
	s.fieldMap["Id"] = s.ID
	s.fieldMap["Sys_Id"] = s.SysID
	s.fieldMap["Permission_Group_Type"] = s.PermissionGroupType
	s.fieldMap["Permission_Group_Desc"] = s.PermissionGroupDesc
	s.fieldMap["Insert_Date_Time"] = s.InsertDateTime
	s.fieldMap["Insert_User"] = s.InsertUser
	s.fieldMap["Update_Date_Time"] = s.UpdateDateTime
	s.fieldMap["Update_User"] = s.UpdateUser
	s.fieldMap["Version"] = s.Version
	s.fieldMap["Deleted"] = s.Deleted
	s.fieldMap["Compay_Id"] = s.CompayID
}

func (s sysPermissionGroup) clone(db *gorm.DB) sysPermissionGroup {
	s.sysPermissionGroupDo.ReplaceDB(db)
	return s
}

type sysPermissionGroupDo struct{ gen.DO }

func (s sysPermissionGroupDo) Debug() *sysPermissionGroupDo {
	return s.withDO(s.DO.Debug())
}

func (s sysPermissionGroupDo) WithContext(ctx context.Context) *sysPermissionGroupDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysPermissionGroupDo) ReadDB() *sysPermissionGroupDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysPermissionGroupDo) WriteDB() *sysPermissionGroupDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysPermissionGroupDo) Clauses(conds ...clause.Expression) *sysPermissionGroupDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysPermissionGroupDo) Returning(value interface{}, columns ...string) *sysPermissionGroupDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysPermissionGroupDo) Not(conds ...gen.Condition) *sysPermissionGroupDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysPermissionGroupDo) Or(conds ...gen.Condition) *sysPermissionGroupDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysPermissionGroupDo) Select(conds ...field.Expr) *sysPermissionGroupDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysPermissionGroupDo) Where(conds ...gen.Condition) *sysPermissionGroupDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysPermissionGroupDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *sysPermissionGroupDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysPermissionGroupDo) Order(conds ...field.Expr) *sysPermissionGroupDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysPermissionGroupDo) Distinct(cols ...field.Expr) *sysPermissionGroupDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysPermissionGroupDo) Omit(cols ...field.Expr) *sysPermissionGroupDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysPermissionGroupDo) Join(table schema.Tabler, on ...field.Expr) *sysPermissionGroupDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysPermissionGroupDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sysPermissionGroupDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysPermissionGroupDo) RightJoin(table schema.Tabler, on ...field.Expr) *sysPermissionGroupDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysPermissionGroupDo) Group(cols ...field.Expr) *sysPermissionGroupDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysPermissionGroupDo) Having(conds ...gen.Condition) *sysPermissionGroupDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysPermissionGroupDo) Limit(limit int) *sysPermissionGroupDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysPermissionGroupDo) Offset(offset int) *sysPermissionGroupDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysPermissionGroupDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sysPermissionGroupDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysPermissionGroupDo) Unscoped() *sysPermissionGroupDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysPermissionGroupDo) Create(values ...*domain.SysPermissionGroup) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysPermissionGroupDo) CreateInBatches(values []*domain.SysPermissionGroup, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysPermissionGroupDo) Save(values ...*domain.SysPermissionGroup) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysPermissionGroupDo) First() (*domain.SysPermissionGroup, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*domain.SysPermissionGroup), nil
	}
}

func (s sysPermissionGroupDo) Take() (*domain.SysPermissionGroup, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*domain.SysPermissionGroup), nil
	}
}

func (s sysPermissionGroupDo) Last() (*domain.SysPermissionGroup, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*domain.SysPermissionGroup), nil
	}
}

func (s sysPermissionGroupDo) Find() ([]*domain.SysPermissionGroup, error) {
	result, err := s.DO.Find()
	return result.([]*domain.SysPermissionGroup), err
}

func (s sysPermissionGroupDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*domain.SysPermissionGroup, err error) {
	buf := make([]*domain.SysPermissionGroup, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysPermissionGroupDo) FindInBatches(result *[]*domain.SysPermissionGroup, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysPermissionGroupDo) Attrs(attrs ...field.AssignExpr) *sysPermissionGroupDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysPermissionGroupDo) Assign(attrs ...field.AssignExpr) *sysPermissionGroupDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysPermissionGroupDo) Joins(fields ...field.RelationField) *sysPermissionGroupDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysPermissionGroupDo) Preload(fields ...field.RelationField) *sysPermissionGroupDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysPermissionGroupDo) FirstOrInit() (*domain.SysPermissionGroup, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*domain.SysPermissionGroup), nil
	}
}

func (s sysPermissionGroupDo) FirstOrCreate() (*domain.SysPermissionGroup, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*domain.SysPermissionGroup), nil
	}
}

func (s sysPermissionGroupDo) FindByPage(offset int, limit int) (result []*domain.SysPermissionGroup, count int64, err error) {
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

func (s sysPermissionGroupDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysPermissionGroupDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysPermissionGroupDo) Delete(models ...*domain.SysPermissionGroup) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysPermissionGroupDo) withDO(do gen.Dao) *sysPermissionGroupDo {
	s.DO = *do.(*gen.DO)
	return s
}
