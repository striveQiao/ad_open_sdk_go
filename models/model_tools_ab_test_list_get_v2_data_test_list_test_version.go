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

// ToolsAbTestListGetV2DataTestListTestVersion
type ToolsAbTestListGetV2DataTestListTestVersion string

// List of tools_ab_test_list_get_v2_data_test_list_test_version
const (
	V1_ToolsAbTestListGetV2DataTestListTestVersion ToolsAbTestListGetV2DataTestListTestVersion = "V1"
	V2_ToolsAbTestListGetV2DataTestListTestVersion ToolsAbTestListGetV2DataTestListTestVersion = "V2"
)

// All allowed values of ToolsAbTestListGetV2DataTestListTestVersion enum
var AllowedToolsAbTestListGetV2DataTestListTestVersionEnumValues = []ToolsAbTestListGetV2DataTestListTestVersion{
	"V1",
	"V2",
}

// NewToolsAbTestListGetV2DataTestListTestVersionFromValue returns a pointer to a valid ToolsAbTestListGetV2DataTestListTestVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAbTestListGetV2DataTestListTestVersionFromValue(v string) (*ToolsAbTestListGetV2DataTestListTestVersion, error) {
	ev := ToolsAbTestListGetV2DataTestListTestVersion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAbTestListGetV2DataTestListTestVersion: valid values are %v", v, AllowedToolsAbTestListGetV2DataTestListTestVersionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAbTestListGetV2DataTestListTestVersion) IsValid() bool {
	for _, existing := range AllowedToolsAbTestListGetV2DataTestListTestVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ab_test_list_get_v2_data_test_list_test_version value
func (v ToolsAbTestListGetV2DataTestListTestVersion) Ptr() *ToolsAbTestListGetV2DataTestListTestVersion {
	return &v
}
