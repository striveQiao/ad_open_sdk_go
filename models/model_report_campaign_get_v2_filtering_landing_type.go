/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportCampaignGetV2FilteringLandingType
type ReportCampaignGetV2FilteringLandingType string

// List of report_campaign_get_v2_filtering_landing_type
const (
	APP_ReportCampaignGetV2FilteringLandingType       ReportCampaignGetV2FilteringLandingType = "APP"
	DPA_ReportCampaignGetV2FilteringLandingType       ReportCampaignGetV2FilteringLandingType = "DPA"
	SHOP_ReportCampaignGetV2FilteringLandingType      ReportCampaignGetV2FilteringLandingType = "SHOP"
	AWEME_ReportCampaignGetV2FilteringLandingType     ReportCampaignGetV2FilteringLandingType = "AWEME"
	LINK_ReportCampaignGetV2FilteringLandingType      ReportCampaignGetV2FilteringLandingType = "LINK"
	QUICK_APP_ReportCampaignGetV2FilteringLandingType ReportCampaignGetV2FilteringLandingType = "QUICK_APP"
	LIVE_ReportCampaignGetV2FilteringLandingType      ReportCampaignGetV2FilteringLandingType = "LIVE"
	STORE_ReportCampaignGetV2FilteringLandingType     ReportCampaignGetV2FilteringLandingType = "STORE"
	GOODS_ReportCampaignGetV2FilteringLandingType     ReportCampaignGetV2FilteringLandingType = "GOODS"
	ARTICLE_ReportCampaignGetV2FilteringLandingType   ReportCampaignGetV2FilteringLandingType = "ARTICLE"
)

// All allowed values of ReportCampaignGetV2FilteringLandingType enum
var AllowedReportCampaignGetV2FilteringLandingTypeEnumValues = []ReportCampaignGetV2FilteringLandingType{
	"APP",
	"DPA",
	"SHOP",
	"AWEME",
	"LINK",
	"QUICK_APP",
	"LIVE",
	"STORE",
	"GOODS",
	"ARTICLE",
}

// NewReportCampaignGetV2FilteringLandingTypeFromValue returns a pointer to a valid ReportCampaignGetV2FilteringLandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCampaignGetV2FilteringLandingTypeFromValue(v string) (*ReportCampaignGetV2FilteringLandingType, error) {
	ev := ReportCampaignGetV2FilteringLandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCampaignGetV2FilteringLandingType: valid values are %v", v, AllowedReportCampaignGetV2FilteringLandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCampaignGetV2FilteringLandingType) IsValid() bool {
	for _, existing := range AllowedReportCampaignGetV2FilteringLandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_campaign_get_v2_filtering_landing_type value
func (v ReportCampaignGetV2FilteringLandingType) Ptr() *ReportCampaignGetV2FilteringLandingType {
	return &v
}
