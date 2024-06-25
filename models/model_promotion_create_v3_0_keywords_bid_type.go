/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionCreateV30KeywordsBidType
type PromotionCreateV30KeywordsBidType string

// List of promotion_create_v3.0_keywords_bid_type
const (
	CUSTOM_PromotionCreateV30KeywordsBidType         PromotionCreateV30KeywordsBidType = "CUSTOM"
	WITH_PROMOTION_PromotionCreateV30KeywordsBidType PromotionCreateV30KeywordsBidType = "WITH_PROMOTION"
)

// All allowed values of PromotionCreateV30KeywordsBidType enum
var AllowedPromotionCreateV30KeywordsBidTypeEnumValues = []PromotionCreateV30KeywordsBidType{
	"CUSTOM",
	"WITH_PROMOTION",
}

// NewPromotionCreateV30KeywordsBidTypeFromValue returns a pointer to a valid PromotionCreateV30KeywordsBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCreateV30KeywordsBidTypeFromValue(v string) (*PromotionCreateV30KeywordsBidType, error) {
	ev := PromotionCreateV30KeywordsBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCreateV30KeywordsBidType: valid values are %v", v, AllowedPromotionCreateV30KeywordsBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCreateV30KeywordsBidType) IsValid() bool {
	for _, existing := range AllowedPromotionCreateV30KeywordsBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_create_v3.0_keywords_bid_type value
func (v PromotionCreateV30KeywordsBidType) Ptr() *PromotionCreateV30KeywordsBidType {
	return &v
}
