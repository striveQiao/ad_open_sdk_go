/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAdConvertQueryV2LandingType
type ToolsAdConvertQueryV2LandingType string

// List of tools_ad_convert_query_v2_landing_type
const (
	APP_ToolsAdConvertQueryV2LandingType       ToolsAdConvertQueryV2LandingType = "APP"
	DPA_ToolsAdConvertQueryV2LandingType       ToolsAdConvertQueryV2LandingType = "DPA"
	SHOP_ToolsAdConvertQueryV2LandingType      ToolsAdConvertQueryV2LandingType = "SHOP"
	AWEME_ToolsAdConvertQueryV2LandingType     ToolsAdConvertQueryV2LandingType = "AWEME"
	LINK_ToolsAdConvertQueryV2LandingType      ToolsAdConvertQueryV2LandingType = "LINK"
	QUICK_APP_ToolsAdConvertQueryV2LandingType ToolsAdConvertQueryV2LandingType = "QUICK_APP"
	LIVE_ToolsAdConvertQueryV2LandingType      ToolsAdConvertQueryV2LandingType = "LIVE"
	STORE_ToolsAdConvertQueryV2LandingType     ToolsAdConvertQueryV2LandingType = "STORE"
	GOODS_ToolsAdConvertQueryV2LandingType     ToolsAdConvertQueryV2LandingType = "GOODS"
	ARTICLE_ToolsAdConvertQueryV2LandingType   ToolsAdConvertQueryV2LandingType = "ARTICLE"
)

// All allowed values of ToolsAdConvertQueryV2LandingType enum
var AllowedToolsAdConvertQueryV2LandingTypeEnumValues = []ToolsAdConvertQueryV2LandingType{
	"APP",
	"DPA",
	"SHOP",
	"AWEME",
	"LINK",
	"QUICK_APP",
	"LIVE",
	"STORE",
	"GOODS",
	"ARTICLE",
}

// NewToolsAdConvertQueryV2LandingTypeFromValue returns a pointer to a valid ToolsAdConvertQueryV2LandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdConvertQueryV2LandingTypeFromValue(v string) (*ToolsAdConvertQueryV2LandingType, error) {
	ev := ToolsAdConvertQueryV2LandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdConvertQueryV2LandingType: valid values are %v", v, AllowedToolsAdConvertQueryV2LandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdConvertQueryV2LandingType) IsValid() bool {
	for _, existing := range AllowedToolsAdConvertQueryV2LandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ad_convert_query_v2_landing_type value
func (v ToolsAdConvertQueryV2LandingType) Ptr() *ToolsAdConvertQueryV2LandingType {
	return &v
}
