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

// QianchuanUniPromotionAdStatusUpdateV10OptStatus
type QianchuanUniPromotionAdStatusUpdateV10OptStatus string

// List of qianchuan_uni_promotion_ad_status_update_v1.0_opt_status
const (
	DISABLE_QianchuanUniPromotionAdStatusUpdateV10OptStatus QianchuanUniPromotionAdStatusUpdateV10OptStatus = "DISABLE"
	ENABLE_QianchuanUniPromotionAdStatusUpdateV10OptStatus  QianchuanUniPromotionAdStatusUpdateV10OptStatus = "ENABLE"
)

// All allowed values of QianchuanUniPromotionAdStatusUpdateV10OptStatus enum
var AllowedQianchuanUniPromotionAdStatusUpdateV10OptStatusEnumValues = []QianchuanUniPromotionAdStatusUpdateV10OptStatus{
	"DISABLE",
	"ENABLE",
}

// NewQianchuanUniPromotionAdStatusUpdateV10OptStatusFromValue returns a pointer to a valid QianchuanUniPromotionAdStatusUpdateV10OptStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanUniPromotionAdStatusUpdateV10OptStatusFromValue(v string) (*QianchuanUniPromotionAdStatusUpdateV10OptStatus, error) {
	ev := QianchuanUniPromotionAdStatusUpdateV10OptStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanUniPromotionAdStatusUpdateV10OptStatus: valid values are %v", v, AllowedQianchuanUniPromotionAdStatusUpdateV10OptStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanUniPromotionAdStatusUpdateV10OptStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanUniPromotionAdStatusUpdateV10OptStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_uni_promotion_ad_status_update_v1.0_opt_status value
func (v QianchuanUniPromotionAdStatusUpdateV10OptStatus) Ptr() *QianchuanUniPromotionAdStatusUpdateV10OptStatus {
	return &v
}
