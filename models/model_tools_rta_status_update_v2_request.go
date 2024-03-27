/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsRtaStatusUpdateV2Request struct for ToolsRtaStatusUpdateV2Request
type ToolsRtaStatusUpdateV2Request struct {
	AccountType *ToolsRtaStatusUpdateV2AccountType `json:"account_type,omitempty"`
	// 广告账户id
	AdvertiserId int64 `json:"advertiser_id"`
	// 预期设置的rta策略id
	RtaIds []int64                      `json:"rta_ids"`
	Status ToolsRtaStatusUpdateV2Status `json:"status"`
}
