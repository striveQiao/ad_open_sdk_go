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

// ReportCreativeGetV2FilteringLandingTypes
type ReportCreativeGetV2FilteringLandingTypes string

// List of report_creative_get_v2_filtering_landing_types
const (
	LIVE_ReportCreativeGetV2FilteringLandingTypes      ReportCreativeGetV2FilteringLandingTypes = "LIVE"
	ARTICLE_ReportCreativeGetV2FilteringLandingTypes   ReportCreativeGetV2FilteringLandingTypes = "ARTICLE"
	LINK_ReportCreativeGetV2FilteringLandingTypes      ReportCreativeGetV2FilteringLandingTypes = "LINK"
	SHOP_ReportCreativeGetV2FilteringLandingTypes      ReportCreativeGetV2FilteringLandingTypes = "SHOP"
	GOODS_ReportCreativeGetV2FilteringLandingTypes     ReportCreativeGetV2FilteringLandingTypes = "GOODS"
	AWEME_ReportCreativeGetV2FilteringLandingTypes     ReportCreativeGetV2FilteringLandingTypes = "AWEME"
	QUICK_APP_ReportCreativeGetV2FilteringLandingTypes ReportCreativeGetV2FilteringLandingTypes = "QUICK_APP"
	DPA_ReportCreativeGetV2FilteringLandingTypes       ReportCreativeGetV2FilteringLandingTypes = "DPA"
	STORE_ReportCreativeGetV2FilteringLandingTypes     ReportCreativeGetV2FilteringLandingTypes = "STORE"
	APP_ReportCreativeGetV2FilteringLandingTypes       ReportCreativeGetV2FilteringLandingTypes = "APP"
)

// All allowed values of ReportCreativeGetV2FilteringLandingTypes enum
var AllowedReportCreativeGetV2FilteringLandingTypesEnumValues = []ReportCreativeGetV2FilteringLandingTypes{
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

// NewReportCreativeGetV2FilteringLandingTypesFromValue returns a pointer to a valid ReportCreativeGetV2FilteringLandingTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCreativeGetV2FilteringLandingTypesFromValue(v string) (*ReportCreativeGetV2FilteringLandingTypes, error) {
	ev := ReportCreativeGetV2FilteringLandingTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCreativeGetV2FilteringLandingTypes: valid values are %v", v, AllowedReportCreativeGetV2FilteringLandingTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCreativeGetV2FilteringLandingTypes) IsValid() bool {
	for _, existing := range AllowedReportCreativeGetV2FilteringLandingTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_creative_get_v2_filtering_landing_types value
func (v ReportCreativeGetV2FilteringLandingTypes) Ptr() *ReportCreativeGetV2FilteringLandingTypes {
	return &v
}
