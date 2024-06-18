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

// ToolsRtaSetScopeV2TargetType
type ToolsRtaSetScopeV2TargetType string

// List of tools_rta_set_scope_v2_target_type
const (
	ADV_ToolsRtaSetScopeV2TargetType     ToolsRtaSetScopeV2TargetType = "ADV"
	PROJECT_ToolsRtaSetScopeV2TargetType ToolsRtaSetScopeV2TargetType = "PROJECT"
)

// All allowed values of ToolsRtaSetScopeV2TargetType enum
var AllowedToolsRtaSetScopeV2TargetTypeEnumValues = []ToolsRtaSetScopeV2TargetType{
	"ADV",
	"PROJECT",
}

// NewToolsRtaSetScopeV2TargetTypeFromValue returns a pointer to a valid ToolsRtaSetScopeV2TargetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsRtaSetScopeV2TargetTypeFromValue(v string) (*ToolsRtaSetScopeV2TargetType, error) {
	ev := ToolsRtaSetScopeV2TargetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsRtaSetScopeV2TargetType: valid values are %v", v, AllowedToolsRtaSetScopeV2TargetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsRtaSetScopeV2TargetType) IsValid() bool {
	for _, existing := range AllowedToolsRtaSetScopeV2TargetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_rta_set_scope_v2_target_type value
func (v ToolsRtaSetScopeV2TargetType) Ptr() *ToolsRtaSetScopeV2TargetType {
	return &v
}
