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

// ToolsClueFormGetV2DataListSubType
type ToolsClueFormGetV2DataListSubType string

// List of tools_clue_form_get_v2_data_list_sub_type
const (
	ADVANCED_CREATIVE_FORM_ToolsClueFormGetV2DataListSubType ToolsClueFormGetV2DataListSubType = "ADVANCED_CREATIVE_FORM"
	NATIVE_FORM_ToolsClueFormGetV2DataListSubType            ToolsClueFormGetV2DataListSubType = "NATIVE_FORM"
	NORMAL_FORM_ToolsClueFormGetV2DataListSubType            ToolsClueFormGetV2DataListSubType = "NORMAL_FORM"
)

// All allowed values of ToolsClueFormGetV2DataListSubType enum
var AllowedToolsClueFormGetV2DataListSubTypeEnumValues = []ToolsClueFormGetV2DataListSubType{
	"ADVANCED_CREATIVE_FORM",
	"NATIVE_FORM",
	"NORMAL_FORM",
}

// NewToolsClueFormGetV2DataListSubTypeFromValue returns a pointer to a valid ToolsClueFormGetV2DataListSubType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsClueFormGetV2DataListSubTypeFromValue(v string) (*ToolsClueFormGetV2DataListSubType, error) {
	ev := ToolsClueFormGetV2DataListSubType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsClueFormGetV2DataListSubType: valid values are %v", v, AllowedToolsClueFormGetV2DataListSubTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsClueFormGetV2DataListSubType) IsValid() bool {
	for _, existing := range AllowedToolsClueFormGetV2DataListSubTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_clue_form_get_v2_data_list_sub_type value
func (v ToolsClueFormGetV2DataListSubType) Ptr() *ToolsClueFormGetV2DataListSubType {
	return &v
}
