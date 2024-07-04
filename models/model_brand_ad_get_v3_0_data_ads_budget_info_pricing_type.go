/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// BrandAdGetV30DataAdsBudgetInfoPricingType
type BrandAdGetV30DataAdsBudgetInfoPricingType int64

// List of brand_ad_get_v3.0_data_ads_budget_info_pricing_type
const (
	Enum_1_BrandAdGetV30DataAdsBudgetInfoPricingType  BrandAdGetV30DataAdsBudgetInfoPricingType = 1
	Enum_2_BrandAdGetV30DataAdsBudgetInfoPricingType  BrandAdGetV30DataAdsBudgetInfoPricingType = 2
	Enum_3_BrandAdGetV30DataAdsBudgetInfoPricingType  BrandAdGetV30DataAdsBudgetInfoPricingType = 3
	Enum_6_BrandAdGetV30DataAdsBudgetInfoPricingType  BrandAdGetV30DataAdsBudgetInfoPricingType = 6
	Enum_7_BrandAdGetV30DataAdsBudgetInfoPricingType  BrandAdGetV30DataAdsBudgetInfoPricingType = 7
	Enum_8_BrandAdGetV30DataAdsBudgetInfoPricingType  BrandAdGetV30DataAdsBudgetInfoPricingType = 8
	Enum_9_BrandAdGetV30DataAdsBudgetInfoPricingType  BrandAdGetV30DataAdsBudgetInfoPricingType = 9
	Enum_10_BrandAdGetV30DataAdsBudgetInfoPricingType BrandAdGetV30DataAdsBudgetInfoPricingType = 10
)

// All allowed values of BrandAdGetV30DataAdsBudgetInfoPricingType enum
var AllowedBrandAdGetV30DataAdsBudgetInfoPricingTypeEnumValues = []BrandAdGetV30DataAdsBudgetInfoPricingType{
	1,
	2,
	3,
	6,
	7,
	8,
	9,
	10,
}

// NewBrandAdGetV30DataAdsBudgetInfoPricingTypeFromValue returns a pointer to a valid BrandAdGetV30DataAdsBudgetInfoPricingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdGetV30DataAdsBudgetInfoPricingTypeFromValue(v int64) (*BrandAdGetV30DataAdsBudgetInfoPricingType, error) {
	ev := BrandAdGetV30DataAdsBudgetInfoPricingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdGetV30DataAdsBudgetInfoPricingType: valid values are %v", v, AllowedBrandAdGetV30DataAdsBudgetInfoPricingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdGetV30DataAdsBudgetInfoPricingType) IsValid() bool {
	for _, existing := range AllowedBrandAdGetV30DataAdsBudgetInfoPricingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_get_v3.0_data_ads_budget_info_pricing_type value
func (v BrandAdGetV30DataAdsBudgetInfoPricingType) Ptr() *BrandAdGetV30DataAdsBudgetInfoPricingType {
	return &v
}
