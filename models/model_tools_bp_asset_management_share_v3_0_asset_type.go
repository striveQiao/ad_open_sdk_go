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

// ToolsBpAssetManagementShareV30AssetType
type ToolsBpAssetManagementShareV30AssetType string

// List of tools_bp_asset_management_share_v3.0_asset_type
const (
	APPLETS_ToolsBpAssetManagementShareV30AssetType       ToolsBpAssetManagementShareV30AssetType = "APPLETS"
	BYTED_APPLETS_ToolsBpAssetManagementShareV30AssetType ToolsBpAssetManagementShareV30AssetType = "BYTED_APPLETS"
	BYTED_GAME_ToolsBpAssetManagementShareV30AssetType    ToolsBpAssetManagementShareV30AssetType = "BYTED_GAME"
	WECHAT_GAME_ToolsBpAssetManagementShareV30AssetType   ToolsBpAssetManagementShareV30AssetType = "WECHAT_GAME"
)

// All allowed values of ToolsBpAssetManagementShareV30AssetType enum
var AllowedToolsBpAssetManagementShareV30AssetTypeEnumValues = []ToolsBpAssetManagementShareV30AssetType{
	"APPLETS",
	"BYTED_APPLETS",
	"BYTED_GAME",
	"WECHAT_GAME",
}

// NewToolsBpAssetManagementShareV30AssetTypeFromValue returns a pointer to a valid ToolsBpAssetManagementShareV30AssetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBpAssetManagementShareV30AssetTypeFromValue(v string) (*ToolsBpAssetManagementShareV30AssetType, error) {
	ev := ToolsBpAssetManagementShareV30AssetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBpAssetManagementShareV30AssetType: valid values are %v", v, AllowedToolsBpAssetManagementShareV30AssetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBpAssetManagementShareV30AssetType) IsValid() bool {
	for _, existing := range AllowedToolsBpAssetManagementShareV30AssetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bp_asset_management_share_v3.0_asset_type value
func (v ToolsBpAssetManagementShareV30AssetType) Ptr() *ToolsBpAssetManagementShareV30AssetType {
	return &v
}
