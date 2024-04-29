/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCustomGetV30ResponseData
type ReportCustomGetV30ResponseData struct {
	PageInfo *ReportCustomGetV30ResponseDataPageInfo `json:"page_info,omitempty"`
	//
	Rows []*ReportCustomGetV30ResponseDataRowsInner `json:"rows"`
	// 指标汇总数据
	TotalMetrics map[string]string `json:"total_metrics"`
}
