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

// ReportBrandAdGetV30LandingType
type ReportBrandAdGetV30LandingType int64

// List of report_brand_ad_get_v3.0_landing_type
const (
	Enum_1_ReportBrandAdGetV30LandingType  ReportBrandAdGetV30LandingType = 1
	Enum_10_ReportBrandAdGetV30LandingType ReportBrandAdGetV30LandingType = 10
	Enum_11_ReportBrandAdGetV30LandingType ReportBrandAdGetV30LandingType = 11
	Enum_12_ReportBrandAdGetV30LandingType ReportBrandAdGetV30LandingType = 12
	Enum_13_ReportBrandAdGetV30LandingType ReportBrandAdGetV30LandingType = 13
	Enum_2_ReportBrandAdGetV30LandingType  ReportBrandAdGetV30LandingType = 2
	Enum_3_ReportBrandAdGetV30LandingType  ReportBrandAdGetV30LandingType = 3
	Enum_4_ReportBrandAdGetV30LandingType  ReportBrandAdGetV30LandingType = 4
	Enum_5_ReportBrandAdGetV30LandingType  ReportBrandAdGetV30LandingType = 5
	Enum_6_ReportBrandAdGetV30LandingType  ReportBrandAdGetV30LandingType = 6
	Enum_7_ReportBrandAdGetV30LandingType  ReportBrandAdGetV30LandingType = 7
	Enum_8_ReportBrandAdGetV30LandingType  ReportBrandAdGetV30LandingType = 8
	Enum_9_ReportBrandAdGetV30LandingType  ReportBrandAdGetV30LandingType = 9
)

// All allowed values of ReportBrandAdGetV30LandingType enum
var AllowedReportBrandAdGetV30LandingTypeEnumValues = []ReportBrandAdGetV30LandingType{
	1,
	10,
	11,
	12,
	13,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
}

// NewReportBrandAdGetV30LandingTypeFromValue returns a pointer to a valid ReportBrandAdGetV30LandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportBrandAdGetV30LandingTypeFromValue(v int64) (*ReportBrandAdGetV30LandingType, error) {
	ev := ReportBrandAdGetV30LandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportBrandAdGetV30LandingType: valid values are %v", v, AllowedReportBrandAdGetV30LandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportBrandAdGetV30LandingType) IsValid() bool {
	for _, existing := range AllowedReportBrandAdGetV30LandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_brand_ad_get_v3.0_landing_type value
func (v ReportBrandAdGetV30LandingType) Ptr() *ReportBrandAdGetV30LandingType {
	return &v
}
