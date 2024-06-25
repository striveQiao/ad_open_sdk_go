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

// ToolsEstimateAudienceV2InterestActionMode
type ToolsEstimateAudienceV2InterestActionMode string

// List of tools_estimate_audience_v2_interest_action_mode
const (
	CUSTOM_ToolsEstimateAudienceV2InterestActionMode    ToolsEstimateAudienceV2InterestActionMode = "CUSTOM"
	RECOMMEND_ToolsEstimateAudienceV2InterestActionMode ToolsEstimateAudienceV2InterestActionMode = "RECOMMEND"
	UNLIMITED_ToolsEstimateAudienceV2InterestActionMode ToolsEstimateAudienceV2InterestActionMode = "UNLIMITED"
)

// All allowed values of ToolsEstimateAudienceV2InterestActionMode enum
var AllowedToolsEstimateAudienceV2InterestActionModeEnumValues = []ToolsEstimateAudienceV2InterestActionMode{
	"CUSTOM",
	"RECOMMEND",
	"UNLIMITED",
}

// NewToolsEstimateAudienceV2InterestActionModeFromValue returns a pointer to a valid ToolsEstimateAudienceV2InterestActionMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEstimateAudienceV2InterestActionModeFromValue(v string) (*ToolsEstimateAudienceV2InterestActionMode, error) {
	ev := ToolsEstimateAudienceV2InterestActionMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEstimateAudienceV2InterestActionMode: valid values are %v", v, AllowedToolsEstimateAudienceV2InterestActionModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEstimateAudienceV2InterestActionMode) IsValid() bool {
	for _, existing := range AllowedToolsEstimateAudienceV2InterestActionModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_estimate_audience_v2_interest_action_mode value
func (v ToolsEstimateAudienceV2InterestActionMode) Ptr() *ToolsEstimateAudienceV2InterestActionMode {
	return &v
}
