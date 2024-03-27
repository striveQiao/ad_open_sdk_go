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

// ReportCustomCreativeGetV30OrderByType
type ReportCustomCreativeGetV30OrderByType string

// List of report_custom_creative_get_v3.0_order_by_type
const (
	ASC_ReportCustomCreativeGetV30OrderByType  ReportCustomCreativeGetV30OrderByType = "ASC"
	DESC_ReportCustomCreativeGetV30OrderByType ReportCustomCreativeGetV30OrderByType = "DESC"
)

// All allowed values of ReportCustomCreativeGetV30OrderByType enum
var AllowedReportCustomCreativeGetV30OrderByTypeEnumValues = []ReportCustomCreativeGetV30OrderByType{
	"ASC",
	"DESC",
}

// NewReportCustomCreativeGetV30OrderByTypeFromValue returns a pointer to a valid ReportCustomCreativeGetV30OrderByType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCustomCreativeGetV30OrderByTypeFromValue(v string) (*ReportCustomCreativeGetV30OrderByType, error) {
	ev := ReportCustomCreativeGetV30OrderByType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCustomCreativeGetV30OrderByType: valid values are %v", v, AllowedReportCustomCreativeGetV30OrderByTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCustomCreativeGetV30OrderByType) IsValid() bool {
	for _, existing := range AllowedReportCustomCreativeGetV30OrderByTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_custom_creative_get_v3.0_order_by_type value
func (v ReportCustomCreativeGetV30OrderByType) Ptr() *ReportCustomCreativeGetV30OrderByType {
	return &v
}
