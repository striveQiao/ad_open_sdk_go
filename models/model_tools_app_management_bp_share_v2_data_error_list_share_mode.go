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

// ToolsAppManagementBpShareV2DataErrorListShareMode
type ToolsAppManagementBpShareV2DataErrorListShareMode string

// List of tools_app_management_bp_share_v2_data_error_list_share_mode
const (
	ALL_ToolsAppManagementBpShareV2DataErrorListShareMode     ToolsAppManagementBpShareV2DataErrorListShareMode = "ALL"
	COMPANY_ToolsAppManagementBpShareV2DataErrorListShareMode ToolsAppManagementBpShareV2DataErrorListShareMode = "COMPANY"
	PART_ToolsAppManagementBpShareV2DataErrorListShareMode    ToolsAppManagementBpShareV2DataErrorListShareMode = "PART"
)

// All allowed values of ToolsAppManagementBpShareV2DataErrorListShareMode enum
var AllowedToolsAppManagementBpShareV2DataErrorListShareModeEnumValues = []ToolsAppManagementBpShareV2DataErrorListShareMode{
	"ALL",
	"COMPANY",
	"PART",
}

// NewToolsAppManagementBpShareV2DataErrorListShareModeFromValue returns a pointer to a valid ToolsAppManagementBpShareV2DataErrorListShareMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementBpShareV2DataErrorListShareModeFromValue(v string) (*ToolsAppManagementBpShareV2DataErrorListShareMode, error) {
	ev := ToolsAppManagementBpShareV2DataErrorListShareMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementBpShareV2DataErrorListShareMode: valid values are %v", v, AllowedToolsAppManagementBpShareV2DataErrorListShareModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementBpShareV2DataErrorListShareMode) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementBpShareV2DataErrorListShareModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_bp_share_v2_data_error_list_share_mode value
func (v ToolsAppManagementBpShareV2DataErrorListShareMode) Ptr() *ToolsAppManagementBpShareV2DataErrorListShareMode {
	return &v
}
