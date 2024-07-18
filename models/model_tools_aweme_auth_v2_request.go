/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAwemeAuthV2Request struct for ToolsAwemeAuthV2Request
type ToolsAwemeAuthV2Request struct {
	// 广告主ID
	AdvertiserId int64                    `json:"advertiser_id"`
	AuthType     ToolsAwemeAuthV2AuthType `json:"auth_type"`
	// 抖音号
	AwemeId string `json:"aweme_id"`
	// 授权码
	Code string `json:"code"`
	// 授权结束时间
	EndTime string `json:"end_time"`
	// 抖音短视频链接
	ItemUrl *string `json:"item_url,omitempty"`
	// 授权开始时间
	StartTime string `json:"start_time"`
}
