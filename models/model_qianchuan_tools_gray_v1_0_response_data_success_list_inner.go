/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanToolsGrayV10ResponseDataSuccessListInner struct for QianchuanToolsGrayV10ResponseDataSuccessListInner
type QianchuanToolsGrayV10ResponseDataSuccessListInner struct {
	//
	AwemeIds []int64 `json:"aweme_ids,omitempty"`
	//
	GrayKey string `json:"gray_key"`
	//
	InGray *bool `json:"in_gray,omitempty"`
	//
	InWhitelist *bool `json:"in_whitelist,omitempty"`
}
