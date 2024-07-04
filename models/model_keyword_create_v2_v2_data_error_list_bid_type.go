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

// KeywordCreateV2V2DataErrorListBidType
type KeywordCreateV2V2DataErrorListBidType string

// List of keyword_create_v2_v2_data_error_list_bid_type
const (
	SUGGEST_KeywordCreateV2V2DataErrorListBidType        KeywordCreateV2V2DataErrorListBidType = "SUGGEST"
	FEED_TO_SEARCH_KeywordCreateV2V2DataErrorListBidType KeywordCreateV2V2DataErrorListBidType = "FEED_TO_SEARCH"
	BRAND_AD_KeywordCreateV2V2DataErrorListBidType       KeywordCreateV2V2DataErrorListBidType = "BRAND_AD"
	CUSTOM_KeywordCreateV2V2DataErrorListBidType         KeywordCreateV2V2DataErrorListBidType = "CUSTOM"
	WITH_AD_KeywordCreateV2V2DataErrorListBidType        KeywordCreateV2V2DataErrorListBidType = "WITH_AD"
)

// All allowed values of KeywordCreateV2V2DataErrorListBidType enum
var AllowedKeywordCreateV2V2DataErrorListBidTypeEnumValues = []KeywordCreateV2V2DataErrorListBidType{
	"SUGGEST",
	"FEED_TO_SEARCH",
	"BRAND_AD",
	"CUSTOM",
	"WITH_AD",
}

// NewKeywordCreateV2V2DataErrorListBidTypeFromValue returns a pointer to a valid KeywordCreateV2V2DataErrorListBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeywordCreateV2V2DataErrorListBidTypeFromValue(v string) (*KeywordCreateV2V2DataErrorListBidType, error) {
	ev := KeywordCreateV2V2DataErrorListBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeywordCreateV2V2DataErrorListBidType: valid values are %v", v, AllowedKeywordCreateV2V2DataErrorListBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeywordCreateV2V2DataErrorListBidType) IsValid() bool {
	for _, existing := range AllowedKeywordCreateV2V2DataErrorListBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to keyword_create_v2_v2_data_error_list_bid_type value
func (v KeywordCreateV2V2DataErrorListBidType) Ptr() *KeywordCreateV2V2DataErrorListBidType {
	return &v
}
