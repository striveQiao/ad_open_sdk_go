/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAdConvertQueryV2AppType
type ToolsAdConvertQueryV2AppType string

// List of tools_ad_convert_query_v2_app_type
const (
	APP_IOS_ToolsAdConvertQueryV2AppType     ToolsAdConvertQueryV2AppType = "APP_IOS"
	APP_ANDROID_ToolsAdConvertQueryV2AppType ToolsAdConvertQueryV2AppType = "APP_ANDROID"
)

// All allowed values of ToolsAdConvertQueryV2AppType enum
var AllowedToolsAdConvertQueryV2AppTypeEnumValues = []ToolsAdConvertQueryV2AppType{
	"APP_IOS",
	"APP_ANDROID",
}

// NewToolsAdConvertQueryV2AppTypeFromValue returns a pointer to a valid ToolsAdConvertQueryV2AppType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdConvertQueryV2AppTypeFromValue(v string) (*ToolsAdConvertQueryV2AppType, error) {
	ev := ToolsAdConvertQueryV2AppType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdConvertQueryV2AppType: valid values are %v", v, AllowedToolsAdConvertQueryV2AppTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdConvertQueryV2AppType) IsValid() bool {
	for _, existing := range AllowedToolsAdConvertQueryV2AppTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ad_convert_query_v2_app_type value
func (v ToolsAdConvertQueryV2AppType) Ptr() *ToolsAdConvertQueryV2AppType {
	return &v
}
