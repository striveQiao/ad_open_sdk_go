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

// ReportCampaignGetV2DataListRealRecallMatchType
type ReportCampaignGetV2DataListRealRecallMatchType string

// List of report_campaign_get_v2_data_list_real_recall_match_type
const (
	EXTENSIVE_ReportCampaignGetV2DataListRealRecallMatchType ReportCampaignGetV2DataListRealRecallMatchType = "EXTENSIVE"
	PHRASE_ReportCampaignGetV2DataListRealRecallMatchType    ReportCampaignGetV2DataListRealRecallMatchType = "PHRASE"
	PRECISION_ReportCampaignGetV2DataListRealRecallMatchType ReportCampaignGetV2DataListRealRecallMatchType = "PRECISION"
)

// All allowed values of ReportCampaignGetV2DataListRealRecallMatchType enum
var AllowedReportCampaignGetV2DataListRealRecallMatchTypeEnumValues = []ReportCampaignGetV2DataListRealRecallMatchType{
	"EXTENSIVE",
	"PHRASE",
	"PRECISION",
}

// NewReportCampaignGetV2DataListRealRecallMatchTypeFromValue returns a pointer to a valid ReportCampaignGetV2DataListRealRecallMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCampaignGetV2DataListRealRecallMatchTypeFromValue(v string) (*ReportCampaignGetV2DataListRealRecallMatchType, error) {
	ev := ReportCampaignGetV2DataListRealRecallMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCampaignGetV2DataListRealRecallMatchType: valid values are %v", v, AllowedReportCampaignGetV2DataListRealRecallMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCampaignGetV2DataListRealRecallMatchType) IsValid() bool {
	for _, existing := range AllowedReportCampaignGetV2DataListRealRecallMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_campaign_get_v2_data_list_real_recall_match_type value
func (v ReportCampaignGetV2DataListRealRecallMatchType) Ptr() *ReportCampaignGetV2DataListRealRecallMatchType {
	return &v
}
