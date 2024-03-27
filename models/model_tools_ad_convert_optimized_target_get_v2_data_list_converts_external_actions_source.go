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

// ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsSource
type ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsSource string

// List of tools_ad_convert_optimized_target_get_v2_data_list_converts_external_actions_source
const (
	SOURCE_CONFIG_ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsSource ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsSource = "SOURCE_CONFIG"
	SOURCE_TETRIS_ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsSource ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsSource = "SOURCE_TETRIS"
	SOURCE_METEOR_ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsSource ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsSource = "SOURCE_METEOR"
)

// All allowed values of ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsSource enum
var AllowedToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsSourceEnumValues = []ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsSource{
	"SOURCE_CONFIG",
	"SOURCE_TETRIS",
	"SOURCE_METEOR",
}

// NewToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsSourceFromValue returns a pointer to a valid ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsSourceFromValue(v string) (*ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsSource, error) {
	ev := ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsSource: valid values are %v", v, AllowedToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsSource) IsValid() bool {
	for _, existing := range AllowedToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ad_convert_optimized_target_get_v2_data_list_converts_external_actions_source value
func (v ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsSource) Ptr() *ToolsAdConvertOptimizedTargetGetV2DataListConvertsExternalActionsSource {
	return &v
}
