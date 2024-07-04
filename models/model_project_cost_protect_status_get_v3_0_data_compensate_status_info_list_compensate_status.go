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

// ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus
type ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus string

// List of project_cost_protect_status_get_v3.0_data_compensate_status_info_list_compensate_status
const (
	CONFIRMING_ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus  ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus = "CONFIRMING"
	DEFAULT_ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus     ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus = "DEFAULT"
	ENDED_ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus       ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus = "ENDED"
	FAILED_ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus      ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus = "FAILED"
	INVALID_ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus     ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus = "INVALID"
	IN_EFFECT_ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus   ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus = "IN_EFFECT"
	PAID_ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus        ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus = "PAID"
	UNSUPPORTED_ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus = "UNSUPPORTED"
)

// All allowed values of ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus enum
var AllowedProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatusEnumValues = []ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus{
	"CONFIRMING",
	"DEFAULT",
	"ENDED",
	"FAILED",
	"INVALID",
	"IN_EFFECT",
	"PAID",
	"UNSUPPORTED",
}

// NewProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatusFromValue returns a pointer to a valid ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatusFromValue(v string) (*ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus, error) {
	ev := ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus: valid values are %v", v, AllowedProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus) IsValid() bool {
	for _, existing := range AllowedProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_cost_protect_status_get_v3.0_data_compensate_status_info_list_compensate_status value
func (v ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus) Ptr() *ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus {
	return &v
}
