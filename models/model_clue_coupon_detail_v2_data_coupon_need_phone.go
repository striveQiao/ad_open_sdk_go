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

// ClueCouponDetailV2DataCouponNeedPhone
type ClueCouponDetailV2DataCouponNeedPhone string

// List of clue_coupon_detail_v2_data_coupon_need_phone
const (
	Enum_0_ClueCouponDetailV2DataCouponNeedPhone ClueCouponDetailV2DataCouponNeedPhone = "0"
	Enum_1_ClueCouponDetailV2DataCouponNeedPhone ClueCouponDetailV2DataCouponNeedPhone = "1"
)

// All allowed values of ClueCouponDetailV2DataCouponNeedPhone enum
var AllowedClueCouponDetailV2DataCouponNeedPhoneEnumValues = []ClueCouponDetailV2DataCouponNeedPhone{
	"0",
	"1",
}

// NewClueCouponDetailV2DataCouponNeedPhoneFromValue returns a pointer to a valid ClueCouponDetailV2DataCouponNeedPhone
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponDetailV2DataCouponNeedPhoneFromValue(v string) (*ClueCouponDetailV2DataCouponNeedPhone, error) {
	ev := ClueCouponDetailV2DataCouponNeedPhone(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponDetailV2DataCouponNeedPhone: valid values are %v", v, AllowedClueCouponDetailV2DataCouponNeedPhoneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponDetailV2DataCouponNeedPhone) IsValid() bool {
	for _, existing := range AllowedClueCouponDetailV2DataCouponNeedPhoneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_detail_v2_data_coupon_need_phone value
func (v ClueCouponDetailV2DataCouponNeedPhone) Ptr() *ClueCouponDetailV2DataCouponNeedPhone {
	return &v
}
