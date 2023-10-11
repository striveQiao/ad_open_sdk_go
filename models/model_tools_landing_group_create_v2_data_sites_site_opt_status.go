/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsLandingGroupCreateV2DataSitesSiteOptStatus
type ToolsLandingGroupCreateV2DataSitesSiteOptStatus string

// List of tools_landing_group_create_v2_data_sites_site_opt_status
const (
	SITE_OPT_STATUS_DISABLE_ToolsLandingGroupCreateV2DataSitesSiteOptStatus ToolsLandingGroupCreateV2DataSitesSiteOptStatus = "SITE_OPT_STATUS_DISABLE"
	SITE_OPT_STATUS_ENABLE_ToolsLandingGroupCreateV2DataSitesSiteOptStatus  ToolsLandingGroupCreateV2DataSitesSiteOptStatus = "SITE_OPT_STATUS_ENABLE"
)

// All allowed values of ToolsLandingGroupCreateV2DataSitesSiteOptStatus enum
var AllowedToolsLandingGroupCreateV2DataSitesSiteOptStatusEnumValues = []ToolsLandingGroupCreateV2DataSitesSiteOptStatus{
	"SITE_OPT_STATUS_DISABLE",
	"SITE_OPT_STATUS_ENABLE",
}

// NewToolsLandingGroupCreateV2DataSitesSiteOptStatusFromValue returns a pointer to a valid ToolsLandingGroupCreateV2DataSitesSiteOptStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsLandingGroupCreateV2DataSitesSiteOptStatusFromValue(v string) (*ToolsLandingGroupCreateV2DataSitesSiteOptStatus, error) {
	ev := ToolsLandingGroupCreateV2DataSitesSiteOptStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsLandingGroupCreateV2DataSitesSiteOptStatus: valid values are %v", v, AllowedToolsLandingGroupCreateV2DataSitesSiteOptStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsLandingGroupCreateV2DataSitesSiteOptStatus) IsValid() bool {
	for _, existing := range AllowedToolsLandingGroupCreateV2DataSitesSiteOptStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_landing_group_create_v2_data_sites_site_opt_status value
func (v ToolsLandingGroupCreateV2DataSitesSiteOptStatus) Ptr() *ToolsLandingGroupCreateV2DataSitesSiteOptStatus {
	return &v
}
