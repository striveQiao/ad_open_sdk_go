/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionListV30ResponseDataListInnerNativeSetting
type PromotionListV30ResponseDataListInnerNativeSetting struct {
	AnchorRelatedType *PromotionListV30DataListNativeSettingAnchorRelatedType `json:"anchor_related_type,omitempty"`
	//
	AwemeId         *string                                               `json:"aweme_id,omitempty"`
	IsFeedAndFavSee *PromotionListV30DataListNativeSettingIsFeedAndFavSee `json:"is_feed_and_fav_see,omitempty"`
}
