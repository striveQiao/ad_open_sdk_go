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

// ToolsRegionGetV2RegionType
type ToolsRegionGetV2RegionType string

// List of tools_region_get_v2_region_type
const (
	BUSINESS_DISTRICT_ToolsRegionGetV2RegionType ToolsRegionGetV2RegionType = "BUSINESS_DISTRICT"
)

// All allowed values of ToolsRegionGetV2RegionType enum
var AllowedToolsRegionGetV2RegionTypeEnumValues = []ToolsRegionGetV2RegionType{
	"BUSINESS_DISTRICT",
}

// NewToolsRegionGetV2RegionTypeFromValue returns a pointer to a valid ToolsRegionGetV2RegionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsRegionGetV2RegionTypeFromValue(v string) (*ToolsRegionGetV2RegionType, error) {
	ev := ToolsRegionGetV2RegionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsRegionGetV2RegionType: valid values are %v", v, AllowedToolsRegionGetV2RegionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsRegionGetV2RegionType) IsValid() bool {
	for _, existing := range AllowedToolsRegionGetV2RegionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_region_get_v2_region_type value
func (v ToolsRegionGetV2RegionType) Ptr() *ToolsRegionGetV2RegionType {
	return &v
}
