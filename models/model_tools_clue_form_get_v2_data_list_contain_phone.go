/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsClueFormGetV2DataListContainPhone
type ToolsClueFormGetV2DataListContainPhone int64

// List of tools_clue_form_get_v2_data_list_contain_phone
const (
	Enum_0_ToolsClueFormGetV2DataListContainPhone ToolsClueFormGetV2DataListContainPhone = 0
	Enum_1_ToolsClueFormGetV2DataListContainPhone ToolsClueFormGetV2DataListContainPhone = 1
)

// All allowed values of ToolsClueFormGetV2DataListContainPhone enum
var AllowedToolsClueFormGetV2DataListContainPhoneEnumValues = []ToolsClueFormGetV2DataListContainPhone{
	0,
	1,
}

// NewToolsClueFormGetV2DataListContainPhoneFromValue returns a pointer to a valid ToolsClueFormGetV2DataListContainPhone
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsClueFormGetV2DataListContainPhoneFromValue(v int64) (*ToolsClueFormGetV2DataListContainPhone, error) {
	ev := ToolsClueFormGetV2DataListContainPhone(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsClueFormGetV2DataListContainPhone: valid values are %v", v, AllowedToolsClueFormGetV2DataListContainPhoneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsClueFormGetV2DataListContainPhone) IsValid() bool {
	for _, existing := range AllowedToolsClueFormGetV2DataListContainPhoneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_clue_form_get_v2_data_list_contain_phone value
func (v ToolsClueFormGetV2DataListContainPhone) Ptr() *ToolsClueFormGetV2DataListContainPhone {
	return &v
}
