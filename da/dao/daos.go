// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/plugin/dbresolver"
)

var (
	Q       = new(Query)
	Company *company
)

func SetDefault(db *gorm.DB) {
	*Q = *Use(db)
	Company = &Q.Company
}
func Use(db *gorm.DB) *Query {
	return &Query{
		db:                     db,
		Area:                   newArea(db),
		BillingOrder:           newBillingOrder(db),
		BillsOrdersItem:        newBillsOrdersItem(db),
		Company:                newCompany(db),
		Customer:               newCustomer(db),
		CustomerPoject:         newCustomerPoject(db),
		Employee:               newEmployee(db),
		LogisticsCompany:       newLogisticsCompany(db),
		Order:                  newOrder(db),
		OrderStatus:            newOrderStatus(db),
		OrdersFee:              newOrdersFee(db),
		OrdersFinance:          newOrdersFinance(db),
		OrdersGood:             newOrdersGood(db),
		Organization:           newOrganization(db),
		Store:                  newStore(db),
		SysOpenUser:            newSysOpenUser(db),
		SysPermission:          newSysPermission(db),
		SysPermissionGroup:     newSysPermissionGroup(db),
		SysPermissionItem:      newSysPermissionItem(db),
		SysReourceExt:          newSysReourceExt(db),
		SysResource:            newSysResource(db),
		SysRole:                newSysRole(db),
		SysRolePermissionGroup: newSysRolePermissionGroup(db),
		SysUser:                newSysUser(db),
		SysUserGroup:           newSysUserGroup(db),
		SysUserGroupRole:       newSysUserGroupRole(db),
		SysUserRole:            newSysUserRole(db),
		Vehicle:                newVehicle(db),
		IDGENERATOR:            newIDGENERATOR(db),
	}
}

type Query struct {
	db *gorm.DB

	Area                   area
	BillingOrder           billingOrder
	BillsOrdersItem        billsOrdersItem
	Company                company
	Customer               customer
	CustomerPoject         customerPoject
	Employee               employee
	LogisticsCompany       logisticsCompany
	Order                  order
	OrderStatus            orderStatus
	OrdersFee              ordersFee
	OrdersFinance          ordersFinance
	OrdersGood             ordersGood
	Organization           organization
	Store                  store
	SysOpenUser            sysOpenUser
	SysPermission          sysPermission
	SysPermissionGroup     sysPermissionGroup
	SysPermissionItem      sysPermissionItem
	SysReourceExt          sysReourceExt
	SysResource            sysResource
	SysRole                sysRole
	SysRolePermissionGroup sysRolePermissionGroup
	SysUser                sysUser
	SysUserGroup           sysUserGroup
	SysUserGroupRole       sysUserGroupRole
	SysUserRole            sysUserRole
	Vehicle                vehicle
	IDGENERATOR            iDGENERATOR
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                     db,
		Area:                   q.Area.clone(db),
		BillingOrder:           q.BillingOrder.clone(db),
		BillsOrdersItem:        q.BillsOrdersItem.clone(db),
		Company:                q.Company.clone(db),
		Customer:               q.Customer.clone(db),
		CustomerPoject:         q.CustomerPoject.clone(db),
		Employee:               q.Employee.clone(db),
		LogisticsCompany:       q.LogisticsCompany.clone(db),
		Order:                  q.Order.clone(db),
		OrderStatus:            q.OrderStatus.clone(db),
		IDGENERATOR:            q.IDGENERATOR.clone(db),
		OrdersFee:              q.OrdersFee.clone(db),
		OrdersFinance:          q.OrdersFinance.clone(db),
		OrdersGood:             q.OrdersGood.clone(db),
		Organization:           q.Organization.clone(db),
		Store:                  q.Store.clone(db),
		SysOpenUser:            q.SysOpenUser.clone(db),
		SysPermission:          q.SysPermission.clone(db),
		SysPermissionGroup:     q.SysPermissionGroup.clone(db),
		SysPermissionItem:      q.SysPermissionItem.clone(db),
		SysReourceExt:          q.SysReourceExt.clone(db),
		SysResource:            q.SysResource.clone(db),
		SysRole:                q.SysRole.clone(db),
		SysRolePermissionGroup: q.SysRolePermissionGroup.clone(db),
		SysUser:                q.SysUser.clone(db),
		SysUserGroup:           q.SysUserGroup.clone(db),
		SysUserGroupRole:       q.SysUserGroupRole.clone(db),
		SysUserRole:            q.SysUserRole.clone(db),
		Vehicle:                q.Vehicle.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.clone(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.clone(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return q.clone(db)
}

type queryCtx struct {
	Area                   *areaDo
	BillingOrder           *billingOrderDo
	BillsOrdersItem        *billsOrdersItemDo
	Company                *companyDo
	Customer               *customerDo
	CustomerPoject         *customerPojectDo
	Employee               *employeeDo
	LogisticsCompany       *logisticsCompanyDo
	Order                  *orderDo
	OrderStatus            *orderStatusDo
	OrdersFee              *ordersFeeDo
	OrdersFinance          *ordersFinanceDo
	OrdersGood             *ordersGoodDo
	Organization           *organizationDo
	Store                  *storeDo
	SysOpenUser            *sysOpenUserDo
	SysPermission          *sysPermissionDo
	SysPermissionGroup     *sysPermissionGroupDo
	SysPermissionItem      *sysPermissionItemDo
	SysReourceExt          *sysReourceExtDo
	SysResource            *sysResourceDo
	SysRole                *sysRoleDo
	SysRolePermissionGroup *sysRolePermissionGroupDo
	SysUser                *sysUserDo
	SysUserGroup           *sysUserGroupDo
	SysUserGroupRole       *sysUserGroupRoleDo
	SysUserRole            *sysUserRoleDo
	Vehicle                *vehicleDo
	IDGENERATOR            *iDGENERATORDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Area:                   q.Area.WithContext(ctx),
		BillingOrder:           q.BillingOrder.WithContext(ctx),
		BillsOrdersItem:        q.BillsOrdersItem.WithContext(ctx),
		Company:                q.Company.WithContext(ctx),
		Customer:               q.Customer.WithContext(ctx),
		CustomerPoject:         q.CustomerPoject.WithContext(ctx),
		Employee:               q.Employee.WithContext(ctx),
		LogisticsCompany:       q.LogisticsCompany.WithContext(ctx),
		Order:                  q.Order.WithContext(ctx),
		OrderStatus:            q.OrderStatus.WithContext(ctx),
		OrdersFee:              q.OrdersFee.WithContext(ctx),
		OrdersFinance:          q.OrdersFinance.WithContext(ctx),
		OrdersGood:             q.OrdersGood.WithContext(ctx),
		Organization:           q.Organization.WithContext(ctx),
		Store:                  q.Store.WithContext(ctx),
		SysOpenUser:            q.SysOpenUser.WithContext(ctx),
		SysPermission:          q.SysPermission.WithContext(ctx),
		SysPermissionGroup:     q.SysPermissionGroup.WithContext(ctx),
		SysPermissionItem:      q.SysPermissionItem.WithContext(ctx),
		SysReourceExt:          q.SysReourceExt.WithContext(ctx),
		SysResource:            q.SysResource.WithContext(ctx),
		SysRole:                q.SysRole.WithContext(ctx),
		SysRolePermissionGroup: q.SysRolePermissionGroup.WithContext(ctx),
		SysUser:                q.SysUser.WithContext(ctx),
		SysUserGroup:           q.SysUserGroup.WithContext(ctx),
		SysUserGroupRole:       q.SysUserGroupRole.WithContext(ctx),
		SysUserRole:            q.SysUserRole.WithContext(ctx),
		Vehicle:                q.Vehicle.WithContext(ctx),
		IDGENERATOR:            q.IDGENERATOR.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
