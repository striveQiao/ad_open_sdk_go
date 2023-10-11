/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueCouponDetailV2DataFormValidateType
type ClueCouponDetailV2DataFormValidateType string

// List of clue_coupon_detail_v2_data_form_validate_type
const (
	VALIDITY_PRIORITY_ClueCouponDetailV2DataFormValidateType ClueCouponDetailV2DataFormValidateType = "VALIDITY_PRIORITY"
	ALL_VERIFICATION_ClueCouponDetailV2DataFormValidateType  ClueCouponDetailV2DataFormValidateType = "ALL_VERIFICATION"
	AUTO_VERIFICATION_ClueCouponDetailV2DataFormValidateType ClueCouponDetailV2DataFormValidateType = "AUTO_VERIFICATION"
	NONE_VERIFICATION_ClueCouponDetailV2DataFormValidateType ClueCouponDetailV2DataFormValidateType = "NONE_VERIFICATION"
	CLUE_PRIORITY_ClueCouponDetailV2DataFormValidateType     ClueCouponDetailV2DataFormValidateType = "CLUE_PRIORITY"
)

// All allowed values of ClueCouponDetailV2DataFormValidateType enum
var AllowedClueCouponDetailV2DataFormValidateTypeEnumValues = []ClueCouponDetailV2DataFormValidateType{
	"VALIDITY_PRIORITY",
	"ALL_VERIFICATION",
	"AUTO_VERIFICATION",
	"NONE_VERIFICATION",
	"CLUE_PRIORITY",
}

// NewClueCouponDetailV2DataFormValidateTypeFromValue returns a pointer to a valid ClueCouponDetailV2DataFormValidateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponDetailV2DataFormValidateTypeFromValue(v string) (*ClueCouponDetailV2DataFormValidateType, error) {
	ev := ClueCouponDetailV2DataFormValidateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponDetailV2DataFormValidateType: valid values are %v", v, AllowedClueCouponDetailV2DataFormValidateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponDetailV2DataFormValidateType) IsValid() bool {
	for _, existing := range AllowedClueCouponDetailV2DataFormValidateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_detail_v2_data_form_validate_type value
func (v ClueCouponDetailV2DataFormValidateType) Ptr() *ClueCouponDetailV2DataFormValidateType {
	return &v
}
