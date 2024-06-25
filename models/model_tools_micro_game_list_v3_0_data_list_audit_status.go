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

// ToolsMicroGameListV30DataListAuditStatus
type ToolsMicroGameListV30DataListAuditStatus string

// List of tools_micro_game_list_v3.0_data_list_audit_status
const (
	AUDIT_ACCEPTED_ToolsMicroGameListV30DataListAuditStatus ToolsMicroGameListV30DataListAuditStatus = "AUDIT_ACCEPTED"
	AUDITING_ToolsMicroGameListV30DataListAuditStatus       ToolsMicroGameListV30DataListAuditStatus = "AUDITING"
	AUDIT_REJECTED_ToolsMicroGameListV30DataListAuditStatus ToolsMicroGameListV30DataListAuditStatus = "AUDIT_REJECTED"
)

// All allowed values of ToolsMicroGameListV30DataListAuditStatus enum
var AllowedToolsMicroGameListV30DataListAuditStatusEnumValues = []ToolsMicroGameListV30DataListAuditStatus{
	"AUDIT_ACCEPTED",
	"AUDITING",
	"AUDIT_REJECTED",
}

// NewToolsMicroGameListV30DataListAuditStatusFromValue returns a pointer to a valid ToolsMicroGameListV30DataListAuditStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsMicroGameListV30DataListAuditStatusFromValue(v string) (*ToolsMicroGameListV30DataListAuditStatus, error) {
	ev := ToolsMicroGameListV30DataListAuditStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsMicroGameListV30DataListAuditStatus: valid values are %v", v, AllowedToolsMicroGameListV30DataListAuditStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsMicroGameListV30DataListAuditStatus) IsValid() bool {
	for _, existing := range AllowedToolsMicroGameListV30DataListAuditStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_micro_game_list_v3.0_data_list_audit_status value
func (v ToolsMicroGameListV30DataListAuditStatus) Ptr() *ToolsMicroGameListV30DataListAuditStatus {
	return &v
}
