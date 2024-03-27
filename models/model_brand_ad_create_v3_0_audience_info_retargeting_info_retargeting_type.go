/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// BrandAdCreateV30AudienceInfoRetargetingInfoRetargetingType
type BrandAdCreateV30AudienceInfoRetargetingInfoRetargetingType string

// List of brand_ad_create_v3.0_audience_info_retargeting_info_retargeting_type
const (
	ECOMMERCE_BrandAdCreateV30AudienceInfoRetargetingInfoRetargetingType BrandAdCreateV30AudienceInfoRetargetingInfoRetargetingType = "ECOMMERCE"
	EXCLUDE_BrandAdCreateV30AudienceInfoRetargetingInfoRetargetingType   BrandAdCreateV30AudienceInfoRetargetingInfoRetargetingType = "EXCLUDE"
	INCLUDE_BrandAdCreateV30AudienceInfoRetargetingInfoRetargetingType   BrandAdCreateV30AudienceInfoRetargetingInfoRetargetingType = "INCLUDE"
	UNLIMITED_BrandAdCreateV30AudienceInfoRetargetingInfoRetargetingType BrandAdCreateV30AudienceInfoRetargetingInfoRetargetingType = "UNLIMITED"
)

// All allowed values of BrandAdCreateV30AudienceInfoRetargetingInfoRetargetingType enum
var AllowedBrandAdCreateV30AudienceInfoRetargetingInfoRetargetingTypeEnumValues = []BrandAdCreateV30AudienceInfoRetargetingInfoRetargetingType{
	"ECOMMERCE",
	"EXCLUDE",
	"INCLUDE",
	"UNLIMITED",
}

// NewBrandAdCreateV30AudienceInfoRetargetingInfoRetargetingTypeFromValue returns a pointer to a valid BrandAdCreateV30AudienceInfoRetargetingInfoRetargetingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdCreateV30AudienceInfoRetargetingInfoRetargetingTypeFromValue(v string) (*BrandAdCreateV30AudienceInfoRetargetingInfoRetargetingType, error) {
	ev := BrandAdCreateV30AudienceInfoRetargetingInfoRetargetingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdCreateV30AudienceInfoRetargetingInfoRetargetingType: valid values are %v", v, AllowedBrandAdCreateV30AudienceInfoRetargetingInfoRetargetingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdCreateV30AudienceInfoRetargetingInfoRetargetingType) IsValid() bool {
	for _, existing := range AllowedBrandAdCreateV30AudienceInfoRetargetingInfoRetargetingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_create_v3.0_audience_info_retargeting_info_retargeting_type value
func (v BrandAdCreateV30AudienceInfoRetargetingInfoRetargetingType) Ptr() *BrandAdCreateV30AudienceInfoRetargetingInfoRetargetingType {
	return &v
}
