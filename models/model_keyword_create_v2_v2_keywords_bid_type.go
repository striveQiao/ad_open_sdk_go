/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// KeywordCreateV2V2KeywordsBidType
type KeywordCreateV2V2KeywordsBidType string

// List of keyword_create_v2_v2_keywords_bid_type
const (
	WITH_AD_KeywordCreateV2V2KeywordsBidType        KeywordCreateV2V2KeywordsBidType = "WITH_AD"
	SUGGEST_KeywordCreateV2V2KeywordsBidType        KeywordCreateV2V2KeywordsBidType = "SUGGEST"
	FEED_TO_SEARCH_KeywordCreateV2V2KeywordsBidType KeywordCreateV2V2KeywordsBidType = "FEED_TO_SEARCH"
	BRAND_AD_KeywordCreateV2V2KeywordsBidType       KeywordCreateV2V2KeywordsBidType = "BRAND_AD"
	CUSTOM_KeywordCreateV2V2KeywordsBidType         KeywordCreateV2V2KeywordsBidType = "CUSTOM"
)

// All allowed values of KeywordCreateV2V2KeywordsBidType enum
var AllowedKeywordCreateV2V2KeywordsBidTypeEnumValues = []KeywordCreateV2V2KeywordsBidType{
	"WITH_AD",
	"SUGGEST",
	"FEED_TO_SEARCH",
	"BRAND_AD",
	"CUSTOM",
}

// NewKeywordCreateV2V2KeywordsBidTypeFromValue returns a pointer to a valid KeywordCreateV2V2KeywordsBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeywordCreateV2V2KeywordsBidTypeFromValue(v string) (*KeywordCreateV2V2KeywordsBidType, error) {
	ev := KeywordCreateV2V2KeywordsBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeywordCreateV2V2KeywordsBidType: valid values are %v", v, AllowedKeywordCreateV2V2KeywordsBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeywordCreateV2V2KeywordsBidType) IsValid() bool {
	for _, existing := range AllowedKeywordCreateV2V2KeywordsBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to keyword_create_v2_v2_keywords_bid_type value
func (v KeywordCreateV2V2KeywordsBidType) Ptr() *KeywordCreateV2V2KeywordsBidType {
	return &v
}
