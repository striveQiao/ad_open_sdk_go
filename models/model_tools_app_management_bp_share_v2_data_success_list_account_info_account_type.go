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

// ToolsAppManagementBpShareV2DataSuccessListAccountInfoAccountType
type ToolsAppManagementBpShareV2DataSuccessListAccountInfoAccountType string

// List of tools_app_management_bp_share_v2_data_success_list_account_info_account_type
const (
	BP_ToolsAppManagementBpShareV2DataSuccessListAccountInfoAccountType   ToolsAppManagementBpShareV2DataSuccessListAccountInfoAccountType = "BP"
	AD_ToolsAppManagementBpShareV2DataSuccessListAccountInfoAccountType   ToolsAppManagementBpShareV2DataSuccessListAccountInfoAccountType = "AD"
	STAR_ToolsAppManagementBpShareV2DataSuccessListAccountInfoAccountType ToolsAppManagementBpShareV2DataSuccessListAccountInfoAccountType = "STAR"
)

// All allowed values of ToolsAppManagementBpShareV2DataSuccessListAccountInfoAccountType enum
var AllowedToolsAppManagementBpShareV2DataSuccessListAccountInfoAccountTypeEnumValues = []ToolsAppManagementBpShareV2DataSuccessListAccountInfoAccountType{
	"BP",
	"AD",
	"STAR",
}

// NewToolsAppManagementBpShareV2DataSuccessListAccountInfoAccountTypeFromValue returns a pointer to a valid ToolsAppManagementBpShareV2DataSuccessListAccountInfoAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementBpShareV2DataSuccessListAccountInfoAccountTypeFromValue(v string) (*ToolsAppManagementBpShareV2DataSuccessListAccountInfoAccountType, error) {
	ev := ToolsAppManagementBpShareV2DataSuccessListAccountInfoAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementBpShareV2DataSuccessListAccountInfoAccountType: valid values are %v", v, AllowedToolsAppManagementBpShareV2DataSuccessListAccountInfoAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementBpShareV2DataSuccessListAccountInfoAccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementBpShareV2DataSuccessListAccountInfoAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_bp_share_v2_data_success_list_account_info_account_type value
func (v ToolsAppManagementBpShareV2DataSuccessListAccountInfoAccountType) Ptr() *ToolsAppManagementBpShareV2DataSuccessListAccountInfoAccountType {
	return &v
}
