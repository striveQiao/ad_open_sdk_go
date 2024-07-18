/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsIndustryGetV2Level
type ToolsIndustryGetV2Level int64

// List of tools_industry_get_v2_level
const (
	Enum_1_ToolsIndustryGetV2Level ToolsIndustryGetV2Level = 1
	Enum_2_ToolsIndustryGetV2Level ToolsIndustryGetV2Level = 2
	Enum_3_ToolsIndustryGetV2Level ToolsIndustryGetV2Level = 3
)

// All allowed values of ToolsIndustryGetV2Level enum
var AllowedToolsIndustryGetV2LevelEnumValues = []ToolsIndustryGetV2Level{
	1,
	2,
	3,
}

// NewToolsIndustryGetV2LevelFromValue returns a pointer to a valid ToolsIndustryGetV2Level
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsIndustryGetV2LevelFromValue(v int64) (*ToolsIndustryGetV2Level, error) {
	ev := ToolsIndustryGetV2Level(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsIndustryGetV2Level: valid values are %v", v, AllowedToolsIndustryGetV2LevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsIndustryGetV2Level) IsValid() bool {
	for _, existing := range AllowedToolsIndustryGetV2LevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_industry_get_v2_level value
func (v ToolsIndustryGetV2Level) Ptr() *ToolsIndustryGetV2Level {
	return &v
}
