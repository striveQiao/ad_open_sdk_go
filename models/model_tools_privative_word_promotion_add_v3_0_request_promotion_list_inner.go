/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordPromotionAddV30RequestPromotionListInner struct for ToolsPrivativeWordPromotionAddV30RequestPromotionListInner
type ToolsPrivativeWordPromotionAddV30RequestPromotionListInner struct {
	// 短语否定词列表
	PhraseWords []string `json:"phrase_words,omitempty"`
	// 精确否定词列表
	PreciseWords []string `json:"precise_words,omitempty"`
	// 项目id
	PromotionId int64 `json:"promotion_id"`
}
