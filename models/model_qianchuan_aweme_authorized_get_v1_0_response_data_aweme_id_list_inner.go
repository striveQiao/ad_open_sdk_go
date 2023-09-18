/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeAuthorizedGetV10ResponseDataAwemeIdListInner struct for QianchuanAwemeAuthorizedGetV10ResponseDataAwemeIdListInner
type QianchuanAwemeAuthorizedGetV10ResponseDataAwemeIdListInner struct {
	//
	AwemeAvatar *string `json:"aweme_avatar,omitempty"`
	//
	AwemeHasLivePermission *bool `json:"aweme_has_live_permission,omitempty"`
	//
	AwemeHasUniProm *bool `json:"aweme_has_uni_prom,omitempty"`
	//
	AwemeHasVideoPermission *bool `json:"aweme_has_video_permission,omitempty"`
	//
	AwemeId *int64 `json:"aweme_id,omitempty"`
	//
	AwemeName *string `json:"aweme_name,omitempty"`
	//
	AwemeShowId *string                                                   `json:"aweme_show_id,omitempty"`
	AwemeStatus *QianchuanAwemeAuthorizedGetV10DataAwemeIdListAwemeStatus `json:"aweme_status,omitempty"`
	//
	BindType []*QianchuanAwemeAuthorizedGetV10DataAwemeIdListBindType `json:"bind_type,omitempty"`
}
