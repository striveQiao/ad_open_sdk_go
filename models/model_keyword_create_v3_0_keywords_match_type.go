/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// KeywordCreateV30KeywordsMatchType
type KeywordCreateV30KeywordsMatchType string

// List of keyword_create_v3.0_keywords_match_type
const (
	EXTENSIVE_KeywordCreateV30KeywordsMatchType KeywordCreateV30KeywordsMatchType = "EXTENSIVE"
	PHRASE_KeywordCreateV30KeywordsMatchType    KeywordCreateV30KeywordsMatchType = "PHRASE"
	PRECISION_KeywordCreateV30KeywordsMatchType KeywordCreateV30KeywordsMatchType = "PRECISION"
)

// All allowed values of KeywordCreateV30KeywordsMatchType enum
var AllowedKeywordCreateV30KeywordsMatchTypeEnumValues = []KeywordCreateV30KeywordsMatchType{
	"EXTENSIVE",
	"PHRASE",
	"PRECISION",
}

// NewKeywordCreateV30KeywordsMatchTypeFromValue returns a pointer to a valid KeywordCreateV30KeywordsMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeywordCreateV30KeywordsMatchTypeFromValue(v string) (*KeywordCreateV30KeywordsMatchType, error) {
	ev := KeywordCreateV30KeywordsMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeywordCreateV30KeywordsMatchType: valid values are %v", v, AllowedKeywordCreateV30KeywordsMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeywordCreateV30KeywordsMatchType) IsValid() bool {
	for _, existing := range AllowedKeywordCreateV30KeywordsMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to keyword_create_v3.0_keywords_match_type value
func (v KeywordCreateV30KeywordsMatchType) Ptr() *KeywordCreateV30KeywordsMatchType {
	return &v
}
