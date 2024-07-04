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

// ClueCouponCreateV2CouponNeedPhone
type ClueCouponCreateV2CouponNeedPhone string

// List of clue_coupon_create_v2_coupon_need_phone
const (
	Enum_0_ClueCouponCreateV2CouponNeedPhone ClueCouponCreateV2CouponNeedPhone = "0"
	Enum_1_ClueCouponCreateV2CouponNeedPhone ClueCouponCreateV2CouponNeedPhone = "1"
)

// All allowed values of ClueCouponCreateV2CouponNeedPhone enum
var AllowedClueCouponCreateV2CouponNeedPhoneEnumValues = []ClueCouponCreateV2CouponNeedPhone{
	"0",
	"1",
}

// NewClueCouponCreateV2CouponNeedPhoneFromValue returns a pointer to a valid ClueCouponCreateV2CouponNeedPhone
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponCreateV2CouponNeedPhoneFromValue(v string) (*ClueCouponCreateV2CouponNeedPhone, error) {
	ev := ClueCouponCreateV2CouponNeedPhone(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponCreateV2CouponNeedPhone: valid values are %v", v, AllowedClueCouponCreateV2CouponNeedPhoneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponCreateV2CouponNeedPhone) IsValid() bool {
	for _, existing := range AllowedClueCouponCreateV2CouponNeedPhoneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_create_v2_coupon_need_phone value
func (v ClueCouponCreateV2CouponNeedPhone) Ptr() *ClueCouponCreateV2CouponNeedPhone {
	return &v
}
