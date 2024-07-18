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

// ToolsBpAssetManagementShareCancelV30ShareMode
type ToolsBpAssetManagementShareCancelV30ShareMode string

// List of tools_bp_asset_management_share_cancel_v3.0_share_mode
const (
	BP_ALL_ACCOUNTS_ToolsBpAssetManagementShareCancelV30ShareMode      ToolsBpAssetManagementShareCancelV30ShareMode = "BP_ALL_ACCOUNTS"
	COMPANY_ALL_ACCOUNTS_ToolsBpAssetManagementShareCancelV30ShareMode ToolsBpAssetManagementShareCancelV30ShareMode = "COMPANY_ALL_ACCOUNTS"
	PART_ToolsBpAssetManagementShareCancelV30ShareMode                 ToolsBpAssetManagementShareCancelV30ShareMode = "PART"
)

// All allowed values of ToolsBpAssetManagementShareCancelV30ShareMode enum
var AllowedToolsBpAssetManagementShareCancelV30ShareModeEnumValues = []ToolsBpAssetManagementShareCancelV30ShareMode{
	"BP_ALL_ACCOUNTS",
	"COMPANY_ALL_ACCOUNTS",
	"PART",
}

// NewToolsBpAssetManagementShareCancelV30ShareModeFromValue returns a pointer to a valid ToolsBpAssetManagementShareCancelV30ShareMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBpAssetManagementShareCancelV30ShareModeFromValue(v string) (*ToolsBpAssetManagementShareCancelV30ShareMode, error) {
	ev := ToolsBpAssetManagementShareCancelV30ShareMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBpAssetManagementShareCancelV30ShareMode: valid values are %v", v, AllowedToolsBpAssetManagementShareCancelV30ShareModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBpAssetManagementShareCancelV30ShareMode) IsValid() bool {
	for _, existing := range AllowedToolsBpAssetManagementShareCancelV30ShareModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bp_asset_management_share_cancel_v3.0_share_mode value
func (v ToolsBpAssetManagementShareCancelV30ShareMode) Ptr() *ToolsBpAssetManagementShareCancelV30ShareMode {
	return &v
}
