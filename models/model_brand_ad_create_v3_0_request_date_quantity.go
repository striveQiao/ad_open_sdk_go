/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandAdCreateV30RequestDateQuantity 预订量
type BrandAdCreateV30RequestDateQuantity struct {
	// 预算
	Budget float64 `json:"budget"`
	// 每天的预订量
	DailyQuantity            []*BrandAdCreateV30RequestDateQuantityDailyQuantityInner `json:"daily_quantity"`
	StockIncreasePackageType *BrandAdCreateV30DateQuantityStockIncreasePackageType    `json:"stock_increase_package_type,omitempty"`
}
