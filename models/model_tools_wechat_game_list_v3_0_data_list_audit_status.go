/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsWechatGameListV30DataListAuditStatus
type ToolsWechatGameListV30DataListAuditStatus string

// List of tools_wechat_game_list_v3.0_data_list_audit_status
const (
	AUDIT_ACCEPTED_ToolsWechatGameListV30DataListAuditStatus ToolsWechatGameListV30DataListAuditStatus = "AUDIT_ACCEPTED"
	AUDITING_ToolsWechatGameListV30DataListAuditStatus       ToolsWechatGameListV30DataListAuditStatus = "AUDITING"
	AUDIT_REJECTED_ToolsWechatGameListV30DataListAuditStatus ToolsWechatGameListV30DataListAuditStatus = "AUDIT_REJECTED"
)

// All allowed values of ToolsWechatGameListV30DataListAuditStatus enum
var AllowedToolsWechatGameListV30DataListAuditStatusEnumValues = []ToolsWechatGameListV30DataListAuditStatus{
	"AUDIT_ACCEPTED",
	"AUDITING",
	"AUDIT_REJECTED",
}

// NewToolsWechatGameListV30DataListAuditStatusFromValue returns a pointer to a valid ToolsWechatGameListV30DataListAuditStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsWechatGameListV30DataListAuditStatusFromValue(v string) (*ToolsWechatGameListV30DataListAuditStatus, error) {
	ev := ToolsWechatGameListV30DataListAuditStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsWechatGameListV30DataListAuditStatus: valid values are %v", v, AllowedToolsWechatGameListV30DataListAuditStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsWechatGameListV30DataListAuditStatus) IsValid() bool {
	for _, existing := range AllowedToolsWechatGameListV30DataListAuditStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_wechat_game_list_v3.0_data_list_audit_status value
func (v ToolsWechatGameListV30DataListAuditStatus) Ptr() *ToolsWechatGameListV30DataListAuditStatus {
	return &v
}
