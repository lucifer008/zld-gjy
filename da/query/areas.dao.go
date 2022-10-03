// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"zld-jy/da/model"
)

func newArea(db *gorm.DB) area {
	_area := area{}

	_area.areaDo.UseDB(db)
	_area.areaDo.UseModel(&model.Area{})

	tableName := _area.areaDo.TableName()
	_area.ALL = field.NewAsterisk(tableName)
	_area.ID = field.NewInt64(tableName, "id")
	_area.Code = field.NewString(tableName, "code")
	_area.Name = field.NewString(tableName, "name")
	_area.CityCode = field.NewString(tableName, "city_code")
	_area.Sort = field.NewInt32(tableName, "sort")
	_area.ShortName = field.NewString(tableName, "short_name")
	_area.Status = field.NewInt32(tableName, "status")
	_area.ShowShortName = field.NewString(tableName, "show_short_name")
	_area.Custom = field.NewInt32(tableName, "custom")

	_area.fillFieldMap()

	return _area
}

type area struct {
	areaDo areaDo

	ALL           field.Asterisk
	ID            field.Int64  // 区域Id
	Code          field.String // 区域code
	Name          field.String // 区域名称
	CityCode      field.String // 城市Code
	Sort          field.Int32  // 排序索引
	ShortName     field.String // 简称
	Status        field.Int32  // 状态(0:启用 1 停用)
	ShowShortName field.String // 显示简称名称
	Custom        field.Int32  // 是否自定义(0:否 1 是)

	fieldMap map[string]field.Expr
}

func (a area) Table(newTableName string) *area {
	a.areaDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a area) As(alias string) *area {
	a.areaDo.DO = *(a.areaDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *area) updateTableName(table string) *area {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.Code = field.NewString(table, "code")
	a.Name = field.NewString(table, "name")
	a.CityCode = field.NewString(table, "city_code")
	a.Sort = field.NewInt32(table, "sort")
	a.ShortName = field.NewString(table, "short_name")
	a.Status = field.NewInt32(table, "status")
	a.ShowShortName = field.NewString(table, "show_short_name")
	a.Custom = field.NewInt32(table, "custom")

	a.fillFieldMap()

	return a
}

func (a *area) WithContext(ctx context.Context) *areaDo { return a.areaDo.WithContext(ctx) }

func (a area) TableName() string { return a.areaDo.TableName() }

func (a area) Alias() string { return a.areaDo.Alias() }

func (a *area) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *area) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 9)
	a.fieldMap["id"] = a.ID
	a.fieldMap["code"] = a.Code
	a.fieldMap["name"] = a.Name
	a.fieldMap["city_code"] = a.CityCode
	a.fieldMap["sort"] = a.Sort
	a.fieldMap["short_name"] = a.ShortName
	a.fieldMap["status"] = a.Status
	a.fieldMap["show_short_name"] = a.ShowShortName
	a.fieldMap["custom"] = a.Custom
}

func (a area) clone(db *gorm.DB) area {
	a.areaDo.ReplaceDB(db)
	return a
}

type areaDo struct{ gen.DO }

func (a areaDo) Debug() *areaDo {
	return a.withDO(a.DO.Debug())
}

func (a areaDo) WithContext(ctx context.Context) *areaDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a areaDo) ReadDB() *areaDo {
	return a.Clauses(dbresolver.Read)
}

func (a areaDo) WriteDB() *areaDo {
	return a.Clauses(dbresolver.Write)
}

func (a areaDo) Clauses(conds ...clause.Expression) *areaDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a areaDo) Returning(value interface{}, columns ...string) *areaDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a areaDo) Not(conds ...gen.Condition) *areaDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a areaDo) Or(conds ...gen.Condition) *areaDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a areaDo) Select(conds ...field.Expr) *areaDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a areaDo) Where(conds ...gen.Condition) *areaDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a areaDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *areaDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a areaDo) Order(conds ...field.Expr) *areaDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a areaDo) Distinct(cols ...field.Expr) *areaDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a areaDo) Omit(cols ...field.Expr) *areaDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a areaDo) Join(table schema.Tabler, on ...field.Expr) *areaDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a areaDo) LeftJoin(table schema.Tabler, on ...field.Expr) *areaDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a areaDo) RightJoin(table schema.Tabler, on ...field.Expr) *areaDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a areaDo) Group(cols ...field.Expr) *areaDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a areaDo) Having(conds ...gen.Condition) *areaDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a areaDo) Limit(limit int) *areaDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a areaDo) Offset(offset int) *areaDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a areaDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *areaDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a areaDo) Unscoped() *areaDo {
	return a.withDO(a.DO.Unscoped())
}

func (a areaDo) Create(values ...*model.Area) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a areaDo) CreateInBatches(values []*model.Area, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a areaDo) Save(values ...*model.Area) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a areaDo) First() (*model.Area, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Area), nil
	}
}

func (a areaDo) Take() (*model.Area, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Area), nil
	}
}

func (a areaDo) Last() (*model.Area, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Area), nil
	}
}

func (a areaDo) Find() ([]*model.Area, error) {
	result, err := a.DO.Find()
	return result.([]*model.Area), err
}

func (a areaDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Area, err error) {
	buf := make([]*model.Area, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a areaDo) FindInBatches(result *[]*model.Area, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a areaDo) Attrs(attrs ...field.AssignExpr) *areaDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a areaDo) Assign(attrs ...field.AssignExpr) *areaDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a areaDo) Joins(fields ...field.RelationField) *areaDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a areaDo) Preload(fields ...field.RelationField) *areaDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a areaDo) FirstOrInit() (*model.Area, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Area), nil
	}
}

func (a areaDo) FirstOrCreate() (*model.Area, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Area), nil
	}
}

func (a areaDo) FindByPage(offset int, limit int) (result []*model.Area, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a areaDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a areaDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a areaDo) Delete(models ...*model.Area) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *areaDo) withDO(do gen.Dao) *areaDo {
	a.DO = *do.(*gen.DO)
	return a
}
