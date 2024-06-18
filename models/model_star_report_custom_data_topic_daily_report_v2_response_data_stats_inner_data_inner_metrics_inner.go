/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarReportCustomDataTopicDailyReportV2ResponseDataStatsInnerDataInnerMetricsInner struct for StarReportCustomDataTopicDailyReportV2ResponseDataStatsInnerDataInnerMetricsInner
type StarReportCustomDataTopicDailyReportV2ResponseDataStatsInnerDataInnerMetricsInner struct {
	// 指标描述
	Description *string `json:"description,omitempty"`
	// 指标字段，仅包含英文与下划线等符号
	Field string `json:"field"`
	// 指标名称，中文或英文描述字段名称
	Name string                                                         `json:"name"`
	Type StarReportCustomDataTopicDailyReportV2DataStatsDataMetricsType `json:"type"`
	// 指标值
	Value *string `json:"value,omitempty"`
}
