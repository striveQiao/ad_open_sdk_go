/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsEstimateAudienceV2DpaLocalAudience
type ToolsEstimateAudienceV2DpaLocalAudience int64

// List of tools_estimate_audience_v2_dpa_local_audience
const (
	Enum_0_ToolsEstimateAudienceV2DpaLocalAudience ToolsEstimateAudienceV2DpaLocalAudience = 0
	Enum_1_ToolsEstimateAudienceV2DpaLocalAudience ToolsEstimateAudienceV2DpaLocalAudience = 1
)

// All allowed values of ToolsEstimateAudienceV2DpaLocalAudience enum
var AllowedToolsEstimateAudienceV2DpaLocalAudienceEnumValues = []ToolsEstimateAudienceV2DpaLocalAudience{
	0,
	1,
}

// NewToolsEstimateAudienceV2DpaLocalAudienceFromValue returns a pointer to a valid ToolsEstimateAudienceV2DpaLocalAudience
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEstimateAudienceV2DpaLocalAudienceFromValue(v int64) (*ToolsEstimateAudienceV2DpaLocalAudience, error) {
	ev := ToolsEstimateAudienceV2DpaLocalAudience(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEstimateAudienceV2DpaLocalAudience: valid values are %v", v, AllowedToolsEstimateAudienceV2DpaLocalAudienceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEstimateAudienceV2DpaLocalAudience) IsValid() bool {
	for _, existing := range AllowedToolsEstimateAudienceV2DpaLocalAudienceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_estimate_audience_v2_dpa_local_audience value
func (v ToolsEstimateAudienceV2DpaLocalAudience) Ptr() *ToolsEstimateAudienceV2DpaLocalAudience {
	return &v
}
