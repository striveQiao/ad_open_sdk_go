/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBidSuggestV2ConvertedTimeDuration
type ToolsBidSuggestV2ConvertedTimeDuration string

// List of tools_bid_suggest_v2_converted_time_duration
const (
	THREE_MONTH_ToolsBidSuggestV2ConvertedTimeDuration  ToolsBidSuggestV2ConvertedTimeDuration = "THREE_MONTH"
	NONE_ToolsBidSuggestV2ConvertedTimeDuration         ToolsBidSuggestV2ConvertedTimeDuration = "NONE"
	SIX_MONTH_ToolsBidSuggestV2ConvertedTimeDuration    ToolsBidSuggestV2ConvertedTimeDuration = "SIX_MONTH"
	SEVEN_DAY_ToolsBidSuggestV2ConvertedTimeDuration    ToolsBidSuggestV2ConvertedTimeDuration = "SEVEN_DAY"
	ONE_MONTH_ToolsBidSuggestV2ConvertedTimeDuration    ToolsBidSuggestV2ConvertedTimeDuration = "ONE_MONTH"
	TWELVE_MONTH_ToolsBidSuggestV2ConvertedTimeDuration ToolsBidSuggestV2ConvertedTimeDuration = "TWELVE_MONTH"
	TODAY_ToolsBidSuggestV2ConvertedTimeDuration        ToolsBidSuggestV2ConvertedTimeDuration = "TODAY"
)

// All allowed values of ToolsBidSuggestV2ConvertedTimeDuration enum
var AllowedToolsBidSuggestV2ConvertedTimeDurationEnumValues = []ToolsBidSuggestV2ConvertedTimeDuration{
	"THREE_MONTH",
	"NONE",
	"SIX_MONTH",
	"SEVEN_DAY",
	"ONE_MONTH",
	"TWELVE_MONTH",
	"TODAY",
}

// NewToolsBidSuggestV2ConvertedTimeDurationFromValue returns a pointer to a valid ToolsBidSuggestV2ConvertedTimeDuration
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2ConvertedTimeDurationFromValue(v string) (*ToolsBidSuggestV2ConvertedTimeDuration, error) {
	ev := ToolsBidSuggestV2ConvertedTimeDuration(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2ConvertedTimeDuration: valid values are %v", v, AllowedToolsBidSuggestV2ConvertedTimeDurationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2ConvertedTimeDuration) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2ConvertedTimeDurationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_converted_time_duration value
func (v ToolsBidSuggestV2ConvertedTimeDuration) Ptr() *ToolsBidSuggestV2ConvertedTimeDuration {
	return &v
}
