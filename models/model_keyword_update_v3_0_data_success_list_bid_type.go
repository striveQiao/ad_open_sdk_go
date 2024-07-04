/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// KeywordUpdateV30DataSuccessListBidType
type KeywordUpdateV30DataSuccessListBidType string

// List of keyword_update_v3.0_data_success_list_bid_type
const (
	CUSTOM_KeywordUpdateV30DataSuccessListBidType         KeywordUpdateV30DataSuccessListBidType = "CUSTOM"
	FEED_TO_SEARCH_KeywordUpdateV30DataSuccessListBidType KeywordUpdateV30DataSuccessListBidType = "FEED_TO_SEARCH"
	WITH_PROMOTION_KeywordUpdateV30DataSuccessListBidType KeywordUpdateV30DataSuccessListBidType = "WITH_PROMOTION"
)

// All allowed values of KeywordUpdateV30DataSuccessListBidType enum
var AllowedKeywordUpdateV30DataSuccessListBidTypeEnumValues = []KeywordUpdateV30DataSuccessListBidType{
	"CUSTOM",
	"FEED_TO_SEARCH",
	"WITH_PROMOTION",
}

// NewKeywordUpdateV30DataSuccessListBidTypeFromValue returns a pointer to a valid KeywordUpdateV30DataSuccessListBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeywordUpdateV30DataSuccessListBidTypeFromValue(v string) (*KeywordUpdateV30DataSuccessListBidType, error) {
	ev := KeywordUpdateV30DataSuccessListBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeywordUpdateV30DataSuccessListBidType: valid values are %v", v, AllowedKeywordUpdateV30DataSuccessListBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeywordUpdateV30DataSuccessListBidType) IsValid() bool {
	for _, existing := range AllowedKeywordUpdateV30DataSuccessListBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to keyword_update_v3.0_data_success_list_bid_type value
func (v KeywordUpdateV30DataSuccessListBidType) Ptr() *KeywordUpdateV30DataSuccessListBidType {
	return &v
}
