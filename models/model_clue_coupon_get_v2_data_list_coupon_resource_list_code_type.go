/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueCouponGetV2DataListCouponResourceListCodeType
type ClueCouponGetV2DataListCouponResourceListCodeType string

// List of clue_coupon_get_v2_data_list_coupon_resource_list_code_type
const (
	COMMON_ClueCouponGetV2DataListCouponResourceListCodeType   ClueCouponGetV2DataListCouponResourceListCodeType = "COMMON"
	API_ClueCouponGetV2DataListCouponResourceListCodeType      ClueCouponGetV2DataListCouponResourceListCodeType = "API"
	PLATFORM_ClueCouponGetV2DataListCouponResourceListCodeType ClueCouponGetV2DataListCouponResourceListCodeType = "PLATFORM"
	MERCHANT_ClueCouponGetV2DataListCouponResourceListCodeType ClueCouponGetV2DataListCouponResourceListCodeType = "MERCHANT"
)

// All allowed values of ClueCouponGetV2DataListCouponResourceListCodeType enum
var AllowedClueCouponGetV2DataListCouponResourceListCodeTypeEnumValues = []ClueCouponGetV2DataListCouponResourceListCodeType{
	"COMMON",
	"API",
	"PLATFORM",
	"MERCHANT",
}

// NewClueCouponGetV2DataListCouponResourceListCodeTypeFromValue returns a pointer to a valid ClueCouponGetV2DataListCouponResourceListCodeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponGetV2DataListCouponResourceListCodeTypeFromValue(v string) (*ClueCouponGetV2DataListCouponResourceListCodeType, error) {
	ev := ClueCouponGetV2DataListCouponResourceListCodeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponGetV2DataListCouponResourceListCodeType: valid values are %v", v, AllowedClueCouponGetV2DataListCouponResourceListCodeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponGetV2DataListCouponResourceListCodeType) IsValid() bool {
	for _, existing := range AllowedClueCouponGetV2DataListCouponResourceListCodeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_get_v2_data_list_coupon_resource_list_code_type value
func (v ClueCouponGetV2DataListCouponResourceListCodeType) Ptr() *ClueCouponGetV2DataListCouponResourceListCodeType {
	return &v
}
