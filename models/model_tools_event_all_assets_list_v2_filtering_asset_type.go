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

// ToolsEventAllAssetsListV2FilteringAssetType
type ToolsEventAllAssetsListV2FilteringAssetType string

// List of tools_event_all_assets_list_v2_filtering_asset_type
const (
	THIRD_EXTERNAL_ToolsEventAllAssetsListV2FilteringAssetType  ToolsEventAllAssetsListV2FilteringAssetType = "THIRD_EXTERNAL"
	TETRIS_EXTERNAL_ToolsEventAllAssetsListV2FilteringAssetType ToolsEventAllAssetsListV2FilteringAssetType = "TETRIS_EXTERNAL"
	APP_ToolsEventAllAssetsListV2FilteringAssetType             ToolsEventAllAssetsListV2FilteringAssetType = "APP"
	QUICK_APP_ToolsEventAllAssetsListV2FilteringAssetType       ToolsEventAllAssetsListV2FilteringAssetType = "QUICK_APP"
	MINI_PROGRAME_ToolsEventAllAssetsListV2FilteringAssetType   ToolsEventAllAssetsListV2FilteringAssetType = "MINI_PROGRAME"
)

// All allowed values of ToolsEventAllAssetsListV2FilteringAssetType enum
var AllowedToolsEventAllAssetsListV2FilteringAssetTypeEnumValues = []ToolsEventAllAssetsListV2FilteringAssetType{
	"THIRD_EXTERNAL",
	"TETRIS_EXTERNAL",
	"APP",
	"QUICK_APP",
	"MINI_PROGRAME",
}

// NewToolsEventAllAssetsListV2FilteringAssetTypeFromValue returns a pointer to a valid ToolsEventAllAssetsListV2FilteringAssetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEventAllAssetsListV2FilteringAssetTypeFromValue(v string) (*ToolsEventAllAssetsListV2FilteringAssetType, error) {
	ev := ToolsEventAllAssetsListV2FilteringAssetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEventAllAssetsListV2FilteringAssetType: valid values are %v", v, AllowedToolsEventAllAssetsListV2FilteringAssetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEventAllAssetsListV2FilteringAssetType) IsValid() bool {
	for _, existing := range AllowedToolsEventAllAssetsListV2FilteringAssetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_event_all_assets_list_v2_filtering_asset_type value
func (v ToolsEventAllAssetsListV2FilteringAssetType) Ptr() *ToolsEventAllAssetsListV2FilteringAssetType {
	return &v
}
