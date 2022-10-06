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

func newCustomerPoject(db *gorm.DB) customerPoject {
	_customerPoject := customerPoject{}

	_customerPoject.customerPojectDo.UseDB(db)
	_customerPoject.customerPojectDo.UseModel(&domain.CustomerPoject{})

	tableName := _customerPoject.customerPojectDo.TableName()
	_customerPoject.ALL = field.NewAsterisk(tableName)
	_customerPoject.ID = field.NewInt64(tableName, "Id")
	_customerPoject.CusID = field.NewInt64(tableName, "Cus_Id")
	_customerPoject.ProjectName = field.NewString(tableName, "Project_Name")
	_customerPoject.ProjectDesc = field.NewString(tableName, "Project_Desc")
	_customerPoject.ProjectMemo = field.NewString(tableName, "Project_Memo")

	_customerPoject.fillFieldMap()

	return _customerPoject
}

type customerPoject struct {
	customerPojectDo customerPojectDo

	ALL         field.Asterisk
	ID          field.Int64  // 项目Id
	CusID       field.Int64  // 客户Id
	ProjectName field.String // 项目名称
	ProjectDesc field.String // 项目描述
	ProjectMemo field.String // 说明

	fieldMap map[string]field.Expr
}

func (c customerPoject) Table(newTableName string) *customerPoject {
	c.customerPojectDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c customerPoject) As(alias string) *customerPoject {
	c.customerPojectDo.DO = *(c.customerPojectDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *customerPoject) updateTableName(table string) *customerPoject {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "Id")
	c.CusID = field.NewInt64(table, "Cus_Id")
	c.ProjectName = field.NewString(table, "Project_Name")
	c.ProjectDesc = field.NewString(table, "Project_Desc")
	c.ProjectMemo = field.NewString(table, "Project_Memo")

	c.fillFieldMap()

	return c
}

func (c *customerPoject) WithContext(ctx context.Context) *customerPojectDo {
	return c.customerPojectDo.WithContext(ctx)
}

func (c customerPoject) TableName() string { return c.customerPojectDo.TableName() }

func (c customerPoject) Alias() string { return c.customerPojectDo.Alias() }

func (c *customerPoject) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *customerPoject) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 5)
	c.fieldMap["Id"] = c.ID
	c.fieldMap["Cus_Id"] = c.CusID
	c.fieldMap["Project_Name"] = c.ProjectName
	c.fieldMap["Project_Desc"] = c.ProjectDesc
	c.fieldMap["Project_Memo"] = c.ProjectMemo
}

func (c customerPoject) clone(db *gorm.DB) customerPoject {
	c.customerPojectDo.ReplaceDB(db)
	return c
}

type customerPojectDo struct{ gen.DO }

func (c customerPojectDo) Debug() *customerPojectDo {
	return c.withDO(c.DO.Debug())
}

func (c customerPojectDo) WithContext(ctx context.Context) *customerPojectDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c customerPojectDo) ReadDB() *customerPojectDo {
	return c.Clauses(dbresolver.Read)
}

func (c customerPojectDo) WriteDB() *customerPojectDo {
	return c.Clauses(dbresolver.Write)
}

func (c customerPojectDo) Clauses(conds ...clause.Expression) *customerPojectDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c customerPojectDo) Returning(value interface{}, columns ...string) *customerPojectDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c customerPojectDo) Not(conds ...gen.Condition) *customerPojectDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c customerPojectDo) Or(conds ...gen.Condition) *customerPojectDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c customerPojectDo) Select(conds ...field.Expr) *customerPojectDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c customerPojectDo) Where(conds ...gen.Condition) *customerPojectDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c customerPojectDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *customerPojectDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c customerPojectDo) Order(conds ...field.Expr) *customerPojectDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c customerPojectDo) Distinct(cols ...field.Expr) *customerPojectDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c customerPojectDo) Omit(cols ...field.Expr) *customerPojectDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c customerPojectDo) Join(table schema.Tabler, on ...field.Expr) *customerPojectDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c customerPojectDo) LeftJoin(table schema.Tabler, on ...field.Expr) *customerPojectDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c customerPojectDo) RightJoin(table schema.Tabler, on ...field.Expr) *customerPojectDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c customerPojectDo) Group(cols ...field.Expr) *customerPojectDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c customerPojectDo) Having(conds ...gen.Condition) *customerPojectDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c customerPojectDo) Limit(limit int) *customerPojectDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c customerPojectDo) Offset(offset int) *customerPojectDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c customerPojectDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *customerPojectDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c customerPojectDo) Unscoped() *customerPojectDo {
	return c.withDO(c.DO.Unscoped())
}

func (c customerPojectDo) Create(values ...*domain.CustomerPoject) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c customerPojectDo) CreateInBatches(values []*domain.CustomerPoject, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c customerPojectDo) Save(values ...*domain.CustomerPoject) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c customerPojectDo) First() (*domain.CustomerPoject, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*domain.CustomerPoject), nil
	}
}

func (c customerPojectDo) Take() (*domain.CustomerPoject, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*domain.CustomerPoject), nil
	}
}

func (c customerPojectDo) Last() (*domain.CustomerPoject, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*domain.CustomerPoject), nil
	}
}

func (c customerPojectDo) Find() ([]*domain.CustomerPoject, error) {
	result, err := c.DO.Find()
	return result.([]*domain.CustomerPoject), err
}

func (c customerPojectDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*domain.CustomerPoject, err error) {
	buf := make([]*domain.CustomerPoject, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c customerPojectDo) FindInBatches(result *[]*domain.CustomerPoject, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c customerPojectDo) Attrs(attrs ...field.AssignExpr) *customerPojectDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c customerPojectDo) Assign(attrs ...field.AssignExpr) *customerPojectDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c customerPojectDo) Joins(fields ...field.RelationField) *customerPojectDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c customerPojectDo) Preload(fields ...field.RelationField) *customerPojectDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c customerPojectDo) FirstOrInit() (*domain.CustomerPoject, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*domain.CustomerPoject), nil
	}
}

func (c customerPojectDo) FirstOrCreate() (*domain.CustomerPoject, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*domain.CustomerPoject), nil
	}
}

func (c customerPojectDo) FindByPage(offset int, limit int) (result []*domain.CustomerPoject, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c customerPojectDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c customerPojectDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c customerPojectDo) Delete(models ...*domain.CustomerPoject) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *customerPojectDo) withDO(do gen.Dao) *customerPojectDo {
	c.DO = *do.(*gen.DO)
	return c
}
