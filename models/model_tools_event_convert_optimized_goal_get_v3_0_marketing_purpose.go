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

// ToolsEventConvertOptimizedGoalGetV30MarketingPurpose
type ToolsEventConvertOptimizedGoalGetV30MarketingPurpose string

// List of tools_event_convert_optimized_goal_get_v3.0_marketing_purpose
const (
	ACKNOWLEDGE_ToolsEventConvertOptimizedGoalGetV30MarketingPurpose ToolsEventConvertOptimizedGoalGetV30MarketingPurpose = "ACKNOWLEDGE"
	CONVERSION_ToolsEventConvertOptimizedGoalGetV30MarketingPurpose  ToolsEventConvertOptimizedGoalGetV30MarketingPurpose = "CONVERSION"
	INTENTION_ToolsEventConvertOptimizedGoalGetV30MarketingPurpose   ToolsEventConvertOptimizedGoalGetV30MarketingPurpose = "INTENTION"
)

// All allowed values of ToolsEventConvertOptimizedGoalGetV30MarketingPurpose enum
var AllowedToolsEventConvertOptimizedGoalGetV30MarketingPurposeEnumValues = []ToolsEventConvertOptimizedGoalGetV30MarketingPurpose{
	"ACKNOWLEDGE",
	"CONVERSION",
	"INTENTION",
}

// NewToolsEventConvertOptimizedGoalGetV30MarketingPurposeFromValue returns a pointer to a valid ToolsEventConvertOptimizedGoalGetV30MarketingPurpose
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEventConvertOptimizedGoalGetV30MarketingPurposeFromValue(v string) (*ToolsEventConvertOptimizedGoalGetV30MarketingPurpose, error) {
	ev := ToolsEventConvertOptimizedGoalGetV30MarketingPurpose(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEventConvertOptimizedGoalGetV30MarketingPurpose: valid values are %v", v, AllowedToolsEventConvertOptimizedGoalGetV30MarketingPurposeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEventConvertOptimizedGoalGetV30MarketingPurpose) IsValid() bool {
	for _, existing := range AllowedToolsEventConvertOptimizedGoalGetV30MarketingPurposeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_event_convert_optimized_goal_get_v3.0_marketing_purpose value
func (v ToolsEventConvertOptimizedGoalGetV30MarketingPurpose) Ptr() *ToolsEventConvertOptimizedGoalGetV30MarketingPurpose {
	return &v
}
