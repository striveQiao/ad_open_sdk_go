/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportCampaignGetV2TimeGranularity
type ReportCampaignGetV2TimeGranularity string

// List of report_campaign_get_v2_time_granularity
const (
	STAT_TIME_GRANULARITY_MONTH_ReportCampaignGetV2TimeGranularity  ReportCampaignGetV2TimeGranularity = "STAT_TIME_GRANULARITY_MONTH"
	STAT_TIME_GRANULARITY_WEEK_ReportCampaignGetV2TimeGranularity   ReportCampaignGetV2TimeGranularity = "STAT_TIME_GRANULARITY_WEEK"
	STAT_TIME_GRANULARITY_DAILY_ReportCampaignGetV2TimeGranularity  ReportCampaignGetV2TimeGranularity = "STAT_TIME_GRANULARITY_DAILY"
	STAT_TIME_GRANULARITY_HOURLY_ReportCampaignGetV2TimeGranularity ReportCampaignGetV2TimeGranularity = "STAT_TIME_GRANULARITY_HOURLY"
)

// All allowed values of ReportCampaignGetV2TimeGranularity enum
var AllowedReportCampaignGetV2TimeGranularityEnumValues = []ReportCampaignGetV2TimeGranularity{
	"STAT_TIME_GRANULARITY_MONTH",
	"STAT_TIME_GRANULARITY_WEEK",
	"STAT_TIME_GRANULARITY_DAILY",
	"STAT_TIME_GRANULARITY_HOURLY",
}

// NewReportCampaignGetV2TimeGranularityFromValue returns a pointer to a valid ReportCampaignGetV2TimeGranularity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCampaignGetV2TimeGranularityFromValue(v string) (*ReportCampaignGetV2TimeGranularity, error) {
	ev := ReportCampaignGetV2TimeGranularity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCampaignGetV2TimeGranularity: valid values are %v", v, AllowedReportCampaignGetV2TimeGranularityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCampaignGetV2TimeGranularity) IsValid() bool {
	for _, existing := range AllowedReportCampaignGetV2TimeGranularityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_campaign_get_v2_time_granularity value
func (v ReportCampaignGetV2TimeGranularity) Ptr() *ReportCampaignGetV2TimeGranularity {
	return &v
}
