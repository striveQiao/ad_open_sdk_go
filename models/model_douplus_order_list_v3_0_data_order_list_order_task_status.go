/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DouplusOrderListV30DataOrderListOrderTaskStatus
type DouplusOrderListV30DataOrderListOrderTaskStatus string

// List of douplus_order_list_v3.0_data_order_list_order_task_status
const (
	AUDITING_DouplusOrderListV30DataOrderListOrderTaskStatus      DouplusOrderListV30DataOrderListOrderTaskStatus = "AUDITING"
	AUDIT_PAUSE_DouplusOrderListV30DataOrderListOrderTaskStatus   DouplusOrderListV30DataOrderListOrderTaskStatus = "AUDIT_PAUSE"
	DELIVERIED_DouplusOrderListV30DataOrderListOrderTaskStatus    DouplusOrderListV30DataOrderListOrderTaskStatus = "DELIVERIED"
	DELIVERING_DouplusOrderListV30DataOrderListOrderTaskStatus    DouplusOrderListV30DataOrderListOrderTaskStatus = "DELIVERING"
	OFFLINE_AUDIT_DouplusOrderListV30DataOrderListOrderTaskStatus DouplusOrderListV30DataOrderListOrderTaskStatus = "OFFLINE_AUDIT"
	TIME_NO_REACH_DouplusOrderListV30DataOrderListOrderTaskStatus DouplusOrderListV30DataOrderListOrderTaskStatus = "TIME_NO_REACH"
	UNDELIVERIED_DouplusOrderListV30DataOrderListOrderTaskStatus  DouplusOrderListV30DataOrderListOrderTaskStatus = "UNDELIVERIED"
	UNPAID_DouplusOrderListV30DataOrderListOrderTaskStatus        DouplusOrderListV30DataOrderListOrderTaskStatus = "UNPAID"
	WAIT_TO_START_DouplusOrderListV30DataOrderListOrderTaskStatus DouplusOrderListV30DataOrderListOrderTaskStatus = "WAIT_TO_START"
)

// All allowed values of DouplusOrderListV30DataOrderListOrderTaskStatus enum
var AllowedDouplusOrderListV30DataOrderListOrderTaskStatusEnumValues = []DouplusOrderListV30DataOrderListOrderTaskStatus{
	"AUDITING",
	"AUDIT_PAUSE",
	"DELIVERIED",
	"DELIVERING",
	"OFFLINE_AUDIT",
	"TIME_NO_REACH",
	"UNDELIVERIED",
	"UNPAID",
	"WAIT_TO_START",
}

// NewDouplusOrderListV30DataOrderListOrderTaskStatusFromValue returns a pointer to a valid DouplusOrderListV30DataOrderListOrderTaskStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDouplusOrderListV30DataOrderListOrderTaskStatusFromValue(v string) (*DouplusOrderListV30DataOrderListOrderTaskStatus, error) {
	ev := DouplusOrderListV30DataOrderListOrderTaskStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DouplusOrderListV30DataOrderListOrderTaskStatus: valid values are %v", v, AllowedDouplusOrderListV30DataOrderListOrderTaskStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DouplusOrderListV30DataOrderListOrderTaskStatus) IsValid() bool {
	for _, existing := range AllowedDouplusOrderListV30DataOrderListOrderTaskStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to douplus_order_list_v3.0_data_order_list_order_task_status value
func (v DouplusOrderListV30DataOrderListOrderTaskStatus) Ptr() *DouplusOrderListV30DataOrderListOrderTaskStatus {
	return &v
}
