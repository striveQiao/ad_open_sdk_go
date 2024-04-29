/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdRecommendKeywordsGetV10OrderType
type QianchuanAdRecommendKeywordsGetV10OrderType string

// List of qianchuan_ad_recommend_keywords_get_v1.0_order_type
const (
	ASC_QianchuanAdRecommendKeywordsGetV10OrderType  QianchuanAdRecommendKeywordsGetV10OrderType = "ASC"
	DESC_QianchuanAdRecommendKeywordsGetV10OrderType QianchuanAdRecommendKeywordsGetV10OrderType = "DESC"
)

// All allowed values of QianchuanAdRecommendKeywordsGetV10OrderType enum
var AllowedQianchuanAdRecommendKeywordsGetV10OrderTypeEnumValues = []QianchuanAdRecommendKeywordsGetV10OrderType{
	"ASC",
	"DESC",
}

// NewQianchuanAdRecommendKeywordsGetV10OrderTypeFromValue returns a pointer to a valid QianchuanAdRecommendKeywordsGetV10OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdRecommendKeywordsGetV10OrderTypeFromValue(v string) (*QianchuanAdRecommendKeywordsGetV10OrderType, error) {
	ev := QianchuanAdRecommendKeywordsGetV10OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdRecommendKeywordsGetV10OrderType: valid values are %v", v, AllowedQianchuanAdRecommendKeywordsGetV10OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdRecommendKeywordsGetV10OrderType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdRecommendKeywordsGetV10OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_recommend_keywords_get_v1.0_order_type value
func (v QianchuanAdRecommendKeywordsGetV10OrderType) Ptr() *QianchuanAdRecommendKeywordsGetV10OrderType {
	return &v
}
