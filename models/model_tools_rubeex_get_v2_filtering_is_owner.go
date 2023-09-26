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

// ToolsRubeexGetV2FilteringIsOwner
type ToolsRubeexGetV2FilteringIsOwner int64

// List of tools_rubeex_get_v2_filtering_is_owner
const (
	Enum_0_ToolsRubeexGetV2FilteringIsOwner ToolsRubeexGetV2FilteringIsOwner = 0
	Enum_1_ToolsRubeexGetV2FilteringIsOwner ToolsRubeexGetV2FilteringIsOwner = 1
)

// All allowed values of ToolsRubeexGetV2FilteringIsOwner enum
var AllowedToolsRubeexGetV2FilteringIsOwnerEnumValues = []ToolsRubeexGetV2FilteringIsOwner{
	0,
	1,
}

// NewToolsRubeexGetV2FilteringIsOwnerFromValue returns a pointer to a valid ToolsRubeexGetV2FilteringIsOwner
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsRubeexGetV2FilteringIsOwnerFromValue(v int64) (*ToolsRubeexGetV2FilteringIsOwner, error) {
	ev := ToolsRubeexGetV2FilteringIsOwner(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsRubeexGetV2FilteringIsOwner: valid values are %v", v, AllowedToolsRubeexGetV2FilteringIsOwnerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsRubeexGetV2FilteringIsOwner) IsValid() bool {
	for _, existing := range AllowedToolsRubeexGetV2FilteringIsOwnerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_rubeex_get_v2_filtering_is_owner value
func (v ToolsRubeexGetV2FilteringIsOwner) Ptr() *ToolsRubeexGetV2FilteringIsOwner {
	return &v
}
