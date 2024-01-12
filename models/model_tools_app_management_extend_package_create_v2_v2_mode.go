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

// ToolsAppManagementExtendPackageCreateV2V2Mode
type ToolsAppManagementExtendPackageCreateV2V2Mode string

// List of tools_app_management_extend_package_create_v2_v2_mode
const (
	AUTO_ToolsAppManagementExtendPackageCreateV2V2Mode      ToolsAppManagementExtendPackageCreateV2V2Mode = "Auto"
	CUSTOMIZE_ToolsAppManagementExtendPackageCreateV2V2Mode ToolsAppManagementExtendPackageCreateV2V2Mode = "Customize"
	MANUAL_ToolsAppManagementExtendPackageCreateV2V2Mode    ToolsAppManagementExtendPackageCreateV2V2Mode = "Manual"
)

// All allowed values of ToolsAppManagementExtendPackageCreateV2V2Mode enum
var AllowedToolsAppManagementExtendPackageCreateV2V2ModeEnumValues = []ToolsAppManagementExtendPackageCreateV2V2Mode{
	"Auto",
	"Customize",
	"Manual",
}

// NewToolsAppManagementExtendPackageCreateV2V2ModeFromValue returns a pointer to a valid ToolsAppManagementExtendPackageCreateV2V2Mode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementExtendPackageCreateV2V2ModeFromValue(v string) (*ToolsAppManagementExtendPackageCreateV2V2Mode, error) {
	ev := ToolsAppManagementExtendPackageCreateV2V2Mode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementExtendPackageCreateV2V2Mode: valid values are %v", v, AllowedToolsAppManagementExtendPackageCreateV2V2ModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementExtendPackageCreateV2V2Mode) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementExtendPackageCreateV2V2ModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_extend_package_create_v2_v2_mode value
func (v ToolsAppManagementExtendPackageCreateV2V2Mode) Ptr() *ToolsAppManagementExtendPackageCreateV2V2Mode {
	return &v
}
