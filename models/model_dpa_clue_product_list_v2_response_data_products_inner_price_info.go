/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaClueProductListV2ResponseDataProductsInnerPriceInfo 价格信息
type DpaClueProductListV2ResponseDataProductsInnerPriceInfo struct {
	// 日付
	DailyMortgage *string `json:"daily_mortgage,omitempty"`
	// 折扣
	Discount *float64 `json:"discount,omitempty"`
	// 月付
	DownPayment *string `json:"down_payment,omitempty"`
	// 首付
	Mortgage *string `json:"mortgage,omitempty"`
	// 商品现价
	Price *float64 `json:"price,omitempty"`
	// 价格单位
	PriceUnit *string `json:"price_unit,omitempty"`
	// 促销活动
	SalesPromotion *string `json:"sales_promotion,omitempty"`
	// 减价
	Saving *float64 `json:"saving,omitempty"`
	// 商品原价
	Value *float64 `json:"value,omitempty"`
}
