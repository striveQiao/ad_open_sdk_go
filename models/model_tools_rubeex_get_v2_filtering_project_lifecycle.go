/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsRubeexGetV2FilteringProjectLifecycle
type ToolsRubeexGetV2FilteringProjectLifecycle string

// List of tools_rubeex_get_v2_filtering_project_lifecycle
const (
	AUDIT_SUCCESS_ToolsRubeexGetV2FilteringProjectLifecycle ToolsRubeexGetV2FilteringProjectLifecycle = "AUDIT_SUCCESS"
	LAUNCHED_ToolsRubeexGetV2FilteringProjectLifecycle      ToolsRubeexGetV2FilteringProjectLifecycle = "LAUNCHED"
	EDITING_ToolsRubeexGetV2FilteringProjectLifecycle       ToolsRubeexGetV2FilteringProjectLifecycle = "EDITING"
	SYNC_AD_ToolsRubeexGetV2FilteringProjectLifecycle       ToolsRubeexGetV2FilteringProjectLifecycle = "SYNC_AD"
	RELAT_PLAN_ToolsRubeexGetV2FilteringProjectLifecycle    ToolsRubeexGetV2FilteringProjectLifecycle = "RELAT_PLAN"
)

// All allowed values of ToolsRubeexGetV2FilteringProjectLifecycle enum
var AllowedToolsRubeexGetV2FilteringProjectLifecycleEnumValues = []ToolsRubeexGetV2FilteringProjectLifecycle{
	"AUDIT_SUCCESS",
	"LAUNCHED",
	"EDITING",
	"SYNC_AD",
	"RELAT_PLAN",
}

// NewToolsRubeexGetV2FilteringProjectLifecycleFromValue returns a pointer to a valid ToolsRubeexGetV2FilteringProjectLifecycle
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsRubeexGetV2FilteringProjectLifecycleFromValue(v string) (*ToolsRubeexGetV2FilteringProjectLifecycle, error) {
	ev := ToolsRubeexGetV2FilteringProjectLifecycle(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsRubeexGetV2FilteringProjectLifecycle: valid values are %v", v, AllowedToolsRubeexGetV2FilteringProjectLifecycleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsRubeexGetV2FilteringProjectLifecycle) IsValid() bool {
	for _, existing := range AllowedToolsRubeexGetV2FilteringProjectLifecycleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_rubeex_get_v2_filtering_project_lifecycle value
func (v ToolsRubeexGetV2FilteringProjectLifecycle) Ptr() *ToolsRubeexGetV2FilteringProjectLifecycle {
	return &v
}
