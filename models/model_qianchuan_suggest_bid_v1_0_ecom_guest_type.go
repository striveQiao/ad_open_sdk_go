/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanSuggestBidV10EcomGuestType
type QianchuanSuggestBidV10EcomGuestType string

// List of qianchuan_suggest_bid_v1.0_ecom_guest_type
const (
	NO_BUY_QianchuanSuggestBidV10EcomGuestType        QianchuanSuggestBidV10EcomGuestType = "NO_BUY"
	NO_BUY_BRAND_QianchuanSuggestBidV10EcomGuestType  QianchuanSuggestBidV10EcomGuestType = "NO_BUY_BRAND"
	NO_BUY_DOUYIN_QianchuanSuggestBidV10EcomGuestType QianchuanSuggestBidV10EcomGuestType = "NO_BUY_DOUYIN"
)

// All allowed values of QianchuanSuggestBidV10EcomGuestType enum
var AllowedQianchuanSuggestBidV10EcomGuestTypeEnumValues = []QianchuanSuggestBidV10EcomGuestType{
	"NO_BUY",
	"NO_BUY_BRAND",
	"NO_BUY_DOUYIN",
}

// NewQianchuanSuggestBidV10EcomGuestTypeFromValue returns a pointer to a valid QianchuanSuggestBidV10EcomGuestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanSuggestBidV10EcomGuestTypeFromValue(v string) (*QianchuanSuggestBidV10EcomGuestType, error) {
	ev := QianchuanSuggestBidV10EcomGuestType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanSuggestBidV10EcomGuestType: valid values are %v", v, AllowedQianchuanSuggestBidV10EcomGuestTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanSuggestBidV10EcomGuestType) IsValid() bool {
	for _, existing := range AllowedQianchuanSuggestBidV10EcomGuestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_suggest_bid_v1.0_ecom_guest_type value
func (v QianchuanSuggestBidV10EcomGuestType) Ptr() *QianchuanSuggestBidV10EcomGuestType {
	return &v
}
