/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionBudgetUpdateV30ResponseData
type PromotionBudgetUpdateV30ResponseData struct {
	// 更新失败的广告ID列表
	Errors []*PromotionBudgetUpdateV30ResponseDataErrorsInner `json:"errors,omitempty"`
	// 更新成功的广告ID列表
	PromotionIds []int64 `json:"promotion_ids,omitempty"`
}
