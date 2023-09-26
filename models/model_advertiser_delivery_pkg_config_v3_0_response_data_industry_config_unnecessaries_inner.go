/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserDeliveryPkgConfigV30ResponseDataIndustryConfigUnnecessariesInner struct for AdvertiserDeliveryPkgConfigV30ResponseDataIndustryConfigUnnecessariesInner
type AdvertiserDeliveryPkgConfigV30ResponseDataIndustryConfigUnnecessariesInner struct {
	// 规则组合id
	CombineId int64 `json:"combine_id"`
	// 选填资质场景描述，用于引导用户提交
	Description *string `json:"description,omitempty"`
	// 具体的资质规则
	Rules []*AdvertiserDeliveryPkgConfigV30ResponseDataIndustryConfigUnnecessariesInnerRulesInner `json:"rules,omitempty"`
}
