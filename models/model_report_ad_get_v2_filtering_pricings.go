/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAdGetV2FilteringPricings
type ReportAdGetV2FilteringPricings string

// List of report_ad_get_v2_filtering_pricings
const (
	PRICING_CPC_ReportAdGetV2FilteringPricings  ReportAdGetV2FilteringPricings = "PRICING_CPC"
	PRICING_OCPM_ReportAdGetV2FilteringPricings ReportAdGetV2FilteringPricings = "PRICING_OCPM"
	PRICING_CPV_ReportAdGetV2FilteringPricings  ReportAdGetV2FilteringPricings = "PRICING_CPV"
	PRICING_CPA_ReportAdGetV2FilteringPricings  ReportAdGetV2FilteringPricings = "PRICING_CPA"
	PRICING_CPM_ReportAdGetV2FilteringPricings  ReportAdGetV2FilteringPricings = "PRICING_CPM"
	PRICING_OCPC_ReportAdGetV2FilteringPricings ReportAdGetV2FilteringPricings = "PRICING_OCPC"
	PRICING_ECPC_ReportAdGetV2FilteringPricings ReportAdGetV2FilteringPricings = "PRICING_ECPC"
)

// Ptr returns reference to report_ad_get_v2_filtering_pricings value
func (v ReportAdGetV2FilteringPricings) Ptr() *ReportAdGetV2FilteringPricings {
	return &v
}
