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

// ReportCampaignGetV2GroupBy
type ReportCampaignGetV2GroupBy string

// List of report_campaign_get_v2_group_by
const (
	STAT_GROUP_BY_CREATIVE_COMPONENT_ID_ReportCampaignGetV2GroupBy  ReportCampaignGetV2GroupBy = "STAT_GROUP_BY_CREATIVE_COMPONENT_ID"
	STAT_GROUP_BY_AC_ReportCampaignGetV2GroupBy                     ReportCampaignGetV2GroupBy = "STAT_GROUP_BY_AC"
	STAT_GROUP_BY_CITY_NAME_ReportCampaignGetV2GroupBy              ReportCampaignGetV2GroupBy = "STAT_GROUP_BY_CITY_NAME"
	STAT_GROUP_BY_MATERIAL_ID_ReportCampaignGetV2GroupBy            ReportCampaignGetV2GroupBy = "STAT_GROUP_BY_MATERIAL_ID"
	STAT_GROUP_BY_EXTERNAL_ACTION_ReportCampaignGetV2GroupBy        ReportCampaignGetV2GroupBy = "STAT_GROUP_BY_EXTERNAL_ACTION"
	STAT_GROUP_BY_FIELD_ID_ReportCampaignGetV2GroupBy               ReportCampaignGetV2GroupBy = "STAT_GROUP_BY_FIELD_ID"
	STAT_GROUP_BY_CAMPAIGN_TYPE_ReportCampaignGetV2GroupBy          ReportCampaignGetV2GroupBy = "STAT_GROUP_BY_CAMPAIGN_TYPE"
	STAT_GROUP_BY_PROVINCE_NAME_ReportCampaignGetV2GroupBy          ReportCampaignGetV2GroupBy = "STAT_GROUP_BY_PROVINCE_NAME"
	STAT_GROUP_BY_PLATFORM_ReportCampaignGetV2GroupBy               ReportCampaignGetV2GroupBy = "STAT_GROUP_BY_PLATFORM"
	STAT_GROUP_BY_PRICING_CATEGORY_ReportCampaignGetV2GroupBy       ReportCampaignGetV2GroupBy = "STAT_GROUP_BY_PRICING_CATEGORY"
	STAT_GROUP_BY_AGE_ReportCampaignGetV2GroupBy                    ReportCampaignGetV2GroupBy = "STAT_GROUP_BY_AGE"
	STAT_GROUP_BY_PRICING_ReportCampaignGetV2GroupBy                ReportCampaignGetV2GroupBy = "STAT_GROUP_BY_PRICING"
	STAT_GROUP_BY_FIELD_STAT_TIME_ReportCampaignGetV2GroupBy        ReportCampaignGetV2GroupBy = "STAT_GROUP_BY_FIELD_STAT_TIME"
	STAT_GROUP_BY_GENDER_ReportCampaignGetV2GroupBy                 ReportCampaignGetV2GroupBy = "STAT_GROUP_BY_GENDER"
	STAT_GROUP_BY_LANDING_TYPE_ReportCampaignGetV2GroupBy           ReportCampaignGetV2GroupBy = "STAT_GROUP_BY_LANDING_TYPE"
	STAT_GROUP_BY_CREATIVE_MATERIAL_MODE_ReportCampaignGetV2GroupBy ReportCampaignGetV2GroupBy = "STAT_GROUP_BY_CREATIVE_MATERIAL_MODE"
	STAT_GROUP_BY_INVENTORY_ReportCampaignGetV2GroupBy              ReportCampaignGetV2GroupBy = "STAT_GROUP_BY_INVENTORY"
	STAT_GROUP_BY_IMAGE_MODE_ReportCampaignGetV2GroupBy             ReportCampaignGetV2GroupBy = "STAT_GROUP_BY_IMAGE_MODE"
)

// All allowed values of ReportCampaignGetV2GroupBy enum
var AllowedReportCampaignGetV2GroupByEnumValues = []ReportCampaignGetV2GroupBy{
	"STAT_GROUP_BY_CREATIVE_COMPONENT_ID",
	"STAT_GROUP_BY_AC",
	"STAT_GROUP_BY_CITY_NAME",
	"STAT_GROUP_BY_MATERIAL_ID",
	"STAT_GROUP_BY_EXTERNAL_ACTION",
	"STAT_GROUP_BY_FIELD_ID",
	"STAT_GROUP_BY_CAMPAIGN_TYPE",
	"STAT_GROUP_BY_PROVINCE_NAME",
	"STAT_GROUP_BY_PLATFORM",
	"STAT_GROUP_BY_PRICING_CATEGORY",
	"STAT_GROUP_BY_AGE",
	"STAT_GROUP_BY_PRICING",
	"STAT_GROUP_BY_FIELD_STAT_TIME",
	"STAT_GROUP_BY_GENDER",
	"STAT_GROUP_BY_LANDING_TYPE",
	"STAT_GROUP_BY_CREATIVE_MATERIAL_MODE",
	"STAT_GROUP_BY_INVENTORY",
	"STAT_GROUP_BY_IMAGE_MODE",
}

// NewReportCampaignGetV2GroupByFromValue returns a pointer to a valid ReportCampaignGetV2GroupBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCampaignGetV2GroupByFromValue(v string) (*ReportCampaignGetV2GroupBy, error) {
	ev := ReportCampaignGetV2GroupBy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCampaignGetV2GroupBy: valid values are %v", v, AllowedReportCampaignGetV2GroupByEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCampaignGetV2GroupBy) IsValid() bool {
	for _, existing := range AllowedReportCampaignGetV2GroupByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_campaign_get_v2_group_by value
func (v ReportCampaignGetV2GroupBy) Ptr() *ReportCampaignGetV2GroupBy {
	return &v
}
