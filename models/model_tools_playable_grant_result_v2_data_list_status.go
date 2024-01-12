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

// ToolsPlayableGrantResultV2DataListStatus
type ToolsPlayableGrantResultV2DataListStatus string

// List of tools_playable_grant_result_v2_data_list_status
const (
	FAILED_ToolsPlayableGrantResultV2DataListStatus  ToolsPlayableGrantResultV2DataListStatus = "FAILED"
	RUNNING_ToolsPlayableGrantResultV2DataListStatus ToolsPlayableGrantResultV2DataListStatus = "RUNNING"
	SUCCESS_ToolsPlayableGrantResultV2DataListStatus ToolsPlayableGrantResultV2DataListStatus = "SUCCESS"
)

// All allowed values of ToolsPlayableGrantResultV2DataListStatus enum
var AllowedToolsPlayableGrantResultV2DataListStatusEnumValues = []ToolsPlayableGrantResultV2DataListStatus{
	"FAILED",
	"RUNNING",
	"SUCCESS",
}

// NewToolsPlayableGrantResultV2DataListStatusFromValue returns a pointer to a valid ToolsPlayableGrantResultV2DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPlayableGrantResultV2DataListStatusFromValue(v string) (*ToolsPlayableGrantResultV2DataListStatus, error) {
	ev := ToolsPlayableGrantResultV2DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPlayableGrantResultV2DataListStatus: valid values are %v", v, AllowedToolsPlayableGrantResultV2DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPlayableGrantResultV2DataListStatus) IsValid() bool {
	for _, existing := range AllowedToolsPlayableGrantResultV2DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_playable_grant_result_v2_data_list_status value
func (v ToolsPlayableGrantResultV2DataListStatus) Ptr() *ToolsPlayableGrantResultV2DataListStatus {
	return &v
}
