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

// StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus
type StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus string

// List of star_vas_get_boost_item_group_detail_v2_data_task_info_status
const (
	AUDITING_StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus     StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus = "AUDITING"
	CANCEL_StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus       StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus = "CANCEL"
	CLOSE_StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus        StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus = "CLOSE"
	DOING_StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus        StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus = "DOING"
	FAIL_CREATE_StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus  StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus = "FAIL_CREATE"
	FINISH_StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus       StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus = "FINISH"
	PREPARE_StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus      StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus = "PREPARE"
	REJECT_AUDIT_StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus = "REJECT_AUDIT"
	SETTLE_StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus       StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus = "SETTLE"
	WAIT_CANCEL_StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus  StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus = "WAIT_CANCEL"
)

// All allowed values of StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus enum
var AllowedStarVasGetBoostItemGroupDetailV2DataTaskInfoStatusEnumValues = []StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus{
	"AUDITING",
	"CANCEL",
	"CLOSE",
	"DOING",
	"FAIL_CREATE",
	"FINISH",
	"PREPARE",
	"REJECT_AUDIT",
	"SETTLE",
	"WAIT_CANCEL",
}

// NewStarVasGetBoostItemGroupDetailV2DataTaskInfoStatusFromValue returns a pointer to a valid StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarVasGetBoostItemGroupDetailV2DataTaskInfoStatusFromValue(v string) (*StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus, error) {
	ev := StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus: valid values are %v", v, AllowedStarVasGetBoostItemGroupDetailV2DataTaskInfoStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus) IsValid() bool {
	for _, existing := range AllowedStarVasGetBoostItemGroupDetailV2DataTaskInfoStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_vas_get_boost_item_group_detail_v2_data_task_info_status value
func (v StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus) Ptr() *StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus {
	return &v
}
