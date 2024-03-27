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

// KeywordUpdateV2V2KeywordsMatchType
type KeywordUpdateV2V2KeywordsMatchType string

// List of keyword_update_v2_v2_keywords_match_type
const (
	PHRASE_KeywordUpdateV2V2KeywordsMatchType    KeywordUpdateV2V2KeywordsMatchType = "PHRASE"
	EXTENSIVE_KeywordUpdateV2V2KeywordsMatchType KeywordUpdateV2V2KeywordsMatchType = "EXTENSIVE"
	PRECISION_KeywordUpdateV2V2KeywordsMatchType KeywordUpdateV2V2KeywordsMatchType = "PRECISION"
)

// All allowed values of KeywordUpdateV2V2KeywordsMatchType enum
var AllowedKeywordUpdateV2V2KeywordsMatchTypeEnumValues = []KeywordUpdateV2V2KeywordsMatchType{
	"PHRASE",
	"EXTENSIVE",
	"PRECISION",
}

// NewKeywordUpdateV2V2KeywordsMatchTypeFromValue returns a pointer to a valid KeywordUpdateV2V2KeywordsMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeywordUpdateV2V2KeywordsMatchTypeFromValue(v string) (*KeywordUpdateV2V2KeywordsMatchType, error) {
	ev := KeywordUpdateV2V2KeywordsMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeywordUpdateV2V2KeywordsMatchType: valid values are %v", v, AllowedKeywordUpdateV2V2KeywordsMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeywordUpdateV2V2KeywordsMatchType) IsValid() bool {
	for _, existing := range AllowedKeywordUpdateV2V2KeywordsMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to keyword_update_v2_v2_keywords_match_type value
func (v KeywordUpdateV2V2KeywordsMatchType) Ptr() *KeywordUpdateV2V2KeywordsMatchType {
	return &v
}
