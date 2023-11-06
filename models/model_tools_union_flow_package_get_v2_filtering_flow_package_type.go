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

// ToolsUnionFlowPackageGetV2FilteringFlowPackageType
type ToolsUnionFlowPackageGetV2FilteringFlowPackageType string

// List of tools_union_flow_package_get_v2_filtering_flow_package_type
const (
	CUSTOMIZE_ToolsUnionFlowPackageGetV2FilteringFlowPackageType ToolsUnionFlowPackageGetV2FilteringFlowPackageType = "CUSTOMIZE"
	FEATURED_ToolsUnionFlowPackageGetV2FilteringFlowPackageType  ToolsUnionFlowPackageGetV2FilteringFlowPackageType = "FEATURED"
	SYSTEM_ToolsUnionFlowPackageGetV2FilteringFlowPackageType    ToolsUnionFlowPackageGetV2FilteringFlowPackageType = "SYSTEM"
)

// All allowed values of ToolsUnionFlowPackageGetV2FilteringFlowPackageType enum
var AllowedToolsUnionFlowPackageGetV2FilteringFlowPackageTypeEnumValues = []ToolsUnionFlowPackageGetV2FilteringFlowPackageType{
	"CUSTOMIZE",
	"FEATURED",
	"SYSTEM",
}

// NewToolsUnionFlowPackageGetV2FilteringFlowPackageTypeFromValue returns a pointer to a valid ToolsUnionFlowPackageGetV2FilteringFlowPackageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsUnionFlowPackageGetV2FilteringFlowPackageTypeFromValue(v string) (*ToolsUnionFlowPackageGetV2FilteringFlowPackageType, error) {
	ev := ToolsUnionFlowPackageGetV2FilteringFlowPackageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsUnionFlowPackageGetV2FilteringFlowPackageType: valid values are %v", v, AllowedToolsUnionFlowPackageGetV2FilteringFlowPackageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsUnionFlowPackageGetV2FilteringFlowPackageType) IsValid() bool {
	for _, existing := range AllowedToolsUnionFlowPackageGetV2FilteringFlowPackageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_union_flow_package_get_v2_filtering_flow_package_type value
func (v ToolsUnionFlowPackageGetV2FilteringFlowPackageType) Ptr() *ToolsUnionFlowPackageGetV2FilteringFlowPackageType {
	return &v
}
