/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAdvertiserStoreSearchV2Type
type ToolsAdvertiserStoreSearchV2Type string

// List of tools_advertiser_store_search_v2_type
const (
	STORE_THIRT_PARTY_ToolsAdvertiserStoreSearchV2Type ToolsAdvertiserStoreSearchV2Type = "STORE_THIRT_PARTY"
	STORE_NORMAL_ToolsAdvertiserStoreSearchV2Type      ToolsAdvertiserStoreSearchV2Type = "STORE_NORMAL"
	STORE_LANDING_ToolsAdvertiserStoreSearchV2Type     ToolsAdvertiserStoreSearchV2Type = "STORE_LANDING"
	STORE_DOUYIN_ToolsAdvertiserStoreSearchV2Type      ToolsAdvertiserStoreSearchV2Type = "STORE_DOUYIN"
)

// All allowed values of ToolsAdvertiserStoreSearchV2Type enum
var AllowedToolsAdvertiserStoreSearchV2TypeEnumValues = []ToolsAdvertiserStoreSearchV2Type{
	"STORE_THIRT_PARTY",
	"STORE_NORMAL",
	"STORE_LANDING",
	"STORE_DOUYIN",
}

// NewToolsAdvertiserStoreSearchV2TypeFromValue returns a pointer to a valid ToolsAdvertiserStoreSearchV2Type
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdvertiserStoreSearchV2TypeFromValue(v string) (*ToolsAdvertiserStoreSearchV2Type, error) {
	ev := ToolsAdvertiserStoreSearchV2Type(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdvertiserStoreSearchV2Type: valid values are %v", v, AllowedToolsAdvertiserStoreSearchV2TypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdvertiserStoreSearchV2Type) IsValid() bool {
	for _, existing := range AllowedToolsAdvertiserStoreSearchV2TypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_advertiser_store_search_v2_type value
func (v ToolsAdvertiserStoreSearchV2Type) Ptr() *ToolsAdvertiserStoreSearchV2Type {
	return &v
}
