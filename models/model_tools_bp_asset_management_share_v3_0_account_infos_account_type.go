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

// ToolsBpAssetManagementShareV30AccountInfosAccountType
type ToolsBpAssetManagementShareV30AccountInfosAccountType string

// List of tools_bp_asset_management_share_v3.0_account_infos_account_type
const (
	AD_ToolsBpAssetManagementShareV30AccountInfosAccountType ToolsBpAssetManagementShareV30AccountInfosAccountType = "AD"
	BP_ToolsBpAssetManagementShareV30AccountInfosAccountType ToolsBpAssetManagementShareV30AccountInfosAccountType = "BP"
)

// All allowed values of ToolsBpAssetManagementShareV30AccountInfosAccountType enum
var AllowedToolsBpAssetManagementShareV30AccountInfosAccountTypeEnumValues = []ToolsBpAssetManagementShareV30AccountInfosAccountType{
	"AD",
	"BP",
}

// NewToolsBpAssetManagementShareV30AccountInfosAccountTypeFromValue returns a pointer to a valid ToolsBpAssetManagementShareV30AccountInfosAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBpAssetManagementShareV30AccountInfosAccountTypeFromValue(v string) (*ToolsBpAssetManagementShareV30AccountInfosAccountType, error) {
	ev := ToolsBpAssetManagementShareV30AccountInfosAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBpAssetManagementShareV30AccountInfosAccountType: valid values are %v", v, AllowedToolsBpAssetManagementShareV30AccountInfosAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBpAssetManagementShareV30AccountInfosAccountType) IsValid() bool {
	for _, existing := range AllowedToolsBpAssetManagementShareV30AccountInfosAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bp_asset_management_share_v3.0_account_infos_account_type value
func (v ToolsBpAssetManagementShareV30AccountInfosAccountType) Ptr() *ToolsBpAssetManagementShareV30AccountInfosAccountType {
	return &v
}
