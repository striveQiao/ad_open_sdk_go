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

// ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountByCompanyAccountType
type ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountByCompanyAccountType string

// List of tools_app_management_bp_share_cancel_v2_data_success_list_all_account_by_company_account_type
const (
	STAR_ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountByCompanyAccountType ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountByCompanyAccountType = "STAR"
	BP_ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountByCompanyAccountType   ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountByCompanyAccountType = "BP"
	AD_ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountByCompanyAccountType   ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountByCompanyAccountType = "AD"
)

// All allowed values of ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountByCompanyAccountType enum
var AllowedToolsAppManagementBpShareCancelV2DataSuccessListAllAccountByCompanyAccountTypeEnumValues = []ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountByCompanyAccountType{
	"STAR",
	"BP",
	"AD",
}

// NewToolsAppManagementBpShareCancelV2DataSuccessListAllAccountByCompanyAccountTypeFromValue returns a pointer to a valid ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountByCompanyAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementBpShareCancelV2DataSuccessListAllAccountByCompanyAccountTypeFromValue(v string) (*ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountByCompanyAccountType, error) {
	ev := ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountByCompanyAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountByCompanyAccountType: valid values are %v", v, AllowedToolsAppManagementBpShareCancelV2DataSuccessListAllAccountByCompanyAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountByCompanyAccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementBpShareCancelV2DataSuccessListAllAccountByCompanyAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_bp_share_cancel_v2_data_success_list_all_account_by_company_account_type value
func (v ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountByCompanyAccountType) Ptr() *ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountByCompanyAccountType {
	return &v
}
