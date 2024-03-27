/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAppManagementAndroidBasicPackageUpdateV2FileOption
type ToolsAppManagementAndroidBasicPackageUpdateV2FileOption string

// List of tools_app_management_android_basic_package_update_v2_file_option
const (
	USE_LAST_IMAGE_VIDEO_ToolsAppManagementAndroidBasicPackageUpdateV2FileOption ToolsAppManagementAndroidBasicPackageUpdateV2FileOption = "USE_LAST_IMAGE_VIDEO"
	USE_NEW_ToolsAppManagementAndroidBasicPackageUpdateV2FileOption              ToolsAppManagementAndroidBasicPackageUpdateV2FileOption = "USE_NEW"
)

// All allowed values of ToolsAppManagementAndroidBasicPackageUpdateV2FileOption enum
var AllowedToolsAppManagementAndroidBasicPackageUpdateV2FileOptionEnumValues = []ToolsAppManagementAndroidBasicPackageUpdateV2FileOption{
	"USE_LAST_IMAGE_VIDEO",
	"USE_NEW",
}

// NewToolsAppManagementAndroidBasicPackageUpdateV2FileOptionFromValue returns a pointer to a valid ToolsAppManagementAndroidBasicPackageUpdateV2FileOption
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementAndroidBasicPackageUpdateV2FileOptionFromValue(v string) (*ToolsAppManagementAndroidBasicPackageUpdateV2FileOption, error) {
	ev := ToolsAppManagementAndroidBasicPackageUpdateV2FileOption(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementAndroidBasicPackageUpdateV2FileOption: valid values are %v", v, AllowedToolsAppManagementAndroidBasicPackageUpdateV2FileOptionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementAndroidBasicPackageUpdateV2FileOption) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementAndroidBasicPackageUpdateV2FileOptionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_android_basic_package_update_v2_file_option value
func (v ToolsAppManagementAndroidBasicPackageUpdateV2FileOption) Ptr() *ToolsAppManagementAndroidBasicPackageUpdateV2FileOption {
	return &v
}
