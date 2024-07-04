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

// ClueFormCreateV2ExtendInfoCountConfigCountType
type ClueFormCreateV2ExtendInfoCountConfigCountType string

// List of clue_form_create_v2_extend_info_count_config_count_type
const (
	COUNT_TYPE_DECREMENT_ClueFormCreateV2ExtendInfoCountConfigCountType ClueFormCreateV2ExtendInfoCountConfigCountType = "COUNT_TYPE_DECREMENT"
	COUNT_TYPE_INCREMENT_ClueFormCreateV2ExtendInfoCountConfigCountType ClueFormCreateV2ExtendInfoCountConfigCountType = "COUNT_TYPE_INCREMENT"
)

// All allowed values of ClueFormCreateV2ExtendInfoCountConfigCountType enum
var AllowedClueFormCreateV2ExtendInfoCountConfigCountTypeEnumValues = []ClueFormCreateV2ExtendInfoCountConfigCountType{
	"COUNT_TYPE_DECREMENT",
	"COUNT_TYPE_INCREMENT",
}

// NewClueFormCreateV2ExtendInfoCountConfigCountTypeFromValue returns a pointer to a valid ClueFormCreateV2ExtendInfoCountConfigCountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueFormCreateV2ExtendInfoCountConfigCountTypeFromValue(v string) (*ClueFormCreateV2ExtendInfoCountConfigCountType, error) {
	ev := ClueFormCreateV2ExtendInfoCountConfigCountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueFormCreateV2ExtendInfoCountConfigCountType: valid values are %v", v, AllowedClueFormCreateV2ExtendInfoCountConfigCountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueFormCreateV2ExtendInfoCountConfigCountType) IsValid() bool {
	for _, existing := range AllowedClueFormCreateV2ExtendInfoCountConfigCountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_form_create_v2_extend_info_count_config_count_type value
func (v ClueFormCreateV2ExtendInfoCountConfigCountType) Ptr() *ClueFormCreateV2ExtendInfoCountConfigCountType {
	return &v
}
