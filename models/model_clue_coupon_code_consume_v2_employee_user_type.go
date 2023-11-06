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

// ClueCouponCodeConsumeV2EmployeeUserType
type ClueCouponCodeConsumeV2EmployeeUserType string

// List of clue_coupon_code_consume_v2_employee_user_type
const (
	TOUTIAO_ClueCouponCodeConsumeV2EmployeeUserType ClueCouponCodeConsumeV2EmployeeUserType = "TOUTIAO"
	DOUYIN_ClueCouponCodeConsumeV2EmployeeUserType  ClueCouponCodeConsumeV2EmployeeUserType = "DOUYIN"
)

// All allowed values of ClueCouponCodeConsumeV2EmployeeUserType enum
var AllowedClueCouponCodeConsumeV2EmployeeUserTypeEnumValues = []ClueCouponCodeConsumeV2EmployeeUserType{
	"TOUTIAO",
	"DOUYIN",
}

// NewClueCouponCodeConsumeV2EmployeeUserTypeFromValue returns a pointer to a valid ClueCouponCodeConsumeV2EmployeeUserType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponCodeConsumeV2EmployeeUserTypeFromValue(v string) (*ClueCouponCodeConsumeV2EmployeeUserType, error) {
	ev := ClueCouponCodeConsumeV2EmployeeUserType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponCodeConsumeV2EmployeeUserType: valid values are %v", v, AllowedClueCouponCodeConsumeV2EmployeeUserTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponCodeConsumeV2EmployeeUserType) IsValid() bool {
	for _, existing := range AllowedClueCouponCodeConsumeV2EmployeeUserTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_code_consume_v2_employee_user_type value
func (v ClueCouponCodeConsumeV2EmployeeUserType) Ptr() *ClueCouponCodeConsumeV2EmployeeUserType {
	return &v
}
