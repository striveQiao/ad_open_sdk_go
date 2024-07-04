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

// BrandQueryStockV30AudienceInfoPlatform
type BrandQueryStockV30AudienceInfoPlatform string

// List of brand_query_stock_v3.0_audience_info_platform
const (
	ANDROID_BrandQueryStockV30AudienceInfoPlatform   BrandQueryStockV30AudienceInfoPlatform = "ANDROID"
	IOS_BrandQueryStockV30AudienceInfoPlatform       BrandQueryStockV30AudienceInfoPlatform = "IOS"
	UNLIMITED_BrandQueryStockV30AudienceInfoPlatform BrandQueryStockV30AudienceInfoPlatform = "UNLIMITED"
)

// All allowed values of BrandQueryStockV30AudienceInfoPlatform enum
var AllowedBrandQueryStockV30AudienceInfoPlatformEnumValues = []BrandQueryStockV30AudienceInfoPlatform{
	"ANDROID",
	"IOS",
	"UNLIMITED",
}

// NewBrandQueryStockV30AudienceInfoPlatformFromValue returns a pointer to a valid BrandQueryStockV30AudienceInfoPlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandQueryStockV30AudienceInfoPlatformFromValue(v string) (*BrandQueryStockV30AudienceInfoPlatform, error) {
	ev := BrandQueryStockV30AudienceInfoPlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandQueryStockV30AudienceInfoPlatform: valid values are %v", v, AllowedBrandQueryStockV30AudienceInfoPlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandQueryStockV30AudienceInfoPlatform) IsValid() bool {
	for _, existing := range AllowedBrandQueryStockV30AudienceInfoPlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_query_stock_v3.0_audience_info_platform value
func (v BrandQueryStockV30AudienceInfoPlatform) Ptr() *BrandQueryStockV30AudienceInfoPlatform {
	return &v
}
