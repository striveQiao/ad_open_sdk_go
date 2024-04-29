/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// BrandAdGetV30DataAdsAudienceInfoPlatformV2
type BrandAdGetV30DataAdsAudienceInfoPlatformV2 int64

// List of brand_ad_get_v3.0_data_ads_audience_info_platform_v2
const (
	Enum_1_BrandAdGetV30DataAdsAudienceInfoPlatformV2 BrandAdGetV30DataAdsAudienceInfoPlatformV2 = 1
	Enum_2_BrandAdGetV30DataAdsAudienceInfoPlatformV2 BrandAdGetV30DataAdsAudienceInfoPlatformV2 = 2
	Enum_3_BrandAdGetV30DataAdsAudienceInfoPlatformV2 BrandAdGetV30DataAdsAudienceInfoPlatformV2 = 3
	Enum_4_BrandAdGetV30DataAdsAudienceInfoPlatformV2 BrandAdGetV30DataAdsAudienceInfoPlatformV2 = 4
	Enum_5_BrandAdGetV30DataAdsAudienceInfoPlatformV2 BrandAdGetV30DataAdsAudienceInfoPlatformV2 = 5
)

// All allowed values of BrandAdGetV30DataAdsAudienceInfoPlatformV2 enum
var AllowedBrandAdGetV30DataAdsAudienceInfoPlatformV2EnumValues = []BrandAdGetV30DataAdsAudienceInfoPlatformV2{
	1,
	2,
	3,
	4,
	5,
}

// NewBrandAdGetV30DataAdsAudienceInfoPlatformV2FromValue returns a pointer to a valid BrandAdGetV30DataAdsAudienceInfoPlatformV2
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdGetV30DataAdsAudienceInfoPlatformV2FromValue(v int64) (*BrandAdGetV30DataAdsAudienceInfoPlatformV2, error) {
	ev := BrandAdGetV30DataAdsAudienceInfoPlatformV2(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdGetV30DataAdsAudienceInfoPlatformV2: valid values are %v", v, AllowedBrandAdGetV30DataAdsAudienceInfoPlatformV2EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdGetV30DataAdsAudienceInfoPlatformV2) IsValid() bool {
	for _, existing := range AllowedBrandAdGetV30DataAdsAudienceInfoPlatformV2EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_get_v3.0_data_ads_audience_info_platform_v2 value
func (v BrandAdGetV30DataAdsAudienceInfoPlatformV2) Ptr() *BrandAdGetV30DataAdsAudienceInfoPlatformV2 {
	return &v
}
