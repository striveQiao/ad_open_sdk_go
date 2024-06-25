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

// ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsBelong
type ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsBelong string

// List of tools_ad_convert_optimized_target_get_v2_data_list_converts_external_actions_belong
const (
	ADVANCED_CREATIVE_ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsBelong ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsBelong = "ADVANCED_CREATIVE"
	EXTERNAL_URL_ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsBelong      ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsBelong = "EXTERNAL_URL"
	MICRO_APP_ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsBelong         ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsBelong = "MICRO_APP"
)

// All allowed values of ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsBelong enum
var AllowedToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsBelongEnumValues = []ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsBelong{
	"ADVANCED_CREATIVE",
	"EXTERNAL_URL",
	"MICRO_APP",
}

// NewToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsBelongFromValue returns a pointer to a valid ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsBelong
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsBelongFromValue(v string) (*ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsBelong, error) {
	ev := ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsBelong(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsBelong: valid values are %v", v, AllowedToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsBelongEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsBelong) IsValid() bool {
	for _, existing := range AllowedToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsBelongEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ad_convert_optimized_target_get_v2_data_list_converts_external_actions_belong value
func (v ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsBelong) Ptr() *ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsBelong {
	return &v
}
