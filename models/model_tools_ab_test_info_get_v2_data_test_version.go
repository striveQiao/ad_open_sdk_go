/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAbTestInfoGetV2DataTestVersion
type ToolsAbTestInfoGetV2DataTestVersion string

// List of tools_ab_test_info_get_v2_data_test_version
const (
	V1_ToolsAbTestInfoGetV2DataTestVersion ToolsAbTestInfoGetV2DataTestVersion = "V1"
	V2_ToolsAbTestInfoGetV2DataTestVersion ToolsAbTestInfoGetV2DataTestVersion = "V2"
)

// All allowed values of ToolsAbTestInfoGetV2DataTestVersion enum
var AllowedToolsAbTestInfoGetV2DataTestVersionEnumValues = []ToolsAbTestInfoGetV2DataTestVersion{
	"V1",
	"V2",
}

// NewToolsAbTestInfoGetV2DataTestVersionFromValue returns a pointer to a valid ToolsAbTestInfoGetV2DataTestVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAbTestInfoGetV2DataTestVersionFromValue(v string) (*ToolsAbTestInfoGetV2DataTestVersion, error) {
	ev := ToolsAbTestInfoGetV2DataTestVersion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAbTestInfoGetV2DataTestVersion: valid values are %v", v, AllowedToolsAbTestInfoGetV2DataTestVersionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAbTestInfoGetV2DataTestVersion) IsValid() bool {
	for _, existing := range AllowedToolsAbTestInfoGetV2DataTestVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ab_test_info_get_v2_data_test_version value
func (v ToolsAbTestInfoGetV2DataTestVersion) Ptr() *ToolsAbTestInfoGetV2DataTestVersion {
	return &v
}
