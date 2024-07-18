/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType
type BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType int64

// List of brand_creative_get_v3.0_data_creatives_creative_creative_info_creative_type
const (
	Enum_0_BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType  BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType = 0
	Enum_1_BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType  BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType = 1
	Enum_2_BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType  BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType = 2
	Enum_3_BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType  BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType = 3
	Enum_4_BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType  BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType = 4
	Enum_47_BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType = 47
	Enum_18_BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType = 18
	Enum_44_BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType = 44
	Enum_14_BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType = 14
	Enum_48_BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType = 48
	Enum_49_BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType = 49
	Enum_50_BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType = 50
	Enum_25_BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType = 25
)

// All allowed values of BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType enum
var AllowedBrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeTypeEnumValues = []BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType{
	0,
	1,
	2,
	3,
	4,
	47,
	18,
	44,
	14,
	48,
	49,
	50,
	25,
}

// NewBrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeTypeFromValue returns a pointer to a valid BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeTypeFromValue(v int64) (*BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType, error) {
	ev := BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType: valid values are %v", v, AllowedBrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType) IsValid() bool {
	for _, existing := range AllowedBrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_creative_get_v3.0_data_creatives_creative_creative_info_creative_type value
func (v BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType) Ptr() *BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType {
	return &v
}
