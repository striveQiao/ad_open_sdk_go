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

// ToolsEstimateAudienceV2AutoExtendEnabled
type ToolsEstimateAudienceV2AutoExtendEnabled int64

// List of tools_estimate_audience_v2_auto_extend_enabled
const (
	Enum_0_ToolsEstimateAudienceV2AutoExtendEnabled ToolsEstimateAudienceV2AutoExtendEnabled = 0
	Enum_1_ToolsEstimateAudienceV2AutoExtendEnabled ToolsEstimateAudienceV2AutoExtendEnabled = 1
)

// All allowed values of ToolsEstimateAudienceV2AutoExtendEnabled enum
var AllowedToolsEstimateAudienceV2AutoExtendEnabledEnumValues = []ToolsEstimateAudienceV2AutoExtendEnabled{
	0,
	1,
}

// NewToolsEstimateAudienceV2AutoExtendEnabledFromValue returns a pointer to a valid ToolsEstimateAudienceV2AutoExtendEnabled
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEstimateAudienceV2AutoExtendEnabledFromValue(v int64) (*ToolsEstimateAudienceV2AutoExtendEnabled, error) {
	ev := ToolsEstimateAudienceV2AutoExtendEnabled(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEstimateAudienceV2AutoExtendEnabled: valid values are %v", v, AllowedToolsEstimateAudienceV2AutoExtendEnabledEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEstimateAudienceV2AutoExtendEnabled) IsValid() bool {
	for _, existing := range AllowedToolsEstimateAudienceV2AutoExtendEnabledEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_estimate_audience_v2_auto_extend_enabled value
func (v ToolsEstimateAudienceV2AutoExtendEnabled) Ptr() *ToolsEstimateAudienceV2AutoExtendEnabled {
	return &v
}
