/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsLevel
type ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsLevel string

// List of tools_admin_info_v2_data_districts_sub_districts_sub_districts_level
const (
	FOUR_LEVEL_ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsLevel  ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsLevel = "FOUR_LEVEL"
	NONE_ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsLevel        ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsLevel = "NONE"
	ONE_LEVEL_ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsLevel   ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsLevel = "ONE_LEVEL"
	THREE_LEVEL_ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsLevel ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsLevel = "THREE_LEVEL"
	TWO_LEVEL_ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsLevel   ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsLevel = "TWO_LEVEL"
)

// All allowed values of ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsLevel enum
var AllowedToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsLevelEnumValues = []ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsLevel{
	"FOUR_LEVEL",
	"NONE",
	"ONE_LEVEL",
	"THREE_LEVEL",
	"TWO_LEVEL",
}

// NewToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsLevelFromValue returns a pointer to a valid ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsLevelFromValue(v string) (*ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsLevel, error) {
	ev := ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsLevel: valid values are %v", v, AllowedToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsLevel) IsValid() bool {
	for _, existing := range AllowedToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_admin_info_v2_data_districts_sub_districts_sub_districts_level value
func (v ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsLevel) Ptr() *ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsLevel {
	return &v
}
