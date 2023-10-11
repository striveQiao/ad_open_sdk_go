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

// AdlabGroupListV30DataGroupsOptStatus
type AdlabGroupListV30DataGroupsOptStatus string

// List of adlab_group_list_v3.0_data_groups_opt_status
const (
	DELETED_AdlabGroupListV30DataGroupsOptStatus           AdlabGroupListV30DataGroupsOptStatus = "DELETED"
	DISABLED_AdlabGroupListV30DataGroupsOptStatus          AdlabGroupListV30DataGroupsOptStatus = "DISABLED"
	ENABLED_AdlabGroupListV30DataGroupsOptStatus           AdlabGroupListV30DataGroupsOptStatus = "ENABLED"
	REACH_QUOTA_LIMIT_AdlabGroupListV30DataGroupsOptStatus AdlabGroupListV30DataGroupsOptStatus = "REACH_QUOTA_LIMIT"
)

// All allowed values of AdlabGroupListV30DataGroupsOptStatus enum
var AllowedAdlabGroupListV30DataGroupsOptStatusEnumValues = []AdlabGroupListV30DataGroupsOptStatus{
	"DELETED",
	"DISABLED",
	"ENABLED",
	"REACH_QUOTA_LIMIT",
}

// NewAdlabGroupListV30DataGroupsOptStatusFromValue returns a pointer to a valid AdlabGroupListV30DataGroupsOptStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupListV30DataGroupsOptStatusFromValue(v string) (*AdlabGroupListV30DataGroupsOptStatus, error) {
	ev := AdlabGroupListV30DataGroupsOptStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupListV30DataGroupsOptStatus: valid values are %v", v, AllowedAdlabGroupListV30DataGroupsOptStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupListV30DataGroupsOptStatus) IsValid() bool {
	for _, existing := range AllowedAdlabGroupListV30DataGroupsOptStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_list_v3.0_data_groups_opt_status value
func (v AdlabGroupListV30DataGroupsOptStatus) Ptr() *AdlabGroupListV30DataGroupsOptStatus {
	return &v
}
