/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsClueFormDetailV2DataElementsAllowEmpty
type ToolsClueFormDetailV2DataElementsAllowEmpty int64

// List of tools_clue_form_detail_v2_data_elements_allow_empty
const (
	Enum_0_ToolsClueFormDetailV2DataElementsAllowEmpty ToolsClueFormDetailV2DataElementsAllowEmpty = 0
	Enum_1_ToolsClueFormDetailV2DataElementsAllowEmpty ToolsClueFormDetailV2DataElementsAllowEmpty = 1
)

// All allowed values of ToolsClueFormDetailV2DataElementsAllowEmpty enum
var AllowedToolsClueFormDetailV2DataElementsAllowEmptyEnumValues = []ToolsClueFormDetailV2DataElementsAllowEmpty{
	0,
	1,
}

// NewToolsClueFormDetailV2DataElementsAllowEmptyFromValue returns a pointer to a valid ToolsClueFormDetailV2DataElementsAllowEmpty
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsClueFormDetailV2DataElementsAllowEmptyFromValue(v int64) (*ToolsClueFormDetailV2DataElementsAllowEmpty, error) {
	ev := ToolsClueFormDetailV2DataElementsAllowEmpty(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsClueFormDetailV2DataElementsAllowEmpty: valid values are %v", v, AllowedToolsClueFormDetailV2DataElementsAllowEmptyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsClueFormDetailV2DataElementsAllowEmpty) IsValid() bool {
	for _, existing := range AllowedToolsClueFormDetailV2DataElementsAllowEmptyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_clue_form_detail_v2_data_elements_allow_empty value
func (v ToolsClueFormDetailV2DataElementsAllowEmpty) Ptr() *ToolsClueFormDetailV2DataElementsAllowEmpty {
	return &v
}
