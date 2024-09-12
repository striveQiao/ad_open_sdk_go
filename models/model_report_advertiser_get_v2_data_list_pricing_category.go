/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAdvertiserGetV2DataListPricingCategory
type ReportAdvertiserGetV2DataListPricingCategory string

// List of report_advertiser_get_v2_data_list_pricing_category
const (
	PRICING_CATEGORY_BRAND_ReportAdvertiserGetV2DataListPricingCategory             ReportAdvertiserGetV2DataListPricingCategory = "PRICING_CATEGORY_BRAND"
	PRICING_CATEGORY_BRAND_AND_PRICING_ReportAdvertiserGetV2DataListPricingCategory ReportAdvertiserGetV2DataListPricingCategory = "PRICING_CATEGORY_BRAND_AND_PRICING"
	PRICING_CATEGORY_FREE_ReportAdvertiserGetV2DataListPricingCategory              ReportAdvertiserGetV2DataListPricingCategory = "PRICING_CATEGORY_FREE"
	PRICING_CATEGORY_NOC_ReportAdvertiserGetV2DataListPricingCategory               ReportAdvertiserGetV2DataListPricingCategory = "PRICING_CATEGORY_NOC"
	PRICING_CATEGORY_BID_ReportAdvertiserGetV2DataListPricingCategory               ReportAdvertiserGetV2DataListPricingCategory = "PRICING_CATEGORY_BID"
)

// Ptr returns reference to report_advertiser_get_v2_data_list_pricing_category value
func (v ReportAdvertiserGetV2DataListPricingCategory) Ptr() *ReportAdvertiserGetV2DataListPricingCategory {
	return &v
}
