/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportAdGetV10TimeGranularity
type QianchuanReportAdGetV10TimeGranularity string

// List of qianchuan_report_ad_get_v1.0_time_granularity
const (
	TIME_GRANULARITY_DAILY_QianchuanReportAdGetV10TimeGranularity  QianchuanReportAdGetV10TimeGranularity = "TIME_GRANULARITY_DAILY"
	TIME_GRANULARITY_HOURLY_QianchuanReportAdGetV10TimeGranularity QianchuanReportAdGetV10TimeGranularity = "TIME_GRANULARITY_HOURLY"
)

// All allowed values of QianchuanReportAdGetV10TimeGranularity enum
var AllowedQianchuanReportAdGetV10TimeGranularityEnumValues = []QianchuanReportAdGetV10TimeGranularity{
	"TIME_GRANULARITY_DAILY",
	"TIME_GRANULARITY_HOURLY",
}

// NewQianchuanReportAdGetV10TimeGranularityFromValue returns a pointer to a valid QianchuanReportAdGetV10TimeGranularity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportAdGetV10TimeGranularityFromValue(v string) (*QianchuanReportAdGetV10TimeGranularity, error) {
	ev := QianchuanReportAdGetV10TimeGranularity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportAdGetV10TimeGranularity: valid values are %v", v, AllowedQianchuanReportAdGetV10TimeGranularityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportAdGetV10TimeGranularity) IsValid() bool {
	for _, existing := range AllowedQianchuanReportAdGetV10TimeGranularityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_ad_get_v1.0_time_granularity value
func (v QianchuanReportAdGetV10TimeGranularity) Ptr() *QianchuanReportAdGetV10TimeGranularity {
	return &v
}
