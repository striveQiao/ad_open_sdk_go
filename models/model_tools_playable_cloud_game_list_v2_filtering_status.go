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

// ToolsPlayableCloudGameListV2FilteringStatus
type ToolsPlayableCloudGameListV2FilteringStatus string

// List of tools_playable_cloud_game_list_v2_filtering_status
const (
	AUDIT_FAIL_ToolsPlayableCloudGameListV2FilteringStatus    ToolsPlayableCloudGameListV2FilteringStatus = "AUDIT_FAIL"
	OFF_SHELF_ToolsPlayableCloudGameListV2FilteringStatus     ToolsPlayableCloudGameListV2FilteringStatus = "OFF_SHELF"
	AUDIT_SUCCESS_ToolsPlayableCloudGameListV2FilteringStatus ToolsPlayableCloudGameListV2FilteringStatus = "AUDIT_SUCCESS"
	ON_SHELF_ToolsPlayableCloudGameListV2FilteringStatus      ToolsPlayableCloudGameListV2FilteringStatus = "ON_SHELF"
	UPLOAD_STATUS_ToolsPlayableCloudGameListV2FilteringStatus ToolsPlayableCloudGameListV2FilteringStatus = "UPLOAD_STATUS"
)

// All allowed values of ToolsPlayableCloudGameListV2FilteringStatus enum
var AllowedToolsPlayableCloudGameListV2FilteringStatusEnumValues = []ToolsPlayableCloudGameListV2FilteringStatus{
	"AUDIT_FAIL",
	"OFF_SHELF",
	"AUDIT_SUCCESS",
	"ON_SHELF",
	"UPLOAD_STATUS",
}

// NewToolsPlayableCloudGameListV2FilteringStatusFromValue returns a pointer to a valid ToolsPlayableCloudGameListV2FilteringStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPlayableCloudGameListV2FilteringStatusFromValue(v string) (*ToolsPlayableCloudGameListV2FilteringStatus, error) {
	ev := ToolsPlayableCloudGameListV2FilteringStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPlayableCloudGameListV2FilteringStatus: valid values are %v", v, AllowedToolsPlayableCloudGameListV2FilteringStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPlayableCloudGameListV2FilteringStatus) IsValid() bool {
	for _, existing := range AllowedToolsPlayableCloudGameListV2FilteringStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_playable_cloud_game_list_v2_filtering_status value
func (v ToolsPlayableCloudGameListV2FilteringStatus) Ptr() *ToolsPlayableCloudGameListV2FilteringStatus {
	return &v
}
