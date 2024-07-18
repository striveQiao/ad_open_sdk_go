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

// BrandAdGetV30DataAdsBudgetInfoGdSendType
type BrandAdGetV30DataAdsBudgetInfoGdSendType int64

// List of brand_ad_get_v3.0_data_ads_budget_info_gd_send_type
const (
	Enum_1_BrandAdGetV30DataAdsBudgetInfoGdSendType  BrandAdGetV30DataAdsBudgetInfoGdSendType = 1
	Enum_2_BrandAdGetV30DataAdsBudgetInfoGdSendType  BrandAdGetV30DataAdsBudgetInfoGdSendType = 2
	Enum_3_BrandAdGetV30DataAdsBudgetInfoGdSendType  BrandAdGetV30DataAdsBudgetInfoGdSendType = 3
	Enum_4_BrandAdGetV30DataAdsBudgetInfoGdSendType  BrandAdGetV30DataAdsBudgetInfoGdSendType = 4
	Enum_5_BrandAdGetV30DataAdsBudgetInfoGdSendType  BrandAdGetV30DataAdsBudgetInfoGdSendType = 5
	Enum_6_BrandAdGetV30DataAdsBudgetInfoGdSendType  BrandAdGetV30DataAdsBudgetInfoGdSendType = 6
	Enum_7_BrandAdGetV30DataAdsBudgetInfoGdSendType  BrandAdGetV30DataAdsBudgetInfoGdSendType = 7
	Enum_8_BrandAdGetV30DataAdsBudgetInfoGdSendType  BrandAdGetV30DataAdsBudgetInfoGdSendType = 8
	Enum_9_BrandAdGetV30DataAdsBudgetInfoGdSendType  BrandAdGetV30DataAdsBudgetInfoGdSendType = 9
	Enum_10_BrandAdGetV30DataAdsBudgetInfoGdSendType BrandAdGetV30DataAdsBudgetInfoGdSendType = 10
	Enum_11_BrandAdGetV30DataAdsBudgetInfoGdSendType BrandAdGetV30DataAdsBudgetInfoGdSendType = 11
	Enum_12_BrandAdGetV30DataAdsBudgetInfoGdSendType BrandAdGetV30DataAdsBudgetInfoGdSendType = 12
	Enum_13_BrandAdGetV30DataAdsBudgetInfoGdSendType BrandAdGetV30DataAdsBudgetInfoGdSendType = 13
	Enum_14_BrandAdGetV30DataAdsBudgetInfoGdSendType BrandAdGetV30DataAdsBudgetInfoGdSendType = 14
	Enum_15_BrandAdGetV30DataAdsBudgetInfoGdSendType BrandAdGetV30DataAdsBudgetInfoGdSendType = 15
	Enum_16_BrandAdGetV30DataAdsBudgetInfoGdSendType BrandAdGetV30DataAdsBudgetInfoGdSendType = 16
	Enum_17_BrandAdGetV30DataAdsBudgetInfoGdSendType BrandAdGetV30DataAdsBudgetInfoGdSendType = 17
	Enum_18_BrandAdGetV30DataAdsBudgetInfoGdSendType BrandAdGetV30DataAdsBudgetInfoGdSendType = 18
	Enum_19_BrandAdGetV30DataAdsBudgetInfoGdSendType BrandAdGetV30DataAdsBudgetInfoGdSendType = 19
	Enum_20_BrandAdGetV30DataAdsBudgetInfoGdSendType BrandAdGetV30DataAdsBudgetInfoGdSendType = 20
	Enum_21_BrandAdGetV30DataAdsBudgetInfoGdSendType BrandAdGetV30DataAdsBudgetInfoGdSendType = 21
	Enum_22_BrandAdGetV30DataAdsBudgetInfoGdSendType BrandAdGetV30DataAdsBudgetInfoGdSendType = 22
	Enum_23_BrandAdGetV30DataAdsBudgetInfoGdSendType BrandAdGetV30DataAdsBudgetInfoGdSendType = 23
	Enum_24_BrandAdGetV30DataAdsBudgetInfoGdSendType BrandAdGetV30DataAdsBudgetInfoGdSendType = 24
	Enum_25_BrandAdGetV30DataAdsBudgetInfoGdSendType BrandAdGetV30DataAdsBudgetInfoGdSendType = 25
	Enum_26_BrandAdGetV30DataAdsBudgetInfoGdSendType BrandAdGetV30DataAdsBudgetInfoGdSendType = 26
	Enum_27_BrandAdGetV30DataAdsBudgetInfoGdSendType BrandAdGetV30DataAdsBudgetInfoGdSendType = 27
	Enum_28_BrandAdGetV30DataAdsBudgetInfoGdSendType BrandAdGetV30DataAdsBudgetInfoGdSendType = 28
	Enum_29_BrandAdGetV30DataAdsBudgetInfoGdSendType BrandAdGetV30DataAdsBudgetInfoGdSendType = 29
	Enum_40_BrandAdGetV30DataAdsBudgetInfoGdSendType BrandAdGetV30DataAdsBudgetInfoGdSendType = 40
)

// All allowed values of BrandAdGetV30DataAdsBudgetInfoGdSendType enum
var AllowedBrandAdGetV30DataAdsBudgetInfoGdSendTypeEnumValues = []BrandAdGetV30DataAdsBudgetInfoGdSendType{
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
	10,
	11,
	12,
	13,
	14,
	15,
	16,
	17,
	18,
	19,
	20,
	21,
	22,
	23,
	24,
	25,
	26,
	27,
	28,
	29,
	40,
}

// NewBrandAdGetV30DataAdsBudgetInfoGdSendTypeFromValue returns a pointer to a valid BrandAdGetV30DataAdsBudgetInfoGdSendType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdGetV30DataAdsBudgetInfoGdSendTypeFromValue(v int64) (*BrandAdGetV30DataAdsBudgetInfoGdSendType, error) {
	ev := BrandAdGetV30DataAdsBudgetInfoGdSendType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdGetV30DataAdsBudgetInfoGdSendType: valid values are %v", v, AllowedBrandAdGetV30DataAdsBudgetInfoGdSendTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdGetV30DataAdsBudgetInfoGdSendType) IsValid() bool {
	for _, existing := range AllowedBrandAdGetV30DataAdsBudgetInfoGdSendTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_get_v3.0_data_ads_budget_info_gd_send_type value
func (v BrandAdGetV30DataAdsBudgetInfoGdSendType) Ptr() *BrandAdGetV30DataAdsBudgetInfoGdSendType {
	return &v
}
