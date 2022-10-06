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

func newCustomer(db *gorm.DB) customer {
	_customer := customer{}

	_customer.customerDo.UseDB(db)
	_customer.customerDo.UseModel(&domain.Customer{})

	tableName := _customer.customerDo.TableName()
	_customer.ALL = field.NewAsterisk(tableName)
	_customer.ID = field.NewInt64(tableName, "Id")
	_customer.CustomerName = field.NewString(tableName, "Customer_Name")
	_customer.CustomerAddress = field.NewString(tableName, "Customer_Address")
	_customer.ContactMan = field.NewString(tableName, "Contact_Man")
	_customer.ContactTel = field.NewString(tableName, "Contact_Tel")
	_customer.EnterTel = field.NewString(tableName, "Enter_Tel")
	_customer.ProvinceCode = field.NewString(tableName, "Province_Code")
	_customer.CityCode = field.NewString(tableName, "City_Code")
	_customer.AreaCode = field.NewString(tableName, "Area_Code")
	_customer.Memo = field.NewString(tableName, "Memo")
	_customer.InsertDateTime = field.NewTime(tableName, "Insert_Date_Time")
	_customer.InsertUser = field.NewInt64(tableName, "Insert_User")
	_customer.UpdateDateTime = field.NewTime(tableName, "Update_Date_Time")
	_customer.UpdateUser = field.NewInt64(tableName, "Update_User")
	_customer.Version = field.NewInt32(tableName, "Version")
	_customer.Deleted = field.NewString(tableName, "Deleted")
	_customer.CompayID = field.NewInt64(tableName, "Compay_Id")

	_customer.fillFieldMap()

	return _customer
}

type customer struct {
	customerDo customerDo

	ALL             field.Asterisk
	ID              field.Int64  // 客户Id
	CustomerName    field.String // 客户名称
	CustomerAddress field.String // 客户地址
	ContactMan      field.String // 客户联系人
	ContactTel      field.String // 客户电话
	EnterTel        field.String // 企业联系人
	ProvinceCode    field.String // 省code
	CityCode        field.String // 市code
	AreaCode        field.String // 区code
	Memo            field.String // 备注
	InsertDateTime  field.Time   // 插入时间
	InsertUser      field.Int64  // 插入用户
	UpdateDateTime  field.Time   // 更新时间
	UpdateUser      field.Int64  // 更新用户
	Version         field.Int32  // 版本
	Deleted         field.String // 删除标志(0:正常 1 删除)
	CompayID        field.Int64  // 公司

	fieldMap map[string]field.Expr
}

func (c customer) Table(newTableName string) *customer {
	c.customerDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c customer) As(alias string) *customer {
	c.customerDo.DO = *(c.customerDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *customer) updateTableName(table string) *customer {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "Id")
	c.CustomerName = field.NewString(table, "Customer_Name")
	c.CustomerAddress = field.NewString(table, "Customer_Address")
	c.ContactMan = field.NewString(table, "Contact_Man")
	c.ContactTel = field.NewString(table, "Contact_Tel")
	c.EnterTel = field.NewString(table, "Enter_Tel")
	c.ProvinceCode = field.NewString(table, "Province_Code")
	c.CityCode = field.NewString(table, "City_Code")
	c.AreaCode = field.NewString(table, "Area_Code")
	c.Memo = field.NewString(table, "Memo")
	c.InsertDateTime = field.NewTime(table, "Insert_Date_Time")
	c.InsertUser = field.NewInt64(table, "Insert_User")
	c.UpdateDateTime = field.NewTime(table, "Update_Date_Time")
	c.UpdateUser = field.NewInt64(table, "Update_User")
	c.Version = field.NewInt32(table, "Version")
	c.Deleted = field.NewString(table, "Deleted")
	c.CompayID = field.NewInt64(table, "Compay_Id")

	c.fillFieldMap()

	return c
}

func (c *customer) WithContext(ctx context.Context) *customerDo { return c.customerDo.WithContext(ctx) }

func (c customer) TableName() string { return c.customerDo.TableName() }

func (c customer) Alias() string { return c.customerDo.Alias() }

func (c *customer) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *customer) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 17)
	c.fieldMap["Id"] = c.ID
	c.fieldMap["Customer_Name"] = c.CustomerName
	c.fieldMap["Customer_Address"] = c.CustomerAddress
	c.fieldMap["Contact_Man"] = c.ContactMan
	c.fieldMap["Contact_Tel"] = c.ContactTel
	c.fieldMap["Enter_Tel"] = c.EnterTel
	c.fieldMap["Province_Code"] = c.ProvinceCode
	c.fieldMap["City_Code"] = c.CityCode
	c.fieldMap["Area_Code"] = c.AreaCode
	c.fieldMap["Memo"] = c.Memo
	c.fieldMap["Insert_Date_Time"] = c.InsertDateTime
	c.fieldMap["Insert_User"] = c.InsertUser
	c.fieldMap["Update_Date_Time"] = c.UpdateDateTime
	c.fieldMap["Update_User"] = c.UpdateUser
	c.fieldMap["Version"] = c.Version
	c.fieldMap["Deleted"] = c.Deleted
	c.fieldMap["Compay_Id"] = c.CompayID
}

func (c customer) clone(db *gorm.DB) customer {
	c.customerDo.ReplaceDB(db)
	return c
}

type customerDo struct{ gen.DO }

func (c customerDo) Debug() *customerDo {
	return c.withDO(c.DO.Debug())
}

func (c customerDo) WithContext(ctx context.Context) *customerDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c customerDo) ReadDB() *customerDo {
	return c.Clauses(dbresolver.Read)
}

func (c customerDo) WriteDB() *customerDo {
	return c.Clauses(dbresolver.Write)
}

func (c customerDo) Clauses(conds ...clause.Expression) *customerDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c customerDo) Returning(value interface{}, columns ...string) *customerDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c customerDo) Not(conds ...gen.Condition) *customerDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c customerDo) Or(conds ...gen.Condition) *customerDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c customerDo) Select(conds ...field.Expr) *customerDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c customerDo) Where(conds ...gen.Condition) *customerDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c customerDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *customerDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c customerDo) Order(conds ...field.Expr) *customerDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c customerDo) Distinct(cols ...field.Expr) *customerDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c customerDo) Omit(cols ...field.Expr) *customerDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c customerDo) Join(table schema.Tabler, on ...field.Expr) *customerDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c customerDo) LeftJoin(table schema.Tabler, on ...field.Expr) *customerDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c customerDo) RightJoin(table schema.Tabler, on ...field.Expr) *customerDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c customerDo) Group(cols ...field.Expr) *customerDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c customerDo) Having(conds ...gen.Condition) *customerDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c customerDo) Limit(limit int) *customerDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c customerDo) Offset(offset int) *customerDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c customerDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *customerDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c customerDo) Unscoped() *customerDo {
	return c.withDO(c.DO.Unscoped())
}

func (c customerDo) Create(values ...*domain.Customer) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c customerDo) CreateInBatches(values []*domain.Customer, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c customerDo) Save(values ...*domain.Customer) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c customerDo) First() (*domain.Customer, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Customer), nil
	}
}

func (c customerDo) Take() (*domain.Customer, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Customer), nil
	}
}

func (c customerDo) Last() (*domain.Customer, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Customer), nil
	}
}

func (c customerDo) Find() ([]*domain.Customer, error) {
	result, err := c.DO.Find()
	return result.([]*domain.Customer), err
}

func (c customerDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*domain.Customer, err error) {
	buf := make([]*domain.Customer, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c customerDo) FindInBatches(result *[]*domain.Customer, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c customerDo) Attrs(attrs ...field.AssignExpr) *customerDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c customerDo) Assign(attrs ...field.AssignExpr) *customerDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c customerDo) Joins(fields ...field.RelationField) *customerDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c customerDo) Preload(fields ...field.RelationField) *customerDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c customerDo) FirstOrInit() (*domain.Customer, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Customer), nil
	}
}

func (c customerDo) FirstOrCreate() (*domain.Customer, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Customer), nil
	}
}

func (c customerDo) FindByPage(offset int, limit int) (result []*domain.Customer, count int64, err error) {
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

func (c customerDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c customerDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c customerDo) Delete(models ...*domain.Customer) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *customerDo) withDO(do gen.Dao) *customerDo {
	c.DO = *do.(*gen.DO)
	return c
}
