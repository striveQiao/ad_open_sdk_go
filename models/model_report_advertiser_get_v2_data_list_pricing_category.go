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

// ReportAdvertiserGetV2DataListPricingCategory
type ReportAdvertiserGetV2DataListPricingCategory string

// List of report_advertiser_get_v2_data_list_pricing_category
const (
	PRICING_CATEGORY_BRAND_AND_PRICING_ReportAdvertiserGetV2DataListPricingCategory ReportAdvertiserGetV2DataListPricingCategory = "PRICING_CATEGORY_BRAND_AND_PRICING"
	PRICING_CATEGORY_BRAND_ReportAdvertiserGetV2DataListPricingCategory             ReportAdvertiserGetV2DataListPricingCategory = "PRICING_CATEGORY_BRAND"
	PRICING_CATEGORY_BID_ReportAdvertiserGetV2DataListPricingCategory               ReportAdvertiserGetV2DataListPricingCategory = "PRICING_CATEGORY_BID"
	PRICING_CATEGORY_FREE_ReportAdvertiserGetV2DataListPricingCategory              ReportAdvertiserGetV2DataListPricingCategory = "PRICING_CATEGORY_FREE"
	PRICING_CATEGORY_NOC_ReportAdvertiserGetV2DataListPricingCategory               ReportAdvertiserGetV2DataListPricingCategory = "PRICING_CATEGORY_NOC"
)

// All allowed values of ReportAdvertiserGetV2DataListPricingCategory enum
var AllowedReportAdvertiserGetV2DataListPricingCategoryEnumValues = []ReportAdvertiserGetV2DataListPricingCategory{
	"PRICING_CATEGORY_BRAND_AND_PRICING",
	"PRICING_CATEGORY_BRAND",
	"PRICING_CATEGORY_BID",
	"PRICING_CATEGORY_FREE",
	"PRICING_CATEGORY_NOC",
}

// NewReportAdvertiserGetV2DataListPricingCategoryFromValue returns a pointer to a valid ReportAdvertiserGetV2DataListPricingCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdvertiserGetV2DataListPricingCategoryFromValue(v string) (*ReportAdvertiserGetV2DataListPricingCategory, error) {
	ev := ReportAdvertiserGetV2DataListPricingCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdvertiserGetV2DataListPricingCategory: valid values are %v", v, AllowedReportAdvertiserGetV2DataListPricingCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdvertiserGetV2DataListPricingCategory) IsValid() bool {
	for _, existing := range AllowedReportAdvertiserGetV2DataListPricingCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_advertiser_get_v2_data_list_pricing_category value
func (v ReportAdvertiserGetV2DataListPricingCategory) Ptr() *ReportAdvertiserGetV2DataListPricingCategory {
	return &v
}
