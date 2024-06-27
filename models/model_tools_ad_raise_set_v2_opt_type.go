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

// ToolsAdRaiseSetV2OptType
type ToolsAdRaiseSetV2OptType string

// List of tools_ad_raise_set_v2_opt_type
const (
	CLICK_RAISE_ToolsAdRaiseSetV2OptType ToolsAdRaiseSetV2OptType = "CLICK_RAISE"
	STOP_RAISE_ToolsAdRaiseSetV2OptType  ToolsAdRaiseSetV2OptType = "STOP_RAISE"
)

// All allowed values of ToolsAdRaiseSetV2OptType enum
var AllowedToolsAdRaiseSetV2OptTypeEnumValues = []ToolsAdRaiseSetV2OptType{
	"CLICK_RAISE",
	"STOP_RAISE",
}

// NewToolsAdRaiseSetV2OptTypeFromValue returns a pointer to a valid ToolsAdRaiseSetV2OptType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdRaiseSetV2OptTypeFromValue(v string) (*ToolsAdRaiseSetV2OptType, error) {
	ev := ToolsAdRaiseSetV2OptType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdRaiseSetV2OptType: valid values are %v", v, AllowedToolsAdRaiseSetV2OptTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdRaiseSetV2OptType) IsValid() bool {
	for _, existing := range AllowedToolsAdRaiseSetV2OptTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ad_raise_set_v2_opt_type value
func (v ToolsAdRaiseSetV2OptType) Ptr() *ToolsAdRaiseSetV2OptType {
	return &v
}
