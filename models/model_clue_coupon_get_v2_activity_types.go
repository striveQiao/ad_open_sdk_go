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

// ClueCouponGetV2ActivityTypes
type ClueCouponGetV2ActivityTypes string

// List of clue_coupon_get_v2_activity_types
const (
	DIRECT_NEED_PHONE_ClueCouponGetV2ActivityTypes     ClueCouponGetV2ActivityTypes = "DIRECT_NEED_PHONE"
	DIRECT_NOT_NEED_PHONE_ClueCouponGetV2ActivityTypes ClueCouponGetV2ActivityTypes = "DIRECT_NOT_NEED_PHONE"
)

// All allowed values of ClueCouponGetV2ActivityTypes enum
var AllowedClueCouponGetV2ActivityTypesEnumValues = []ClueCouponGetV2ActivityTypes{
	"DIRECT_NEED_PHONE",
	"DIRECT_NOT_NEED_PHONE",
}

// NewClueCouponGetV2ActivityTypesFromValue returns a pointer to a valid ClueCouponGetV2ActivityTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponGetV2ActivityTypesFromValue(v string) (*ClueCouponGetV2ActivityTypes, error) {
	ev := ClueCouponGetV2ActivityTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponGetV2ActivityTypes: valid values are %v", v, AllowedClueCouponGetV2ActivityTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponGetV2ActivityTypes) IsValid() bool {
	for _, existing := range AllowedClueCouponGetV2ActivityTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_get_v2_activity_types value
func (v ClueCouponGetV2ActivityTypes) Ptr() *ClueCouponGetV2ActivityTypes {
	return &v
}
