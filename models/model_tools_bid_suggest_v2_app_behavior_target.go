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

// ToolsBidSuggestV2AppBehaviorTarget
type ToolsBidSuggestV2AppBehaviorTarget string

// List of tools_bid_suggest_v2_app_behavior_target
const (
	CATEGORY_ToolsBidSuggestV2AppBehaviorTarget ToolsBidSuggestV2AppBehaviorTarget = "CATEGORY"
	APP_ToolsBidSuggestV2AppBehaviorTarget      ToolsBidSuggestV2AppBehaviorTarget = "APP"
	NONE_ToolsBidSuggestV2AppBehaviorTarget     ToolsBidSuggestV2AppBehaviorTarget = "NONE"
)

// All allowed values of ToolsBidSuggestV2AppBehaviorTarget enum
var AllowedToolsBidSuggestV2AppBehaviorTargetEnumValues = []ToolsBidSuggestV2AppBehaviorTarget{
	"CATEGORY",
	"APP",
	"NONE",
}

// NewToolsBidSuggestV2AppBehaviorTargetFromValue returns a pointer to a valid ToolsBidSuggestV2AppBehaviorTarget
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2AppBehaviorTargetFromValue(v string) (*ToolsBidSuggestV2AppBehaviorTarget, error) {
	ev := ToolsBidSuggestV2AppBehaviorTarget(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2AppBehaviorTarget: valid values are %v", v, AllowedToolsBidSuggestV2AppBehaviorTargetEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2AppBehaviorTarget) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2AppBehaviorTargetEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_app_behavior_target value
func (v ToolsBidSuggestV2AppBehaviorTarget) Ptr() *ToolsBidSuggestV2AppBehaviorTarget {
	return &v
}
