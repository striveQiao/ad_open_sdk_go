/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandQueryStockV30ScheduleInfoScheduleListInner struct for BrandQueryStockV30ScheduleInfoScheduleListInner
type BrandQueryStockV30ScheduleInfoScheduleListInner struct {
	// 日期
	Date string `json:"date"`
	// 时段
	Period []string `json:"period,omitempty"`
}
