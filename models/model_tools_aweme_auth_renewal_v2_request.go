/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAwemeAuthRenewalV2Request struct for ToolsAwemeAuthRenewalV2Request
type ToolsAwemeAuthRenewalV2Request struct {
	//
	AdvertiserId *int64                          `json:"advertiser_id,omitempty"`
	AuthType     ToolsAwemeAuthRenewalV2AuthType `json:"auth_type"`
	// 抖音号，如果authType是直播授权或抖音号授权，则该字段必填
	AwemeId *string `json:"aweme_id,omitempty"`
	// 新的授权截止日期，格式：“yyyy-MM-dd HH:mm:ss”。例如：\"2024-01-12 10:00:00\"
	EndTime string `json:"end_time"`
	// 抖音短视频ID，如果authType是短视频授权，则该字段必填
	ItemId *int64 `json:"item_id,omitempty"`
}
