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

// ReportCustomAsyncTaskCreateV30DataTopic
type ReportCustomAsyncTaskCreateV30DataTopic string

// List of report_custom_async_task_create_v3.0_data_topic
const (
	BASIC_DATA_ReportCustomAsyncTaskCreateV30DataTopic    ReportCustomAsyncTaskCreateV30DataTopic = "BASIC_DATA"
	BIDWORD_DATA_ReportCustomAsyncTaskCreateV30DataTopic  ReportCustomAsyncTaskCreateV30DataTopic = "BIDWORD_DATA"
	MATERIAL_DATA_ReportCustomAsyncTaskCreateV30DataTopic ReportCustomAsyncTaskCreateV30DataTopic = "MATERIAL_DATA"
	PRODUCT_DATA_ReportCustomAsyncTaskCreateV30DataTopic  ReportCustomAsyncTaskCreateV30DataTopic = "PRODUCT_DATA"
	QUERY_DATA_ReportCustomAsyncTaskCreateV30DataTopic    ReportCustomAsyncTaskCreateV30DataTopic = "QUERY_DATA"
)

// All allowed values of ReportCustomAsyncTaskCreateV30DataTopic enum
var AllowedReportCustomAsyncTaskCreateV30DataTopicEnumValues = []ReportCustomAsyncTaskCreateV30DataTopic{
	"BASIC_DATA",
	"BIDWORD_DATA",
	"MATERIAL_DATA",
	"PRODUCT_DATA",
	"QUERY_DATA",
}

// NewReportCustomAsyncTaskCreateV30DataTopicFromValue returns a pointer to a valid ReportCustomAsyncTaskCreateV30DataTopic
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCustomAsyncTaskCreateV30DataTopicFromValue(v string) (*ReportCustomAsyncTaskCreateV30DataTopic, error) {
	ev := ReportCustomAsyncTaskCreateV30DataTopic(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCustomAsyncTaskCreateV30DataTopic: valid values are %v", v, AllowedReportCustomAsyncTaskCreateV30DataTopicEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCustomAsyncTaskCreateV30DataTopic) IsValid() bool {
	for _, existing := range AllowedReportCustomAsyncTaskCreateV30DataTopicEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_custom_async_task_create_v3.0_data_topic value
func (v ReportCustomAsyncTaskCreateV30DataTopic) Ptr() *ReportCustomAsyncTaskCreateV30DataTopic {
	return &v
}
