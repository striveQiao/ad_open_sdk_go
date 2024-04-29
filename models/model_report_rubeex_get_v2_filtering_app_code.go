/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportRubeexGetV2FilteringAppCode
type ReportRubeexGetV2FilteringAppCode int64

// List of report_rubeex_get_v2_filtering_app_code
const (
	Enum_2_ReportRubeexGetV2FilteringAppCode  ReportRubeexGetV2FilteringAppCode = 2
	Enum_3_ReportRubeexGetV2FilteringAppCode  ReportRubeexGetV2FilteringAppCode = 3
	Enum_4_ReportRubeexGetV2FilteringAppCode  ReportRubeexGetV2FilteringAppCode = 4
	Enum_8_ReportRubeexGetV2FilteringAppCode  ReportRubeexGetV2FilteringAppCode = 8
	Enum_9_ReportRubeexGetV2FilteringAppCode  ReportRubeexGetV2FilteringAppCode = 9
	Enum_14_ReportRubeexGetV2FilteringAppCode ReportRubeexGetV2FilteringAppCode = 14
	Enum_28_ReportRubeexGetV2FilteringAppCode ReportRubeexGetV2FilteringAppCode = 28
)

// All allowed values of ReportRubeexGetV2FilteringAppCode enum
var AllowedReportRubeexGetV2FilteringAppCodeEnumValues = []ReportRubeexGetV2FilteringAppCode{
	2,
	3,
	4,
	8,
	9,
	14,
	28,
}

// NewReportRubeexGetV2FilteringAppCodeFromValue returns a pointer to a valid ReportRubeexGetV2FilteringAppCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportRubeexGetV2FilteringAppCodeFromValue(v int64) (*ReportRubeexGetV2FilteringAppCode, error) {
	ev := ReportRubeexGetV2FilteringAppCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportRubeexGetV2FilteringAppCode: valid values are %v", v, AllowedReportRubeexGetV2FilteringAppCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportRubeexGetV2FilteringAppCode) IsValid() bool {
	for _, existing := range AllowedReportRubeexGetV2FilteringAppCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_rubeex_get_v2_filtering_app_code value
func (v ReportRubeexGetV2FilteringAppCode) Ptr() *ReportRubeexGetV2FilteringAppCode {
	return &v
}
