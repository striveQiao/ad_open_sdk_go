/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportBrandAdvertiserGetV30ResponseData
type ReportBrandAdvertiserGetV30ResponseData struct {
	//
	DataReports []*ReportBrandAdvertiserGetV30ResponseDataDataReportsInner `json:"data_reports,omitempty"`
	// 数据总数
	TotalCount *int64 `json:"total_count,omitempty"`
}
