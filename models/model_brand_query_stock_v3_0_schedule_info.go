/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandQueryStockV30ScheduleInfo
type BrandQueryStockV30ScheduleInfo struct {
	// 需要查询的日期列表
	ScheduleList             []*BrandQueryStockV30ScheduleInfoScheduleListInner      `json:"schedule_list"`
	StockIncreasePackageType *BrandQueryStockV30ScheduleInfoStockIncreasePackageType `json:"stock_increase_package_type,omitempty"`
}
