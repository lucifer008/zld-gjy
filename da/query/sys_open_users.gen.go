// Code generated by "zld-jy/da/field". DO NOT EDIT.
// Code generated by "zld-jy/da/field". DO NOT EDIT.
// Code generated by "zld-jy/da/field". DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"zld-jy/da/field"

	"gorm.io/plugin/dbresolver"

	"zld-jy/da/model"
)

func newSysOpenUser(db *gorm.DB) sysOpenUser {
	_sysOpenUser := sysOpenUser{}

	_sysOpenUser.sysOpenUserDo.UseDB(db)
	_sysOpenUser.sysOpenUserDo.UseModel(&model.SysOpenUser{})

	tableName := _sysOpenUser.sysOpenUserDo.TableName()
	_sysOpenUser.ALL = field.NewAsterisk(tableName)
	_sysOpenUser.Id13 = field.NewInt64(tableName, "Id13")
	_sysOpenUser.OpenIdentity = field.NewString(tableName, "Open_Identity")
	_sysOpenUser.OpenName = field.NewString(tableName, "Open_Name")
	_sysOpenUser.OpenType = field.NewInt32(tableName, "Open_Type")
	_sysOpenUser.OpenMemo1 = field.NewString(tableName, "Open_Memo1")

	_sysOpenUser.fillFieldMap()

	return _sysOpenUser
}

type sysOpenUser struct {
	sysOpenUserDo sysOpenUserDo

	ALL          field.Asterisk
	Id13         field.Int64  // Id
	OpenIdentity field.String // 平台标示
	OpenName     field.String // 平台名称
	OpenType     field.Int32  // 开发平台类型(1:钉钉 2: 微信 )
	OpenMemo1    field.String // 平台扩展

	fieldMap map[string]field.Expr
}

func (s sysOpenUser) Table(newTableName string) *sysOpenUser {
	s.sysOpenUserDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysOpenUser) As(alias string) *sysOpenUser {
	s.sysOpenUserDo.DO = *(s.sysOpenUserDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysOpenUser) updateTableName(table string) *sysOpenUser {
	s.ALL = field.NewAsterisk(table)
	s.Id13 = field.NewInt64(table, "Id13")
	s.OpenIdentity = field.NewString(table, "Open_Identity")
	s.OpenName = field.NewString(table, "Open_Name")
	s.OpenType = field.NewInt32(table, "Open_Type")
	s.OpenMemo1 = field.NewString(table, "Open_Memo1")

	s.fillFieldMap()

	return s
}

func (s *sysOpenUser) WithContext(ctx context.Context) *sysOpenUserDo {
	return s.sysOpenUserDo.WithContext(ctx)
}

func (s sysOpenUser) TableName() string { return s.sysOpenUserDo.TableName() }

func (s sysOpenUser) Alias() string { return s.sysOpenUserDo.Alias() }

func (s *sysOpenUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysOpenUser) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 5)
	s.fieldMap["Id13"] = s.Id13
	s.fieldMap["Open_Identity"] = s.OpenIdentity
	s.fieldMap["Open_Name"] = s.OpenName
	s.fieldMap["Open_Type"] = s.OpenType
	s.fieldMap["Open_Memo1"] = s.OpenMemo1
}

func (s sysOpenUser) clone(db *gorm.DB) sysOpenUser {
	s.sysOpenUserDo.ReplaceDB(db)
	return s
}

type sysOpenUserDo struct{ gen.DO }

func (s sysOpenUserDo) Debug() *sysOpenUserDo {
	return s.withDO(s.DO.Debug())
}

func (s sysOpenUserDo) WithContext(ctx context.Context) *sysOpenUserDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysOpenUserDo) ReadDB() *sysOpenUserDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysOpenUserDo) WriteDB() *sysOpenUserDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysOpenUserDo) Clauses(conds ...clause.Expression) *sysOpenUserDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysOpenUserDo) Returning(value interface{}, columns ...string) *sysOpenUserDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysOpenUserDo) Not(conds ...gen.Condition) *sysOpenUserDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysOpenUserDo) Or(conds ...gen.Condition) *sysOpenUserDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysOpenUserDo) Select(conds ...field.Expr) *sysOpenUserDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysOpenUserDo) Where(conds ...gen.Condition) *sysOpenUserDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysOpenUserDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *sysOpenUserDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysOpenUserDo) Order(conds ...field.Expr) *sysOpenUserDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysOpenUserDo) Distinct(cols ...field.Expr) *sysOpenUserDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysOpenUserDo) Omit(cols ...field.Expr) *sysOpenUserDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysOpenUserDo) Join(table schema.Tabler, on ...field.Expr) *sysOpenUserDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysOpenUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sysOpenUserDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysOpenUserDo) RightJoin(table schema.Tabler, on ...field.Expr) *sysOpenUserDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysOpenUserDo) Group(cols ...field.Expr) *sysOpenUserDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysOpenUserDo) Having(conds ...gen.Condition) *sysOpenUserDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysOpenUserDo) Limit(limit int) *sysOpenUserDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysOpenUserDo) Offset(offset int) *sysOpenUserDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysOpenUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sysOpenUserDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysOpenUserDo) Unscoped() *sysOpenUserDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysOpenUserDo) Create(values ...*model.SysOpenUser) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysOpenUserDo) CreateInBatches(values []*model.SysOpenUser, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysOpenUserDo) Save(values ...*model.SysOpenUser) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysOpenUserDo) First() (*model.SysOpenUser, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOpenUser), nil
	}
}

func (s sysOpenUserDo) Take() (*model.SysOpenUser, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOpenUser), nil
	}
}

func (s sysOpenUserDo) Last() (*model.SysOpenUser, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOpenUser), nil
	}
}

func (s sysOpenUserDo) Find() ([]*model.SysOpenUser, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysOpenUser), err
}

func (s sysOpenUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysOpenUser, err error) {
	buf := make([]*model.SysOpenUser, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysOpenUserDo) FindInBatches(result *[]*model.SysOpenUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysOpenUserDo) Attrs(attrs ...field.AssignExpr) *sysOpenUserDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysOpenUserDo) Assign(attrs ...field.AssignExpr) *sysOpenUserDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysOpenUserDo) Joins(fields ...field.RelationField) *sysOpenUserDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysOpenUserDo) Preload(fields ...field.RelationField) *sysOpenUserDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysOpenUserDo) FirstOrInit() (*model.SysOpenUser, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOpenUser), nil
	}
}

func (s sysOpenUserDo) FirstOrCreate() (*model.SysOpenUser, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOpenUser), nil
	}
}

func (s sysOpenUserDo) FindByPage(offset int, limit int) (result []*model.SysOpenUser, count int64, err error) {
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

func (s sysOpenUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysOpenUserDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysOpenUserDo) Delete(models ...*model.SysOpenUser) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysOpenUserDo) withDO(do gen.Dao) *sysOpenUserDo {
	s.DO = *do.(*gen.DO)
	return s
}
