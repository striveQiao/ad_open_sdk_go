/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniPromotionListV10ResponseDataAdListInner struct for QianchuanUniPromotionListV10ResponseDataAdListInner
type QianchuanUniPromotionListV10ResponseDataAdListInner struct {
	AdInfo *QianchuanUniPromotionListV10ResponseDataAdListInnerAdInfo `json:"ad_info,omitempty"`
	//
	RoomInfo  []*QianchuanUniPromotionListV10ResponseDataAdListInnerRoomInfoInner `json:"room_info,omitempty"`
	StatsInfo *QianchuanUniPromotionListV10ResponseDataAdListInnerStatsInfo       `json:"stats_info,omitempty"`
}
