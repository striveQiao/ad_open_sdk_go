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

// ToolsAppManagementBpShareCancelV2DataErrorListAccountInfoAccountType
type ToolsAppManagementBpShareCancelV2DataErrorListAccountInfoAccountType string

// List of tools_app_management_bp_share_cancel_v2_data_error_list_account_info_account_type
const (
	STAR_ToolsAppManagementBpShareCancelV2DataErrorListAccountInfoAccountType ToolsAppManagementBpShareCancelV2DataErrorListAccountInfoAccountType = "STAR"
	BP_ToolsAppManagementBpShareCancelV2DataErrorListAccountInfoAccountType   ToolsAppManagementBpShareCancelV2DataErrorListAccountInfoAccountType = "BP"
	AD_ToolsAppManagementBpShareCancelV2DataErrorListAccountInfoAccountType   ToolsAppManagementBpShareCancelV2DataErrorListAccountInfoAccountType = "AD"
)

// All allowed values of ToolsAppManagementBpShareCancelV2DataErrorListAccountInfoAccountType enum
var AllowedToolsAppManagementBpShareCancelV2DataErrorListAccountInfoAccountTypeEnumValues = []ToolsAppManagementBpShareCancelV2DataErrorListAccountInfoAccountType{
	"STAR",
	"BP",
	"AD",
}

// NewToolsAppManagementBpShareCancelV2DataErrorListAccountInfoAccountTypeFromValue returns a pointer to a valid ToolsAppManagementBpShareCancelV2DataErrorListAccountInfoAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementBpShareCancelV2DataErrorListAccountInfoAccountTypeFromValue(v string) (*ToolsAppManagementBpShareCancelV2DataErrorListAccountInfoAccountType, error) {
	ev := ToolsAppManagementBpShareCancelV2DataErrorListAccountInfoAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementBpShareCancelV2DataErrorListAccountInfoAccountType: valid values are %v", v, AllowedToolsAppManagementBpShareCancelV2DataErrorListAccountInfoAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementBpShareCancelV2DataErrorListAccountInfoAccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementBpShareCancelV2DataErrorListAccountInfoAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_bp_share_cancel_v2_data_error_list_account_info_account_type value
func (v ToolsAppManagementBpShareCancelV2DataErrorListAccountInfoAccountType) Ptr() *ToolsAppManagementBpShareCancelV2DataErrorListAccountInfoAccountType {
	return &v
}
