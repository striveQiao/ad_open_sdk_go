/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EnterpriseBindListGetV10ResponseDataListInner struct for EnterpriseBindListGetV10ResponseDataListInner
type EnterpriseBindListGetV10ResponseDataListInner struct {
	//
	AuthorizeTimes []*EnterpriseBindListGetV10ResponseDataListInnerAuthorizeTimesInner `json:"authorize_times,omitempty"`
	//
	AwemeAvatar *string `json:"aweme_avatar,omitempty"`
	//
	AwemeId *string `json:"aweme_id,omitempty"`
	//
	AwemeName *string `json:"aweme_name,omitempty"`
	//
	OpenId *string `json:"open_id,omitempty"`
}
