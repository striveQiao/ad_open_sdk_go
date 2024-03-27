/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsEventConvertOptimizedGoalGetV30LandingType
type ToolsEventConvertOptimizedGoalGetV30LandingType string

// List of tools_event_convert_optimized_goal_get_v3.0_landing_type
const (
	LINK_ToolsEventConvertOptimizedGoalGetV30LandingType ToolsEventConvertOptimizedGoalGetV30LandingType = "LINK"
)

// All allowed values of ToolsEventConvertOptimizedGoalGetV30LandingType enum
var AllowedToolsEventConvertOptimizedGoalGetV30LandingTypeEnumValues = []ToolsEventConvertOptimizedGoalGetV30LandingType{
	"LINK",
}

// NewToolsEventConvertOptimizedGoalGetV30LandingTypeFromValue returns a pointer to a valid ToolsEventConvertOptimizedGoalGetV30LandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEventConvertOptimizedGoalGetV30LandingTypeFromValue(v string) (*ToolsEventConvertOptimizedGoalGetV30LandingType, error) {
	ev := ToolsEventConvertOptimizedGoalGetV30LandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEventConvertOptimizedGoalGetV30LandingType: valid values are %v", v, AllowedToolsEventConvertOptimizedGoalGetV30LandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEventConvertOptimizedGoalGetV30LandingType) IsValid() bool {
	for _, existing := range AllowedToolsEventConvertOptimizedGoalGetV30LandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_event_convert_optimized_goal_get_v3.0_landing_type value
func (v ToolsEventConvertOptimizedGoalGetV30LandingType) Ptr() *ToolsEventConvertOptimizedGoalGetV30LandingType {
	return &v
}
