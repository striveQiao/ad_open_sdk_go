/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportCreativeGetV2DataListLandingType
type ReportCreativeGetV2DataListLandingType string

// List of report_creative_get_v2_data_list_landing_type
const (
	LIVE_ReportCreativeGetV2DataListLandingType      ReportCreativeGetV2DataListLandingType = "LIVE"
	ARTICLE_ReportCreativeGetV2DataListLandingType   ReportCreativeGetV2DataListLandingType = "ARTICLE"
	LINK_ReportCreativeGetV2DataListLandingType      ReportCreativeGetV2DataListLandingType = "LINK"
	SHOP_ReportCreativeGetV2DataListLandingType      ReportCreativeGetV2DataListLandingType = "SHOP"
	GOODS_ReportCreativeGetV2DataListLandingType     ReportCreativeGetV2DataListLandingType = "GOODS"
	AWEME_ReportCreativeGetV2DataListLandingType     ReportCreativeGetV2DataListLandingType = "AWEME"
	QUICK_APP_ReportCreativeGetV2DataListLandingType ReportCreativeGetV2DataListLandingType = "QUICK_APP"
	DPA_ReportCreativeGetV2DataListLandingType       ReportCreativeGetV2DataListLandingType = "DPA"
	STORE_ReportCreativeGetV2DataListLandingType     ReportCreativeGetV2DataListLandingType = "STORE"
	APP_ReportCreativeGetV2DataListLandingType       ReportCreativeGetV2DataListLandingType = "APP"
)

// All allowed values of ReportCreativeGetV2DataListLandingType enum
var AllowedReportCreativeGetV2DataListLandingTypeEnumValues = []ReportCreativeGetV2DataListLandingType{
	"LIVE",
	"ARTICLE",
	"LINK",
	"SHOP",
	"GOODS",
	"AWEME",
	"QUICK_APP",
	"DPA",
	"STORE",
	"APP",
}

// NewReportCreativeGetV2DataListLandingTypeFromValue returns a pointer to a valid ReportCreativeGetV2DataListLandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCreativeGetV2DataListLandingTypeFromValue(v string) (*ReportCreativeGetV2DataListLandingType, error) {
	ev := ReportCreativeGetV2DataListLandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCreativeGetV2DataListLandingType: valid values are %v", v, AllowedReportCreativeGetV2DataListLandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCreativeGetV2DataListLandingType) IsValid() bool {
	for _, existing := range AllowedReportCreativeGetV2DataListLandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_creative_get_v2_data_list_landing_type value
func (v ReportCreativeGetV2DataListLandingType) Ptr() *ReportCreativeGetV2DataListLandingType {
	return &v
}
