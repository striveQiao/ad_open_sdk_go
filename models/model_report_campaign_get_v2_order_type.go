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

// ReportCampaignGetV2OrderType
type ReportCampaignGetV2OrderType string

// List of report_campaign_get_v2_order_type
const (
	ASC_ReportCampaignGetV2OrderType  ReportCampaignGetV2OrderType = "ASC"
	DESC_ReportCampaignGetV2OrderType ReportCampaignGetV2OrderType = "DESC"
)

// All allowed values of ReportCampaignGetV2OrderType enum
var AllowedReportCampaignGetV2OrderTypeEnumValues = []ReportCampaignGetV2OrderType{
	"ASC",
	"DESC",
}

// NewReportCampaignGetV2OrderTypeFromValue returns a pointer to a valid ReportCampaignGetV2OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCampaignGetV2OrderTypeFromValue(v string) (*ReportCampaignGetV2OrderType, error) {
	ev := ReportCampaignGetV2OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCampaignGetV2OrderType: valid values are %v", v, AllowedReportCampaignGetV2OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCampaignGetV2OrderType) IsValid() bool {
	for _, existing := range AllowedReportCampaignGetV2OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_campaign_get_v2_order_type value
func (v ReportCampaignGetV2OrderType) Ptr() *ReportCampaignGetV2OrderType {
	return &v
}
