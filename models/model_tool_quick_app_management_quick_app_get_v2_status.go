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

// ToolQuickAppManagementQuickAppGetV2Status
type ToolQuickAppManagementQuickAppGetV2Status string

// List of tool_quick_app_management_quick_app_get_v2_status
const (
	AUDIT_ACCEPTED_ToolQuickAppManagementQuickAppGetV2Status      ToolQuickAppManagementQuickAppGetV2Status = "AUDIT_ACCEPTED"
	AUDIT_DOING_ToolQuickAppManagementQuickAppGetV2Status         ToolQuickAppManagementQuickAppGetV2Status = "AUDIT_DOING"
	AUDIT_REJECTED_ToolQuickAppManagementQuickAppGetV2Status      ToolQuickAppManagementQuickAppGetV2Status = "AUDIT_REJECTED"
	AUDIT_SEND_REJECTED_ToolQuickAppManagementQuickAppGetV2Status ToolQuickAppManagementQuickAppGetV2Status = "AUDIT_SEND_REJECTED"
	REMOVED_ToolQuickAppManagementQuickAppGetV2Status             ToolQuickAppManagementQuickAppGetV2Status = "REMOVED"
)

// All allowed values of ToolQuickAppManagementQuickAppGetV2Status enum
var AllowedToolQuickAppManagementQuickAppGetV2StatusEnumValues = []ToolQuickAppManagementQuickAppGetV2Status{
	"AUDIT_ACCEPTED",
	"AUDIT_DOING",
	"AUDIT_REJECTED",
	"AUDIT_SEND_REJECTED",
	"REMOVED",
}

// NewToolQuickAppManagementQuickAppGetV2StatusFromValue returns a pointer to a valid ToolQuickAppManagementQuickAppGetV2Status
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolQuickAppManagementQuickAppGetV2StatusFromValue(v string) (*ToolQuickAppManagementQuickAppGetV2Status, error) {
	ev := ToolQuickAppManagementQuickAppGetV2Status(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolQuickAppManagementQuickAppGetV2Status: valid values are %v", v, AllowedToolQuickAppManagementQuickAppGetV2StatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolQuickAppManagementQuickAppGetV2Status) IsValid() bool {
	for _, existing := range AllowedToolQuickAppManagementQuickAppGetV2StatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tool_quick_app_management_quick_app_get_v2_status value
func (v ToolQuickAppManagementQuickAppGetV2Status) Ptr() *ToolQuickAppManagementQuickAppGetV2Status {
	return &v
}
