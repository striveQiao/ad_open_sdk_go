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

// ToolsBpAssetManagementShareCancelV30AccountInfosAccountType
type ToolsBpAssetManagementShareCancelV30AccountInfosAccountType string

// List of tools_bp_asset_management_share_cancel_v3.0_account_infos_account_type
const (
	AD_ToolsBpAssetManagementShareCancelV30AccountInfosAccountType ToolsBpAssetManagementShareCancelV30AccountInfosAccountType = "AD"
	BP_ToolsBpAssetManagementShareCancelV30AccountInfosAccountType ToolsBpAssetManagementShareCancelV30AccountInfosAccountType = "BP"
)

// All allowed values of ToolsBpAssetManagementShareCancelV30AccountInfosAccountType enum
var AllowedToolsBpAssetManagementShareCancelV30AccountInfosAccountTypeEnumValues = []ToolsBpAssetManagementShareCancelV30AccountInfosAccountType{
	"AD",
	"BP",
}

// NewToolsBpAssetManagementShareCancelV30AccountInfosAccountTypeFromValue returns a pointer to a valid ToolsBpAssetManagementShareCancelV30AccountInfosAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBpAssetManagementShareCancelV30AccountInfosAccountTypeFromValue(v string) (*ToolsBpAssetManagementShareCancelV30AccountInfosAccountType, error) {
	ev := ToolsBpAssetManagementShareCancelV30AccountInfosAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBpAssetManagementShareCancelV30AccountInfosAccountType: valid values are %v", v, AllowedToolsBpAssetManagementShareCancelV30AccountInfosAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBpAssetManagementShareCancelV30AccountInfosAccountType) IsValid() bool {
	for _, existing := range AllowedToolsBpAssetManagementShareCancelV30AccountInfosAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bp_asset_management_share_cancel_v3.0_account_infos_account_type value
func (v ToolsBpAssetManagementShareCancelV30AccountInfosAccountType) Ptr() *ToolsBpAssetManagementShareCancelV30AccountInfosAccountType {
	return &v
}
