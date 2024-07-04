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

// ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType
type ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType string

// List of tools_ad_convert_optimized_target_get_v2_data_list_converts_convert_type
const (
	AD_CONVERT_SOURCE_TYPE_H5_API_ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType              ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType = "AD_CONVERT_SOURCE_TYPE_H5_API"
	AD_CONVERT_SOURCE_TYPE_CONFIG_ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType              ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType = "AD_CONVERT_SOURCE_TYPE_CONFIG"
	AD_CONVERT_SOURCE_TYPE_SDK_ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType                 ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType = "AD_CONVERT_SOURCE_TYPE_SDK"
	AD_CONVERT_SOURCE_TYPE_OPEN_URL_ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType            ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType = "AD_CONVERT_SOURCE_TYPE_OPEN_URL"
	AD_CONVERT_TYPE_MICRO_APP_SDK_ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType              ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType = "AD_CONVERT_TYPE_MICRO_APP_SDK"
	AD_CONVERT_SOURCE_TYPE_NORMAL_APP_DOWNLOAD_ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType = "AD_CONVERT_SOURCE_TYPE_NORMAL_APP_DOWNLOAD"
	AD_CONVERT_SOURCE_TYPE_APP_DOWNLOAD_ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType        ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType = "AD_CONVERT_SOURCE_TYPE_APP_DOWNLOAD"
	AD_CONVERT_SOURCE_TYPE_API_ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType                 ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType = "AD_CONVERT_SOURCE_TYPE_API"
	AD_CONVERT_SOURCE_TYPE_JS_ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType                  ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType = "AD_CONVERT_SOURCE_TYPE_JS"
	AD_CONVERT_SOURCE_TYPE_XPATH_ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType               ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType = "AD_CONVERT_SOURCE_TYPE_XPATH"
	AD_CONVERT_SOURCE_TYPE_APP_API_TEMAI_ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType       ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType = "AD_CONVERT_SOURCE_TYPE_APP_API_TEMAI"
	AD_CONVERT_SOURCE_TYPE_INTERNAL_ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType            ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType = "AD_CONVERT_SOURCE_TYPE_INTERNAL"
	AD_CONVERT_TYPE_NATIVE_PROMOTION_ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType           ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType = "AD_CONVERT_TYPE_NATIVE_PROMOTION"
	AD_CONVERT_SOURCE_TYPE_CPS_GAME_ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType            ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType = "AD_CONVERT_SOURCE_TYPE_CPS_GAME"
	AD_CONVERT_TYPE_MICRO_APP_API_ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType              ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType = "AD_CONVERT_TYPE_MICRO_APP_API"
	AD_CONVERT_SOURCE_TYPE_MULTI_CONVERT_ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType       ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType = "AD_CONVERT_SOURCE_TYPE_MULTI_CONVERT"
)

// All allowed values of ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType enum
var AllowedToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertTypeEnumValues = []ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType{
	"AD_CONVERT_SOURCE_TYPE_H5_API",
	"AD_CONVERT_SOURCE_TYPE_CONFIG",
	"AD_CONVERT_SOURCE_TYPE_SDK",
	"AD_CONVERT_SOURCE_TYPE_OPEN_URL",
	"AD_CONVERT_TYPE_MICRO_APP_SDK",
	"AD_CONVERT_SOURCE_TYPE_NORMAL_APP_DOWNLOAD",
	"AD_CONVERT_SOURCE_TYPE_APP_DOWNLOAD",
	"AD_CONVERT_SOURCE_TYPE_API",
	"AD_CONVERT_SOURCE_TYPE_JS",
	"AD_CONVERT_SOURCE_TYPE_XPATH",
	"AD_CONVERT_SOURCE_TYPE_APP_API_TEMAI",
	"AD_CONVERT_SOURCE_TYPE_INTERNAL",
	"AD_CONVERT_TYPE_NATIVE_PROMOTION",
	"AD_CONVERT_SOURCE_TYPE_CPS_GAME",
	"AD_CONVERT_TYPE_MICRO_APP_API",
	"AD_CONVERT_SOURCE_TYPE_MULTI_CONVERT",
}

// NewToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertTypeFromValue returns a pointer to a valid ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertTypeFromValue(v string) (*ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType, error) {
	ev := ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType: valid values are %v", v, AllowedToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType) IsValid() bool {
	for _, existing := range AllowedToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ad_convert_optimized_target_get_v2_data_list_converts_convert_type value
func (v ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType) Ptr() *ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType {
	return &v
}
