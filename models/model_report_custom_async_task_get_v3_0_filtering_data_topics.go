/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportCustomAsyncTaskGetV30FilteringDataTopics
type ReportCustomAsyncTaskGetV30FilteringDataTopics string

// List of report_custom_async_task_get_v3.0_filtering_data_topics
const (
	BASIC_DATA_ReportCustomAsyncTaskGetV30FilteringDataTopics         ReportCustomAsyncTaskGetV30FilteringDataTopics = "BASIC_DATA"
	BIDWORD_DATA_ReportCustomAsyncTaskGetV30FilteringDataTopics       ReportCustomAsyncTaskGetV30FilteringDataTopics = "BIDWORD_DATA"
	MATERIAL_DATA_ReportCustomAsyncTaskGetV30FilteringDataTopics      ReportCustomAsyncTaskGetV30FilteringDataTopics = "MATERIAL_DATA"
	ONE_KEY_BOOST_DATA_ReportCustomAsyncTaskGetV30FilteringDataTopics ReportCustomAsyncTaskGetV30FilteringDataTopics = "ONE_KEY_BOOST_DATA"
	PRODUCT_DATA_ReportCustomAsyncTaskGetV30FilteringDataTopics       ReportCustomAsyncTaskGetV30FilteringDataTopics = "PRODUCT_DATA"
	QUERY_DATA_ReportCustomAsyncTaskGetV30FilteringDataTopics         ReportCustomAsyncTaskGetV30FilteringDataTopics = "QUERY_DATA"
)

// All allowed values of ReportCustomAsyncTaskGetV30FilteringDataTopics enum
var AllowedReportCustomAsyncTaskGetV30FilteringDataTopicsEnumValues = []ReportCustomAsyncTaskGetV30FilteringDataTopics{
	"BASIC_DATA",
	"BIDWORD_DATA",
	"MATERIAL_DATA",
	"ONE_KEY_BOOST_DATA",
	"PRODUCT_DATA",
	"QUERY_DATA",
}

// NewReportCustomAsyncTaskGetV30FilteringDataTopicsFromValue returns a pointer to a valid ReportCustomAsyncTaskGetV30FilteringDataTopics
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCustomAsyncTaskGetV30FilteringDataTopicsFromValue(v string) (*ReportCustomAsyncTaskGetV30FilteringDataTopics, error) {
	ev := ReportCustomAsyncTaskGetV30FilteringDataTopics(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCustomAsyncTaskGetV30FilteringDataTopics: valid values are %v", v, AllowedReportCustomAsyncTaskGetV30FilteringDataTopicsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCustomAsyncTaskGetV30FilteringDataTopics) IsValid() bool {
	for _, existing := range AllowedReportCustomAsyncTaskGetV30FilteringDataTopicsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_custom_async_task_get_v3.0_filtering_data_topics value
func (v ReportCustomAsyncTaskGetV30FilteringDataTopics) Ptr() *ReportCustomAsyncTaskGetV30FilteringDataTopics {
	return &v
}
