/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionListV30DataListKeywordsBidType
type PromotionListV30DataListKeywordsBidType string

// List of promotion_list_v3.0_data_list_keywords_bid_type
const (
	CUSTOM_PromotionListV30DataListKeywordsBidType         PromotionListV30DataListKeywordsBidType = "CUSTOM"
	WITH_PROMOTION_PromotionListV30DataListKeywordsBidType PromotionListV30DataListKeywordsBidType = "WITH_PROMOTION"
)

// All allowed values of PromotionListV30DataListKeywordsBidType enum
var AllowedPromotionListV30DataListKeywordsBidTypeEnumValues = []PromotionListV30DataListKeywordsBidType{
	"CUSTOM",
	"WITH_PROMOTION",
}

// NewPromotionListV30DataListKeywordsBidTypeFromValue returns a pointer to a valid PromotionListV30DataListKeywordsBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListKeywordsBidTypeFromValue(v string) (*PromotionListV30DataListKeywordsBidType, error) {
	ev := PromotionListV30DataListKeywordsBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListKeywordsBidType: valid values are %v", v, AllowedPromotionListV30DataListKeywordsBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListKeywordsBidType) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListKeywordsBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_keywords_bid_type value
func (v PromotionListV30DataListKeywordsBidType) Ptr() *PromotionListV30DataListKeywordsBidType {
	return &v
}
