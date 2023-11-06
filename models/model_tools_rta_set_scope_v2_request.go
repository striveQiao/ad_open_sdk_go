/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsRtaSetScopeV2Request struct for ToolsRtaSetScopeV2Request
type ToolsRtaSetScopeV2Request struct {
	// 广告账户id
	AdvertiserId int64 `json:"advertiser_id"`
	// 预期设置的rta策略id
	RtaId int64 `json:"rta_id"`
	// 生效ID列表，当target_type = CAMPAIGN 或 PROJECT 有效
	TargetIds  []int64                       `json:"target_ids,omitempty"`
	TargetType *ToolsRtaSetScopeV2TargetType `json:"target_type,omitempty"`
}
