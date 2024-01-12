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

// ToolsAppManagementAndroidAppListV2AccountType
type ToolsAppManagementAndroidAppListV2AccountType string

// List of tools_app_management_android_app_list_v2_account_type
const (
	AD_ToolsAppManagementAndroidAppListV2AccountType   ToolsAppManagementAndroidAppListV2AccountType = "AD"
	BP_ToolsAppManagementAndroidAppListV2AccountType   ToolsAppManagementAndroidAppListV2AccountType = "BP"
	STAR_ToolsAppManagementAndroidAppListV2AccountType ToolsAppManagementAndroidAppListV2AccountType = "STAR"
)

// All allowed values of ToolsAppManagementAndroidAppListV2AccountType enum
var AllowedToolsAppManagementAndroidAppListV2AccountTypeEnumValues = []ToolsAppManagementAndroidAppListV2AccountType{
	"AD",
	"BP",
	"STAR",
}

// NewToolsAppManagementAndroidAppListV2AccountTypeFromValue returns a pointer to a valid ToolsAppManagementAndroidAppListV2AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementAndroidAppListV2AccountTypeFromValue(v string) (*ToolsAppManagementAndroidAppListV2AccountType, error) {
	ev := ToolsAppManagementAndroidAppListV2AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementAndroidAppListV2AccountType: valid values are %v", v, AllowedToolsAppManagementAndroidAppListV2AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementAndroidAppListV2AccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementAndroidAppListV2AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_android_app_list_v2_account_type value
func (v ToolsAppManagementAndroidAppListV2AccountType) Ptr() *ToolsAppManagementAndroidAppListV2AccountType {
	return &v
}
