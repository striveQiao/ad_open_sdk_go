/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanReportSearchWordGetV10FilteringRange 关键词/搜索词 只支持部分指标过滤
type QianchuanReportSearchWordGetV10FilteringRange struct {
	Calculation *QianchuanReportSearchWordGetV10FilteringRangeCalculation `json:"calculation,omitempty"`
	//
	Field *string `json:"field,omitempty"`
	//
	Value *string `json:"value,omitempty"`
}
