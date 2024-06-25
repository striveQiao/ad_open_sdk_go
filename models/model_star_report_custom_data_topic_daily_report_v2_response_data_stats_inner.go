/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarReportCustomDataTopicDailyReportV2ResponseDataStatsInner struct for StarReportCustomDataTopicDailyReportV2ResponseDataStatsInner
type StarReportCustomDataTopicDailyReportV2ResponseDataStatsInner struct {
	// 详细增量指标数据列表
	Data []*StarReportCustomDataTopicDailyReportV2ResponseDataStatsInnerDataInner `json:"data,omitempty"`
	// 该指标数据的yyyy-MM-dd日期时间, 按天聚合。注意⚠️：BasicData作为通用指标，在增量请求中date为 \"0\" 表示通用数据
	Date *string `json:"date,omitempty"`
	// 请求的任务Id
	DemandId int64 `json:"demand_id"`
	// 请求的作品交付物Id
	WorkId int64 `json:"work_id"`
}
