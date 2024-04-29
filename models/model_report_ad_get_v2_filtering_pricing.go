/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportAdGetV2FilteringPricing
type ReportAdGetV2FilteringPricing string

// List of report_ad_get_v2_filtering_pricing
const (
	PRICING_CPC_ReportAdGetV2FilteringPricing  ReportAdGetV2FilteringPricing = "PRICING_CPC"
	PRICING_CPV_ReportAdGetV2FilteringPricing  ReportAdGetV2FilteringPricing = "PRICING_CPV"
	PRICING_ECPC_ReportAdGetV2FilteringPricing ReportAdGetV2FilteringPricing = "PRICING_ECPC"
	PRICING_OCPM_ReportAdGetV2FilteringPricing ReportAdGetV2FilteringPricing = "PRICING_OCPM"
	PRICING_OCPC_ReportAdGetV2FilteringPricing ReportAdGetV2FilteringPricing = "PRICING_OCPC"
	PRICING_CPA_ReportAdGetV2FilteringPricing  ReportAdGetV2FilteringPricing = "PRICING_CPA"
	PRICING_CPM_ReportAdGetV2FilteringPricing  ReportAdGetV2FilteringPricing = "PRICING_CPM"
)

// All allowed values of ReportAdGetV2FilteringPricing enum
var AllowedReportAdGetV2FilteringPricingEnumValues = []ReportAdGetV2FilteringPricing{
	"PRICING_CPC",
	"PRICING_CPV",
	"PRICING_ECPC",
	"PRICING_OCPM",
	"PRICING_OCPC",
	"PRICING_CPA",
	"PRICING_CPM",
}

// NewReportAdGetV2FilteringPricingFromValue returns a pointer to a valid ReportAdGetV2FilteringPricing
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdGetV2FilteringPricingFromValue(v string) (*ReportAdGetV2FilteringPricing, error) {
	ev := ReportAdGetV2FilteringPricing(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdGetV2FilteringPricing: valid values are %v", v, AllowedReportAdGetV2FilteringPricingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdGetV2FilteringPricing) IsValid() bool {
	for _, existing := range AllowedReportAdGetV2FilteringPricingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_ad_get_v2_filtering_pricing value
func (v ReportAdGetV2FilteringPricing) Ptr() *ReportAdGetV2FilteringPricing {
	return &v
}
