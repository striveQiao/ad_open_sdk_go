/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAbTestListGetV2FilterStatus
type ToolsAbTestListGetV2FilterStatus string

// List of tools_ab_test_list_get_v2_filter_status
const (
	CREATED_ToolsAbTestListGetV2FilterStatus    ToolsAbTestListGetV2FilterStatus = "CREATED"
	FAILED_ToolsAbTestListGetV2FilterStatus     ToolsAbTestListGetV2FilterStatus = "FAILED"
	FINISH_ToolsAbTestListGetV2FilterStatus     ToolsAbTestListGetV2FilterStatus = "FINISH"
	PROCESSING_ToolsAbTestListGetV2FilterStatus ToolsAbTestListGetV2FilterStatus = "PROCESSING"
)

// All allowed values of ToolsAbTestListGetV2FilterStatus enum
var AllowedToolsAbTestListGetV2FilterStatusEnumValues = []ToolsAbTestListGetV2FilterStatus{
	"CREATED",
	"FAILED",
	"FINISH",
	"PROCESSING",
}

// NewToolsAbTestListGetV2FilterStatusFromValue returns a pointer to a valid ToolsAbTestListGetV2FilterStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAbTestListGetV2FilterStatusFromValue(v string) (*ToolsAbTestListGetV2FilterStatus, error) {
	ev := ToolsAbTestListGetV2FilterStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAbTestListGetV2FilterStatus: valid values are %v", v, AllowedToolsAbTestListGetV2FilterStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAbTestListGetV2FilterStatus) IsValid() bool {
	for _, existing := range AllowedToolsAbTestListGetV2FilterStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ab_test_list_get_v2_filter_status value
func (v ToolsAbTestListGetV2FilterStatus) Ptr() *ToolsAbTestListGetV2FilterStatus {
	return &v
}
