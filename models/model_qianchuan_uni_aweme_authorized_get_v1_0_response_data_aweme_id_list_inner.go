/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniAwemeAuthorizedGetV10ResponseDataAwemeIdListInner struct for QianchuanUniAwemeAuthorizedGetV10ResponseDataAwemeIdListInner
type QianchuanUniAwemeAuthorizedGetV10ResponseDataAwemeIdListInner struct {
	//
	AwemeAvatar *string `json:"aweme_avatar,omitempty"`
	//
	AwemeId *int64 `json:"aweme_id,omitempty"`
	//
	AwemeName *string `json:"aweme_name,omitempty"`
	//
	AwemeShowId *string `json:"aweme_show_id,omitempty"`
	//
	HasRoi2DeliveryLimit *bool `json:"has_roi2_delivery_limit,omitempty"`
	//
	HasRoi2GroupCreated *bool `json:"has_roi2_group_created,omitempty"`
}
