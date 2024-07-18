/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAppManagementExtendPackageCreateV2V2AccountType
type ToolsAppManagementExtendPackageCreateV2V2AccountType string

// List of tools_app_management_extend_package_create_v2_v2_account_type
const (
	AD_ToolsAppManagementExtendPackageCreateV2V2AccountType   ToolsAppManagementExtendPackageCreateV2V2AccountType = "AD"
	BP_ToolsAppManagementExtendPackageCreateV2V2AccountType   ToolsAppManagementExtendPackageCreateV2V2AccountType = "BP"
	STAR_ToolsAppManagementExtendPackageCreateV2V2AccountType ToolsAppManagementExtendPackageCreateV2V2AccountType = "STAR"
)

// All allowed values of ToolsAppManagementExtendPackageCreateV2V2AccountType enum
var AllowedToolsAppManagementExtendPackageCreateV2V2AccountTypeEnumValues = []ToolsAppManagementExtendPackageCreateV2V2AccountType{
	"AD",
	"BP",
	"STAR",
}

// NewToolsAppManagementExtendPackageCreateV2V2AccountTypeFromValue returns a pointer to a valid ToolsAppManagementExtendPackageCreateV2V2AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementExtendPackageCreateV2V2AccountTypeFromValue(v string) (*ToolsAppManagementExtendPackageCreateV2V2AccountType, error) {
	ev := ToolsAppManagementExtendPackageCreateV2V2AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementExtendPackageCreateV2V2AccountType: valid values are %v", v, AllowedToolsAppManagementExtendPackageCreateV2V2AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementExtendPackageCreateV2V2AccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementExtendPackageCreateV2V2AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_extend_package_create_v2_v2_account_type value
func (v ToolsAppManagementExtendPackageCreateV2V2AccountType) Ptr() *ToolsAppManagementExtendPackageCreateV2V2AccountType {
	return &v
}
