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

func newOrganization(db *gorm.DB) organization {
	_organization := organization{}

	_organization.organizationDo.UseDB(db)
	_organization.organizationDo.UseModel(&domain.Organization{})

	tableName := _organization.organizationDo.TableName()
	_organization.ALL = field.NewAsterisk(tableName)
	_organization.ID = field.NewInt64(tableName, "Id")
	_organization.ComID = field.NewInt64(tableName, "Com_Id")
	_organization.OrgID = field.NewInt64(tableName, "Org_Id")
	_organization.OrgName = field.NewString(tableName, "Org_Name")
	_organization.OrgCode = field.NewString(tableName, "Org_Code")
	_organization.OrgLevel = field.NewInt32(tableName, "Org_Level")
	_organization.BeginDate = field.NewTime(tableName, "Begin_Date")
	_organization.EndDate = field.NewTime(tableName, "End_Date")
	_organization.InsertDateTime = field.NewTime(tableName, "Insert_Date_Time")
	_organization.InsertUser = field.NewInt64(tableName, "Insert_User")
	_organization.UpdateDateTime = field.NewTime(tableName, "Update_Date_Time")
	_organization.UpdateUser = field.NewInt64(tableName, "Update_User")
	_organization.Version = field.NewInt32(tableName, "Version")
	_organization.Deleted = field.NewString(tableName, "Deleted")

	_organization.fillFieldMap()

	return _organization
}

type organization struct {
	organizationDo organizationDo

	ALL            field.Asterisk
	ID             field.Int64  // Id
	ComID          field.Int64  // Id
	OrgID          field.Int64  // Id
	OrgName        field.String // 组织名称
	OrgCode        field.String // 组织代码
	OrgLevel       field.Int32  // 组织级别
	BeginDate      field.Time   // 开始日期
	EndDate        field.Time   // 结束日期
	InsertDateTime field.Time   // 插入时间
	InsertUser     field.Int64  // 插入用户
	UpdateDateTime field.Time   // 更新时间
	UpdateUser     field.Int64  // 更新用户
	Version        field.Int32  // 版本
	Deleted        field.String // 删除标志(0:正常 1 删除)

	fieldMap map[string]field.Expr
}

func (o organization) Table(newTableName string) *organization {
	o.organizationDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o organization) As(alias string) *organization {
	o.organizationDo.DO = *(o.organizationDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *organization) updateTableName(table string) *organization {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewInt64(table, "Id")
	o.ComID = field.NewInt64(table, "Com_Id")
	o.OrgID = field.NewInt64(table, "Org_Id")
	o.OrgName = field.NewString(table, "Org_Name")
	o.OrgCode = field.NewString(table, "Org_Code")
	o.OrgLevel = field.NewInt32(table, "Org_Level")
	o.BeginDate = field.NewTime(table, "Begin_Date")
	o.EndDate = field.NewTime(table, "End_Date")
	o.InsertDateTime = field.NewTime(table, "Insert_Date_Time")
	o.InsertUser = field.NewInt64(table, "Insert_User")
	o.UpdateDateTime = field.NewTime(table, "Update_Date_Time")
	o.UpdateUser = field.NewInt64(table, "Update_User")
	o.Version = field.NewInt32(table, "Version")
	o.Deleted = field.NewString(table, "Deleted")

	o.fillFieldMap()

	return o
}

func (o *organization) WithContext(ctx context.Context) *organizationDo {
	return o.organizationDo.WithContext(ctx)
}

func (o organization) TableName() string { return o.organizationDo.TableName() }

func (o organization) Alias() string { return o.organizationDo.Alias() }

func (o *organization) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *organization) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 14)
	o.fieldMap["Id"] = o.ID
	o.fieldMap["Com_Id"] = o.ComID
	o.fieldMap["Org_Id"] = o.OrgID
	o.fieldMap["Org_Name"] = o.OrgName
	o.fieldMap["Org_Code"] = o.OrgCode
	o.fieldMap["Org_Level"] = o.OrgLevel
	o.fieldMap["Begin_Date"] = o.BeginDate
	o.fieldMap["End_Date"] = o.EndDate
	o.fieldMap["Insert_Date_Time"] = o.InsertDateTime
	o.fieldMap["Insert_User"] = o.InsertUser
	o.fieldMap["Update_Date_Time"] = o.UpdateDateTime
	o.fieldMap["Update_User"] = o.UpdateUser
	o.fieldMap["Version"] = o.Version
	o.fieldMap["Deleted"] = o.Deleted
}

func (o organization) clone(db *gorm.DB) organization {
	o.organizationDo.ReplaceDB(db)
	return o
}

type organizationDo struct{ gen.DO }

func (o organizationDo) Debug() *organizationDo {
	return o.withDO(o.DO.Debug())
}

func (o organizationDo) WithContext(ctx context.Context) *organizationDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o organizationDo) ReadDB() *organizationDo {
	return o.Clauses(dbresolver.Read)
}

func (o organizationDo) WriteDB() *organizationDo {
	return o.Clauses(dbresolver.Write)
}

func (o organizationDo) Clauses(conds ...clause.Expression) *organizationDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o organizationDo) Returning(value interface{}, columns ...string) *organizationDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o organizationDo) Not(conds ...gen.Condition) *organizationDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o organizationDo) Or(conds ...gen.Condition) *organizationDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o organizationDo) Select(conds ...field.Expr) *organizationDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o organizationDo) Where(conds ...gen.Condition) *organizationDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o organizationDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *organizationDo {
	return o.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (o organizationDo) Order(conds ...field.Expr) *organizationDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o organizationDo) Distinct(cols ...field.Expr) *organizationDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o organizationDo) Omit(cols ...field.Expr) *organizationDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o organizationDo) Join(table schema.Tabler, on ...field.Expr) *organizationDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o organizationDo) LeftJoin(table schema.Tabler, on ...field.Expr) *organizationDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o organizationDo) RightJoin(table schema.Tabler, on ...field.Expr) *organizationDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o organizationDo) Group(cols ...field.Expr) *organizationDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o organizationDo) Having(conds ...gen.Condition) *organizationDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o organizationDo) Limit(limit int) *organizationDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o organizationDo) Offset(offset int) *organizationDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o organizationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *organizationDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o organizationDo) Unscoped() *organizationDo {
	return o.withDO(o.DO.Unscoped())
}

func (o organizationDo) Create(values ...*domain.Organization) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o organizationDo) CreateInBatches(values []*domain.Organization, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o organizationDo) Save(values ...*domain.Organization) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o organizationDo) First() (*domain.Organization, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Organization), nil
	}
}

func (o organizationDo) Take() (*domain.Organization, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Organization), nil
	}
}

func (o organizationDo) Last() (*domain.Organization, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Organization), nil
	}
}

func (o organizationDo) Find() ([]*domain.Organization, error) {
	result, err := o.DO.Find()
	return result.([]*domain.Organization), err
}

func (o organizationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*domain.Organization, err error) {
	buf := make([]*domain.Organization, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o organizationDo) FindInBatches(result *[]*domain.Organization, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o organizationDo) Attrs(attrs ...field.AssignExpr) *organizationDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o organizationDo) Assign(attrs ...field.AssignExpr) *organizationDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o organizationDo) Joins(fields ...field.RelationField) *organizationDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o organizationDo) Preload(fields ...field.RelationField) *organizationDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o organizationDo) FirstOrInit() (*domain.Organization, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Organization), nil
	}
}

func (o organizationDo) FirstOrCreate() (*domain.Organization, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Organization), nil
	}
}

func (o organizationDo) FindByPage(offset int, limit int) (result []*domain.Organization, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o organizationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o organizationDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o organizationDo) Delete(models ...*domain.Organization) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *organizationDo) withDO(do gen.Dao) *organizationDo {
	o.DO = *do.(*gen.DO)
	return o
}
