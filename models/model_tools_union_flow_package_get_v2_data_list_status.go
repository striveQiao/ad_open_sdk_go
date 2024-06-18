/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsUnionFlowPackageGetV2DataListStatus
type ToolsUnionFlowPackageGetV2DataListStatus string

// List of tools_union_flow_package_get_v2_data_list_status
const (
	FLOW_PACKAGE_ENABLE_ToolsUnionFlowPackageGetV2DataListStatus  ToolsUnionFlowPackageGetV2DataListStatus = "FLOW_PACKAGE_ENABLE"
	FLOW_PACKAGE_DISABLE_ToolsUnionFlowPackageGetV2DataListStatus ToolsUnionFlowPackageGetV2DataListStatus = "FLOW_PACKAGE_DISABLE"
)

// All allowed values of ToolsUnionFlowPackageGetV2DataListStatus enum
var AllowedToolsUnionFlowPackageGetV2DataListStatusEnumValues = []ToolsUnionFlowPackageGetV2DataListStatus{
	"FLOW_PACKAGE_ENABLE",
	"FLOW_PACKAGE_DISABLE",
}

// NewToolsUnionFlowPackageGetV2DataListStatusFromValue returns a pointer to a valid ToolsUnionFlowPackageGetV2DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsUnionFlowPackageGetV2DataListStatusFromValue(v string) (*ToolsUnionFlowPackageGetV2DataListStatus, error) {
	ev := ToolsUnionFlowPackageGetV2DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsUnionFlowPackageGetV2DataListStatus: valid values are %v", v, AllowedToolsUnionFlowPackageGetV2DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsUnionFlowPackageGetV2DataListStatus) IsValid() bool {
	for _, existing := range AllowedToolsUnionFlowPackageGetV2DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_union_flow_package_get_v2_data_list_status value
func (v ToolsUnionFlowPackageGetV2DataListStatus) Ptr() *ToolsUnionFlowPackageGetV2DataListStatus {
	return &v
}
