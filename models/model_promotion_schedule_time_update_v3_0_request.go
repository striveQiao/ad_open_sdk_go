/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionScheduleTimeUpdateV30Request struct for PromotionScheduleTimeUpdateV30Request
type PromotionScheduleTimeUpdateV30Request struct {
	// 广告账户id
	AdvertiserId int64 `json:"advertiser_id"`
	// 批量修改投放时段，限制最多批量修改10条广告
	Data []*PromotionScheduleTimeUpdateV30RequestDataInner `json:"data"`
}
