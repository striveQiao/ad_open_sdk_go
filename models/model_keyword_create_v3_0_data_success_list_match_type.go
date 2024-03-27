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

// KeywordCreateV30DataSuccessListMatchType
type KeywordCreateV30DataSuccessListMatchType string

// List of keyword_create_v3.0_data_success_list_match_type
const (
	EXTENSIVE_KeywordCreateV30DataSuccessListMatchType KeywordCreateV30DataSuccessListMatchType = "EXTENSIVE"
	PHRASE_KeywordCreateV30DataSuccessListMatchType    KeywordCreateV30DataSuccessListMatchType = "PHRASE"
	PRECISION_KeywordCreateV30DataSuccessListMatchType KeywordCreateV30DataSuccessListMatchType = "PRECISION"
)

// All allowed values of KeywordCreateV30DataSuccessListMatchType enum
var AllowedKeywordCreateV30DataSuccessListMatchTypeEnumValues = []KeywordCreateV30DataSuccessListMatchType{
	"EXTENSIVE",
	"PHRASE",
	"PRECISION",
}

// NewKeywordCreateV30DataSuccessListMatchTypeFromValue returns a pointer to a valid KeywordCreateV30DataSuccessListMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeywordCreateV30DataSuccessListMatchTypeFromValue(v string) (*KeywordCreateV30DataSuccessListMatchType, error) {
	ev := KeywordCreateV30DataSuccessListMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeywordCreateV30DataSuccessListMatchType: valid values are %v", v, AllowedKeywordCreateV30DataSuccessListMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeywordCreateV30DataSuccessListMatchType) IsValid() bool {
	for _, existing := range AllowedKeywordCreateV30DataSuccessListMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to keyword_create_v3.0_data_success_list_match_type value
func (v KeywordCreateV30DataSuccessListMatchType) Ptr() *KeywordCreateV30DataSuccessListMatchType {
	return &v
}
