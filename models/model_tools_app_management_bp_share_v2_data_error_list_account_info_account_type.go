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

// ToolsAppManagementBpShareV2DataErrorListAccountInfoAccountType
type ToolsAppManagementBpShareV2DataErrorListAccountInfoAccountType string

// List of tools_app_management_bp_share_v2_data_error_list_account_info_account_type
const (
	BP_ToolsAppManagementBpShareV2DataErrorListAccountInfoAccountType   ToolsAppManagementBpShareV2DataErrorListAccountInfoAccountType = "BP"
	STAR_ToolsAppManagementBpShareV2DataErrorListAccountInfoAccountType ToolsAppManagementBpShareV2DataErrorListAccountInfoAccountType = "STAR"
	AD_ToolsAppManagementBpShareV2DataErrorListAccountInfoAccountType   ToolsAppManagementBpShareV2DataErrorListAccountInfoAccountType = "AD"
)

// All allowed values of ToolsAppManagementBpShareV2DataErrorListAccountInfoAccountType enum
var AllowedToolsAppManagementBpShareV2DataErrorListAccountInfoAccountTypeEnumValues = []ToolsAppManagementBpShareV2DataErrorListAccountInfoAccountType{
	"BP",
	"STAR",
	"AD",
}

// NewToolsAppManagementBpShareV2DataErrorListAccountInfoAccountTypeFromValue returns a pointer to a valid ToolsAppManagementBpShareV2DataErrorListAccountInfoAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementBpShareV2DataErrorListAccountInfoAccountTypeFromValue(v string) (*ToolsAppManagementBpShareV2DataErrorListAccountInfoAccountType, error) {
	ev := ToolsAppManagementBpShareV2DataErrorListAccountInfoAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementBpShareV2DataErrorListAccountInfoAccountType: valid values are %v", v, AllowedToolsAppManagementBpShareV2DataErrorListAccountInfoAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementBpShareV2DataErrorListAccountInfoAccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementBpShareV2DataErrorListAccountInfoAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_bp_share_v2_data_error_list_account_info_account_type value
func (v ToolsAppManagementBpShareV2DataErrorListAccountInfoAccountType) Ptr() *ToolsAppManagementBpShareV2DataErrorListAccountInfoAccountType {
	return &v
}
