/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBidSuggestV2DpaRtaSwitch
type ToolsBidSuggestV2DpaRtaSwitch string

// List of tools_bid_suggest_v2_dpa_rta_switch
const (
	ON_ToolsBidSuggestV2DpaRtaSwitch  ToolsBidSuggestV2DpaRtaSwitch = "ON"
	OFF_ToolsBidSuggestV2DpaRtaSwitch ToolsBidSuggestV2DpaRtaSwitch = "OFF"
)

// All allowed values of ToolsBidSuggestV2DpaRtaSwitch enum
var AllowedToolsBidSuggestV2DpaRtaSwitchEnumValues = []ToolsBidSuggestV2DpaRtaSwitch{
	"ON",
	"OFF",
}

// NewToolsBidSuggestV2DpaRtaSwitchFromValue returns a pointer to a valid ToolsBidSuggestV2DpaRtaSwitch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2DpaRtaSwitchFromValue(v string) (*ToolsBidSuggestV2DpaRtaSwitch, error) {
	ev := ToolsBidSuggestV2DpaRtaSwitch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2DpaRtaSwitch: valid values are %v", v, AllowedToolsBidSuggestV2DpaRtaSwitchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2DpaRtaSwitch) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2DpaRtaSwitchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_dpa_rta_switch value
func (v ToolsBidSuggestV2DpaRtaSwitch) Ptr() *ToolsBidSuggestV2DpaRtaSwitch {
	return &v
}
