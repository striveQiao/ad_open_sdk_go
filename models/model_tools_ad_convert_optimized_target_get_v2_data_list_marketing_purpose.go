/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAdConvertOptimizedTargetGetV2DataListMarketingPurpose
type ToolsAdConvertOptimizedTargetGetV2DataListMarketingPurpose string

// List of tools_ad_convert_optimized_target_get_v2_data_list_marketing_purpose
const (
	CONVERSION_ToolsAdConvertOptimizedTargetGetV2DataListMarketingPurpose  ToolsAdConvertOptimizedTargetGetV2DataListMarketingPurpose = "CONVERSION"
	UNLIMITED_ToolsAdConvertOptimizedTargetGetV2DataListMarketingPurpose   ToolsAdConvertOptimizedTargetGetV2DataListMarketingPurpose = "UNLIMITED"
	INTENTION_ToolsAdConvertOptimizedTargetGetV2DataListMarketingPurpose   ToolsAdConvertOptimizedTargetGetV2DataListMarketingPurpose = "INTENTION"
	ACKNOWLEDGE_ToolsAdConvertOptimizedTargetGetV2DataListMarketingPurpose ToolsAdConvertOptimizedTargetGetV2DataListMarketingPurpose = "ACKNOWLEDGE"
)

// All allowed values of ToolsAdConvertOptimizedTargetGetV2DataListMarketingPurpose enum
var AllowedToolsAdConvertOptimizedTargetGetV2DataListMarketingPurposeEnumValues = []ToolsAdConvertOptimizedTargetGetV2DataListMarketingPurpose{
	"CONVERSION",
	"UNLIMITED",
	"INTENTION",
	"ACKNOWLEDGE",
}

// NewToolsAdConvertOptimizedTargetGetV2DataListMarketingPurposeFromValue returns a pointer to a valid ToolsAdConvertOptimizedTargetGetV2DataListMarketingPurpose
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdConvertOptimizedTargetGetV2DataListMarketingPurposeFromValue(v string) (*ToolsAdConvertOptimizedTargetGetV2DataListMarketingPurpose, error) {
	ev := ToolsAdConvertOptimizedTargetGetV2DataListMarketingPurpose(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdConvertOptimizedTargetGetV2DataListMarketingPurpose: valid values are %v", v, AllowedToolsAdConvertOptimizedTargetGetV2DataListMarketingPurposeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdConvertOptimizedTargetGetV2DataListMarketingPurpose) IsValid() bool {
	for _, existing := range AllowedToolsAdConvertOptimizedTargetGetV2DataListMarketingPurposeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ad_convert_optimized_target_get_v2_data_list_marketing_purpose value
func (v ToolsAdConvertOptimizedTargetGetV2DataListMarketingPurpose) Ptr() *ToolsAdConvertOptimizedTargetGetV2DataListMarketingPurpose {
	return &v
}
