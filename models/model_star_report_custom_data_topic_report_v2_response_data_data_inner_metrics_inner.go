/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarReportCustomDataTopicReportV2ResponseDataDataInnerMetricsInner struct for StarReportCustomDataTopicReportV2ResponseDataDataInnerMetricsInner
type StarReportCustomDataTopicReportV2ResponseDataDataInnerMetricsInner struct {
	// 指标描述，建议首次使用前先调用/open_api/2/star/report/data_topic_config/接口详细了解具体数据主题下的Topic配置
	Description *string `json:"description,omitempty"`
	// 指标字段，仅包含英文与下划线等符号
	Field string `json:"field"`
	// 指标名称，中文或英文描述字段名称
	Name string                                               `json:"name"`
	Type StarReportCustomDataTopicReportV2DataDataMetricsType `json:"type"`
	// 指标值
	Value *string `json:"value,omitempty"`
}
