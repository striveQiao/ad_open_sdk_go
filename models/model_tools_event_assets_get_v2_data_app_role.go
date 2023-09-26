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

// ToolsEventAssetsGetV2DataAppRole
type ToolsEventAssetsGetV2DataAppRole string

// List of tools_event_assets_get_v2_data_app_role
const (
	MINE_ToolsEventAssetsGetV2DataAppRole   ToolsEventAssetsGetV2DataAppRole = "MINE"
	SHARED_ToolsEventAssetsGetV2DataAppRole ToolsEventAssetsGetV2DataAppRole = "SHARED"
)

// All allowed values of ToolsEventAssetsGetV2DataAppRole enum
var AllowedToolsEventAssetsGetV2DataAppRoleEnumValues = []ToolsEventAssetsGetV2DataAppRole{
	"MINE",
	"SHARED",
}

// NewToolsEventAssetsGetV2DataAppRoleFromValue returns a pointer to a valid ToolsEventAssetsGetV2DataAppRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEventAssetsGetV2DataAppRoleFromValue(v string) (*ToolsEventAssetsGetV2DataAppRole, error) {
	ev := ToolsEventAssetsGetV2DataAppRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEventAssetsGetV2DataAppRole: valid values are %v", v, AllowedToolsEventAssetsGetV2DataAppRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEventAssetsGetV2DataAppRole) IsValid() bool {
	for _, existing := range AllowedToolsEventAssetsGetV2DataAppRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_event_assets_get_v2_data_app_role value
func (v ToolsEventAssetsGetV2DataAppRole) Ptr() *ToolsEventAssetsGetV2DataAppRole {
	return &v
}
