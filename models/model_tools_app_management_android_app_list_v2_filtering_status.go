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

// ToolsAppManagementAndroidAppListV2FilteringStatus
type ToolsAppManagementAndroidAppListV2FilteringStatus string

// List of tools_app_management_android_app_list_v2_filtering_status
const (
	ALL_ToolsAppManagementAndroidAppListV2FilteringStatus            ToolsAppManagementAndroidAppListV2FilteringStatus = "ALL"
	AUDIT_DOING_ToolsAppManagementAndroidAppListV2FilteringStatus    ToolsAppManagementAndroidAppListV2FilteringStatus = "AUDIT_DOING"
	AUDIT_REJECTED_ToolsAppManagementAndroidAppListV2FilteringStatus ToolsAppManagementAndroidAppListV2FilteringStatus = "AUDIT_REJECTED"
	AUDIT_ACCEPTED_ToolsAppManagementAndroidAppListV2FilteringStatus ToolsAppManagementAndroidAppListV2FilteringStatus = "AUDIT_ACCEPTED"
	ENABLE_ToolsAppManagementAndroidAppListV2FilteringStatus         ToolsAppManagementAndroidAppListV2FilteringStatus = "ENABLE"
)

// All allowed values of ToolsAppManagementAndroidAppListV2FilteringStatus enum
var AllowedToolsAppManagementAndroidAppListV2FilteringStatusEnumValues = []ToolsAppManagementAndroidAppListV2FilteringStatus{
	"ALL",
	"AUDIT_DOING",
	"AUDIT_REJECTED",
	"AUDIT_ACCEPTED",
	"ENABLE",
}

// NewToolsAppManagementAndroidAppListV2FilteringStatusFromValue returns a pointer to a valid ToolsAppManagementAndroidAppListV2FilteringStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementAndroidAppListV2FilteringStatusFromValue(v string) (*ToolsAppManagementAndroidAppListV2FilteringStatus, error) {
	ev := ToolsAppManagementAndroidAppListV2FilteringStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementAndroidAppListV2FilteringStatus: valid values are %v", v, AllowedToolsAppManagementAndroidAppListV2FilteringStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementAndroidAppListV2FilteringStatus) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementAndroidAppListV2FilteringStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_android_app_list_v2_filtering_status value
func (v ToolsAppManagementAndroidAppListV2FilteringStatus) Ptr() *ToolsAppManagementAndroidAppListV2FilteringStatus {
	return &v
}
