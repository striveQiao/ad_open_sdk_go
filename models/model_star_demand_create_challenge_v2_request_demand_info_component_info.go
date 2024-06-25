/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandCreateChallengeV2RequestDemandInfoComponentInfo 组件信息
type StarDemandCreateChallengeV2RequestDemandInfoComponentInfo struct {
	// 行业组件ID 行业组件不可与Link（常规）组件同时使用，目前仅支持游戏/电商/网服类型 考核指标为组件点击数、激活总数、安装完成数量和APP唤起时必填
	IndustryComponentId *int64 `json:"industry_component_id,omitempty"`
	// Link组件ID（落地页组件） 目前只支持1个，仅支持落地页类型
	LinkComponentIds []int64 `json:"link_component_ids,omitempty"`
	// 直播引流组件ID
	LiveAttractComponentId *int64 `json:"live_attract_component_id,omitempty"`
}
