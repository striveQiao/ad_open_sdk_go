/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionScheduleTimeUpdateV30ResponseDataErrorsInner struct for PromotionScheduleTimeUpdateV30ResponseDataErrorsInner
type PromotionScheduleTimeUpdateV30ResponseDataErrorsInner struct {
	// 错误信息
	ErrorMessage *string `json:"error_message,omitempty"`
	// 广告ID
	PromotionId *int64 `json:"promotion_id,omitempty"`
}
