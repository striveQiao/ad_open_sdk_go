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

// ToolsAdConvertOptimizedTargetGetV2MarketingPurpose
type ToolsAdConvertOptimizedTargetGetV2MarketingPurpose string

// List of tools_ad_convert_optimized_target_get_v2_marketing_purpose
const (
	INTENTION_ToolsAdConvertOptimizedTargetGetV2MarketingPurpose   ToolsAdConvertOptimizedTargetGetV2MarketingPurpose = "INTENTION"
	ACKNOWLEDGE_ToolsAdConvertOptimizedTargetGetV2MarketingPurpose ToolsAdConvertOptimizedTargetGetV2MarketingPurpose = "ACKNOWLEDGE"
	CONVERSION_ToolsAdConvertOptimizedTargetGetV2MarketingPurpose  ToolsAdConvertOptimizedTargetGetV2MarketingPurpose = "CONVERSION"
	UNLIMITED_ToolsAdConvertOptimizedTargetGetV2MarketingPurpose   ToolsAdConvertOptimizedTargetGetV2MarketingPurpose = "UNLIMITED"
)

// All allowed values of ToolsAdConvertOptimizedTargetGetV2MarketingPurpose enum
var AllowedToolsAdConvertOptimizedTargetGetV2MarketingPurposeEnumValues = []ToolsAdConvertOptimizedTargetGetV2MarketingPurpose{
	"INTENTION",
	"ACKNOWLEDGE",
	"CONVERSION",
	"UNLIMITED",
}

// NewToolsAdConvertOptimizedTargetGetV2MarketingPurposeFromValue returns a pointer to a valid ToolsAdConvertOptimizedTargetGetV2MarketingPurpose
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdConvertOptimizedTargetGetV2MarketingPurposeFromValue(v string) (*ToolsAdConvertOptimizedTargetGetV2MarketingPurpose, error) {
	ev := ToolsAdConvertOptimizedTargetGetV2MarketingPurpose(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdConvertOptimizedTargetGetV2MarketingPurpose: valid values are %v", v, AllowedToolsAdConvertOptimizedTargetGetV2MarketingPurposeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdConvertOptimizedTargetGetV2MarketingPurpose) IsValid() bool {
	for _, existing := range AllowedToolsAdConvertOptimizedTargetGetV2MarketingPurposeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ad_convert_optimized_target_get_v2_marketing_purpose value
func (v ToolsAdConvertOptimizedTargetGetV2MarketingPurpose) Ptr() *ToolsAdConvertOptimizedTargetGetV2MarketingPurpose {
	return &v
}
