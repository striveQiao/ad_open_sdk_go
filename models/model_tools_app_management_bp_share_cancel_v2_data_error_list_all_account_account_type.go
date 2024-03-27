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

// ToolsAppManagementBpShareCancelV2DataErrorListAllAccountAccountType
type ToolsAppManagementBpShareCancelV2DataErrorListAllAccountAccountType string

// List of tools_app_management_bp_share_cancel_v2_data_error_list_all_account_account_type
const (
	STAR_ToolsAppManagementBpShareCancelV2DataErrorListAllAccountAccountType ToolsAppManagementBpShareCancelV2DataErrorListAllAccountAccountType = "STAR"
	AD_ToolsAppManagementBpShareCancelV2DataErrorListAllAccountAccountType   ToolsAppManagementBpShareCancelV2DataErrorListAllAccountAccountType = "AD"
	BP_ToolsAppManagementBpShareCancelV2DataErrorListAllAccountAccountType   ToolsAppManagementBpShareCancelV2DataErrorListAllAccountAccountType = "BP"
)

// All allowed values of ToolsAppManagementBpShareCancelV2DataErrorListAllAccountAccountType enum
var AllowedToolsAppManagementBpShareCancelV2DataErrorListAllAccountAccountTypeEnumValues = []ToolsAppManagementBpShareCancelV2DataErrorListAllAccountAccountType{
	"STAR",
	"AD",
	"BP",
}

// NewToolsAppManagementBpShareCancelV2DataErrorListAllAccountAccountTypeFromValue returns a pointer to a valid ToolsAppManagementBpShareCancelV2DataErrorListAllAccountAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementBpShareCancelV2DataErrorListAllAccountAccountTypeFromValue(v string) (*ToolsAppManagementBpShareCancelV2DataErrorListAllAccountAccountType, error) {
	ev := ToolsAppManagementBpShareCancelV2DataErrorListAllAccountAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementBpShareCancelV2DataErrorListAllAccountAccountType: valid values are %v", v, AllowedToolsAppManagementBpShareCancelV2DataErrorListAllAccountAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementBpShareCancelV2DataErrorListAllAccountAccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementBpShareCancelV2DataErrorListAllAccountAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_bp_share_cancel_v2_data_error_list_all_account_account_type value
func (v ToolsAppManagementBpShareCancelV2DataErrorListAllAccountAccountType) Ptr() *ToolsAppManagementBpShareCancelV2DataErrorListAllAccountAccountType {
	return &v
}
