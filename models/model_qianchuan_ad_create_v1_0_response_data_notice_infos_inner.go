/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdCreateV10ResponseDataNoticeInfosInner struct for QianchuanAdCreateV10ResponseDataNoticeInfosInner
type QianchuanAdCreateV10ResponseDataNoticeInfosInner struct {
	//
	Code *int64 `json:"code,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	SearchKeywordError []*QianchuanAdCreateV10ResponseDataNoticeInfosInnerSearchKeywordErrorInner `json:"search_keyword_error,omitempty"`
}
