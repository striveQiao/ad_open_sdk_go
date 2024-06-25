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

// KeywordCreateV30DataSuccessListBidType
type KeywordCreateV30DataSuccessListBidType string

// List of keyword_create_v3.0_data_success_list_bid_type
const (
	CUSTOM_KeywordCreateV30DataSuccessListBidType         KeywordCreateV30DataSuccessListBidType = "CUSTOM"
	FEED_TO_SEARCH_KeywordCreateV30DataSuccessListBidType KeywordCreateV30DataSuccessListBidType = "FEED_TO_SEARCH"
	WITH_PROMOTION_KeywordCreateV30DataSuccessListBidType KeywordCreateV30DataSuccessListBidType = "WITH_PROMOTION"
)

// All allowed values of KeywordCreateV30DataSuccessListBidType enum
var AllowedKeywordCreateV30DataSuccessListBidTypeEnumValues = []KeywordCreateV30DataSuccessListBidType{
	"CUSTOM",
	"FEED_TO_SEARCH",
	"WITH_PROMOTION",
}

// NewKeywordCreateV30DataSuccessListBidTypeFromValue returns a pointer to a valid KeywordCreateV30DataSuccessListBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeywordCreateV30DataSuccessListBidTypeFromValue(v string) (*KeywordCreateV30DataSuccessListBidType, error) {
	ev := KeywordCreateV30DataSuccessListBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeywordCreateV30DataSuccessListBidType: valid values are %v", v, AllowedKeywordCreateV30DataSuccessListBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeywordCreateV30DataSuccessListBidType) IsValid() bool {
	for _, existing := range AllowedKeywordCreateV30DataSuccessListBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to keyword_create_v3.0_data_success_list_bid_type value
func (v KeywordCreateV30DataSuccessListBidType) Ptr() *KeywordCreateV30DataSuccessListBidType {
	return &v
}
