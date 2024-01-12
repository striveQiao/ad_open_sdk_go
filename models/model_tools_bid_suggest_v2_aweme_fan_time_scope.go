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

// ToolsBidSuggestV2AwemeFanTimeScope
type ToolsBidSuggestV2AwemeFanTimeScope string

// List of tools_bid_suggest_v2_aweme_fan_time_scope
const (
	FIFTEEN_DAYS_ToolsBidSuggestV2AwemeFanTimeScope ToolsBidSuggestV2AwemeFanTimeScope = "FIFTEEN_DAYS"
	THIRTY_DAYS_ToolsBidSuggestV2AwemeFanTimeScope  ToolsBidSuggestV2AwemeFanTimeScope = "THIRTY_DAYS"
	SIXTY_DAYS_ToolsBidSuggestV2AwemeFanTimeScope   ToolsBidSuggestV2AwemeFanTimeScope = "SIXTY_DAYS"
)

// All allowed values of ToolsBidSuggestV2AwemeFanTimeScope enum
var AllowedToolsBidSuggestV2AwemeFanTimeScopeEnumValues = []ToolsBidSuggestV2AwemeFanTimeScope{
	"FIFTEEN_DAYS",
	"THIRTY_DAYS",
	"SIXTY_DAYS",
}

// NewToolsBidSuggestV2AwemeFanTimeScopeFromValue returns a pointer to a valid ToolsBidSuggestV2AwemeFanTimeScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2AwemeFanTimeScopeFromValue(v string) (*ToolsBidSuggestV2AwemeFanTimeScope, error) {
	ev := ToolsBidSuggestV2AwemeFanTimeScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2AwemeFanTimeScope: valid values are %v", v, AllowedToolsBidSuggestV2AwemeFanTimeScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2AwemeFanTimeScope) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2AwemeFanTimeScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_aweme_fan_time_scope value
func (v ToolsBidSuggestV2AwemeFanTimeScope) Ptr() *ToolsBidSuggestV2AwemeFanTimeScope {
	return &v
}
