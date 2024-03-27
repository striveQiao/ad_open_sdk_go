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

// ClueCouponDetailV2DataActivityType
type ClueCouponDetailV2DataActivityType string

// List of clue_coupon_detail_v2_data_activity_type
const (
	DIRECT_NOT_NEED_PHONE_ClueCouponDetailV2DataActivityType ClueCouponDetailV2DataActivityType = "DIRECT_NOT_NEED_PHONE"
	DIRECT_NEED_PHONE_ClueCouponDetailV2DataActivityType     ClueCouponDetailV2DataActivityType = "DIRECT_NEED_PHONE"
)

// All allowed values of ClueCouponDetailV2DataActivityType enum
var AllowedClueCouponDetailV2DataActivityTypeEnumValues = []ClueCouponDetailV2DataActivityType{
	"DIRECT_NOT_NEED_PHONE",
	"DIRECT_NEED_PHONE",
}

// NewClueCouponDetailV2DataActivityTypeFromValue returns a pointer to a valid ClueCouponDetailV2DataActivityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponDetailV2DataActivityTypeFromValue(v string) (*ClueCouponDetailV2DataActivityType, error) {
	ev := ClueCouponDetailV2DataActivityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponDetailV2DataActivityType: valid values are %v", v, AllowedClueCouponDetailV2DataActivityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponDetailV2DataActivityType) IsValid() bool {
	for _, existing := range AllowedClueCouponDetailV2DataActivityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_detail_v2_data_activity_type value
func (v ClueCouponDetailV2DataActivityType) Ptr() *ClueCouponDetailV2DataActivityType {
	return &v
}
