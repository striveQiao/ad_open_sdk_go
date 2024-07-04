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

// ReportBrandAdvertiserGetV30LandingType
type ReportBrandAdvertiserGetV30LandingType int64

// List of report_brand_advertiser_get_v3.0_landing_type
const (
	Enum_1_ReportBrandAdvertiserGetV30LandingType  ReportBrandAdvertiserGetV30LandingType = 1
	Enum_10_ReportBrandAdvertiserGetV30LandingType ReportBrandAdvertiserGetV30LandingType = 10
	Enum_11_ReportBrandAdvertiserGetV30LandingType ReportBrandAdvertiserGetV30LandingType = 11
	Enum_12_ReportBrandAdvertiserGetV30LandingType ReportBrandAdvertiserGetV30LandingType = 12
	Enum_13_ReportBrandAdvertiserGetV30LandingType ReportBrandAdvertiserGetV30LandingType = 13
	Enum_2_ReportBrandAdvertiserGetV30LandingType  ReportBrandAdvertiserGetV30LandingType = 2
	Enum_3_ReportBrandAdvertiserGetV30LandingType  ReportBrandAdvertiserGetV30LandingType = 3
	Enum_4_ReportBrandAdvertiserGetV30LandingType  ReportBrandAdvertiserGetV30LandingType = 4
	Enum_5_ReportBrandAdvertiserGetV30LandingType  ReportBrandAdvertiserGetV30LandingType = 5
	Enum_6_ReportBrandAdvertiserGetV30LandingType  ReportBrandAdvertiserGetV30LandingType = 6
	Enum_7_ReportBrandAdvertiserGetV30LandingType  ReportBrandAdvertiserGetV30LandingType = 7
	Enum_8_ReportBrandAdvertiserGetV30LandingType  ReportBrandAdvertiserGetV30LandingType = 8
	Enum_9_ReportBrandAdvertiserGetV30LandingType  ReportBrandAdvertiserGetV30LandingType = 9
)

// All allowed values of ReportBrandAdvertiserGetV30LandingType enum
var AllowedReportBrandAdvertiserGetV30LandingTypeEnumValues = []ReportBrandAdvertiserGetV30LandingType{
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

// NewReportBrandAdvertiserGetV30LandingTypeFromValue returns a pointer to a valid ReportBrandAdvertiserGetV30LandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportBrandAdvertiserGetV30LandingTypeFromValue(v int64) (*ReportBrandAdvertiserGetV30LandingType, error) {
	ev := ReportBrandAdvertiserGetV30LandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportBrandAdvertiserGetV30LandingType: valid values are %v", v, AllowedReportBrandAdvertiserGetV30LandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportBrandAdvertiserGetV30LandingType) IsValid() bool {
	for _, existing := range AllowedReportBrandAdvertiserGetV30LandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_brand_advertiser_get_v3.0_landing_type value
func (v ReportBrandAdvertiserGetV30LandingType) Ptr() *ReportBrandAdvertiserGetV30LandingType {
	return &v
}
