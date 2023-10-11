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

// KeywordUpdateV30DataSuccessListMatchType
type KeywordUpdateV30DataSuccessListMatchType string

// List of keyword_update_v3.0_data_success_list_match_type
const (
	EXTENSIVE_KeywordUpdateV30DataSuccessListMatchType KeywordUpdateV30DataSuccessListMatchType = "EXTENSIVE"
	PHRASE_KeywordUpdateV30DataSuccessListMatchType    KeywordUpdateV30DataSuccessListMatchType = "PHRASE"
	PRECISION_KeywordUpdateV30DataSuccessListMatchType KeywordUpdateV30DataSuccessListMatchType = "PRECISION"
)

// All allowed values of KeywordUpdateV30DataSuccessListMatchType enum
var AllowedKeywordUpdateV30DataSuccessListMatchTypeEnumValues = []KeywordUpdateV30DataSuccessListMatchType{
	"EXTENSIVE",
	"PHRASE",
	"PRECISION",
}

// NewKeywordUpdateV30DataSuccessListMatchTypeFromValue returns a pointer to a valid KeywordUpdateV30DataSuccessListMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeywordUpdateV30DataSuccessListMatchTypeFromValue(v string) (*KeywordUpdateV30DataSuccessListMatchType, error) {
	ev := KeywordUpdateV30DataSuccessListMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeywordUpdateV30DataSuccessListMatchType: valid values are %v", v, AllowedKeywordUpdateV30DataSuccessListMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeywordUpdateV30DataSuccessListMatchType) IsValid() bool {
	for _, existing := range AllowedKeywordUpdateV30DataSuccessListMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to keyword_update_v3.0_data_success_list_match_type value
func (v KeywordUpdateV30DataSuccessListMatchType) Ptr() *KeywordUpdateV30DataSuccessListMatchType {
	return &v
}
