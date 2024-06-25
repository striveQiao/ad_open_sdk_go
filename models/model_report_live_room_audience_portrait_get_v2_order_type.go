/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportLiveRoomAudiencePortraitGetV2OrderType
type ReportLiveRoomAudiencePortraitGetV2OrderType string

// List of report_live_room_audience_portrait_get_v2_order_type
const (
	DESC_ReportLiveRoomAudiencePortraitGetV2OrderType ReportLiveRoomAudiencePortraitGetV2OrderType = "DESC"
	ASC_ReportLiveRoomAudiencePortraitGetV2OrderType  ReportLiveRoomAudiencePortraitGetV2OrderType = "ASC"
)

// All allowed values of ReportLiveRoomAudiencePortraitGetV2OrderType enum
var AllowedReportLiveRoomAudiencePortraitGetV2OrderTypeEnumValues = []ReportLiveRoomAudiencePortraitGetV2OrderType{
	"DESC",
	"ASC",
}

// NewReportLiveRoomAudiencePortraitGetV2OrderTypeFromValue returns a pointer to a valid ReportLiveRoomAudiencePortraitGetV2OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportLiveRoomAudiencePortraitGetV2OrderTypeFromValue(v string) (*ReportLiveRoomAudiencePortraitGetV2OrderType, error) {
	ev := ReportLiveRoomAudiencePortraitGetV2OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportLiveRoomAudiencePortraitGetV2OrderType: valid values are %v", v, AllowedReportLiveRoomAudiencePortraitGetV2OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportLiveRoomAudiencePortraitGetV2OrderType) IsValid() bool {
	for _, existing := range AllowedReportLiveRoomAudiencePortraitGetV2OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_live_room_audience_portrait_get_v2_order_type value
func (v ReportLiveRoomAudiencePortraitGetV2OrderType) Ptr() *ReportLiveRoomAudiencePortraitGetV2OrderType {
	return &v
}
