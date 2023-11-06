/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportAdGetV2DataListCampaignType
type ReportAdGetV2DataListCampaignType string

// List of report_ad_get_v2_data_list_campaign_type
const (
	SEARCH_ReportAdGetV2DataListCampaignType ReportAdGetV2DataListCampaignType = "SEARCH"
	FEED_ReportAdGetV2DataListCampaignType   ReportAdGetV2DataListCampaignType = "FEED"
)

// All allowed values of ReportAdGetV2DataListCampaignType enum
var AllowedReportAdGetV2DataListCampaignTypeEnumValues = []ReportAdGetV2DataListCampaignType{
	"SEARCH",
	"FEED",
}

// NewReportAdGetV2DataListCampaignTypeFromValue returns a pointer to a valid ReportAdGetV2DataListCampaignType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdGetV2DataListCampaignTypeFromValue(v string) (*ReportAdGetV2DataListCampaignType, error) {
	ev := ReportAdGetV2DataListCampaignType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdGetV2DataListCampaignType: valid values are %v", v, AllowedReportAdGetV2DataListCampaignTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdGetV2DataListCampaignType) IsValid() bool {
	for _, existing := range AllowedReportAdGetV2DataListCampaignTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_ad_get_v2_data_list_campaign_type value
func (v ReportAdGetV2DataListCampaignType) Ptr() *ReportAdGetV2DataListCampaignType {
	return &v
}
