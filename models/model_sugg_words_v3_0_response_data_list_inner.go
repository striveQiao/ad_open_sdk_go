/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SuggWordsV30ResponseDataListInner struct for SuggWordsV30ResponseDataListInner
type SuggWordsV30ResponseDataListInner struct {
	// 推荐关键词
	Keyword *string `json:"keyword,omitempty"`
	// 月搜索量
	Msv *int64 `json:"msv,omitempty"`
}
