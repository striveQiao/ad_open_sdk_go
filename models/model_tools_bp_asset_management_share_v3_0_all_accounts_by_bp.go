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

// ToolsBpAssetManagementShareV30AllAccountsByBp
type ToolsBpAssetManagementShareV30AllAccountsByBp string

// List of tools_bp_asset_management_share_v3.0_all_accounts_by_bp
const (
	AD_ToolsBpAssetManagementShareV30AllAccountsByBp ToolsBpAssetManagementShareV30AllAccountsByBp = "AD"
)

// All allowed values of ToolsBpAssetManagementShareV30AllAccountsByBp enum
var AllowedToolsBpAssetManagementShareV30AllAccountsByBpEnumValues = []ToolsBpAssetManagementShareV30AllAccountsByBp{
	"AD",
}

// NewToolsBpAssetManagementShareV30AllAccountsByBpFromValue returns a pointer to a valid ToolsBpAssetManagementShareV30AllAccountsByBp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBpAssetManagementShareV30AllAccountsByBpFromValue(v string) (*ToolsBpAssetManagementShareV30AllAccountsByBp, error) {
	ev := ToolsBpAssetManagementShareV30AllAccountsByBp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBpAssetManagementShareV30AllAccountsByBp: valid values are %v", v, AllowedToolsBpAssetManagementShareV30AllAccountsByBpEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBpAssetManagementShareV30AllAccountsByBp) IsValid() bool {
	for _, existing := range AllowedToolsBpAssetManagementShareV30AllAccountsByBpEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bp_asset_management_share_v3.0_all_accounts_by_bp value
func (v ToolsBpAssetManagementShareV30AllAccountsByBp) Ptr() *ToolsBpAssetManagementShareV30AllAccountsByBp {
	return &v
}
