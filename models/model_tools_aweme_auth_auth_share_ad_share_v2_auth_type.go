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

// ToolsAwemeAuthAuthShareAdShareV2AuthType
type ToolsAwemeAuthAuthShareAdShareV2AuthType string

// List of tools_aweme_auth_auth_share_ad_share_v2_auth_type
const (
	AWEME_ACCOUNT_ToolsAwemeAuthAuthShareAdShareV2AuthType  ToolsAwemeAuthAuthShareAdShareV2AuthType = "AWEME_ACCOUNT"
	AWEME_HOMEPAGE_ToolsAwemeAuthAuthShareAdShareV2AuthType ToolsAwemeAuthAuthShareAdShareV2AuthType = "AWEME_HOMEPAGE"
	LIVE_ACCOUNT_ToolsAwemeAuthAuthShareAdShareV2AuthType   ToolsAwemeAuthAuthShareAdShareV2AuthType = "LIVE_ACCOUNT"
	VIDEO_ITEM_ToolsAwemeAuthAuthShareAdShareV2AuthType     ToolsAwemeAuthAuthShareAdShareV2AuthType = "VIDEO_ITEM"
)

// All allowed values of ToolsAwemeAuthAuthShareAdShareV2AuthType enum
var AllowedToolsAwemeAuthAuthShareAdShareV2AuthTypeEnumValues = []ToolsAwemeAuthAuthShareAdShareV2AuthType{
	"AWEME_ACCOUNT",
	"AWEME_HOMEPAGE",
	"LIVE_ACCOUNT",
	"VIDEO_ITEM",
}

// NewToolsAwemeAuthAuthShareAdShareV2AuthTypeFromValue returns a pointer to a valid ToolsAwemeAuthAuthShareAdShareV2AuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAwemeAuthAuthShareAdShareV2AuthTypeFromValue(v string) (*ToolsAwemeAuthAuthShareAdShareV2AuthType, error) {
	ev := ToolsAwemeAuthAuthShareAdShareV2AuthType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAwemeAuthAuthShareAdShareV2AuthType: valid values are %v", v, AllowedToolsAwemeAuthAuthShareAdShareV2AuthTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAwemeAuthAuthShareAdShareV2AuthType) IsValid() bool {
	for _, existing := range AllowedToolsAwemeAuthAuthShareAdShareV2AuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_aweme_auth_auth_share_ad_share_v2_auth_type value
func (v ToolsAwemeAuthAuthShareAdShareV2AuthType) Ptr() *ToolsAwemeAuthAuthShareAdShareV2AuthType {
	return &v
}
