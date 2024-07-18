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

// ClueCouponUpdateV2Status
type ClueCouponUpdateV2Status string

// List of clue_coupon_update_v2_status
const (
	AUDIT_DOING_ClueCouponUpdateV2Status ClueCouponUpdateV2Status = "AUDIT_DOING"
	DELETED_ClueCouponUpdateV2Status     ClueCouponUpdateV2Status = "DELETED"
	OFFLINE_ClueCouponUpdateV2Status     ClueCouponUpdateV2Status = "OFFLINE"
	NORMAL_ClueCouponUpdateV2Status      ClueCouponUpdateV2Status = "NORMAL"
	PAUSE_ClueCouponUpdateV2Status       ClueCouponUpdateV2Status = "PAUSE"
	UNAUDITED_ClueCouponUpdateV2Status   ClueCouponUpdateV2Status = "UNAUDITED"
	AUDIT_FAIL_ClueCouponUpdateV2Status  ClueCouponUpdateV2Status = "AUDIT_FAIL"
)

// All allowed values of ClueCouponUpdateV2Status enum
var AllowedClueCouponUpdateV2StatusEnumValues = []ClueCouponUpdateV2Status{
	"AUDIT_DOING",
	"DELETED",
	"OFFLINE",
	"NORMAL",
	"PAUSE",
	"UNAUDITED",
	"AUDIT_FAIL",
}

// NewClueCouponUpdateV2StatusFromValue returns a pointer to a valid ClueCouponUpdateV2Status
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponUpdateV2StatusFromValue(v string) (*ClueCouponUpdateV2Status, error) {
	ev := ClueCouponUpdateV2Status(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponUpdateV2Status: valid values are %v", v, AllowedClueCouponUpdateV2StatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponUpdateV2Status) IsValid() bool {
	for _, existing := range AllowedClueCouponUpdateV2StatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_update_v2_status value
func (v ClueCouponUpdateV2Status) Ptr() *ClueCouponUpdateV2Status {
	return &v
}
