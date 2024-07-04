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

// ReportLiveRoomAnalysisGetV2OrderType
type ReportLiveRoomAnalysisGetV2OrderType string

// List of report_live_room_analysis_get_v2_order_type
const (
	DESC_ReportLiveRoomAnalysisGetV2OrderType ReportLiveRoomAnalysisGetV2OrderType = "DESC"
	ASC_ReportLiveRoomAnalysisGetV2OrderType  ReportLiveRoomAnalysisGetV2OrderType = "ASC"
)

// All allowed values of ReportLiveRoomAnalysisGetV2OrderType enum
var AllowedReportLiveRoomAnalysisGetV2OrderTypeEnumValues = []ReportLiveRoomAnalysisGetV2OrderType{
	"DESC",
	"ASC",
}

// NewReportLiveRoomAnalysisGetV2OrderTypeFromValue returns a pointer to a valid ReportLiveRoomAnalysisGetV2OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportLiveRoomAnalysisGetV2OrderTypeFromValue(v string) (*ReportLiveRoomAnalysisGetV2OrderType, error) {
	ev := ReportLiveRoomAnalysisGetV2OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportLiveRoomAnalysisGetV2OrderType: valid values are %v", v, AllowedReportLiveRoomAnalysisGetV2OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportLiveRoomAnalysisGetV2OrderType) IsValid() bool {
	for _, existing := range AllowedReportLiveRoomAnalysisGetV2OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_live_room_analysis_get_v2_order_type value
func (v ReportLiveRoomAnalysisGetV2OrderType) Ptr() *ReportLiveRoomAnalysisGetV2OrderType {
	return &v
}
