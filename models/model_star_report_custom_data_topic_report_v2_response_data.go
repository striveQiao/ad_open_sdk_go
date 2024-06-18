/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarReportCustomDataTopicReportV2ResponseData
type StarReportCustomDataTopicReportV2ResponseData struct {
	// 对应请求的数据主题数组
	Data []*StarReportCustomDataTopicReportV2ResponseDataDataInner `json:"data,omitempty"`
	// 请求的任务Id
	DemandId int64 `json:"demand_id"`
	// 请求的交付物Id
	ItemId int64 `json:"item_id"`
}
