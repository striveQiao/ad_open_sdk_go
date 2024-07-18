/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportSitePageV2TimeDuration
type ReportSitePageV2TimeDuration string

// List of report_site_page_v2_time_duration
const (
	DAY_ReportSitePageV2TimeDuration        ReportSitePageV2TimeDuration = "DAY"
	THREE_DAYS_ReportSitePageV2TimeDuration ReportSitePageV2TimeDuration = "THREE_DAYS"
	WEEK_ReportSitePageV2TimeDuration       ReportSitePageV2TimeDuration = "WEEK"
	MONTH_ReportSitePageV2TimeDuration      ReportSitePageV2TimeDuration = "MONTH"
)

// All allowed values of ReportSitePageV2TimeDuration enum
var AllowedReportSitePageV2TimeDurationEnumValues = []ReportSitePageV2TimeDuration{
	"DAY",
	"THREE_DAYS",
	"WEEK",
	"MONTH",
}

// NewReportSitePageV2TimeDurationFromValue returns a pointer to a valid ReportSitePageV2TimeDuration
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportSitePageV2TimeDurationFromValue(v string) (*ReportSitePageV2TimeDuration, error) {
	ev := ReportSitePageV2TimeDuration(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportSitePageV2TimeDuration: valid values are %v", v, AllowedReportSitePageV2TimeDurationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportSitePageV2TimeDuration) IsValid() bool {
	for _, existing := range AllowedReportSitePageV2TimeDurationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_site_page_v2_time_duration value
func (v ReportSitePageV2TimeDuration) Ptr() *ReportSitePageV2TimeDuration {
	return &v
}
