/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportAdGetV2DataListPricingCategory
type ReportAdGetV2DataListPricingCategory string

// List of report_ad_get_v2_data_list_pricing_category
const (
	PRICING_CATEGORY_FREE_ReportAdGetV2DataListPricingCategory              ReportAdGetV2DataListPricingCategory = "PRICING_CATEGORY_FREE"
	PRICING_CATEGORY_BID_ReportAdGetV2DataListPricingCategory               ReportAdGetV2DataListPricingCategory = "PRICING_CATEGORY_BID"
	PRICING_CATEGORY_BRAND_ReportAdGetV2DataListPricingCategory             ReportAdGetV2DataListPricingCategory = "PRICING_CATEGORY_BRAND"
	PRICING_CATEGORY_BRAND_AND_PRICING_ReportAdGetV2DataListPricingCategory ReportAdGetV2DataListPricingCategory = "PRICING_CATEGORY_BRAND_AND_PRICING"
	PRICING_CATEGORY_NOC_ReportAdGetV2DataListPricingCategory               ReportAdGetV2DataListPricingCategory = "PRICING_CATEGORY_NOC"
)

// All allowed values of ReportAdGetV2DataListPricingCategory enum
var AllowedReportAdGetV2DataListPricingCategoryEnumValues = []ReportAdGetV2DataListPricingCategory{
	"PRICING_CATEGORY_FREE",
	"PRICING_CATEGORY_BID",
	"PRICING_CATEGORY_BRAND",
	"PRICING_CATEGORY_BRAND_AND_PRICING",
	"PRICING_CATEGORY_NOC",
}

// NewReportAdGetV2DataListPricingCategoryFromValue returns a pointer to a valid ReportAdGetV2DataListPricingCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdGetV2DataListPricingCategoryFromValue(v string) (*ReportAdGetV2DataListPricingCategory, error) {
	ev := ReportAdGetV2DataListPricingCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdGetV2DataListPricingCategory: valid values are %v", v, AllowedReportAdGetV2DataListPricingCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdGetV2DataListPricingCategory) IsValid() bool {
	for _, existing := range AllowedReportAdGetV2DataListPricingCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_ad_get_v2_data_list_pricing_category value
func (v ReportAdGetV2DataListPricingCategory) Ptr() *ReportAdGetV2DataListPricingCategory {
	return &v
}
