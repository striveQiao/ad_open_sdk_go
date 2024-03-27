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

// ToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByBp
type ToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByBp string

// List of tools_bp_asset_management_share_cancel_v3.0_data_error_list_all_accounts_by_bp
const (
	AD_ToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByBp ToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByBp = "AD"
)

// All allowed values of ToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByBp enum
var AllowedToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByBpEnumValues = []ToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByBp{
	"AD",
}

// NewToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByBpFromValue returns a pointer to a valid ToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByBp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByBpFromValue(v string) (*ToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByBp, error) {
	ev := ToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByBp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByBp: valid values are %v", v, AllowedToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByBpEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByBp) IsValid() bool {
	for _, existing := range AllowedToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByBpEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bp_asset_management_share_cancel_v3.0_data_error_list_all_accounts_by_bp value
func (v ToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByBp) Ptr() *ToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByBp {
	return &v
}
