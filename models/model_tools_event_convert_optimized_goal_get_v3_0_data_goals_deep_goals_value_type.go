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

// ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType
type ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType string

// List of tools_event_convert_optimized_goal_get_v3.0_data_goals_deepGoals_value_type
const (
	DISABLED_ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType              ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType = "Disabled"
	DISCRIMINATE_BY_GROUP_ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType = "DiscriminateByGroup"
	DYNAMIC_VALUE_ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType         ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType = "DynamicValue"
	FIXED_ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType                 ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType = "Fixed"
)

// All allowed values of ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType enum
var AllowedToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueTypeEnumValues = []ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType{
	"Disabled",
	"DiscriminateByGroup",
	"DynamicValue",
	"Fixed",
}

// NewToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueTypeFromValue returns a pointer to a valid ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueTypeFromValue(v string) (*ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType, error) {
	ev := ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType: valid values are %v", v, AllowedToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType) IsValid() bool {
	for _, existing := range AllowedToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_event_convert_optimized_goal_get_v3.0_data_goals_deepGoals_value_type value
func (v ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType) Ptr() *ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType {
	return &v
}
