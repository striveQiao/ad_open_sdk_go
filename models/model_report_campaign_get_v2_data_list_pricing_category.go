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

// ReportCampaignGetV2DataListPricingCategory
type ReportCampaignGetV2DataListPricingCategory string

// List of report_campaign_get_v2_data_list_pricing_category
const (
	PRICING_CATEGORY_NOC_ReportCampaignGetV2DataListPricingCategory               ReportCampaignGetV2DataListPricingCategory = "PRICING_CATEGORY_NOC"
	PRICING_CATEGORY_BID_ReportCampaignGetV2DataListPricingCategory               ReportCampaignGetV2DataListPricingCategory = "PRICING_CATEGORY_BID"
	PRICING_CATEGORY_FREE_ReportCampaignGetV2DataListPricingCategory              ReportCampaignGetV2DataListPricingCategory = "PRICING_CATEGORY_FREE"
	PRICING_CATEGORY_BRAND_ReportCampaignGetV2DataListPricingCategory             ReportCampaignGetV2DataListPricingCategory = "PRICING_CATEGORY_BRAND"
	PRICING_CATEGORY_BRAND_AND_PRICING_ReportCampaignGetV2DataListPricingCategory ReportCampaignGetV2DataListPricingCategory = "PRICING_CATEGORY_BRAND_AND_PRICING"
)

// All allowed values of ReportCampaignGetV2DataListPricingCategory enum
var AllowedReportCampaignGetV2DataListPricingCategoryEnumValues = []ReportCampaignGetV2DataListPricingCategory{
	"PRICING_CATEGORY_NOC",
	"PRICING_CATEGORY_BID",
	"PRICING_CATEGORY_FREE",
	"PRICING_CATEGORY_BRAND",
	"PRICING_CATEGORY_BRAND_AND_PRICING",
}

// NewReportCampaignGetV2DataListPricingCategoryFromValue returns a pointer to a valid ReportCampaignGetV2DataListPricingCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCampaignGetV2DataListPricingCategoryFromValue(v string) (*ReportCampaignGetV2DataListPricingCategory, error) {
	ev := ReportCampaignGetV2DataListPricingCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCampaignGetV2DataListPricingCategory: valid values are %v", v, AllowedReportCampaignGetV2DataListPricingCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCampaignGetV2DataListPricingCategory) IsValid() bool {
	for _, existing := range AllowedReportCampaignGetV2DataListPricingCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_campaign_get_v2_data_list_pricing_category value
func (v ReportCampaignGetV2DataListPricingCategory) Ptr() *ReportCampaignGetV2DataListPricingCategory {
	return &v
}
