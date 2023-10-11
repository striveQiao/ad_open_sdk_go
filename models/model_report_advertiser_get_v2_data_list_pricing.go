/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportAdvertiserGetV2DataListPricing
type ReportAdvertiserGetV2DataListPricing string

// List of report_advertiser_get_v2_data_list_pricing
const (
	PRICING_OCPM_ReportAdvertiserGetV2DataListPricing ReportAdvertiserGetV2DataListPricing = "PRICING_OCPM"
	PRICING_CPA_ReportAdvertiserGetV2DataListPricing  ReportAdvertiserGetV2DataListPricing = "PRICING_CPA"
	PRICING_CPC_ReportAdvertiserGetV2DataListPricing  ReportAdvertiserGetV2DataListPricing = "PRICING_CPC"
	PRICING_CPV_ReportAdvertiserGetV2DataListPricing  ReportAdvertiserGetV2DataListPricing = "PRICING_CPV"
	PRICING_ECPC_ReportAdvertiserGetV2DataListPricing ReportAdvertiserGetV2DataListPricing = "PRICING_ECPC"
	PRICING_OCPC_ReportAdvertiserGetV2DataListPricing ReportAdvertiserGetV2DataListPricing = "PRICING_OCPC"
	PRICING_CPM_ReportAdvertiserGetV2DataListPricing  ReportAdvertiserGetV2DataListPricing = "PRICING_CPM"
)

// All allowed values of ReportAdvertiserGetV2DataListPricing enum
var AllowedReportAdvertiserGetV2DataListPricingEnumValues = []ReportAdvertiserGetV2DataListPricing{
	"PRICING_OCPM",
	"PRICING_CPA",
	"PRICING_CPC",
	"PRICING_CPV",
	"PRICING_ECPC",
	"PRICING_OCPC",
	"PRICING_CPM",
}

// NewReportAdvertiserGetV2DataListPricingFromValue returns a pointer to a valid ReportAdvertiserGetV2DataListPricing
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdvertiserGetV2DataListPricingFromValue(v string) (*ReportAdvertiserGetV2DataListPricing, error) {
	ev := ReportAdvertiserGetV2DataListPricing(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdvertiserGetV2DataListPricing: valid values are %v", v, AllowedReportAdvertiserGetV2DataListPricingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdvertiserGetV2DataListPricing) IsValid() bool {
	for _, existing := range AllowedReportAdvertiserGetV2DataListPricingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_advertiser_get_v2_data_list_pricing value
func (v ReportAdvertiserGetV2DataListPricing) Ptr() *ReportAdvertiserGetV2DataListPricing {
	return &v
}
