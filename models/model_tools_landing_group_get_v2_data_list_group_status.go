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

// ToolsLandingGroupGetV2DataListGroupStatus
type ToolsLandingGroupGetV2DataListGroupStatus string

// List of tools_landing_group_get_v2_data_list_group_status
const (
	LANDING_GROUP_STATUS_AVAILABLE_ToolsLandingGroupGetV2DataListGroupStatus   ToolsLandingGroupGetV2DataListGroupStatus = "LANDING_GROUP_STATUS_AVAILABLE"
	LANDING_GROUP_STATUS_UNAVAILABLE_ToolsLandingGroupGetV2DataListGroupStatus ToolsLandingGroupGetV2DataListGroupStatus = "LANDING_GROUP_STATUS_UNAVAILABLE"
)

// All allowed values of ToolsLandingGroupGetV2DataListGroupStatus enum
var AllowedToolsLandingGroupGetV2DataListGroupStatusEnumValues = []ToolsLandingGroupGetV2DataListGroupStatus{
	"LANDING_GROUP_STATUS_AVAILABLE",
	"LANDING_GROUP_STATUS_UNAVAILABLE",
}

// NewToolsLandingGroupGetV2DataListGroupStatusFromValue returns a pointer to a valid ToolsLandingGroupGetV2DataListGroupStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsLandingGroupGetV2DataListGroupStatusFromValue(v string) (*ToolsLandingGroupGetV2DataListGroupStatus, error) {
	ev := ToolsLandingGroupGetV2DataListGroupStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsLandingGroupGetV2DataListGroupStatus: valid values are %v", v, AllowedToolsLandingGroupGetV2DataListGroupStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsLandingGroupGetV2DataListGroupStatus) IsValid() bool {
	for _, existing := range AllowedToolsLandingGroupGetV2DataListGroupStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_landing_group_get_v2_data_list_group_status value
func (v ToolsLandingGroupGetV2DataListGroupStatus) Ptr() *ToolsLandingGroupGetV2DataListGroupStatus {
	return &v
}
