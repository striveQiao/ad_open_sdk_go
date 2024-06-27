/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportCreativeGetV2DataListCampaignType
type ReportCreativeGetV2DataListCampaignType string

// List of report_creative_get_v2_data_list_campaign_type
const (
	SEARCH_ReportCreativeGetV2DataListCampaignType ReportCreativeGetV2DataListCampaignType = "SEARCH"
	FEED_ReportCreativeGetV2DataListCampaignType   ReportCreativeGetV2DataListCampaignType = "FEED"
)

// All allowed values of ReportCreativeGetV2DataListCampaignType enum
var AllowedReportCreativeGetV2DataListCampaignTypeEnumValues = []ReportCreativeGetV2DataListCampaignType{
	"SEARCH",
	"FEED",
}

// NewReportCreativeGetV2DataListCampaignTypeFromValue returns a pointer to a valid ReportCreativeGetV2DataListCampaignType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCreativeGetV2DataListCampaignTypeFromValue(v string) (*ReportCreativeGetV2DataListCampaignType, error) {
	ev := ReportCreativeGetV2DataListCampaignType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCreativeGetV2DataListCampaignType: valid values are %v", v, AllowedReportCreativeGetV2DataListCampaignTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCreativeGetV2DataListCampaignType) IsValid() bool {
	for _, existing := range AllowedReportCreativeGetV2DataListCampaignTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_creative_get_v2_data_list_campaign_type value
func (v ReportCreativeGetV2DataListCampaignType) Ptr() *ReportCreativeGetV2DataListCampaignType {
	return &v
}
