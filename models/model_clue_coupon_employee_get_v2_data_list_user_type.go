/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueCouponEmployeeGetV2DataListUserType
type ClueCouponEmployeeGetV2DataListUserType string

// List of clue_coupon_employee_get_v2_data_list_user_type
const (
	TOUTIAO_ClueCouponEmployeeGetV2DataListUserType ClueCouponEmployeeGetV2DataListUserType = "TOUTIAO"
	DOUYIN_ClueCouponEmployeeGetV2DataListUserType  ClueCouponEmployeeGetV2DataListUserType = "DOUYIN"
)

// All allowed values of ClueCouponEmployeeGetV2DataListUserType enum
var AllowedClueCouponEmployeeGetV2DataListUserTypeEnumValues = []ClueCouponEmployeeGetV2DataListUserType{
	"TOUTIAO",
	"DOUYIN",
}

// NewClueCouponEmployeeGetV2DataListUserTypeFromValue returns a pointer to a valid ClueCouponEmployeeGetV2DataListUserType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponEmployeeGetV2DataListUserTypeFromValue(v string) (*ClueCouponEmployeeGetV2DataListUserType, error) {
	ev := ClueCouponEmployeeGetV2DataListUserType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponEmployeeGetV2DataListUserType: valid values are %v", v, AllowedClueCouponEmployeeGetV2DataListUserTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponEmployeeGetV2DataListUserType) IsValid() bool {
	for _, existing := range AllowedClueCouponEmployeeGetV2DataListUserTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_employee_get_v2_data_list_user_type value
func (v ClueCouponEmployeeGetV2DataListUserType) Ptr() *ClueCouponEmployeeGetV2DataListUserType {
	return &v
}
