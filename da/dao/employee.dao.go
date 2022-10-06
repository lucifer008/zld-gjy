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

func newEmployee(db *gorm.DB) employee {
	_employee := employee{}

	_employee.employeeDo.UseDB(db)
	_employee.employeeDo.UseModel(&domain.Employee{})

	tableName := _employee.employeeDo.TableName()
	_employee.ALL = field.NewAsterisk(tableName)
	_employee.ID = field.NewInt64(tableName, "Id")
	_employee.OrgID = field.NewInt64(tableName, "Org_Id")
	_employee.EmpNo = field.NewString(tableName, "Emp_No")
	_employee.EmpName = field.NewString(tableName, "Emp_Name")
	_employee.EmpMobile = field.NewString(tableName, "Emp_Mobile")
	_employee.InsertDateTime = field.NewTime(tableName, "Insert_Date_Time")
	_employee.InsertUser = field.NewInt64(tableName, "Insert_User")
	_employee.UpdateDateTime = field.NewTime(tableName, "Update_Date_Time")
	_employee.UpdateUser = field.NewInt64(tableName, "Update_User")
	_employee.Version = field.NewInt32(tableName, "Version")
	_employee.Deleted = field.NewString(tableName, "Deleted")

	_employee.fillFieldMap()

	return _employee
}

type employee struct {
	employeeDo employeeDo

	ALL            field.Asterisk
	ID             field.Int64  // 雇员Id
	OrgID          field.Int64  // Id
	EmpNo          field.String // 雇员编号
	EmpName        field.String // 雇员名称
	EmpMobile      field.String // 雇员手机号
	InsertDateTime field.Time   // 插入时间
	InsertUser     field.Int64  // 插入用户
	UpdateDateTime field.Time   // 更新时间
	UpdateUser     field.Int64  // 更新用户
	Version        field.Int32  // 版本
	Deleted        field.String // 删除标志(0:正常 1 删除)

	fieldMap map[string]field.Expr
}

func (e employee) Table(newTableName string) *employee {
	e.employeeDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e employee) As(alias string) *employee {
	e.employeeDo.DO = *(e.employeeDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *employee) updateTableName(table string) *employee {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt64(table, "Id")
	e.OrgID = field.NewInt64(table, "Org_Id")
	e.EmpNo = field.NewString(table, "Emp_No")
	e.EmpName = field.NewString(table, "Emp_Name")
	e.EmpMobile = field.NewString(table, "Emp_Mobile")
	e.InsertDateTime = field.NewTime(table, "Insert_Date_Time")
	e.InsertUser = field.NewInt64(table, "Insert_User")
	e.UpdateDateTime = field.NewTime(table, "Update_Date_Time")
	e.UpdateUser = field.NewInt64(table, "Update_User")
	e.Version = field.NewInt32(table, "Version")
	e.Deleted = field.NewString(table, "Deleted")

	e.fillFieldMap()

	return e
}

func (e *employee) WithContext(ctx context.Context) *employeeDo { return e.employeeDo.WithContext(ctx) }

func (e employee) TableName() string { return e.employeeDo.TableName() }

func (e employee) Alias() string { return e.employeeDo.Alias() }

func (e *employee) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *employee) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 11)
	e.fieldMap["Id"] = e.ID
	e.fieldMap["Org_Id"] = e.OrgID
	e.fieldMap["Emp_No"] = e.EmpNo
	e.fieldMap["Emp_Name"] = e.EmpName
	e.fieldMap["Emp_Mobile"] = e.EmpMobile
	e.fieldMap["Insert_Date_Time"] = e.InsertDateTime
	e.fieldMap["Insert_User"] = e.InsertUser
	e.fieldMap["Update_Date_Time"] = e.UpdateDateTime
	e.fieldMap["Update_User"] = e.UpdateUser
	e.fieldMap["Version"] = e.Version
	e.fieldMap["Deleted"] = e.Deleted
}

func (e employee) clone(db *gorm.DB) employee {
	e.employeeDo.ReplaceDB(db)
	return e
}

type employeeDo struct{ gen.DO }

func (e employeeDo) Debug() *employeeDo {
	return e.withDO(e.DO.Debug())
}

func (e employeeDo) WithContext(ctx context.Context) *employeeDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e employeeDo) ReadDB() *employeeDo {
	return e.Clauses(dbresolver.Read)
}

func (e employeeDo) WriteDB() *employeeDo {
	return e.Clauses(dbresolver.Write)
}

func (e employeeDo) Clauses(conds ...clause.Expression) *employeeDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e employeeDo) Returning(value interface{}, columns ...string) *employeeDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e employeeDo) Not(conds ...gen.Condition) *employeeDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e employeeDo) Or(conds ...gen.Condition) *employeeDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e employeeDo) Select(conds ...field.Expr) *employeeDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e employeeDo) Where(conds ...gen.Condition) *employeeDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e employeeDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *employeeDo {
	return e.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (e employeeDo) Order(conds ...field.Expr) *employeeDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e employeeDo) Distinct(cols ...field.Expr) *employeeDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e employeeDo) Omit(cols ...field.Expr) *employeeDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e employeeDo) Join(table schema.Tabler, on ...field.Expr) *employeeDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e employeeDo) LeftJoin(table schema.Tabler, on ...field.Expr) *employeeDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e employeeDo) RightJoin(table schema.Tabler, on ...field.Expr) *employeeDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e employeeDo) Group(cols ...field.Expr) *employeeDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e employeeDo) Having(conds ...gen.Condition) *employeeDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e employeeDo) Limit(limit int) *employeeDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e employeeDo) Offset(offset int) *employeeDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e employeeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *employeeDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e employeeDo) Unscoped() *employeeDo {
	return e.withDO(e.DO.Unscoped())
}

func (e employeeDo) Create(values ...*domain.Employee) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e employeeDo) CreateInBatches(values []*domain.Employee, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e employeeDo) Save(values ...*domain.Employee) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e employeeDo) First() (*domain.Employee, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Employee), nil
	}
}

func (e employeeDo) Take() (*domain.Employee, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Employee), nil
	}
}

func (e employeeDo) Last() (*domain.Employee, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Employee), nil
	}
}

func (e employeeDo) Find() ([]*domain.Employee, error) {
	result, err := e.DO.Find()
	return result.([]*domain.Employee), err
}

func (e employeeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*domain.Employee, err error) {
	buf := make([]*domain.Employee, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e employeeDo) FindInBatches(result *[]*domain.Employee, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e employeeDo) Attrs(attrs ...field.AssignExpr) *employeeDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e employeeDo) Assign(attrs ...field.AssignExpr) *employeeDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e employeeDo) Joins(fields ...field.RelationField) *employeeDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e employeeDo) Preload(fields ...field.RelationField) *employeeDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e employeeDo) FirstOrInit() (*domain.Employee, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Employee), nil
	}
}

func (e employeeDo) FirstOrCreate() (*domain.Employee, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Employee), nil
	}
}

func (e employeeDo) FindByPage(offset int, limit int) (result []*domain.Employee, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e employeeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e employeeDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e employeeDo) Delete(models ...*domain.Employee) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *employeeDo) withDO(do gen.Dao) *employeeDo {
	e.DO = *do.(*gen.DO)
	return e
}
