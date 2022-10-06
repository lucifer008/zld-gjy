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

func newSysReourceExt(db *gorm.DB) sysReourceExt {
	_sysReourceExt := sysReourceExt{}

	_sysReourceExt.sysReourceExtDo.UseDB(db)
	_sysReourceExt.sysReourceExtDo.UseModel(&domain.SysReourceExt{})

	tableName := _sysReourceExt.sysReourceExtDo.TableName()
	_sysReourceExt.ALL = field.NewAsterisk(tableName)
	_sysReourceExt.ID = field.NewInt64(tableName, "Id")
	_sysReourceExt.Component = field.NewString(tableName, "Component")
	_sysReourceExt.AlwaysShow = field.NewBool(tableName, "Always_Show")
	_sysReourceExt.MetaTitle = field.NewString(tableName, "Meta_Title")
	_sysReourceExt.MetaLimit = field.NewString(tableName, "Meta_Limit")
	_sysReourceExt.MetaBreadcrumb = field.NewString(tableName, "Meta_Breadcrumb")
	_sysReourceExt.MetaIcon = field.NewString(tableName, "Meta_Icon")
	_sysReourceExt.MetaNoCache = field.NewString(tableName, "Meta_No_Cache")
	_sysReourceExt.MetaHidden = field.NewString(tableName, "Meta_Hidden")
	_sysReourceExt.Redirect = field.NewString(tableName, "Redirect")
	_sysReourceExt.Hidden = field.NewString(tableName, "Hidden")
	_sysReourceExt.Type = field.NewInt32(tableName, "Type")

	_sysReourceExt.fillFieldMap()

	return _sysReourceExt
}

type sysReourceExt struct {
	sysReourceExtDo sysReourceExtDo

	ALL            field.Asterisk
	ID             field.Int64  // 资源Id
	Component      field.String // 组件
	AlwaysShow     field.Bool   // 显示
	MetaTitle      field.String // 标题
	MetaLimit      field.String // 元数据限制
	MetaBreadcrumb field.String // 元数据面包屑
	MetaIcon       field.String // 元数据图标
	MetaNoCache    field.String // 元数据缓存
	MetaHidden     field.String // 元数据隐藏域
	Redirect       field.String // 重定向URL
	Hidden         field.String // 隐藏域
	Type           field.Int32  // 类型

	fieldMap map[string]field.Expr
}

func (s sysReourceExt) Table(newTableName string) *sysReourceExt {
	s.sysReourceExtDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysReourceExt) As(alias string) *sysReourceExt {
	s.sysReourceExtDo.DO = *(s.sysReourceExtDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysReourceExt) updateTableName(table string) *sysReourceExt {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "Id")
	s.Component = field.NewString(table, "Component")
	s.AlwaysShow = field.NewBool(table, "Always_Show")
	s.MetaTitle = field.NewString(table, "Meta_Title")
	s.MetaLimit = field.NewString(table, "Meta_Limit")
	s.MetaBreadcrumb = field.NewString(table, "Meta_Breadcrumb")
	s.MetaIcon = field.NewString(table, "Meta_Icon")
	s.MetaNoCache = field.NewString(table, "Meta_No_Cache")
	s.MetaHidden = field.NewString(table, "Meta_Hidden")
	s.Redirect = field.NewString(table, "Redirect")
	s.Hidden = field.NewString(table, "Hidden")
	s.Type = field.NewInt32(table, "Type")

	s.fillFieldMap()

	return s
}

func (s *sysReourceExt) WithContext(ctx context.Context) *sysReourceExtDo {
	return s.sysReourceExtDo.WithContext(ctx)
}

func (s sysReourceExt) TableName() string { return s.sysReourceExtDo.TableName() }

func (s sysReourceExt) Alias() string { return s.sysReourceExtDo.Alias() }

func (s *sysReourceExt) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysReourceExt) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 12)
	s.fieldMap["Id"] = s.ID
	s.fieldMap["Component"] = s.Component
	s.fieldMap["Always_Show"] = s.AlwaysShow
	s.fieldMap["Meta_Title"] = s.MetaTitle
	s.fieldMap["Meta_Limit"] = s.MetaLimit
	s.fieldMap["Meta_Breadcrumb"] = s.MetaBreadcrumb
	s.fieldMap["Meta_Icon"] = s.MetaIcon
	s.fieldMap["Meta_No_Cache"] = s.MetaNoCache
	s.fieldMap["Meta_Hidden"] = s.MetaHidden
	s.fieldMap["Redirect"] = s.Redirect
	s.fieldMap["Hidden"] = s.Hidden
	s.fieldMap["Type"] = s.Type
}

func (s sysReourceExt) clone(db *gorm.DB) sysReourceExt {
	s.sysReourceExtDo.ReplaceDB(db)
	return s
}

type sysReourceExtDo struct{ gen.DO }

func (s sysReourceExtDo) Debug() *sysReourceExtDo {
	return s.withDO(s.DO.Debug())
}

func (s sysReourceExtDo) WithContext(ctx context.Context) *sysReourceExtDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysReourceExtDo) ReadDB() *sysReourceExtDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysReourceExtDo) WriteDB() *sysReourceExtDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysReourceExtDo) Clauses(conds ...clause.Expression) *sysReourceExtDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysReourceExtDo) Returning(value interface{}, columns ...string) *sysReourceExtDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysReourceExtDo) Not(conds ...gen.Condition) *sysReourceExtDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysReourceExtDo) Or(conds ...gen.Condition) *sysReourceExtDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysReourceExtDo) Select(conds ...field.Expr) *sysReourceExtDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysReourceExtDo) Where(conds ...gen.Condition) *sysReourceExtDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysReourceExtDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *sysReourceExtDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysReourceExtDo) Order(conds ...field.Expr) *sysReourceExtDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysReourceExtDo) Distinct(cols ...field.Expr) *sysReourceExtDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysReourceExtDo) Omit(cols ...field.Expr) *sysReourceExtDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysReourceExtDo) Join(table schema.Tabler, on ...field.Expr) *sysReourceExtDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysReourceExtDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sysReourceExtDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysReourceExtDo) RightJoin(table schema.Tabler, on ...field.Expr) *sysReourceExtDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysReourceExtDo) Group(cols ...field.Expr) *sysReourceExtDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysReourceExtDo) Having(conds ...gen.Condition) *sysReourceExtDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysReourceExtDo) Limit(limit int) *sysReourceExtDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysReourceExtDo) Offset(offset int) *sysReourceExtDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysReourceExtDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sysReourceExtDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysReourceExtDo) Unscoped() *sysReourceExtDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysReourceExtDo) Create(values ...*domain.SysReourceExt) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysReourceExtDo) CreateInBatches(values []*domain.SysReourceExt, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysReourceExtDo) Save(values ...*domain.SysReourceExt) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysReourceExtDo) First() (*domain.SysReourceExt, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*domain.SysReourceExt), nil
	}
}

func (s sysReourceExtDo) Take() (*domain.SysReourceExt, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*domain.SysReourceExt), nil
	}
}

func (s sysReourceExtDo) Last() (*domain.SysReourceExt, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*domain.SysReourceExt), nil
	}
}

func (s sysReourceExtDo) Find() ([]*domain.SysReourceExt, error) {
	result, err := s.DO.Find()
	return result.([]*domain.SysReourceExt), err
}

func (s sysReourceExtDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*domain.SysReourceExt, err error) {
	buf := make([]*domain.SysReourceExt, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysReourceExtDo) FindInBatches(result *[]*domain.SysReourceExt, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysReourceExtDo) Attrs(attrs ...field.AssignExpr) *sysReourceExtDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysReourceExtDo) Assign(attrs ...field.AssignExpr) *sysReourceExtDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysReourceExtDo) Joins(fields ...field.RelationField) *sysReourceExtDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysReourceExtDo) Preload(fields ...field.RelationField) *sysReourceExtDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysReourceExtDo) FirstOrInit() (*domain.SysReourceExt, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*domain.SysReourceExt), nil
	}
}

func (s sysReourceExtDo) FirstOrCreate() (*domain.SysReourceExt, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*domain.SysReourceExt), nil
	}
}

func (s sysReourceExtDo) FindByPage(offset int, limit int) (result []*domain.SysReourceExt, count int64, err error) {
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

func (s sysReourceExtDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysReourceExtDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysReourceExtDo) Delete(models ...*domain.SysReourceExt) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysReourceExtDo) withDO(do gen.Dao) *sysReourceExtDo {
	s.DO = *do.(*gen.DO)
	return s
}
