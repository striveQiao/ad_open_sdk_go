/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueCouponDetailV2DataCouponNeedSmsVerify
type ClueCouponDetailV2DataCouponNeedSmsVerify string

// List of clue_coupon_detail_v2_data_coupon_need_sms_verify
const (
	Enum_0_ClueCouponDetailV2DataCouponNeedSmsVerify ClueCouponDetailV2DataCouponNeedSmsVerify = "0"
	Enum_1_ClueCouponDetailV2DataCouponNeedSmsVerify ClueCouponDetailV2DataCouponNeedSmsVerify = "1"
)

// All allowed values of ClueCouponDetailV2DataCouponNeedSmsVerify enum
var AllowedClueCouponDetailV2DataCouponNeedSmsVerifyEnumValues = []ClueCouponDetailV2DataCouponNeedSmsVerify{
	"0",
	"1",
}

// NewClueCouponDetailV2DataCouponNeedSmsVerifyFromValue returns a pointer to a valid ClueCouponDetailV2DataCouponNeedSmsVerify
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponDetailV2DataCouponNeedSmsVerifyFromValue(v string) (*ClueCouponDetailV2DataCouponNeedSmsVerify, error) {
	ev := ClueCouponDetailV2DataCouponNeedSmsVerify(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponDetailV2DataCouponNeedSmsVerify: valid values are %v", v, AllowedClueCouponDetailV2DataCouponNeedSmsVerifyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponDetailV2DataCouponNeedSmsVerify) IsValid() bool {
	for _, existing := range AllowedClueCouponDetailV2DataCouponNeedSmsVerifyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_detail_v2_data_coupon_need_sms_verify value
func (v ClueCouponDetailV2DataCouponNeedSmsVerify) Ptr() *ClueCouponDetailV2DataCouponNeedSmsVerify {
	return &v
}
