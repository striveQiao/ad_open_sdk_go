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

// ToolsBpAssetManagementShareGetV30DataSharedAccountsShareMode
type ToolsBpAssetManagementShareGetV30DataSharedAccountsShareMode string

// List of tools_bp_asset_management_share_get_v3.0_data_sharedAccounts_share_mode
const (
	BP_ALL_ACCOUNTS_ToolsBpAssetManagementShareGetV30DataSharedAccountsShareMode      ToolsBpAssetManagementShareGetV30DataSharedAccountsShareMode = "BP_ALL_ACCOUNTS"
	COMPANY_ALL_ACCOUNTS_ToolsBpAssetManagementShareGetV30DataSharedAccountsShareMode ToolsBpAssetManagementShareGetV30DataSharedAccountsShareMode = "COMPANY_ALL_ACCOUNTS"
	PART_ToolsBpAssetManagementShareGetV30DataSharedAccountsShareMode                 ToolsBpAssetManagementShareGetV30DataSharedAccountsShareMode = "PART"
)

// All allowed values of ToolsBpAssetManagementShareGetV30DataSharedAccountsShareMode enum
var AllowedToolsBpAssetManagementShareGetV30DataSharedAccountsShareModeEnumValues = []ToolsBpAssetManagementShareGetV30DataSharedAccountsShareMode{
	"BP_ALL_ACCOUNTS",
	"COMPANY_ALL_ACCOUNTS",
	"PART",
}

// NewToolsBpAssetManagementShareGetV30DataSharedAccountsShareModeFromValue returns a pointer to a valid ToolsBpAssetManagementShareGetV30DataSharedAccountsShareMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBpAssetManagementShareGetV30DataSharedAccountsShareModeFromValue(v string) (*ToolsBpAssetManagementShareGetV30DataSharedAccountsShareMode, error) {
	ev := ToolsBpAssetManagementShareGetV30DataSharedAccountsShareMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBpAssetManagementShareGetV30DataSharedAccountsShareMode: valid values are %v", v, AllowedToolsBpAssetManagementShareGetV30DataSharedAccountsShareModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBpAssetManagementShareGetV30DataSharedAccountsShareMode) IsValid() bool {
	for _, existing := range AllowedToolsBpAssetManagementShareGetV30DataSharedAccountsShareModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bp_asset_management_share_get_v3.0_data_sharedAccounts_share_mode value
func (v ToolsBpAssetManagementShareGetV30DataSharedAccountsShareMode) Ptr() *ToolsBpAssetManagementShareGetV30DataSharedAccountsShareMode {
	return &v
}
