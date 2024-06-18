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

// ReportAdGetV2TimeGranularity
type ReportAdGetV2TimeGranularity string

// List of report_ad_get_v2_time_granularity
const (
	STAT_TIME_GRANULARITY_MONTH_ReportAdGetV2TimeGranularity  ReportAdGetV2TimeGranularity = "STAT_TIME_GRANULARITY_MONTH"
	STAT_TIME_GRANULARITY_WEEK_ReportAdGetV2TimeGranularity   ReportAdGetV2TimeGranularity = "STAT_TIME_GRANULARITY_WEEK"
	STAT_TIME_GRANULARITY_DAILY_ReportAdGetV2TimeGranularity  ReportAdGetV2TimeGranularity = "STAT_TIME_GRANULARITY_DAILY"
	STAT_TIME_GRANULARITY_HOURLY_ReportAdGetV2TimeGranularity ReportAdGetV2TimeGranularity = "STAT_TIME_GRANULARITY_HOURLY"
)

// All allowed values of ReportAdGetV2TimeGranularity enum
var AllowedReportAdGetV2TimeGranularityEnumValues = []ReportAdGetV2TimeGranularity{
	"STAT_TIME_GRANULARITY_MONTH",
	"STAT_TIME_GRANULARITY_WEEK",
	"STAT_TIME_GRANULARITY_DAILY",
	"STAT_TIME_GRANULARITY_HOURLY",
}

// NewReportAdGetV2TimeGranularityFromValue returns a pointer to a valid ReportAdGetV2TimeGranularity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdGetV2TimeGranularityFromValue(v string) (*ReportAdGetV2TimeGranularity, error) {
	ev := ReportAdGetV2TimeGranularity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdGetV2TimeGranularity: valid values are %v", v, AllowedReportAdGetV2TimeGranularityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdGetV2TimeGranularity) IsValid() bool {
	for _, existing := range AllowedReportAdGetV2TimeGranularityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_ad_get_v2_time_granularity value
func (v ReportAdGetV2TimeGranularity) Ptr() *ReportAdGetV2TimeGranularity {
	return &v
}
