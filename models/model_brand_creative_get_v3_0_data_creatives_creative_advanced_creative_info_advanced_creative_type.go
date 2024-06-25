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

// BrandCreativeGetV30DataCreativesCreativeAdvancedCreativeInfoAdvancedCreativeType
type BrandCreativeGetV30DataCreativesCreativeAdvancedCreativeInfoAdvancedCreativeType int64

// List of brand_creative_get_v3.0_data_creatives_creative_advanced_creative_info_advanced_creative_type
const (
	Enum_0_BrandCreativeGetV30DataCreativesCreativeAdvancedCreativeInfoAdvancedCreativeType  BrandCreativeGetV30DataCreativesCreativeAdvancedCreativeInfoAdvancedCreativeType = 0
	Enum_1_BrandCreativeGetV30DataCreativesCreativeAdvancedCreativeInfoAdvancedCreativeType  BrandCreativeGetV30DataCreativesCreativeAdvancedCreativeInfoAdvancedCreativeType = 1
	Enum_2_BrandCreativeGetV30DataCreativesCreativeAdvancedCreativeInfoAdvancedCreativeType  BrandCreativeGetV30DataCreativesCreativeAdvancedCreativeInfoAdvancedCreativeType = 2
	Enum_3_BrandCreativeGetV30DataCreativesCreativeAdvancedCreativeInfoAdvancedCreativeType  BrandCreativeGetV30DataCreativesCreativeAdvancedCreativeInfoAdvancedCreativeType = 3
	Enum_50_BrandCreativeGetV30DataCreativesCreativeAdvancedCreativeInfoAdvancedCreativeType BrandCreativeGetV30DataCreativesCreativeAdvancedCreativeInfoAdvancedCreativeType = 50
	Enum_8_BrandCreativeGetV30DataCreativesCreativeAdvancedCreativeInfoAdvancedCreativeType  BrandCreativeGetV30DataCreativesCreativeAdvancedCreativeInfoAdvancedCreativeType = 8
)

// All allowed values of BrandCreativeGetV30DataCreativesCreativeAdvancedCreativeInfoAdvancedCreativeType enum
var AllowedBrandCreativeGetV30DataCreativesCreativeAdvancedCreativeInfoAdvancedCreativeTypeEnumValues = []BrandCreativeGetV30DataCreativesCreativeAdvancedCreativeInfoAdvancedCreativeType{
	0,
	1,
	2,
	3,
	50,
	8,
}

// NewBrandCreativeGetV30DataCreativesCreativeAdvancedCreativeInfoAdvancedCreativeTypeFromValue returns a pointer to a valid BrandCreativeGetV30DataCreativesCreativeAdvancedCreativeInfoAdvancedCreativeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandCreativeGetV30DataCreativesCreativeAdvancedCreativeInfoAdvancedCreativeTypeFromValue(v int64) (*BrandCreativeGetV30DataCreativesCreativeAdvancedCreativeInfoAdvancedCreativeType, error) {
	ev := BrandCreativeGetV30DataCreativesCreativeAdvancedCreativeInfoAdvancedCreativeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandCreativeGetV30DataCreativesCreativeAdvancedCreativeInfoAdvancedCreativeType: valid values are %v", v, AllowedBrandCreativeGetV30DataCreativesCreativeAdvancedCreativeInfoAdvancedCreativeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandCreativeGetV30DataCreativesCreativeAdvancedCreativeInfoAdvancedCreativeType) IsValid() bool {
	for _, existing := range AllowedBrandCreativeGetV30DataCreativesCreativeAdvancedCreativeInfoAdvancedCreativeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_creative_get_v3.0_data_creatives_creative_advanced_creative_info_advanced_creative_type value
func (v BrandCreativeGetV30DataCreativesCreativeAdvancedCreativeInfoAdvancedCreativeType) Ptr() *BrandCreativeGetV30DataCreativesCreativeAdvancedCreativeInfoAdvancedCreativeType {
	return &v
}
