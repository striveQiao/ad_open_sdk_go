/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdKeywordsUpdateV10RequestKeywordsInner struct for QianchuanAdKeywordsUpdateV10RequestKeywordsInner
type QianchuanAdKeywordsUpdateV10RequestKeywordsInner struct {
	// 关键词ID
	Id         int64                                           `json:"id"`
	MatchType  *QianchuanAdKeywordsUpdateV10KeywordsMatchType  `json:"match_type,omitempty"`
	StatusType *QianchuanAdKeywordsUpdateV10KeywordsStatusType `json:"status_type,omitempty"`
}
