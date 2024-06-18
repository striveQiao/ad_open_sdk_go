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

// QianchuanReportLiveGetV10StatsAuthority
type QianchuanReportLiveGetV10StatsAuthority string

// List of qianchuan_report_live_get_v1.0_stats_authority
const (
	CURRENT_QianchuanReportLiveGetV10StatsAuthority       QianchuanReportLiveGetV10StatsAuthority = "CURRENT"
	QUALIFICATION_QianchuanReportLiveGetV10StatsAuthority QianchuanReportLiveGetV10StatsAuthority = "QUALIFICATION"
)

// All allowed values of QianchuanReportLiveGetV10StatsAuthority enum
var AllowedQianchuanReportLiveGetV10StatsAuthorityEnumValues = []QianchuanReportLiveGetV10StatsAuthority{
	"CURRENT",
	"QUALIFICATION",
}

// NewQianchuanReportLiveGetV10StatsAuthorityFromValue returns a pointer to a valid QianchuanReportLiveGetV10StatsAuthority
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportLiveGetV10StatsAuthorityFromValue(v string) (*QianchuanReportLiveGetV10StatsAuthority, error) {
	ev := QianchuanReportLiveGetV10StatsAuthority(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportLiveGetV10StatsAuthority: valid values are %v", v, AllowedQianchuanReportLiveGetV10StatsAuthorityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportLiveGetV10StatsAuthority) IsValid() bool {
	for _, existing := range AllowedQianchuanReportLiveGetV10StatsAuthorityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_live_get_v1.0_stats_authority value
func (v QianchuanReportLiveGetV10StatsAuthority) Ptr() *QianchuanReportLiveGetV10StatsAuthority {
	return &v
}
