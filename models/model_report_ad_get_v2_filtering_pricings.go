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

// ReportAdGetV2FilteringPricings
type ReportAdGetV2FilteringPricings string

// List of report_ad_get_v2_filtering_pricings
const (
	PRICING_OCPM_ReportAdGetV2FilteringPricings ReportAdGetV2FilteringPricings = "PRICING_OCPM"
	PRICING_CPA_ReportAdGetV2FilteringPricings  ReportAdGetV2FilteringPricings = "PRICING_CPA"
	PRICING_CPC_ReportAdGetV2FilteringPricings  ReportAdGetV2FilteringPricings = "PRICING_CPC"
	PRICING_CPV_ReportAdGetV2FilteringPricings  ReportAdGetV2FilteringPricings = "PRICING_CPV"
	PRICING_ECPC_ReportAdGetV2FilteringPricings ReportAdGetV2FilteringPricings = "PRICING_ECPC"
	PRICING_OCPC_ReportAdGetV2FilteringPricings ReportAdGetV2FilteringPricings = "PRICING_OCPC"
	PRICING_CPM_ReportAdGetV2FilteringPricings  ReportAdGetV2FilteringPricings = "PRICING_CPM"
)

// All allowed values of ReportAdGetV2FilteringPricings enum
var AllowedReportAdGetV2FilteringPricingsEnumValues = []ReportAdGetV2FilteringPricings{
	"PRICING_OCPM",
	"PRICING_CPA",
	"PRICING_CPC",
	"PRICING_CPV",
	"PRICING_ECPC",
	"PRICING_OCPC",
	"PRICING_CPM",
}

// NewReportAdGetV2FilteringPricingsFromValue returns a pointer to a valid ReportAdGetV2FilteringPricings
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdGetV2FilteringPricingsFromValue(v string) (*ReportAdGetV2FilteringPricings, error) {
	ev := ReportAdGetV2FilteringPricings(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdGetV2FilteringPricings: valid values are %v", v, AllowedReportAdGetV2FilteringPricingsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdGetV2FilteringPricings) IsValid() bool {
	for _, existing := range AllowedReportAdGetV2FilteringPricingsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_ad_get_v2_filtering_pricings value
func (v ReportAdGetV2FilteringPricings) Ptr() *ReportAdGetV2FilteringPricings {
	return &v
}
