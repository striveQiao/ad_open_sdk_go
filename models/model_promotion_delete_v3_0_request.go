/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionDeleteV30Request struct for PromotionDeleteV30Request
type PromotionDeleteV30Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 待删除的广告ID列表
	PromotionIds []int64 `json:"promotion_ids"`
}
