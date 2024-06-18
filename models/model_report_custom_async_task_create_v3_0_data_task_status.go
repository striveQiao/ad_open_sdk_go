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

// ReportCustomAsyncTaskCreateV30DataTaskStatus
type ReportCustomAsyncTaskCreateV30DataTaskStatus string

// List of report_custom_async_task_create_v3.0_data_task_status
const (
	ASYNC_TASK_STATUS_COMPLETED_ReportCustomAsyncTaskCreateV30DataTaskStatus ReportCustomAsyncTaskCreateV30DataTaskStatus = "ASYNC_TASK_STATUS_COMPLETED"
	ASYNC_TASK_STATUS_CREATED_ReportCustomAsyncTaskCreateV30DataTaskStatus   ReportCustomAsyncTaskCreateV30DataTaskStatus = "ASYNC_TASK_STATUS_CREATED"
	ASYNC_TASK_STATUS_EXECUTING_ReportCustomAsyncTaskCreateV30DataTaskStatus ReportCustomAsyncTaskCreateV30DataTaskStatus = "ASYNC_TASK_STATUS_EXECUTING"
	ASYNC_TASK_STATUS_FAILED_ReportCustomAsyncTaskCreateV30DataTaskStatus    ReportCustomAsyncTaskCreateV30DataTaskStatus = "ASYNC_TASK_STATUS_FAILED"
)

// All allowed values of ReportCustomAsyncTaskCreateV30DataTaskStatus enum
var AllowedReportCustomAsyncTaskCreateV30DataTaskStatusEnumValues = []ReportCustomAsyncTaskCreateV30DataTaskStatus{
	"ASYNC_TASK_STATUS_COMPLETED",
	"ASYNC_TASK_STATUS_CREATED",
	"ASYNC_TASK_STATUS_EXECUTING",
	"ASYNC_TASK_STATUS_FAILED",
}

// NewReportCustomAsyncTaskCreateV30DataTaskStatusFromValue returns a pointer to a valid ReportCustomAsyncTaskCreateV30DataTaskStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCustomAsyncTaskCreateV30DataTaskStatusFromValue(v string) (*ReportCustomAsyncTaskCreateV30DataTaskStatus, error) {
	ev := ReportCustomAsyncTaskCreateV30DataTaskStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCustomAsyncTaskCreateV30DataTaskStatus: valid values are %v", v, AllowedReportCustomAsyncTaskCreateV30DataTaskStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCustomAsyncTaskCreateV30DataTaskStatus) IsValid() bool {
	for _, existing := range AllowedReportCustomAsyncTaskCreateV30DataTaskStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_custom_async_task_create_v3.0_data_task_status value
func (v ReportCustomAsyncTaskCreateV30DataTaskStatus) Ptr() *ReportCustomAsyncTaskCreateV30DataTaskStatus {
	return &v
}
