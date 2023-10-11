/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsRubeexRemarkV2Scene
type ToolsRubeexRemarkV2Scene int64

// List of tools_rubeex_remark_v2_scene
const (
	Enum_1_ToolsRubeexRemarkV2Scene ToolsRubeexRemarkV2Scene = 1
	Enum_2_ToolsRubeexRemarkV2Scene ToolsRubeexRemarkV2Scene = 2
)

// All allowed values of ToolsRubeexRemarkV2Scene enum
var AllowedToolsRubeexRemarkV2SceneEnumValues = []ToolsRubeexRemarkV2Scene{
	1,
	2,
}

// NewToolsRubeexRemarkV2SceneFromValue returns a pointer to a valid ToolsRubeexRemarkV2Scene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsRubeexRemarkV2SceneFromValue(v int64) (*ToolsRubeexRemarkV2Scene, error) {
	ev := ToolsRubeexRemarkV2Scene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsRubeexRemarkV2Scene: valid values are %v", v, AllowedToolsRubeexRemarkV2SceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsRubeexRemarkV2Scene) IsValid() bool {
	for _, existing := range AllowedToolsRubeexRemarkV2SceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_rubeex_remark_v2_scene value
func (v ToolsRubeexRemarkV2Scene) Ptr() *ToolsRubeexRemarkV2Scene {
	return &v
}
