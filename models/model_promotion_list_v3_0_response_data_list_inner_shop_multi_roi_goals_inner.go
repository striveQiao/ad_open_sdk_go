/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionListV30ResponseDataListInnerShopMultiRoiGoalsInner struct for PromotionListV30ResponseDataListInnerShopMultiRoiGoalsInner
type PromotionListV30ResponseDataListInnerShopMultiRoiGoalsInner struct {
	//
	RoiGoal      *float64                                               `json:"roi_goal,omitempty"`
	ShopPlatform *PromotionListV30DataListShopMultiRoiGoalsShopPlatform `json:"shop_platform,omitempty"`
}
