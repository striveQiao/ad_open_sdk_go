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

// BrandQueryStockV30AudienceInfoDistrictInfoCitySelectType
type BrandQueryStockV30AudienceInfoDistrictInfoCitySelectType string

// List of brand_query_stock_v3.0_audience_info_district_info_city_select_type
const (
	EXCLUDE_BrandQueryStockV30AudienceInfoDistrictInfoCitySelectType BrandQueryStockV30AudienceInfoDistrictInfoCitySelectType = "EXCLUDE"
	SELECT_BrandQueryStockV30AudienceInfoDistrictInfoCitySelectType  BrandQueryStockV30AudienceInfoDistrictInfoCitySelectType = "SELECT"
)

// All allowed values of BrandQueryStockV30AudienceInfoDistrictInfoCitySelectType enum
var AllowedBrandQueryStockV30AudienceInfoDistrictInfoCitySelectTypeEnumValues = []BrandQueryStockV30AudienceInfoDistrictInfoCitySelectType{
	"EXCLUDE",
	"SELECT",
}

// NewBrandQueryStockV30AudienceInfoDistrictInfoCitySelectTypeFromValue returns a pointer to a valid BrandQueryStockV30AudienceInfoDistrictInfoCitySelectType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandQueryStockV30AudienceInfoDistrictInfoCitySelectTypeFromValue(v string) (*BrandQueryStockV30AudienceInfoDistrictInfoCitySelectType, error) {
	ev := BrandQueryStockV30AudienceInfoDistrictInfoCitySelectType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandQueryStockV30AudienceInfoDistrictInfoCitySelectType: valid values are %v", v, AllowedBrandQueryStockV30AudienceInfoDistrictInfoCitySelectTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandQueryStockV30AudienceInfoDistrictInfoCitySelectType) IsValid() bool {
	for _, existing := range AllowedBrandQueryStockV30AudienceInfoDistrictInfoCitySelectTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_query_stock_v3.0_audience_info_district_info_city_select_type value
func (v BrandQueryStockV30AudienceInfoDistrictInfoCitySelectType) Ptr() *BrandQueryStockV30AudienceInfoDistrictInfoCitySelectType {
	return &v
}
