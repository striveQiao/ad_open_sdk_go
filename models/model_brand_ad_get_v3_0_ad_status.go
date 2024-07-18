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

// BrandAdGetV30AdStatus
type BrandAdGetV30AdStatus int64

// List of brand_ad_get_v3.0_ad_status
const (
	Enum_1_BrandAdGetV30AdStatus  BrandAdGetV30AdStatus = 1
	Enum_10_BrandAdGetV30AdStatus BrandAdGetV30AdStatus = 10
	Enum_11_BrandAdGetV30AdStatus BrandAdGetV30AdStatus = 11
	Enum_12_BrandAdGetV30AdStatus BrandAdGetV30AdStatus = 12
	Enum_2_BrandAdGetV30AdStatus  BrandAdGetV30AdStatus = 2
	Enum_3_BrandAdGetV30AdStatus  BrandAdGetV30AdStatus = 3
	Enum_4_BrandAdGetV30AdStatus  BrandAdGetV30AdStatus = 4
	Enum_5_BrandAdGetV30AdStatus  BrandAdGetV30AdStatus = 5
	Enum_6_BrandAdGetV30AdStatus  BrandAdGetV30AdStatus = 6
	Enum_7_BrandAdGetV30AdStatus  BrandAdGetV30AdStatus = 7
	Enum_8_BrandAdGetV30AdStatus  BrandAdGetV30AdStatus = 8
	Enum_9_BrandAdGetV30AdStatus  BrandAdGetV30AdStatus = 9
)

// All allowed values of BrandAdGetV30AdStatus enum
var AllowedBrandAdGetV30AdStatusEnumValues = []BrandAdGetV30AdStatus{
	1,
	10,
	11,
	12,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
}

// NewBrandAdGetV30AdStatusFromValue returns a pointer to a valid BrandAdGetV30AdStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdGetV30AdStatusFromValue(v int64) (*BrandAdGetV30AdStatus, error) {
	ev := BrandAdGetV30AdStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdGetV30AdStatus: valid values are %v", v, AllowedBrandAdGetV30AdStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdGetV30AdStatus) IsValid() bool {
	for _, existing := range AllowedBrandAdGetV30AdStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_get_v3.0_ad_status value
func (v BrandAdGetV30AdStatus) Ptr() *BrandAdGetV30AdStatus {
	return &v
}
