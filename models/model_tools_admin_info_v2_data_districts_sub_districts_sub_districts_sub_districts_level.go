/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsLevel
type ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsLevel string

// List of tools_admin_info_v2_data_districts_sub_districts_sub_districts_sub_districts_level
const (
	FOUR_LEVEL_ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsLevel  ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsLevel = "FOUR_LEVEL"
	NONE_ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsLevel        ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsLevel = "NONE"
	ONE_LEVEL_ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsLevel   ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsLevel = "ONE_LEVEL"
	THREE_LEVEL_ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsLevel ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsLevel = "THREE_LEVEL"
	TWO_LEVEL_ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsLevel   ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsLevel = "TWO_LEVEL"
)

// All allowed values of ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsLevel enum
var AllowedToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsLevelEnumValues = []ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsLevel{
	"FOUR_LEVEL",
	"NONE",
	"ONE_LEVEL",
	"THREE_LEVEL",
	"TWO_LEVEL",
}

// NewToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsLevelFromValue returns a pointer to a valid ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsLevelFromValue(v string) (*ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsLevel, error) {
	ev := ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsLevel: valid values are %v", v, AllowedToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsLevel) IsValid() bool {
	for _, existing := range AllowedToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_admin_info_v2_data_districts_sub_districts_sub_districts_sub_districts_level value
func (v ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsLevel) Ptr() *ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsLevel {
	return &v
}
