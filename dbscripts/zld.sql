/*
 Navicat Premium Data Transfer

 Source Server         : 82.157.168.170
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : 82.157.168.170:3306
 Source Schema         : zld

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 04/10/2022 23:42:06
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for Areas
-- ----------------------------
DROP TABLE IF EXISTS `Areas`;
CREATE TABLE `Areas`  (
  `id` bigint(20) NOT NULL COMMENT '区域Id',
  `code` char(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '区域code',
  `name` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '区域名称',
  `city_code` char(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '城市Code',
  `sort` smallint(6) NULL DEFAULT NULL COMMENT '排序索引',
  `short_name` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '简称',
  `status` tinyint(4) NULL DEFAULT 0 COMMENT '状态(0:启用 1 停用)',
  `show_short_name` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '显示简称名称',
  `custom` tinyint(4) NULL DEFAULT 0 COMMENT '是否自定义(0:否 1 是)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '区域' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Billing_Orders
-- ----------------------------
DROP TABLE IF EXISTS `Billing_Orders`;
CREATE TABLE `Billing_Orders`  (
  `Id` bigint(20) NOT NULL COMMENT '开票Id',
  `Cus_Id` bigint(20) NULL DEFAULT NULL COMMENT '客户Id',
  `company_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '公司名称',
  `apply_no` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '发票申请单号',
  `apply_amount` decimal(8, 2) NULL DEFAULT NULL COMMENT '申请金额',
  `apply_date` datetime(0) NULL DEFAULT NULL COMMENT '申请时间',
  `status` tinyint(4) NULL DEFAULT NULL COMMENT '票据状态(0 待申请 1 开票中 2 已开票 3 已驳回 4作废)',
  `billing_no` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '发票编号',
  `taxpayer_ID_number` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '纳税人识别号',
  `tel` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '电话',
  `address` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址',
  `bank_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '开户银行',
  `bank_account` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '开户账户',
  `consignee_person` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '收货人',
  `consignee_province_code` char(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '收货省份Code',
  `consignee_city_code` char(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '收货城市Code',
  `consignee_area_code` char(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '收货地区code',
  `consignee_address` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '收货详细地址',
  `consignee_email` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '收货人邮箱',
  `consignee_tel` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '收货人电话',
  `billing_date` datetime(0) NULL DEFAULT NULL COMMENT '开票时间',
  PRIMARY KEY (`Id`) USING BTREE,
  INDEX `FK_Customer_Billing_Orders`(`Cus_Id`) USING BTREE,
  CONSTRAINT `FK_Customer_Billing_Orders` FOREIGN KEY (`Cus_Id`) REFERENCES `Customers` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '开票申请单' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Bills_Orders_Item
-- ----------------------------
DROP TABLE IF EXISTS `Bills_Orders_Item`;
CREATE TABLE `Bills_Orders_Item`  (
  `Id` bigint(20) NOT NULL COMMENT '开票申请单明细Id',
  `Ord_id` bigint(20) NULL DEFAULT NULL COMMENT '订单Id',
  `Bil_Id` bigint(20) NULL DEFAULT NULL COMMENT '开票Id',
  PRIMARY KEY (`Id`) USING BTREE,
  INDEX `FK_Bills_Orders_Item`(`Bil_Id`) USING BTREE,
  INDEX `FK_Orders_Bills_Orders_Item`(`Ord_id`) USING BTREE,
  CONSTRAINT `FK_Bills_Orders_Item` FOREIGN KEY (`Bil_Id`) REFERENCES `Billing_Orders` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `FK_Orders_Bills_Orders_Item` FOREIGN KEY (`Ord_id`) REFERENCES `Orders` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '开票申请单明细' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Company
-- ----------------------------
DROP TABLE IF EXISTS `Company`;
CREATE TABLE `Company`  (
  `Id` bigint(20) NOT NULL COMMENT 'Id',
  `Company_Name` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '公司名称',
  `Company_Type` smallint(6) NULL DEFAULT NULL COMMENT '公司类型(0:政府 1: 消防 3 :企业 )',
  `Company_Desc` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `Register_Date_Time` datetime(0) NULL DEFAULT NULL COMMENT '注册时间',
  `Status` smallint(6) NULL DEFAULT 0 COMMENT '状态(0: 正常 1:停用)',
  `Insert_Date_Time` datetime(0) NULL DEFAULT NULL COMMENT '插入时间',
  `Insert_User` bigint(20) NULL DEFAULT NULL COMMENT '插入用户',
  `Update_Date_Time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `Update_User` bigint(20) NULL DEFAULT NULL COMMENT '更新用户',
  `Version` int(11) NULL DEFAULT 1 COMMENT '版本',
  `Deleted` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '删除标志(0:正常 1 删除)',
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '公司' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Customer_Pojects
-- ----------------------------
DROP TABLE IF EXISTS `Customer_Pojects`;
CREATE TABLE `Customer_Pojects`  (
  `Id` bigint(20) NOT NULL COMMENT '项目Id',
  `Cus_Id` bigint(20) NULL DEFAULT NULL COMMENT '客户Id',
  `Project_Name` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '项目名称',
  `Project_Desc` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '项目描述',
  `Project_Memo` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '说明',
  PRIMARY KEY (`Id`) USING BTREE,
  INDEX `FK_Customer_Project`(`Cus_Id`) USING BTREE,
  CONSTRAINT `FK_Customer_Project` FOREIGN KEY (`Cus_Id`) REFERENCES `Customers` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '客户项目' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Customers
-- ----------------------------
DROP TABLE IF EXISTS `Customers`;
CREATE TABLE `Customers`  (
  `Id` bigint(20) NOT NULL COMMENT '客户Id',
  `Customer_Name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '客户名称',
  `Customer_Address` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '客户地址',
  `Contact_Man` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '客户联系人',
  `Contact_Tel` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '客户电话',
  `Enter_Tel` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '企业联系人',
  `Province_Code` char(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '省code',
  `City_Code` char(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '市code',
  `Area_Code` char(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '区code',
  `Memo` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `Insert_Date_Time` datetime(0) NOT NULL COMMENT '插入时间',
  `Insert_User` bigint(20) NOT NULL COMMENT '插入用户',
  `Update_Date_Time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `Update_User` bigint(20) NULL DEFAULT NULL COMMENT '更新用户',
  `Version` int(11) NOT NULL COMMENT '版本',
  `Deleted` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '删除标志(0:正常 1 删除)',
  `Compay_Id` bigint(20) NOT NULL COMMENT '公司',
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '客户信息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Employee
-- ----------------------------
DROP TABLE IF EXISTS `Employee`;
CREATE TABLE `Employee`  (
  `Id` bigint(20) NOT NULL COMMENT '雇员Id',
  `Org_Id` bigint(20) NULL DEFAULT NULL COMMENT 'Id',
  `Emp_No` char(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '雇员编号',
  `Emp_Name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '雇员名称',
  `Emp_Mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '雇员手机号',
  `Insert_Date_Time` datetime(0) NULL DEFAULT NULL COMMENT '插入时间',
  `Insert_User` bigint(20) NULL DEFAULT NULL COMMENT '插入用户',
  `Update_Date_Time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `Update_User` bigint(20) NULL DEFAULT NULL COMMENT '更新用户',
  `Version` int(11) NULL DEFAULT 1 COMMENT '版本',
  `Deleted` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '删除标志(0:正常 1 删除)',
  PRIMARY KEY (`Id`) USING BTREE,
  INDEX `FK_Organization_Employee`(`Org_Id`) USING BTREE,
  CONSTRAINT `FK_Organization_Employee` FOREIGN KEY (`Org_Id`) REFERENCES `Organization` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '雇员' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Logistics_Company
-- ----------------------------
DROP TABLE IF EXISTS `Logistics_Company`;
CREATE TABLE `Logistics_Company`  (
  `Id` bigint(20) NOT NULL COMMENT '物流公司Id',
  `Company_Name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '公司名称/车企名称',
  `Corporation` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '法人姓名',
  `Type` smallint(6) NULL DEFAULT 0 COMMENT '类型(0:个人 1 :公司 2 车队)',
  `Tel` char(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机号',
  `Bank_no` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '银行卡',
  `Plate_Total_Income` decimal(10, 2) NULL DEFAULT NULL COMMENT '平台总收益',
  `Account_Quota` decimal(10, 2) NULL DEFAULT NULL COMMENT '账户限额',
  `Register_Date` datetime(0) NULL DEFAULT NULL COMMENT '注册时间',
  `Sign` tinyint(4) NULL DEFAULT 1 COMMENT '是否签约(0:未签约 1 已签约)',
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '物流公司' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Order_Status
-- ----------------------------
DROP TABLE IF EXISTS `Order_Status`;
CREATE TABLE `Order_Status`  (
  `Id` bigint(20) NOT NULL COMMENT '订单费用Id',
  `Ord_id` bigint(20) NULL DEFAULT NULL COMMENT '订单Id',
  `Order_Status` tinyint(4) NOT NULL COMMENT '订单状态(0:草稿 1:已下单 2:已接单 3: 已装货 4: 已签收)',
  `Order_Desc` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '描述',
  `Memo` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`Id`) USING BTREE,
  INDEX `FK_Orders_Status`(`Ord_id`) USING BTREE,
  CONSTRAINT `FK_Orders_Status` FOREIGN KEY (`Ord_id`) REFERENCES `Orders` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '订单状态明细' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Orders
-- ----------------------------
DROP TABLE IF EXISTS `Orders`;
CREATE TABLE `Orders`  (
  `id` bigint(20) NOT NULL COMMENT '订单Id',
  `Log_Id` bigint(20) NULL DEFAULT NULL COMMENT '物流公司Id',
  `order_no` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单号',
  `out_order_no` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '外部订单号',
  `project_name` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '项目名称',
  `shipment_date` datetime(0) NULL DEFAULT NULL COMMENT '托运时间',
  `status` int(11) NULL DEFAULT NULL COMMENT '订单状态(0:草稿 1:已下单 2:已接单 3: 已装货 4: 已签收)',
  `deliver_province_code` char(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '发货省code',
  `deliver_city_code` char(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '发货城市Code',
  `deliver_area_code` char(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '发货地区Code',
  `deliver_address` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '出发地',
  `arrive_province_code` char(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '到达省份Code',
  `arrive_city_code` char(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '到达城市Code',
  `arrive_area_code` char(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '到达地区Code',
  `arrive_address` char(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '到达地',
  `arrive_date` datetime(0) NULL DEFAULT NULL COMMENT '到达时间',
  `receipt_customer` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '收货客户',
  `receipt_customer_tel` char(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '收货客户电话',
  `pricing_method` tinyint(4) NULL DEFAULT NULL COMMENT '计价方式(0: 趟 1 吨 2 方)',
  `return_bill_attachment` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '回单附件',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `FK_Logistics_Company_Orders`(`Log_Id`) USING BTREE,
  CONSTRAINT `FK_Logistics_Company_Orders` FOREIGN KEY (`Log_Id`) REFERENCES `Logistics_Company` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '订单' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Orders_Fee
-- ----------------------------
DROP TABLE IF EXISTS `Orders_Fee`;
CREATE TABLE `Orders_Fee`  (
  `Id` bigint(20) NOT NULL COMMENT '订单费用Id',
  `Ord_id` bigint(20) NULL DEFAULT NULL COMMENT '订单Id',
  `pick_up_fee` decimal(9, 2) NULL DEFAULT NULL COMMENT '提货费',
  `delivery_fee` decimal(9, 2) NULL DEFAULT NULL COMMENT '送货费',
  `transport_fee` decimal(9, 2) NULL DEFAULT NULL COMMENT '干线费',
  `unloading_fee` decimal(9, 2) NULL DEFAULT NULL COMMENT '卸载费',
  `quote_fee` decimal(9, 2) NULL DEFAULT NULL COMMENT '报价金额',
  `premium_fee` decimal(9, 2) NULL DEFAULT NULL COMMENT '保费',
  `non_tax_fee` decimal(9, 2) NULL DEFAULT NULL COMMENT '不含税金额',
  `tax_rate` decimal(9, 2) NULL DEFAULT NULL COMMENT '税率',
  `bonded_fee_total` decimal(9, 2) NULL DEFAULT NULL COMMENT '保税合计',
  PRIMARY KEY (`Id`) USING BTREE,
  INDEX `FK_Orders_Fees`(`Ord_id`) USING BTREE,
  CONSTRAINT `FK_Orders_Fees` FOREIGN KEY (`Ord_id`) REFERENCES `Orders` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '订单费用' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Orders_Finances
-- ----------------------------
DROP TABLE IF EXISTS `Orders_Finances`;
CREATE TABLE `Orders_Finances`  (
  `Id` bigint(20) NOT NULL COMMENT '结算Id',
  `Ord_id` bigint(20) NULL DEFAULT NULL COMMENT '订单Id',
  `type` tinyint(4) NULL DEFAULT NULL COMMENT '结算类型(0 收款 1 付款)',
  `non_tax_amount` decimal(8, 2) NULL DEFAULT NULL COMMENT '不含税金额',
  `tax_rate` decimal(2, 2) NULL DEFAULT NULL COMMENT '税率',
  `total` decimal(8, 2) NULL DEFAULT NULL COMMENT '价税合计',
  `bill_tag` tinyint(4) NULL DEFAULT NULL COMMENT '是否开票',
  `operate_date` datetime(0) NULL DEFAULT NULL COMMENT '收款/付款时间',
  `operate_way` tinyint(4) NULL DEFAULT NULL COMMENT '收款/付款方式(0 线下 2 银行卡 3 微信支付 4 支付宝)',
  PRIMARY KEY (`Id`) USING BTREE,
  INDEX `FK_Orders_Finances`(`Ord_id`) USING BTREE,
  CONSTRAINT `FK_Orders_Finances` FOREIGN KEY (`Ord_id`) REFERENCES `Orders` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '订单结算' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Orders_Goods
-- ----------------------------
DROP TABLE IF EXISTS `Orders_Goods`;
CREATE TABLE `Orders_Goods`  (
  `Id` bigint(20) NOT NULL COMMENT '货物Id',
  `Ord_id` bigint(20) NULL DEFAULT NULL COMMENT '订单Id',
  `goods_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '货物名称',
  `weight` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '重量',
  `volume` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '体积',
  `size` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '尺寸',
  PRIMARY KEY (`Id`) USING BTREE,
  INDEX `FK_Orders_Goods`(`Ord_id`) USING BTREE,
  CONSTRAINT `FK_Orders_Goods` FOREIGN KEY (`Ord_id`) REFERENCES `Orders` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '订单货物' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Organization
-- ----------------------------
DROP TABLE IF EXISTS `Organization`;
CREATE TABLE `Organization`  (
  `Id` bigint(20) NOT NULL COMMENT 'Id',
  `Com_Id` bigint(20) NULL DEFAULT NULL COMMENT 'Id',
  `Org_Id` bigint(20) NULL DEFAULT NULL COMMENT 'Id',
  `Org_Name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '组织名称',
  `Org_Code` char(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '组织代码',
  `Org_Level` int(11) NULL DEFAULT NULL COMMENT '组织级别',
  `Begin_Date` datetime(0) NULL DEFAULT NULL COMMENT '开始日期',
  `End_Date` datetime(0) NULL DEFAULT NULL COMMENT '结束日期',
  `Insert_Date_Time` datetime(0) NULL DEFAULT NULL COMMENT '插入时间',
  `Insert_User` bigint(20) NULL DEFAULT NULL COMMENT '插入用户',
  `Update_Date_Time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `Update_User` bigint(20) NULL DEFAULT NULL COMMENT '更新用户',
  `Version` int(11) NULL DEFAULT 1 COMMENT '版本',
  `Deleted` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '删除标志(0:正常 1 删除)',
  PRIMARY KEY (`Id`) USING BTREE,
  INDEX `FK_Company_Organization`(`Com_Id`) USING BTREE,
  INDEX `FK_Organization_Owe`(`Org_Id`) USING BTREE,
  CONSTRAINT `FK_Company_Organization` FOREIGN KEY (`Com_Id`) REFERENCES `Company` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `FK_Organization_Owe` FOREIGN KEY (`Org_Id`) REFERENCES `Organization` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '组织' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Stores
-- ----------------------------
DROP TABLE IF EXISTS `Stores`;
CREATE TABLE `Stores`  (
  `Id` bigint(20) NOT NULL COMMENT '仓储Id',
  `Store_Name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '仓储名称',
  `Store_Tel` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '联系电话',
  `Store_Address` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址',
  `Memo` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '仓储' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Sys_Open_Users
-- ----------------------------
DROP TABLE IF EXISTS `Sys_Open_Users`;
CREATE TABLE `Sys_Open_Users`  (
  `Id` bigint(20) NOT NULL COMMENT 'Id',
  `Open_Identity` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '平台标示',
  `Open_Name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '平台名称',
  `Open_Type` smallint(6) NULL DEFAULT NULL COMMENT '开发平台类型(1:钉钉 2: 微信 )',
  `Open_Memo1` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '平台扩展',
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '平台用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Sys_Permission
-- ----------------------------
DROP TABLE IF EXISTS `Sys_Permission`;
CREATE TABLE `Sys_Permission`  (
  `Id` bigint(20) NOT NULL COMMENT 'Id',
  `Permission_Val` smallint(6) NULL DEFAULT NULL COMMENT '权限值',
  `Permission_Desc` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '权限描述',
  `App_Permission_Val` smallint(6) NULL DEFAULT NULL COMMENT '平台权限值',
  `Module_Permission_Val` smallint(6) NULL DEFAULT NULL COMMENT '模块权限值',
  `Menu_Permission_Val` smallint(6) NULL DEFAULT NULL COMMENT '菜单权限值',
  `Page_Permission_Val` smallint(6) NULL DEFAULT NULL COMMENT '页面权限值',
  `Common_Operator_Permission_Val` smallint(6) NULL DEFAULT NULL COMMENT '通用操作权限值',
  `Permission_Scope` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '权限空间',
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '权限描述\r\n1.添加权限 按位或（|）\r\n2.校验权限：按位与【&】\r\n3.删除权限:先按' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Sys_Permission_Group
-- ----------------------------
DROP TABLE IF EXISTS `Sys_Permission_Group`;
CREATE TABLE `Sys_Permission_Group`  (
  `Id` bigint(20) NOT NULL COMMENT 'Id',
  `Sys_Id` bigint(20) NULL DEFAULT NULL COMMENT 'Id',
  `Permission_Group_Type` smallint(6) NULL DEFAULT NULL COMMENT '许可组类型',
  `Permission_Group_Desc` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '许可组描述',
  `Insert_Date_Time` datetime(0) NOT NULL COMMENT '插入时间',
  `Insert_User` bigint(20) NOT NULL COMMENT '插入用户',
  `Update_Date_Time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `Update_User` bigint(20) NULL DEFAULT NULL COMMENT '更新用户',
  `Version` int(11) NOT NULL COMMENT '版本',
  `Deleted` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '删除标志(0:正常 1 删除)',
  `Compay_Id` bigint(20) NOT NULL COMMENT '公司',
  PRIMARY KEY (`Id`) USING BTREE,
  INDEX `FK_Sys_Permission_Group_Persmission`(`Sys_Id`) USING BTREE,
  CONSTRAINT `FK_Sys_Permission_Group_Persmission` FOREIGN KEY (`Sys_Id`) REFERENCES `Sys_Permission` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '权限许可组' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Sys_Permission_Item
-- ----------------------------
DROP TABLE IF EXISTS `Sys_Permission_Item`;
CREATE TABLE `Sys_Permission_Item`  (
  `Id` bigint(20) NOT NULL COMMENT 'Id',
  `Sys_Id` bigint(20) NULL DEFAULT NULL COMMENT 'Id',
  `Sys_Id2` bigint(20) NULL DEFAULT NULL COMMENT '资源Id',
  `Permission_Item_Val` smallint(6) NULL DEFAULT NULL COMMENT '权限值',
  `Permission_Type` smallint(6) NULL DEFAULT NULL COMMENT '权限类型',
  `Permission_Desc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '权限描述',
  PRIMARY KEY (`Id`) USING BTREE,
  INDEX `FK_Sys_Permission_Resource_Item`(`Sys_Id`) USING BTREE,
  INDEX `FK_Sys_Resource_Permission_Item`(`Sys_Id2`) USING BTREE,
  CONSTRAINT `FK_Sys_Permission_Resource_Item` FOREIGN KEY (`Sys_Id`) REFERENCES `Sys_Permission` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `FK_Sys_Resource_Permission_Item` FOREIGN KEY (`Sys_Id2`) REFERENCES `Sys_Resources` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Sys_Reource_Ext
-- ----------------------------
DROP TABLE IF EXISTS `Sys_Reource_Ext`;
CREATE TABLE `Sys_Reource_Ext`  (
  `Id` bigint(20) NULL DEFAULT NULL COMMENT '资源Id',
  `Component` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '组件',
  `Always_Show` tinyint(1) NULL DEFAULT NULL COMMENT '显示',
  `Meta_Title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标题',
  `Meta_Limit` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '元数据限制',
  `Meta_Breadcrumb` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '元数据面包屑',
  `Meta_Icon` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '元数据图标',
  `Meta_No_Cache` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '元数据缓存',
  `Meta_Hidden` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '元数据隐藏域',
  `Redirect` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '重定向URL',
  `Hidden` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '隐藏域',
  `Type` smallint(6) NULL DEFAULT NULL COMMENT '类型',
  INDEX `FK_Sys_Resource_Res_Ext`(`Id`) USING BTREE,
  CONSTRAINT `FK_Sys_Resource_Res_Ext` FOREIGN KEY (`Id`) REFERENCES `Sys_Resources` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '资源扩展' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Sys_Resources
-- ----------------------------
DROP TABLE IF EXISTS `Sys_Resources`;
CREATE TABLE `Sys_Resources`  (
  `Id` bigint(20) NOT NULL COMMENT '资源Id',
  `Sys_Id` bigint(20) NULL DEFAULT NULL COMMENT '资源Id',
  `Resource_Identity` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '资源标示',
  `Resource_Type` smallint(6) NOT NULL COMMENT '资源类型(0 app   1 模块  2  菜单 3  页面 4  按钮 5  扩展按钮)',
  `Resource_Status` smallint(6) NULL DEFAULT NULL COMMENT '资源状态(0:正常 1 :停用)',
  `Resource_Actopn_Type` smallint(6) NULL DEFAULT NULL COMMENT '资源动作类型 (-1 无 1.增加;2 删 3 改 4 上传 5导出 6 查询 .)',
  `Resource_Action_Identity` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '资源动作标示(针对扩展按钮，备用字段)',
  `Resource_Desc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '资源描述',
  `Resource_Orders` int(11) NULL DEFAULT NULL COMMENT '资源序序号',
  `Insert_Date_Time` datetime(0) NOT NULL COMMENT '插入时间',
  `Insert_User` bigint(20) NOT NULL COMMENT '插入用户',
  `Update_Date_Time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `Update_User` bigint(20) NULL DEFAULT NULL COMMENT '更新用户',
  `Version` int(11) NOT NULL COMMENT '版本',
  `Deleted` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '删除标志',
  `Compay_Id3` bigint(20) NOT NULL COMMENT '公司',
  PRIMARY KEY (`Id`) USING BTREE,
  INDEX `FK_Sys_Resources_Owe`(`Sys_Id`) USING BTREE,
  CONSTRAINT `FK_Sys_Resources_Owe` FOREIGN KEY (`Sys_Id`) REFERENCES `Sys_Resources` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '资源' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Sys_Role_Permission_Group
-- ----------------------------
DROP TABLE IF EXISTS `Sys_Role_Permission_Group`;
CREATE TABLE `Sys_Role_Permission_Group`  (
  `Id` bigint(20) NOT NULL,
  `Sys_Id` bigint(20) NULL DEFAULT NULL COMMENT 'Id',
  `Sys_Id2` bigint(20) NULL DEFAULT NULL COMMENT '角色Id',
  PRIMARY KEY (`Id`) USING BTREE,
  INDEX `FK_Sys_Permission_Group_Roles`(`Sys_Id`) USING BTREE,
  INDEX `FK_Sys_Role_Permission_Group`(`Sys_Id2`) USING BTREE,
  CONSTRAINT `FK_Sys_Permission_Group_Roles` FOREIGN KEY (`Sys_Id`) REFERENCES `Sys_Permission_Group` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `FK_Sys_Role_Permission_Group` FOREIGN KEY (`Sys_Id2`) REFERENCES `Sys_Roles` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Sys_Roles
-- ----------------------------
DROP TABLE IF EXISTS `Sys_Roles`;
CREATE TABLE `Sys_Roles`  (
  `Id` bigint(20) NOT NULL COMMENT '角色Id',
  `Role_Code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色代码',
  `Role_Name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色名称',
  `Insert_Date_Time` datetime(0) NOT NULL COMMENT '插入时间',
  `Insert_User` bigint(20) NOT NULL COMMENT '插入用户',
  `Update_Date_Time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `Update_User` bigint(20) NULL DEFAULT NULL COMMENT '更新用户',
  `Version` int(11) NOT NULL COMMENT '版本',
  `Deleted` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '删除标志(0:正常 1 删除)',
  `Compay_Id` bigint(20) NOT NULL COMMENT '公司',
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Sys_User_Group
-- ----------------------------
DROP TABLE IF EXISTS `Sys_User_Group`;
CREATE TABLE `Sys_User_Group`  (
  `Id` bigint(20) NOT NULL COMMENT 'Id',
  `Sys_Id` bigint(20) NULL DEFAULT NULL COMMENT 'Id',
  `Sys_Id2` bigint(20) NULL DEFAULT NULL COMMENT 'Id',
  `User_Group_Type` smallint(6) NULL DEFAULT NULL COMMENT '用户组类型',
  `User_Group_Desc` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户组描述',
  `Insert_Date_Time` datetime(0) NOT NULL COMMENT '插入时间',
  `Insert_User` bigint(20) NOT NULL COMMENT '插入用户',
  `Update_Date_Time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `Update_User` bigint(20) NULL DEFAULT NULL COMMENT '更新用户',
  `Version` int(11) NOT NULL COMMENT '版本',
  `Deleted` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '删除标志(0:正常 1 删除)',
  `Compay_Id` bigint(20) NOT NULL COMMENT '公司',
  PRIMARY KEY (`Id`) USING BTREE,
  INDEX `FK_Sys_User_Group_Owe`(`Sys_Id`) USING BTREE,
  INDEX `FK_Sys_User_User_Group`(`Sys_Id2`) USING BTREE,
  CONSTRAINT `FK_Sys_User_Group_Owe` FOREIGN KEY (`Sys_Id`) REFERENCES `Sys_User_Group` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `FK_Sys_User_User_Group` FOREIGN KEY (`Sys_Id2`) REFERENCES `Sys_Users` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户组' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Sys_User_Roles
-- ----------------------------
DROP TABLE IF EXISTS `Sys_User_Roles`;
CREATE TABLE `Sys_User_Roles`  (
  `Id` bigint(20) NOT NULL,
  `Sys_Id` bigint(20) NULL DEFAULT NULL COMMENT '角色Id',
  `Sys_Id2` bigint(20) NULL DEFAULT NULL COMMENT 'Id',
  PRIMARY KEY (`Id`) USING BTREE,
  INDEX `FK_Sys_Roles_Users`(`Sys_Id`) USING BTREE,
  INDEX `FK_Sys_User_Roles`(`Sys_Id2`) USING BTREE,
  CONSTRAINT `FK_Sys_Roles_Users` FOREIGN KEY (`Sys_Id`) REFERENCES `Sys_Roles` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `FK_Sys_User_Roles` FOREIGN KEY (`Sys_Id2`) REFERENCES `Sys_Users` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户角色' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Sys_Users
-- ----------------------------
DROP TABLE IF EXISTS `Sys_Users`;
CREATE TABLE `Sys_Users`  (
  `Id` bigint(20) NOT NULL COMMENT 'Id',
  `Sys_Id` bigint(20) NULL DEFAULT NULL COMMENT 'Id',
  `Emp_Id` bigint(20) NULL DEFAULT NULL COMMENT '雇员Id',
  `User_Name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户名称',
  `User_Email` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户邮箱',
  `User_Password` varchar(344) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '登录密码',
  `User_Status` smallint(6) NULL DEFAULT NULL COMMENT '用户状态(0:正常 1 :停用 )',
  `User_Type` smallint(6) NULL DEFAULT NULL COMMENT '用户类型',
  `Insert_Date_Time` datetime(0) NULL DEFAULT NULL COMMENT '插入时间',
  `Insert_User` bigint(20) NULL DEFAULT NULL COMMENT '插入用户',
  `Update_Date_Time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `Update_User` bigint(20) NULL DEFAULT NULL COMMENT '更新用户',
  `Version` int(11) NULL DEFAULT NULL COMMENT '版本',
  `Deleted` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '删除标志(0:正常 1 删除)',
  `Compay_Id` bigint(20) NULL DEFAULT NULL COMMENT '公司',
  PRIMARY KEY (`Id`) USING BTREE,
  INDEX `FK_Employee_Users`(`Emp_Id`) USING BTREE,
  INDEX `FK_Sys_Open_Users`(`Sys_Id`) USING BTREE,
  CONSTRAINT `FK_Employee_Users` FOREIGN KEY (`Emp_Id`) REFERENCES `Employee` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `FK_Sys_Open_Users` FOREIGN KEY (`Sys_Id`) REFERENCES `Sys_Open_Users` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Vehicles
-- ----------------------------
DROP TABLE IF EXISTS `Vehicles`;
CREATE TABLE `Vehicles`  (
  `Id` bigint(20) NOT NULL COMMENT '车辆Id',
  `Log_Id` bigint(20) NULL DEFAULT NULL COMMENT '物流公司Id',
  `License_Plate` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '车牌',
  `Model` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '车型(厢式/平板/高栏/爬梯)',
  `Len` varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '车长(4.2/6.8/9.6/13/13.75/17.5)',
  `Memo` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`Id`) USING BTREE,
  INDEX `FK_Logistics_Company_Vehicles`(`Log_Id`) USING BTREE,
  CONSTRAINT `FK_Logistics_Company_Vehicles` FOREIGN KEY (`Log_Id`) REFERENCES `Logistics_Company` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '车辆信息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_User_Group_Roles
-- ----------------------------
DROP TABLE IF EXISTS `sys_User_Group_Roles`;
CREATE TABLE `sys_User_Group_Roles`  (
  `Id` bigint(20) NOT NULL,
  `Sys_Id` bigint(20) NULL DEFAULT NULL COMMENT '角色Id',
  `Sys_Id2` bigint(20) NULL DEFAULT NULL COMMENT 'Id',
  PRIMARY KEY (`Id`) USING BTREE,
  INDEX `FK_Sys_Role_User_Group`(`Sys_Id`) USING BTREE,
  INDEX `FK_Sys_User_Group_Roles`(`Sys_Id2`) USING BTREE,
  CONSTRAINT `FK_Sys_Role_User_Group` FOREIGN KEY (`Sys_Id`) REFERENCES `Sys_Roles` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `FK_Sys_User_Group_Roles` FOREIGN KEY (`Sys_Id2`) REFERENCES `Sys_User_Group` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户角色' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
