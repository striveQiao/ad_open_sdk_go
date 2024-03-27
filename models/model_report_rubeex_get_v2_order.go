/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportRubeexGetV2Order
type ReportRubeexGetV2Order string

// List of report_rubeex_get_v2_order
const (
	DESC_ReportRubeexGetV2Order ReportRubeexGetV2Order = "DESC"
	ASC_ReportRubeexGetV2Order  ReportRubeexGetV2Order = "ASC"
)

// All allowed values of ReportRubeexGetV2Order enum
var AllowedReportRubeexGetV2OrderEnumValues = []ReportRubeexGetV2Order{
	"DESC",
	"ASC",
}

// NewReportRubeexGetV2OrderFromValue returns a pointer to a valid ReportRubeexGetV2Order
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportRubeexGetV2OrderFromValue(v string) (*ReportRubeexGetV2Order, error) {
	ev := ReportRubeexGetV2Order(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportRubeexGetV2Order: valid values are %v", v, AllowedReportRubeexGetV2OrderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportRubeexGetV2Order) IsValid() bool {
	for _, existing := range AllowedReportRubeexGetV2OrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_rubeex_get_v2_order value
func (v ReportRubeexGetV2Order) Ptr() *ReportRubeexGetV2Order {
	return &v
}
