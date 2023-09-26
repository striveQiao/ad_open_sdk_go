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

// ReportAdvertiserGetV2DataListLandingType
type ReportAdvertiserGetV2DataListLandingType string

// List of report_advertiser_get_v2_data_list_landing_type
const (
	LIVE_ReportAdvertiserGetV2DataListLandingType      ReportAdvertiserGetV2DataListLandingType = "LIVE"
	ARTICLE_ReportAdvertiserGetV2DataListLandingType   ReportAdvertiserGetV2DataListLandingType = "ARTICLE"
	LINK_ReportAdvertiserGetV2DataListLandingType      ReportAdvertiserGetV2DataListLandingType = "LINK"
	SHOP_ReportAdvertiserGetV2DataListLandingType      ReportAdvertiserGetV2DataListLandingType = "SHOP"
	GOODS_ReportAdvertiserGetV2DataListLandingType     ReportAdvertiserGetV2DataListLandingType = "GOODS"
	AWEME_ReportAdvertiserGetV2DataListLandingType     ReportAdvertiserGetV2DataListLandingType = "AWEME"
	QUICK_APP_ReportAdvertiserGetV2DataListLandingType ReportAdvertiserGetV2DataListLandingType = "QUICK_APP"
	DPA_ReportAdvertiserGetV2DataListLandingType       ReportAdvertiserGetV2DataListLandingType = "DPA"
	STORE_ReportAdvertiserGetV2DataListLandingType     ReportAdvertiserGetV2DataListLandingType = "STORE"
	APP_ReportAdvertiserGetV2DataListLandingType       ReportAdvertiserGetV2DataListLandingType = "APP"
)

// All allowed values of ReportAdvertiserGetV2DataListLandingType enum
var AllowedReportAdvertiserGetV2DataListLandingTypeEnumValues = []ReportAdvertiserGetV2DataListLandingType{
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

// NewReportAdvertiserGetV2DataListLandingTypeFromValue returns a pointer to a valid ReportAdvertiserGetV2DataListLandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdvertiserGetV2DataListLandingTypeFromValue(v string) (*ReportAdvertiserGetV2DataListLandingType, error) {
	ev := ReportAdvertiserGetV2DataListLandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdvertiserGetV2DataListLandingType: valid values are %v", v, AllowedReportAdvertiserGetV2DataListLandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdvertiserGetV2DataListLandingType) IsValid() bool {
	for _, existing := range AllowedReportAdvertiserGetV2DataListLandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_advertiser_get_v2_data_list_landing_type value
func (v ReportAdvertiserGetV2DataListLandingType) Ptr() *ReportAdvertiserGetV2DataListLandingType {
	return &v
}
