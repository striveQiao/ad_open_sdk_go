/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBpAssetManagementShareGetV30ShareType
type ToolsBpAssetManagementShareGetV30ShareType string

// List of tools_bp_asset_management_share_get_v3.0_share_type
const (
	ACCOUNT_ToolsBpAssetManagementShareGetV30ShareType ToolsBpAssetManagementShareGetV30ShareType = "ACCOUNT"
	GROUP_ToolsBpAssetManagementShareGetV30ShareType   ToolsBpAssetManagementShareGetV30ShareType = "GROUP"
)

// All allowed values of ToolsBpAssetManagementShareGetV30ShareType enum
var AllowedToolsBpAssetManagementShareGetV30ShareTypeEnumValues = []ToolsBpAssetManagementShareGetV30ShareType{
	"ACCOUNT",
	"GROUP",
}

// NewToolsBpAssetManagementShareGetV30ShareTypeFromValue returns a pointer to a valid ToolsBpAssetManagementShareGetV30ShareType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBpAssetManagementShareGetV30ShareTypeFromValue(v string) (*ToolsBpAssetManagementShareGetV30ShareType, error) {
	ev := ToolsBpAssetManagementShareGetV30ShareType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBpAssetManagementShareGetV30ShareType: valid values are %v", v, AllowedToolsBpAssetManagementShareGetV30ShareTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBpAssetManagementShareGetV30ShareType) IsValid() bool {
	for _, existing := range AllowedToolsBpAssetManagementShareGetV30ShareTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bp_asset_management_share_get_v3.0_share_type value
func (v ToolsBpAssetManagementShareGetV30ShareType) Ptr() *ToolsBpAssetManagementShareGetV30ShareType {
	return &v
}
