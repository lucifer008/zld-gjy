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

func newOrder(db *gorm.DB) order {
	_order := order{}

	_order.orderDo.UseDB(db)
	_order.orderDo.UseModel(&domain.Order{})

	tableName := _order.orderDo.TableName()
	_order.ALL = field.NewAsterisk(tableName)
	_order.ID = field.NewInt64(tableName, "id")
	_order.LogID = field.NewInt64(tableName, "Log_Id")
	_order.OrderNo = field.NewString(tableName, "order_no")
	_order.OutOrderNo = field.NewString(tableName, "out_order_no")
	_order.ProjectName = field.NewString(tableName, "project_name")
	_order.ShipmentDate = field.NewTime(tableName, "shipment_date")
	_order.Status = field.NewInt32(tableName, "status")
	_order.DeliverProvinceCode = field.NewString(tableName, "deliver_province_code")
	_order.DeliverCityCode = field.NewString(tableName, "deliver_city_code")
	_order.DeliverAreaCode = field.NewString(tableName, "deliver_area_code")
	_order.DeliverAddress = field.NewString(tableName, "deliver_address")
	_order.ArriveProvinceCode = field.NewString(tableName, "arrive_province_code")
	_order.ArriveCityCode = field.NewString(tableName, "arrive_city_code")
	_order.ArriveAreaCode = field.NewString(tableName, "arrive_area_code")
	_order.ArriveAddress = field.NewString(tableName, "arrive_address")
	_order.ArriveDate = field.NewTime(tableName, "arrive_date")
	_order.ReceiptCustomer = field.NewString(tableName, "receipt_customer")
	_order.ReceiptCustomerTel = field.NewString(tableName, "receipt_customer_tel")
	_order.PricingMethod = field.NewInt32(tableName, "pricing_method")
	_order.ReturnBillAttachment = field.NewString(tableName, "return_bill_attachment")

	_order.fillFieldMap()

	return _order
}

type order struct {
	orderDo orderDo

	ALL                  field.Asterisk
	ID                   field.Int64  // 订单Id
	LogID                field.Int64  // 物流公司Id
	OrderNo              field.String // 订单号
	OutOrderNo           field.String // 外部订单号
	ProjectName          field.String // 项目名称
	ShipmentDate         field.Time   // 托运时间
	Status               field.Int32  // 订单状态(0:草稿 1:已下单 2:已接单 3: 已装货 4: 已签收)
	DeliverProvinceCode  field.String // 发货省code
	DeliverCityCode      field.String // 发货城市Code
	DeliverAreaCode      field.String // 发货地区Code
	DeliverAddress       field.String // 出发地
	ArriveProvinceCode   field.String // 到达省份Code
	ArriveCityCode       field.String // 到达城市Code
	ArriveAreaCode       field.String // 到达地区Code
	ArriveAddress        field.String // 到达地
	ArriveDate           field.Time   // 到达时间
	ReceiptCustomer      field.String // 收货客户
	ReceiptCustomerTel   field.String // 收货客户电话
	PricingMethod        field.Int32  // 计价方式(0: 趟 1 吨 2 方)
	ReturnBillAttachment field.String // 回单附件

	fieldMap map[string]field.Expr
}

func (o order) Table(newTableName string) *order {
	o.orderDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o order) As(alias string) *order {
	o.orderDo.DO = *(o.orderDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *order) updateTableName(table string) *order {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewInt64(table, "id")
	o.LogID = field.NewInt64(table, "Log_Id")
	o.OrderNo = field.NewString(table, "order_no")
	o.OutOrderNo = field.NewString(table, "out_order_no")
	o.ProjectName = field.NewString(table, "project_name")
	o.ShipmentDate = field.NewTime(table, "shipment_date")
	o.Status = field.NewInt32(table, "status")
	o.DeliverProvinceCode = field.NewString(table, "deliver_province_code")
	o.DeliverCityCode = field.NewString(table, "deliver_city_code")
	o.DeliverAreaCode = field.NewString(table, "deliver_area_code")
	o.DeliverAddress = field.NewString(table, "deliver_address")
	o.ArriveProvinceCode = field.NewString(table, "arrive_province_code")
	o.ArriveCityCode = field.NewString(table, "arrive_city_code")
	o.ArriveAreaCode = field.NewString(table, "arrive_area_code")
	o.ArriveAddress = field.NewString(table, "arrive_address")
	o.ArriveDate = field.NewTime(table, "arrive_date")
	o.ReceiptCustomer = field.NewString(table, "receipt_customer")
	o.ReceiptCustomerTel = field.NewString(table, "receipt_customer_tel")
	o.PricingMethod = field.NewInt32(table, "pricing_method")
	o.ReturnBillAttachment = field.NewString(table, "return_bill_attachment")

	o.fillFieldMap()

	return o
}

func (o *order) WithContext(ctx context.Context) *orderDo { return o.orderDo.WithContext(ctx) }

func (o order) TableName() string { return o.orderDo.TableName() }

func (o order) Alias() string { return o.orderDo.Alias() }

func (o *order) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *order) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 20)
	o.fieldMap["id"] = o.ID
	o.fieldMap["Log_Id"] = o.LogID
	o.fieldMap["order_no"] = o.OrderNo
	o.fieldMap["out_order_no"] = o.OutOrderNo
	o.fieldMap["project_name"] = o.ProjectName
	o.fieldMap["shipment_date"] = o.ShipmentDate
	o.fieldMap["status"] = o.Status
	o.fieldMap["deliver_province_code"] = o.DeliverProvinceCode
	o.fieldMap["deliver_city_code"] = o.DeliverCityCode
	o.fieldMap["deliver_area_code"] = o.DeliverAreaCode
	o.fieldMap["deliver_address"] = o.DeliverAddress
	o.fieldMap["arrive_province_code"] = o.ArriveProvinceCode
	o.fieldMap["arrive_city_code"] = o.ArriveCityCode
	o.fieldMap["arrive_area_code"] = o.ArriveAreaCode
	o.fieldMap["arrive_address"] = o.ArriveAddress
	o.fieldMap["arrive_date"] = o.ArriveDate
	o.fieldMap["receipt_customer"] = o.ReceiptCustomer
	o.fieldMap["receipt_customer_tel"] = o.ReceiptCustomerTel
	o.fieldMap["pricing_method"] = o.PricingMethod
	o.fieldMap["return_bill_attachment"] = o.ReturnBillAttachment
}

func (o order) clone(db *gorm.DB) order {
	o.orderDo.ReplaceDB(db)
	return o
}

type orderDo struct{ gen.DO }

func (o orderDo) Debug() *orderDo {
	return o.withDO(o.DO.Debug())
}

func (o orderDo) WithContext(ctx context.Context) *orderDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o orderDo) ReadDB() *orderDo {
	return o.Clauses(dbresolver.Read)
}

func (o orderDo) WriteDB() *orderDo {
	return o.Clauses(dbresolver.Write)
}

func (o orderDo) Clauses(conds ...clause.Expression) *orderDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o orderDo) Returning(value interface{}, columns ...string) *orderDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o orderDo) Not(conds ...gen.Condition) *orderDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o orderDo) Or(conds ...gen.Condition) *orderDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o orderDo) Select(conds ...field.Expr) *orderDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o orderDo) Where(conds ...gen.Condition) *orderDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o orderDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *orderDo {
	return o.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (o orderDo) Order(conds ...field.Expr) *orderDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o orderDo) Distinct(cols ...field.Expr) *orderDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o orderDo) Omit(cols ...field.Expr) *orderDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o orderDo) Join(table schema.Tabler, on ...field.Expr) *orderDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o orderDo) LeftJoin(table schema.Tabler, on ...field.Expr) *orderDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o orderDo) RightJoin(table schema.Tabler, on ...field.Expr) *orderDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o orderDo) Group(cols ...field.Expr) *orderDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o orderDo) Having(conds ...gen.Condition) *orderDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o orderDo) Limit(limit int) *orderDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o orderDo) Offset(offset int) *orderDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o orderDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *orderDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o orderDo) Unscoped() *orderDo {
	return o.withDO(o.DO.Unscoped())
}

func (o orderDo) Create(values ...*domain.Order) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o orderDo) CreateInBatches(values []*domain.Order, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o orderDo) Save(values ...*domain.Order) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o orderDo) First() (*domain.Order, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Order), nil
	}
}

func (o orderDo) Take() (*domain.Order, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Order), nil
	}
}

func (o orderDo) Last() (*domain.Order, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Order), nil
	}
}

func (o orderDo) Find() ([]*domain.Order, error) {
	result, err := o.DO.Find()
	return result.([]*domain.Order), err
}

func (o orderDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*domain.Order, err error) {
	buf := make([]*domain.Order, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o orderDo) FindInBatches(result *[]*domain.Order, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o orderDo) Attrs(attrs ...field.AssignExpr) *orderDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o orderDo) Assign(attrs ...field.AssignExpr) *orderDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o orderDo) Joins(fields ...field.RelationField) *orderDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o orderDo) Preload(fields ...field.RelationField) *orderDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o orderDo) FirstOrInit() (*domain.Order, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Order), nil
	}
}

func (o orderDo) FirstOrCreate() (*domain.Order, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Order), nil
	}
}

func (o orderDo) FindByPage(offset int, limit int) (result []*domain.Order, count int64, err error) {
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

func (o orderDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o orderDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o orderDo) Delete(models ...*domain.Order) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *orderDo) withDO(do gen.Dao) *orderDo {
	o.DO = *do.(*gen.DO)
	return o
}
