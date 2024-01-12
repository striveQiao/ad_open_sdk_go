/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportCreativeGetV2TimeGranularity
type ReportCreativeGetV2TimeGranularity string

// List of report_creative_get_v2_time_granularity
const (
	STAT_TIME_GRANULARITY_MONTH_ReportCreativeGetV2TimeGranularity  ReportCreativeGetV2TimeGranularity = "STAT_TIME_GRANULARITY_MONTH"
	STAT_TIME_GRANULARITY_WEEK_ReportCreativeGetV2TimeGranularity   ReportCreativeGetV2TimeGranularity = "STAT_TIME_GRANULARITY_WEEK"
	STAT_TIME_GRANULARITY_DAILY_ReportCreativeGetV2TimeGranularity  ReportCreativeGetV2TimeGranularity = "STAT_TIME_GRANULARITY_DAILY"
	STAT_TIME_GRANULARITY_HOURLY_ReportCreativeGetV2TimeGranularity ReportCreativeGetV2TimeGranularity = "STAT_TIME_GRANULARITY_HOURLY"
)

// All allowed values of ReportCreativeGetV2TimeGranularity enum
var AllowedReportCreativeGetV2TimeGranularityEnumValues = []ReportCreativeGetV2TimeGranularity{
	"STAT_TIME_GRANULARITY_MONTH",
	"STAT_TIME_GRANULARITY_WEEK",
	"STAT_TIME_GRANULARITY_DAILY",
	"STAT_TIME_GRANULARITY_HOURLY",
}

// NewReportCreativeGetV2TimeGranularityFromValue returns a pointer to a valid ReportCreativeGetV2TimeGranularity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCreativeGetV2TimeGranularityFromValue(v string) (*ReportCreativeGetV2TimeGranularity, error) {
	ev := ReportCreativeGetV2TimeGranularity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCreativeGetV2TimeGranularity: valid values are %v", v, AllowedReportCreativeGetV2TimeGranularityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCreativeGetV2TimeGranularity) IsValid() bool {
	for _, existing := range AllowedReportCreativeGetV2TimeGranularityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_creative_get_v2_time_granularity value
func (v ReportCreativeGetV2TimeGranularity) Ptr() *ReportCreativeGetV2TimeGranularity {
	return &v
}
