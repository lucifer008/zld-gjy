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

func newVehicle(db *gorm.DB) vehicle {
	_vehicle := vehicle{}

	_vehicle.vehicleDo.UseDB(db)
	_vehicle.vehicleDo.UseModel(&model.Vehicle{})

	tableName := _vehicle.vehicleDo.TableName()
	_vehicle.ALL = field.NewAsterisk(tableName)
	_vehicle.ID = field.NewInt64(tableName, "Id")
	_vehicle.LogID = field.NewInt64(tableName, "Log_Id")
	_vehicle.LicensePlate = field.NewString(tableName, "License_Plate")
	_vehicle.Model = field.NewString(tableName, "Model")
	_vehicle.Len = field.NewString(tableName, "Len")
	_vehicle.Memo = field.NewString(tableName, "Memo")

	_vehicle.fillFieldMap()

	return _vehicle
}

type vehicle struct {
	vehicleDo vehicleDo

	ALL          field.Asterisk
	ID           field.Int64  // 车辆Id
	LogID        field.Int64  // 物流公司Id
	LicensePlate field.String // 车牌
	Model        field.String // 车型(厢式/平板/高栏/爬梯)
	Len          field.String // 车长(4.2/6.8/9.6/13/13.75/17.5)
	Memo         field.String // 备注

	fieldMap map[string]field.Expr
}

func (v vehicle) Table(newTableName string) *vehicle {
	v.vehicleDo.UseTable(newTableName)
	return v.updateTableName(newTableName)
}

func (v vehicle) As(alias string) *vehicle {
	v.vehicleDo.DO = *(v.vehicleDo.As(alias).(*gen.DO))
	return v.updateTableName(alias)
}

func (v *vehicle) updateTableName(table string) *vehicle {
	v.ALL = field.NewAsterisk(table)
	v.ID = field.NewInt64(table, "Id")
	v.LogID = field.NewInt64(table, "Log_Id")
	v.LicensePlate = field.NewString(table, "License_Plate")
	v.Model = field.NewString(table, "Model")
	v.Len = field.NewString(table, "Len")
	v.Memo = field.NewString(table, "Memo")

	v.fillFieldMap()

	return v
}

func (v *vehicle) WithContext(ctx context.Context) *vehicleDo { return v.vehicleDo.WithContext(ctx) }

func (v vehicle) TableName() string { return v.vehicleDo.TableName() }

func (v vehicle) Alias() string { return v.vehicleDo.Alias() }

func (v *vehicle) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := v.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (v *vehicle) fillFieldMap() {
	v.fieldMap = make(map[string]field.Expr, 6)
	v.fieldMap["Id"] = v.ID
	v.fieldMap["Log_Id"] = v.LogID
	v.fieldMap["License_Plate"] = v.LicensePlate
	v.fieldMap["Model"] = v.Model
	v.fieldMap["Len"] = v.Len
	v.fieldMap["Memo"] = v.Memo
}

func (v vehicle) clone(db *gorm.DB) vehicle {
	v.vehicleDo.ReplaceDB(db)
	return v
}

type vehicleDo struct{ gen.DO }

func (v vehicleDo) Debug() *vehicleDo {
	return v.withDO(v.DO.Debug())
}

func (v vehicleDo) WithContext(ctx context.Context) *vehicleDo {
	return v.withDO(v.DO.WithContext(ctx))
}

func (v vehicleDo) ReadDB() *vehicleDo {
	return v.Clauses(dbresolver.Read)
}

func (v vehicleDo) WriteDB() *vehicleDo {
	return v.Clauses(dbresolver.Write)
}

func (v vehicleDo) Clauses(conds ...clause.Expression) *vehicleDo {
	return v.withDO(v.DO.Clauses(conds...))
}

func (v vehicleDo) Returning(value interface{}, columns ...string) *vehicleDo {
	return v.withDO(v.DO.Returning(value, columns...))
}

func (v vehicleDo) Not(conds ...gen.Condition) *vehicleDo {
	return v.withDO(v.DO.Not(conds...))
}

func (v vehicleDo) Or(conds ...gen.Condition) *vehicleDo {
	return v.withDO(v.DO.Or(conds...))
}

func (v vehicleDo) Select(conds ...field.Expr) *vehicleDo {
	return v.withDO(v.DO.Select(conds...))
}

func (v vehicleDo) Where(conds ...gen.Condition) *vehicleDo {
	return v.withDO(v.DO.Where(conds...))
}

func (v vehicleDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *vehicleDo {
	return v.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (v vehicleDo) Order(conds ...field.Expr) *vehicleDo {
	return v.withDO(v.DO.Order(conds...))
}

func (v vehicleDo) Distinct(cols ...field.Expr) *vehicleDo {
	return v.withDO(v.DO.Distinct(cols...))
}

func (v vehicleDo) Omit(cols ...field.Expr) *vehicleDo {
	return v.withDO(v.DO.Omit(cols...))
}

func (v vehicleDo) Join(table schema.Tabler, on ...field.Expr) *vehicleDo {
	return v.withDO(v.DO.Join(table, on...))
}

func (v vehicleDo) LeftJoin(table schema.Tabler, on ...field.Expr) *vehicleDo {
	return v.withDO(v.DO.LeftJoin(table, on...))
}

func (v vehicleDo) RightJoin(table schema.Tabler, on ...field.Expr) *vehicleDo {
	return v.withDO(v.DO.RightJoin(table, on...))
}

func (v vehicleDo) Group(cols ...field.Expr) *vehicleDo {
	return v.withDO(v.DO.Group(cols...))
}

func (v vehicleDo) Having(conds ...gen.Condition) *vehicleDo {
	return v.withDO(v.DO.Having(conds...))
}

func (v vehicleDo) Limit(limit int) *vehicleDo {
	return v.withDO(v.DO.Limit(limit))
}

func (v vehicleDo) Offset(offset int) *vehicleDo {
	return v.withDO(v.DO.Offset(offset))
}

func (v vehicleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *vehicleDo {
	return v.withDO(v.DO.Scopes(funcs...))
}

func (v vehicleDo) Unscoped() *vehicleDo {
	return v.withDO(v.DO.Unscoped())
}

func (v vehicleDo) Create(values ...*model.Vehicle) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Create(values)
}

func (v vehicleDo) CreateInBatches(values []*model.Vehicle, batchSize int) error {
	return v.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (v vehicleDo) Save(values ...*model.Vehicle) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Save(values)
}

func (v vehicleDo) First() (*model.Vehicle, error) {
	if result, err := v.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Vehicle), nil
	}
}

func (v vehicleDo) Take() (*model.Vehicle, error) {
	if result, err := v.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Vehicle), nil
	}
}

func (v vehicleDo) Last() (*model.Vehicle, error) {
	if result, err := v.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Vehicle), nil
	}
}

func (v vehicleDo) Find() ([]*model.Vehicle, error) {
	result, err := v.DO.Find()
	return result.([]*model.Vehicle), err
}

func (v vehicleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Vehicle, err error) {
	buf := make([]*model.Vehicle, 0, batchSize)
	err = v.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (v vehicleDo) FindInBatches(result *[]*model.Vehicle, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return v.DO.FindInBatches(result, batchSize, fc)
}

func (v vehicleDo) Attrs(attrs ...field.AssignExpr) *vehicleDo {
	return v.withDO(v.DO.Attrs(attrs...))
}

func (v vehicleDo) Assign(attrs ...field.AssignExpr) *vehicleDo {
	return v.withDO(v.DO.Assign(attrs...))
}

func (v vehicleDo) Joins(fields ...field.RelationField) *vehicleDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Joins(_f))
	}
	return &v
}

func (v vehicleDo) Preload(fields ...field.RelationField) *vehicleDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Preload(_f))
	}
	return &v
}

func (v vehicleDo) FirstOrInit() (*model.Vehicle, error) {
	if result, err := v.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Vehicle), nil
	}
}

func (v vehicleDo) FirstOrCreate() (*model.Vehicle, error) {
	if result, err := v.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Vehicle), nil
	}
}

func (v vehicleDo) FindByPage(offset int, limit int) (result []*model.Vehicle, count int64, err error) {
	result, err = v.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = v.Offset(-1).Limit(-1).Count()
	return
}

func (v vehicleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = v.Count()
	if err != nil {
		return
	}

	err = v.Offset(offset).Limit(limit).Scan(result)
	return
}

func (v vehicleDo) Scan(result interface{}) (err error) {
	return v.DO.Scan(result)
}

func (v vehicleDo) Delete(models ...*model.Vehicle) (result gen.ResultInfo, err error) {
	return v.DO.Delete(models)
}

func (v *vehicleDo) withDO(do gen.Dao) *vehicleDo {
	v.DO = *do.(*gen.DO)
	return v
}
