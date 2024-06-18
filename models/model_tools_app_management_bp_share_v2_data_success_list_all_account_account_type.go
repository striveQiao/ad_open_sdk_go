/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAppManagementBpShareV2DataSuccessListAllAccountAccountType
type ToolsAppManagementBpShareV2DataSuccessListAllAccountAccountType string

// List of tools_app_management_bp_share_v2_data_success_list_all_account_account_type
const (
	STAR_ToolsAppManagementBpShareV2DataSuccessListAllAccountAccountType ToolsAppManagementBpShareV2DataSuccessListAllAccountAccountType = "STAR"
	BP_ToolsAppManagementBpShareV2DataSuccessListAllAccountAccountType   ToolsAppManagementBpShareV2DataSuccessListAllAccountAccountType = "BP"
	AD_ToolsAppManagementBpShareV2DataSuccessListAllAccountAccountType   ToolsAppManagementBpShareV2DataSuccessListAllAccountAccountType = "AD"
)

// All allowed values of ToolsAppManagementBpShareV2DataSuccessListAllAccountAccountType enum
var AllowedToolsAppManagementBpShareV2DataSuccessListAllAccountAccountTypeEnumValues = []ToolsAppManagementBpShareV2DataSuccessListAllAccountAccountType{
	"STAR",
	"BP",
	"AD",
}

// NewToolsAppManagementBpShareV2DataSuccessListAllAccountAccountTypeFromValue returns a pointer to a valid ToolsAppManagementBpShareV2DataSuccessListAllAccountAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementBpShareV2DataSuccessListAllAccountAccountTypeFromValue(v string) (*ToolsAppManagementBpShareV2DataSuccessListAllAccountAccountType, error) {
	ev := ToolsAppManagementBpShareV2DataSuccessListAllAccountAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementBpShareV2DataSuccessListAllAccountAccountType: valid values are %v", v, AllowedToolsAppManagementBpShareV2DataSuccessListAllAccountAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementBpShareV2DataSuccessListAllAccountAccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementBpShareV2DataSuccessListAllAccountAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_bp_share_v2_data_success_list_all_account_account_type value
func (v ToolsAppManagementBpShareV2DataSuccessListAllAccountAccountType) Ptr() *ToolsAppManagementBpShareV2DataSuccessListAllAccountAccountType {
	return &v
}
