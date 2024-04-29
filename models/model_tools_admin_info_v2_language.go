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

// ToolsAdminInfoV2Language
type ToolsAdminInfoV2Language string

// List of tools_admin_info_v2_language
const (
	ZH_CN_ToolsAdminInfoV2Language     ToolsAdminInfoV2Language = "ZH_CN"
	ZH_CN_GOV_ToolsAdminInfoV2Language ToolsAdminInfoV2Language = "ZH_CN_GOV"
)

// All allowed values of ToolsAdminInfoV2Language enum
var AllowedToolsAdminInfoV2LanguageEnumValues = []ToolsAdminInfoV2Language{
	"ZH_CN",
	"ZH_CN_GOV",
}

// NewToolsAdminInfoV2LanguageFromValue returns a pointer to a valid ToolsAdminInfoV2Language
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdminInfoV2LanguageFromValue(v string) (*ToolsAdminInfoV2Language, error) {
	ev := ToolsAdminInfoV2Language(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdminInfoV2Language: valid values are %v", v, AllowedToolsAdminInfoV2LanguageEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdminInfoV2Language) IsValid() bool {
	for _, existing := range AllowedToolsAdminInfoV2LanguageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_admin_info_v2_language value
func (v ToolsAdminInfoV2Language) Ptr() *ToolsAdminInfoV2Language {
	return &v
}
