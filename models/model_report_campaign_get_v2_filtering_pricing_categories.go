/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportCampaignGetV2FilteringPricingCategories
type ReportCampaignGetV2FilteringPricingCategories string

// List of report_campaign_get_v2_filtering_pricing_categories
const (
	PRICING_CATEGORY_BRAND_AND_PRICING_ReportCampaignGetV2FilteringPricingCategories ReportCampaignGetV2FilteringPricingCategories = "PRICING_CATEGORY_BRAND_AND_PRICING"
	PRICING_CATEGORY_NOC_ReportCampaignGetV2FilteringPricingCategories               ReportCampaignGetV2FilteringPricingCategories = "PRICING_CATEGORY_NOC"
	PRICING_CATEGORY_FREE_ReportCampaignGetV2FilteringPricingCategories              ReportCampaignGetV2FilteringPricingCategories = "PRICING_CATEGORY_FREE"
	PRICING_CATEGORY_BID_ReportCampaignGetV2FilteringPricingCategories               ReportCampaignGetV2FilteringPricingCategories = "PRICING_CATEGORY_BID"
	PRICING_CATEGORY_BRAND_ReportCampaignGetV2FilteringPricingCategories             ReportCampaignGetV2FilteringPricingCategories = "PRICING_CATEGORY_BRAND"
)

// All allowed values of ReportCampaignGetV2FilteringPricingCategories enum
var AllowedReportCampaignGetV2FilteringPricingCategoriesEnumValues = []ReportCampaignGetV2FilteringPricingCategories{
	"PRICING_CATEGORY_BRAND_AND_PRICING",
	"PRICING_CATEGORY_NOC",
	"PRICING_CATEGORY_FREE",
	"PRICING_CATEGORY_BID",
	"PRICING_CATEGORY_BRAND",
}

// NewReportCampaignGetV2FilteringPricingCategoriesFromValue returns a pointer to a valid ReportCampaignGetV2FilteringPricingCategories
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCampaignGetV2FilteringPricingCategoriesFromValue(v string) (*ReportCampaignGetV2FilteringPricingCategories, error) {
	ev := ReportCampaignGetV2FilteringPricingCategories(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCampaignGetV2FilteringPricingCategories: valid values are %v", v, AllowedReportCampaignGetV2FilteringPricingCategoriesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCampaignGetV2FilteringPricingCategories) IsValid() bool {
	for _, existing := range AllowedReportCampaignGetV2FilteringPricingCategoriesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_campaign_get_v2_filtering_pricing_categories value
func (v ReportCampaignGetV2FilteringPricingCategories) Ptr() *ReportCampaignGetV2FilteringPricingCategories {
	return &v
}
