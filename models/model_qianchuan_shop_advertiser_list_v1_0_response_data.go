/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanShopAdvertiserListV10ResponseData
type QianchuanShopAdvertiserListV10ResponseData struct {
	//
	AdvIdList []*QianchuanShopAdvertiserListV10ResponseDataAdvIdListInner `json:"adv_id_list,omitempty"`
	//
	List     []int64                                             `json:"list,omitempty"`
	PageInfo *QianchuanShopAdvertiserListV10ResponseDataPageInfo `json:"page_info,omitempty"`
}
