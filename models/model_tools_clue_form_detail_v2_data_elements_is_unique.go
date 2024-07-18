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

// ToolsClueFormDetailV2DataElementsIsUnique
type ToolsClueFormDetailV2DataElementsIsUnique int64

// List of tools_clue_form_detail_v2_data_elements_is_unique
const (
	Enum_0_ToolsClueFormDetailV2DataElementsIsUnique ToolsClueFormDetailV2DataElementsIsUnique = 0
	Enum_1_ToolsClueFormDetailV2DataElementsIsUnique ToolsClueFormDetailV2DataElementsIsUnique = 1
)

// All allowed values of ToolsClueFormDetailV2DataElementsIsUnique enum
var AllowedToolsClueFormDetailV2DataElementsIsUniqueEnumValues = []ToolsClueFormDetailV2DataElementsIsUnique{
	0,
	1,
}

// NewToolsClueFormDetailV2DataElementsIsUniqueFromValue returns a pointer to a valid ToolsClueFormDetailV2DataElementsIsUnique
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsClueFormDetailV2DataElementsIsUniqueFromValue(v int64) (*ToolsClueFormDetailV2DataElementsIsUnique, error) {
	ev := ToolsClueFormDetailV2DataElementsIsUnique(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsClueFormDetailV2DataElementsIsUnique: valid values are %v", v, AllowedToolsClueFormDetailV2DataElementsIsUniqueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsClueFormDetailV2DataElementsIsUnique) IsValid() bool {
	for _, existing := range AllowedToolsClueFormDetailV2DataElementsIsUniqueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_clue_form_detail_v2_data_elements_is_unique value
func (v ToolsClueFormDetailV2DataElementsIsUnique) Ptr() *ToolsClueFormDetailV2DataElementsIsUnique {
	return &v
}
