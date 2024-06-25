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

// ToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByCompanyAccountType
type ToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByCompanyAccountType string

// List of tools_bp_asset_management_share_get_v3.0_data_sharedAccounts_all_accounts_by_company_account_type
const (
	AD_ToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByCompanyAccountType ToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByCompanyAccountType = "AD"
)

// All allowed values of ToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByCompanyAccountType enum
var AllowedToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByCompanyAccountTypeEnumValues = []ToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByCompanyAccountType{
	"AD",
}

// NewToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByCompanyAccountTypeFromValue returns a pointer to a valid ToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByCompanyAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByCompanyAccountTypeFromValue(v string) (*ToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByCompanyAccountType, error) {
	ev := ToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByCompanyAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByCompanyAccountType: valid values are %v", v, AllowedToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByCompanyAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByCompanyAccountType) IsValid() bool {
	for _, existing := range AllowedToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByCompanyAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bp_asset_management_share_get_v3.0_data_sharedAccounts_all_accounts_by_company_account_type value
func (v ToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByCompanyAccountType) Ptr() *ToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByCompanyAccountType {
	return &v
}
