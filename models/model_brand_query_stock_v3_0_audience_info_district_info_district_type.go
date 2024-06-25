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

// BrandQueryStockV30AudienceInfoDistrictInfoDistrictType
type BrandQueryStockV30AudienceInfoDistrictInfoDistrictType string

// List of brand_query_stock_v3.0_audience_info_district_info_district_type
const (
	REGION_BrandQueryStockV30AudienceInfoDistrictInfoDistrictType    BrandQueryStockV30AudienceInfoDistrictInfoDistrictType = "REGION"
	UNLIMITED_BrandQueryStockV30AudienceInfoDistrictInfoDistrictType BrandQueryStockV30AudienceInfoDistrictInfoDistrictType = "UNLIMITED"
)

// All allowed values of BrandQueryStockV30AudienceInfoDistrictInfoDistrictType enum
var AllowedBrandQueryStockV30AudienceInfoDistrictInfoDistrictTypeEnumValues = []BrandQueryStockV30AudienceInfoDistrictInfoDistrictType{
	"REGION",
	"UNLIMITED",
}

// NewBrandQueryStockV30AudienceInfoDistrictInfoDistrictTypeFromValue returns a pointer to a valid BrandQueryStockV30AudienceInfoDistrictInfoDistrictType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandQueryStockV30AudienceInfoDistrictInfoDistrictTypeFromValue(v string) (*BrandQueryStockV30AudienceInfoDistrictInfoDistrictType, error) {
	ev := BrandQueryStockV30AudienceInfoDistrictInfoDistrictType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandQueryStockV30AudienceInfoDistrictInfoDistrictType: valid values are %v", v, AllowedBrandQueryStockV30AudienceInfoDistrictInfoDistrictTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandQueryStockV30AudienceInfoDistrictInfoDistrictType) IsValid() bool {
	for _, existing := range AllowedBrandQueryStockV30AudienceInfoDistrictInfoDistrictTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_query_stock_v3.0_audience_info_district_info_district_type value
func (v BrandQueryStockV30AudienceInfoDistrictInfoDistrictType) Ptr() *BrandQueryStockV30AudienceInfoDistrictInfoDistrictType {
	return &v
}
