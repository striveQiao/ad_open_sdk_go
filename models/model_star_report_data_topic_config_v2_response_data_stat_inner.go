/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarReportDataTopicConfigV2ResponseDataStatInner struct for StarReportDataTopicConfigV2ResponseDataStatInner
type StarReportDataTopicConfigV2ResponseDataStatInner struct {
	DataTopic StarReportDataTopicConfigV2DataStatDataTopic `json:"data_topic"`
	// 数据主题下的数据指标，一个数据主题对应一个或多个数据指标。不存在的数据指标将缺省。
	Metrics []*StarReportDataTopicConfigV2ResponseDataStatInnerMetricsInner `json:"metrics"`
}
