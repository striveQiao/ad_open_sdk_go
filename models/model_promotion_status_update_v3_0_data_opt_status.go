/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionStatusUpdateV30DataOptStatus
type PromotionStatusUpdateV30DataOptStatus string

// List of promotion_status_update_v3.0_data_opt_status
const (
	DISABLE_PromotionStatusUpdateV30DataOptStatus PromotionStatusUpdateV30DataOptStatus = "DISABLE"
	ENABLE_PromotionStatusUpdateV30DataOptStatus  PromotionStatusUpdateV30DataOptStatus = "ENABLE"
)

// All allowed values of PromotionStatusUpdateV30DataOptStatus enum
var AllowedPromotionStatusUpdateV30DataOptStatusEnumValues = []PromotionStatusUpdateV30DataOptStatus{
	"DISABLE",
	"ENABLE",
}

// NewPromotionStatusUpdateV30DataOptStatusFromValue returns a pointer to a valid PromotionStatusUpdateV30DataOptStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionStatusUpdateV30DataOptStatusFromValue(v string) (*PromotionStatusUpdateV30DataOptStatus, error) {
	ev := PromotionStatusUpdateV30DataOptStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionStatusUpdateV30DataOptStatus: valid values are %v", v, AllowedPromotionStatusUpdateV30DataOptStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionStatusUpdateV30DataOptStatus) IsValid() bool {
	for _, existing := range AllowedPromotionStatusUpdateV30DataOptStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_status_update_v3.0_data_opt_status value
func (v PromotionStatusUpdateV30DataOptStatus) Ptr() *PromotionStatusUpdateV30DataOptStatus {
	return &v
}
