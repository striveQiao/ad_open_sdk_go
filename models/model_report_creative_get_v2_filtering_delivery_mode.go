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

// ReportCreativeGetV2FilteringDeliveryMode
type ReportCreativeGetV2FilteringDeliveryMode string

// List of report_creative_get_v2_filtering_delivery_mode
const (
	ADLAB_FREE_ReportCreativeGetV2FilteringDeliveryMode ReportCreativeGetV2FilteringDeliveryMode = "ADLAB_FREE"
	STANDARD_ReportCreativeGetV2FilteringDeliveryMode   ReportCreativeGetV2FilteringDeliveryMode = "STANDARD"
)

// All allowed values of ReportCreativeGetV2FilteringDeliveryMode enum
var AllowedReportCreativeGetV2FilteringDeliveryModeEnumValues = []ReportCreativeGetV2FilteringDeliveryMode{
	"ADLAB_FREE",
	"STANDARD",
}

// NewReportCreativeGetV2FilteringDeliveryModeFromValue returns a pointer to a valid ReportCreativeGetV2FilteringDeliveryMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCreativeGetV2FilteringDeliveryModeFromValue(v string) (*ReportCreativeGetV2FilteringDeliveryMode, error) {
	ev := ReportCreativeGetV2FilteringDeliveryMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCreativeGetV2FilteringDeliveryMode: valid values are %v", v, AllowedReportCreativeGetV2FilteringDeliveryModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCreativeGetV2FilteringDeliveryMode) IsValid() bool {
	for _, existing := range AllowedReportCreativeGetV2FilteringDeliveryModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_creative_get_v2_filtering_delivery_mode value
func (v ReportCreativeGetV2FilteringDeliveryMode) Ptr() *ReportCreativeGetV2FilteringDeliveryMode {
	return &v
}
