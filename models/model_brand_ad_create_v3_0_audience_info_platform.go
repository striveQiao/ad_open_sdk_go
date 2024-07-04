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

// BrandAdCreateV30AudienceInfoPlatform
type BrandAdCreateV30AudienceInfoPlatform string

// List of brand_ad_create_v3.0_audience_info_platform
const (
	ANDROID_BrandAdCreateV30AudienceInfoPlatform   BrandAdCreateV30AudienceInfoPlatform = "ANDROID"
	IOS_BrandAdCreateV30AudienceInfoPlatform       BrandAdCreateV30AudienceInfoPlatform = "IOS"
	UNLIMITED_BrandAdCreateV30AudienceInfoPlatform BrandAdCreateV30AudienceInfoPlatform = "UNLIMITED"
)

// All allowed values of BrandAdCreateV30AudienceInfoPlatform enum
var AllowedBrandAdCreateV30AudienceInfoPlatformEnumValues = []BrandAdCreateV30AudienceInfoPlatform{
	"ANDROID",
	"IOS",
	"UNLIMITED",
}

// NewBrandAdCreateV30AudienceInfoPlatformFromValue returns a pointer to a valid BrandAdCreateV30AudienceInfoPlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdCreateV30AudienceInfoPlatformFromValue(v string) (*BrandAdCreateV30AudienceInfoPlatform, error) {
	ev := BrandAdCreateV30AudienceInfoPlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdCreateV30AudienceInfoPlatform: valid values are %v", v, AllowedBrandAdCreateV30AudienceInfoPlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdCreateV30AudienceInfoPlatform) IsValid() bool {
	for _, existing := range AllowedBrandAdCreateV30AudienceInfoPlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_create_v3.0_audience_info_platform value
func (v BrandAdCreateV30AudienceInfoPlatform) Ptr() *BrandAdCreateV30AudienceInfoPlatform {
	return &v
}
