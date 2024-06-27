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

// ToolsAppManagementBpShareCancelV2DataSuccessListShareMode
type ToolsAppManagementBpShareCancelV2DataSuccessListShareMode string

// List of tools_app_management_bp_share_cancel_v2_data_success_list_share_mode
const (
	PART_ToolsAppManagementBpShareCancelV2DataSuccessListShareMode    ToolsAppManagementBpShareCancelV2DataSuccessListShareMode = "PART"
	ALL_ToolsAppManagementBpShareCancelV2DataSuccessListShareMode     ToolsAppManagementBpShareCancelV2DataSuccessListShareMode = "ALL"
	COMPANY_ToolsAppManagementBpShareCancelV2DataSuccessListShareMode ToolsAppManagementBpShareCancelV2DataSuccessListShareMode = "COMPANY"
)

// All allowed values of ToolsAppManagementBpShareCancelV2DataSuccessListShareMode enum
var AllowedToolsAppManagementBpShareCancelV2DataSuccessListShareModeEnumValues = []ToolsAppManagementBpShareCancelV2DataSuccessListShareMode{
	"PART",
	"ALL",
	"COMPANY",
}

// NewToolsAppManagementBpShareCancelV2DataSuccessListShareModeFromValue returns a pointer to a valid ToolsAppManagementBpShareCancelV2DataSuccessListShareMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementBpShareCancelV2DataSuccessListShareModeFromValue(v string) (*ToolsAppManagementBpShareCancelV2DataSuccessListShareMode, error) {
	ev := ToolsAppManagementBpShareCancelV2DataSuccessListShareMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementBpShareCancelV2DataSuccessListShareMode: valid values are %v", v, AllowedToolsAppManagementBpShareCancelV2DataSuccessListShareModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementBpShareCancelV2DataSuccessListShareMode) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementBpShareCancelV2DataSuccessListShareModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_bp_share_cancel_v2_data_success_list_share_mode value
func (v ToolsAppManagementBpShareCancelV2DataSuccessListShareMode) Ptr() *ToolsAppManagementBpShareCancelV2DataSuccessListShareMode {
	return &v
}
