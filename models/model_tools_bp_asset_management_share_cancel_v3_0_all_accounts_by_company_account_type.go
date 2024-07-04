/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBpAssetManagementShareCancelV30AllAccountsByCompanyAccountType
type ToolsBpAssetManagementShareCancelV30AllAccountsByCompanyAccountType string

// List of tools_bp_asset_management_share_cancel_v3.0_all_accounts_by_company_account_type
const (
	AD_ToolsBpAssetManagementShareCancelV30AllAccountsByCompanyAccountType ToolsBpAssetManagementShareCancelV30AllAccountsByCompanyAccountType = "AD"
)

// All allowed values of ToolsBpAssetManagementShareCancelV30AllAccountsByCompanyAccountType enum
var AllowedToolsBpAssetManagementShareCancelV30AllAccountsByCompanyAccountTypeEnumValues = []ToolsBpAssetManagementShareCancelV30AllAccountsByCompanyAccountType{
	"AD",
}

// NewToolsBpAssetManagementShareCancelV30AllAccountsByCompanyAccountTypeFromValue returns a pointer to a valid ToolsBpAssetManagementShareCancelV30AllAccountsByCompanyAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBpAssetManagementShareCancelV30AllAccountsByCompanyAccountTypeFromValue(v string) (*ToolsBpAssetManagementShareCancelV30AllAccountsByCompanyAccountType, error) {
	ev := ToolsBpAssetManagementShareCancelV30AllAccountsByCompanyAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBpAssetManagementShareCancelV30AllAccountsByCompanyAccountType: valid values are %v", v, AllowedToolsBpAssetManagementShareCancelV30AllAccountsByCompanyAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBpAssetManagementShareCancelV30AllAccountsByCompanyAccountType) IsValid() bool {
	for _, existing := range AllowedToolsBpAssetManagementShareCancelV30AllAccountsByCompanyAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bp_asset_management_share_cancel_v3.0_all_accounts_by_company_account_type value
func (v ToolsBpAssetManagementShareCancelV30AllAccountsByCompanyAccountType) Ptr() *ToolsBpAssetManagementShareCancelV30AllAccountsByCompanyAccountType {
	return &v
}
