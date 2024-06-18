/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueFormCreateV2ElementsAllowEmpty
type ClueFormCreateV2ElementsAllowEmpty string

// List of clue_form_create_v2_elements_allow_empty
const (
	Enum_0_ClueFormCreateV2ElementsAllowEmpty ClueFormCreateV2ElementsAllowEmpty = "0"
	Enum_1_ClueFormCreateV2ElementsAllowEmpty ClueFormCreateV2ElementsAllowEmpty = "1"
)

// All allowed values of ClueFormCreateV2ElementsAllowEmpty enum
var AllowedClueFormCreateV2ElementsAllowEmptyEnumValues = []ClueFormCreateV2ElementsAllowEmpty{
	"0",
	"1",
}

// NewClueFormCreateV2ElementsAllowEmptyFromValue returns a pointer to a valid ClueFormCreateV2ElementsAllowEmpty
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueFormCreateV2ElementsAllowEmptyFromValue(v string) (*ClueFormCreateV2ElementsAllowEmpty, error) {
	ev := ClueFormCreateV2ElementsAllowEmpty(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueFormCreateV2ElementsAllowEmpty: valid values are %v", v, AllowedClueFormCreateV2ElementsAllowEmptyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueFormCreateV2ElementsAllowEmpty) IsValid() bool {
	for _, existing := range AllowedClueFormCreateV2ElementsAllowEmptyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_form_create_v2_elements_allow_empty value
func (v ClueFormCreateV2ElementsAllowEmpty) Ptr() *ClueFormCreateV2ElementsAllowEmpty {
	return &v
}
