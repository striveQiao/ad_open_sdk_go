/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSuggestBudgetGetV30ResponseDataListInner struct for ToolsSuggestBudgetGetV30ResponseDataListInner
type ToolsSuggestBudgetGetV30ResponseDataListInner struct {
	// 广告预算
	PromotionBudget *float64 `json:"promotion_budget,omitempty"`
	// 广告ID
	PromotionId *int64 `json:"promotion_id,omitempty"`
	// 建议起量预算
	SuggestBudget *float64 `json:"suggest_budget,omitempty"`
}
