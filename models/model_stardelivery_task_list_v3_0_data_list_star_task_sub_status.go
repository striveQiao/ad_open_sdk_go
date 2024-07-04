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

// StardeliveryTaskListV30DataListStarTaskSubStatus
type StardeliveryTaskListV30DataListStarTaskSubStatus string

// List of stardelivery_task_list_v3.0_data_list_star_task_sub_status
const (
	AWAITING_ISV_ACCEPT_StardeliveryTaskListV30DataListStarTaskSubStatus    StardeliveryTaskListV30DataListStarTaskSubStatus = "AWAITING_ISV_ACCEPT"
	CALCULATING_COSTS_StardeliveryTaskListV30DataListStarTaskSubStatus      StardeliveryTaskListV30DataListStarTaskSubStatus = "CALCULATING_COSTS"
	NO_ISV_ACCEPT_StardeliveryTaskListV30DataListStarTaskSubStatus          StardeliveryTaskListV30DataListStarTaskSubStatus = "NO_ISV_ACCEPT"
	REJECTED_StardeliveryTaskListV30DataListStarTaskSubStatus               StardeliveryTaskListV30DataListStarTaskSubStatus = "REJECTED"
	SUBMISSION_IN_PROGRESS_StardeliveryTaskListV30DataListStarTaskSubStatus StardeliveryTaskListV30DataListStarTaskSubStatus = "SUBMISSION_IN_PROGRESS"
	UNDER_REVIEW_StardeliveryTaskListV30DataListStarTaskSubStatus           StardeliveryTaskListV30DataListStarTaskSubStatus = "UNDER_REVIEW"
)

// All allowed values of StardeliveryTaskListV30DataListStarTaskSubStatus enum
var AllowedStardeliveryTaskListV30DataListStarTaskSubStatusEnumValues = []StardeliveryTaskListV30DataListStarTaskSubStatus{
	"AWAITING_ISV_ACCEPT",
	"CALCULATING_COSTS",
	"NO_ISV_ACCEPT",
	"REJECTED",
	"SUBMISSION_IN_PROGRESS",
	"UNDER_REVIEW",
}

// NewStardeliveryTaskListV30DataListStarTaskSubStatusFromValue returns a pointer to a valid StardeliveryTaskListV30DataListStarTaskSubStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStardeliveryTaskListV30DataListStarTaskSubStatusFromValue(v string) (*StardeliveryTaskListV30DataListStarTaskSubStatus, error) {
	ev := StardeliveryTaskListV30DataListStarTaskSubStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StardeliveryTaskListV30DataListStarTaskSubStatus: valid values are %v", v, AllowedStardeliveryTaskListV30DataListStarTaskSubStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StardeliveryTaskListV30DataListStarTaskSubStatus) IsValid() bool {
	for _, existing := range AllowedStardeliveryTaskListV30DataListStarTaskSubStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to stardelivery_task_list_v3.0_data_list_star_task_sub_status value
func (v StardeliveryTaskListV30DataListStarTaskSubStatus) Ptr() *StardeliveryTaskListV30DataListStarTaskSubStatus {
	return &v
}
