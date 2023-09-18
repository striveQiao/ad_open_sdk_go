/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAbTestUpdateV2Request struct for ToolsAbTestUpdateV2Request
type ToolsAbTestUpdateV2Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 实验结束时间，格式为：\"2020-12-25 07:12:08\"，开始与结束时间跨度至少选择1天，最多选择30天。
	TestEndTime *string `json:"test_end_time,omitempty"`
	// 实验ID
	TestId int64 `json:"test_id"`
	// 实验名称，长度1-100字符
	TestName *string `json:"test_name,omitempty"`
	// 实验开始时间，格式为：\"2020-12-25 07:12:08\"，开始时间不能早于当前时间
	TestStartTime *string `json:"test_start_time,omitempty"`
}
