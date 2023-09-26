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

// ReportCampaignGetV2DataListPricing
type ReportCampaignGetV2DataListPricing string

// List of report_campaign_get_v2_data_list_pricing
const (
	PRICING_OCPM_ReportCampaignGetV2DataListPricing ReportCampaignGetV2DataListPricing = "PRICING_OCPM"
	PRICING_CPA_ReportCampaignGetV2DataListPricing  ReportCampaignGetV2DataListPricing = "PRICING_CPA"
	PRICING_CPC_ReportCampaignGetV2DataListPricing  ReportCampaignGetV2DataListPricing = "PRICING_CPC"
	PRICING_CPV_ReportCampaignGetV2DataListPricing  ReportCampaignGetV2DataListPricing = "PRICING_CPV"
	PRICING_ECPC_ReportCampaignGetV2DataListPricing ReportCampaignGetV2DataListPricing = "PRICING_ECPC"
	PRICING_OCPC_ReportCampaignGetV2DataListPricing ReportCampaignGetV2DataListPricing = "PRICING_OCPC"
	PRICING_CPM_ReportCampaignGetV2DataListPricing  ReportCampaignGetV2DataListPricing = "PRICING_CPM"
)

// All allowed values of ReportCampaignGetV2DataListPricing enum
var AllowedReportCampaignGetV2DataListPricingEnumValues = []ReportCampaignGetV2DataListPricing{
	"PRICING_OCPM",
	"PRICING_CPA",
	"PRICING_CPC",
	"PRICING_CPV",
	"PRICING_ECPC",
	"PRICING_OCPC",
	"PRICING_CPM",
}

// NewReportCampaignGetV2DataListPricingFromValue returns a pointer to a valid ReportCampaignGetV2DataListPricing
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCampaignGetV2DataListPricingFromValue(v string) (*ReportCampaignGetV2DataListPricing, error) {
	ev := ReportCampaignGetV2DataListPricing(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCampaignGetV2DataListPricing: valid values are %v", v, AllowedReportCampaignGetV2DataListPricingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCampaignGetV2DataListPricing) IsValid() bool {
	for _, existing := range AllowedReportCampaignGetV2DataListPricingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_campaign_get_v2_data_list_pricing value
func (v ReportCampaignGetV2DataListPricing) Ptr() *ReportCampaignGetV2DataListPricing {
	return &v
}
