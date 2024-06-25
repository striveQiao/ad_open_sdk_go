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

// ToolsAppManagementExtendPackageListV2V2FilteringStatus
type ToolsAppManagementExtendPackageListV2V2FilteringStatus string

// List of tools_app_management_extend_package_list_v2_v2_filtering_status
const (
	ALL_ToolsAppManagementExtendPackageListV2V2FilteringStatus             ToolsAppManagementExtendPackageListV2V2FilteringStatus = "ALL"
	NOT_UPDATE_ToolsAppManagementExtendPackageListV2V2FilteringStatus      ToolsAppManagementExtendPackageListV2V2FilteringStatus = "NOT_UPDATE"
	CREATING_ToolsAppManagementExtendPackageListV2V2FilteringStatus        ToolsAppManagementExtendPackageListV2V2FilteringStatus = "CREATING"
	UPDATING_ToolsAppManagementExtendPackageListV2V2FilteringStatus        ToolsAppManagementExtendPackageListV2V2FilteringStatus = "UPDATING"
	ENABLE_ToolsAppManagementExtendPackageListV2V2FilteringStatus          ToolsAppManagementExtendPackageListV2V2FilteringStatus = "ENABLE"
	CREATION_FAILED_ToolsAppManagementExtendPackageListV2V2FilteringStatus ToolsAppManagementExtendPackageListV2V2FilteringStatus = "CREATION_FAILED"
	UPDATE_FAILED_ToolsAppManagementExtendPackageListV2V2FilteringStatus   ToolsAppManagementExtendPackageListV2V2FilteringStatus = "UPDATE_FAILED"
	PUBLISHED_ToolsAppManagementExtendPackageListV2V2FilteringStatus       ToolsAppManagementExtendPackageListV2V2FilteringStatus = "PUBLISHED"
)

// All allowed values of ToolsAppManagementExtendPackageListV2V2FilteringStatus enum
var AllowedToolsAppManagementExtendPackageListV2V2FilteringStatusEnumValues = []ToolsAppManagementExtendPackageListV2V2FilteringStatus{
	"ALL",
	"NOT_UPDATE",
	"CREATING",
	"UPDATING",
	"ENABLE",
	"CREATION_FAILED",
	"UPDATE_FAILED",
	"PUBLISHED",
}

// NewToolsAppManagementExtendPackageListV2V2FilteringStatusFromValue returns a pointer to a valid ToolsAppManagementExtendPackageListV2V2FilteringStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementExtendPackageListV2V2FilteringStatusFromValue(v string) (*ToolsAppManagementExtendPackageListV2V2FilteringStatus, error) {
	ev := ToolsAppManagementExtendPackageListV2V2FilteringStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementExtendPackageListV2V2FilteringStatus: valid values are %v", v, AllowedToolsAppManagementExtendPackageListV2V2FilteringStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementExtendPackageListV2V2FilteringStatus) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementExtendPackageListV2V2FilteringStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_extend_package_list_v2_v2_filtering_status value
func (v ToolsAppManagementExtendPackageListV2V2FilteringStatus) Ptr() *ToolsAppManagementExtendPackageListV2V2FilteringStatus {
	return &v
}
