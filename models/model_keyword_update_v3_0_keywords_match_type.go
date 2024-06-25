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

// KeywordUpdateV30KeywordsMatchType
type KeywordUpdateV30KeywordsMatchType string

// List of keyword_update_v3.0_keywords_match_type
const (
	EXTENSIVE_KeywordUpdateV30KeywordsMatchType KeywordUpdateV30KeywordsMatchType = "EXTENSIVE"
	PHRASE_KeywordUpdateV30KeywordsMatchType    KeywordUpdateV30KeywordsMatchType = "PHRASE"
	PRECISION_KeywordUpdateV30KeywordsMatchType KeywordUpdateV30KeywordsMatchType = "PRECISION"
)

// All allowed values of KeywordUpdateV30KeywordsMatchType enum
var AllowedKeywordUpdateV30KeywordsMatchTypeEnumValues = []KeywordUpdateV30KeywordsMatchType{
	"EXTENSIVE",
	"PHRASE",
	"PRECISION",
}

// NewKeywordUpdateV30KeywordsMatchTypeFromValue returns a pointer to a valid KeywordUpdateV30KeywordsMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeywordUpdateV30KeywordsMatchTypeFromValue(v string) (*KeywordUpdateV30KeywordsMatchType, error) {
	ev := KeywordUpdateV30KeywordsMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeywordUpdateV30KeywordsMatchType: valid values are %v", v, AllowedKeywordUpdateV30KeywordsMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeywordUpdateV30KeywordsMatchType) IsValid() bool {
	for _, existing := range AllowedKeywordUpdateV30KeywordsMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to keyword_update_v3.0_keywords_match_type value
func (v KeywordUpdateV30KeywordsMatchType) Ptr() *KeywordUpdateV30KeywordsMatchType {
	return &v
}
