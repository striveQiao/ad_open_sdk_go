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

// ToolsAwemeBannedListV30BannedType
type ToolsAwemeBannedListV30BannedType string

// List of tools_aweme_banned_list_v3.0_banned_type
const (
	AWEME_TYPE_ToolsAwemeBannedListV30BannedType  ToolsAwemeBannedListV30BannedType = "AWEME_TYPE"
	CUSTOM_TYPE_ToolsAwemeBannedListV30BannedType ToolsAwemeBannedListV30BannedType = "CUSTOM_TYPE"
)

// All allowed values of ToolsAwemeBannedListV30BannedType enum
var AllowedToolsAwemeBannedListV30BannedTypeEnumValues = []ToolsAwemeBannedListV30BannedType{
	"AWEME_TYPE",
	"CUSTOM_TYPE",
}

// NewToolsAwemeBannedListV30BannedTypeFromValue returns a pointer to a valid ToolsAwemeBannedListV30BannedType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAwemeBannedListV30BannedTypeFromValue(v string) (*ToolsAwemeBannedListV30BannedType, error) {
	ev := ToolsAwemeBannedListV30BannedType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAwemeBannedListV30BannedType: valid values are %v", v, AllowedToolsAwemeBannedListV30BannedTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAwemeBannedListV30BannedType) IsValid() bool {
	for _, existing := range AllowedToolsAwemeBannedListV30BannedTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_aweme_banned_list_v3.0_banned_type value
func (v ToolsAwemeBannedListV30BannedType) Ptr() *ToolsAwemeBannedListV30BannedType {
	return &v
}
