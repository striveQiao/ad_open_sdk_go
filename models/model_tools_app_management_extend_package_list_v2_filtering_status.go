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

// ToolsAppManagementExtendPackageListV2FilteringStatus
type ToolsAppManagementExtendPackageListV2FilteringStatus string

// List of tools_app_management_extend_package_list_v2_filtering_status
const (
	CREATION_FAILED_ToolsAppManagementExtendPackageListV2FilteringStatus ToolsAppManagementExtendPackageListV2FilteringStatus = "CREATION_FAILED"
	UPDATING_ToolsAppManagementExtendPackageListV2FilteringStatus        ToolsAppManagementExtendPackageListV2FilteringStatus = "UPDATING"
	ALL_ToolsAppManagementExtendPackageListV2FilteringStatus             ToolsAppManagementExtendPackageListV2FilteringStatus = "ALL"
	CREATING_ToolsAppManagementExtendPackageListV2FilteringStatus        ToolsAppManagementExtendPackageListV2FilteringStatus = "CREATING"
	PUBLISHED_ToolsAppManagementExtendPackageListV2FilteringStatus       ToolsAppManagementExtendPackageListV2FilteringStatus = "PUBLISHED"
	UPDATE_FAILED_ToolsAppManagementExtendPackageListV2FilteringStatus   ToolsAppManagementExtendPackageListV2FilteringStatus = "UPDATE_FAILED"
	NOT_UPDATE_ToolsAppManagementExtendPackageListV2FilteringStatus      ToolsAppManagementExtendPackageListV2FilteringStatus = "NOT_UPDATE"
)

// All allowed values of ToolsAppManagementExtendPackageListV2FilteringStatus enum
var AllowedToolsAppManagementExtendPackageListV2FilteringStatusEnumValues = []ToolsAppManagementExtendPackageListV2FilteringStatus{
	"CREATION_FAILED",
	"UPDATING",
	"ALL",
	"CREATING",
	"PUBLISHED",
	"UPDATE_FAILED",
	"NOT_UPDATE",
}

// NewToolsAppManagementExtendPackageListV2FilteringStatusFromValue returns a pointer to a valid ToolsAppManagementExtendPackageListV2FilteringStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementExtendPackageListV2FilteringStatusFromValue(v string) (*ToolsAppManagementExtendPackageListV2FilteringStatus, error) {
	ev := ToolsAppManagementExtendPackageListV2FilteringStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementExtendPackageListV2FilteringStatus: valid values are %v", v, AllowedToolsAppManagementExtendPackageListV2FilteringStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementExtendPackageListV2FilteringStatus) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementExtendPackageListV2FilteringStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_extend_package_list_v2_filtering_status value
func (v ToolsAppManagementExtendPackageListV2FilteringStatus) Ptr() *ToolsAppManagementExtendPackageListV2FilteringStatus {
	return &v
}
