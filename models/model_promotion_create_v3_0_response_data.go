/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionCreateV30ResponseData
type PromotionCreateV30ResponseData struct {
	//
	ErrorKeywordsList []*PromotionCreateV30ResponseDataErrorKeywordsListInner `json:"error_keywords_list,omitempty"`
	//
	PromotionId *int64 `json:"promotion_id,omitempty"`
}
