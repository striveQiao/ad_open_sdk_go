/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportAdvertiserGetV2DataListAc
type ReportAdvertiserGetV2DataListAc string

// List of report_advertiser_get_v2_data_list_ac
const (
	Enum_5_G_ReportAdvertiserGetV2DataListAc ReportAdvertiserGetV2DataListAc = "5G"
	UNKNOWN_ReportAdvertiserGetV2DataListAc  ReportAdvertiserGetV2DataListAc = "unknown"
	Enum_2_G_ReportAdvertiserGetV2DataListAc ReportAdvertiserGetV2DataListAc = "2G"
	WIFI_ReportAdvertiserGetV2DataListAc     ReportAdvertiserGetV2DataListAc = "WIFI"
	Enum_3_G_ReportAdvertiserGetV2DataListAc ReportAdvertiserGetV2DataListAc = "3G"
	Enum_4_G_ReportAdvertiserGetV2DataListAc ReportAdvertiserGetV2DataListAc = "4G"
)

// All allowed values of ReportAdvertiserGetV2DataListAc enum
var AllowedReportAdvertiserGetV2DataListAcEnumValues = []ReportAdvertiserGetV2DataListAc{
	"5G",
	"unknown",
	"2G",
	"WIFI",
	"3G",
	"4G",
}

// NewReportAdvertiserGetV2DataListAcFromValue returns a pointer to a valid ReportAdvertiserGetV2DataListAc
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdvertiserGetV2DataListAcFromValue(v string) (*ReportAdvertiserGetV2DataListAc, error) {
	ev := ReportAdvertiserGetV2DataListAc(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdvertiserGetV2DataListAc: valid values are %v", v, AllowedReportAdvertiserGetV2DataListAcEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdvertiserGetV2DataListAc) IsValid() bool {
	for _, existing := range AllowedReportAdvertiserGetV2DataListAcEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_advertiser_get_v2_data_list_ac value
func (v ReportAdvertiserGetV2DataListAc) Ptr() *ReportAdvertiserGetV2DataListAc {
	return &v
}
