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

// AdGetV2DataPromotionType
type AdGetV2DataPromotionType string

// List of ad_get_v2_data_promotion_type
const (
	LIVE_AdGetV2DataPromotionType              AdGetV2DataPromotionType = "LIVE"
	AWEME_HOME_PAGE_AdGetV2DataPromotionType   AdGetV2DataPromotionType = "AWEME_HOME_PAGE"
	GOODS_AdGetV2DataPromotionType             AdGetV2DataPromotionType = "GOODS"
	LANDING_PAGE_LINK_AdGetV2DataPromotionType AdGetV2DataPromotionType = "LANDING_PAGE_LINK"
)

// All allowed values of AdGetV2DataPromotionType enum
var AllowedAdGetV2DataPromotionTypeEnumValues = []AdGetV2DataPromotionType{
	"LIVE",
	"AWEME_HOME_PAGE",
	"GOODS",
	"LANDING_PAGE_LINK",
}

// NewAdGetV2DataPromotionTypeFromValue returns a pointer to a valid AdGetV2DataPromotionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataPromotionTypeFromValue(v string) (*AdGetV2DataPromotionType, error) {
	ev := AdGetV2DataPromotionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataPromotionType: valid values are %v", v, AllowedAdGetV2DataPromotionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataPromotionType) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataPromotionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_promotion_type value
func (v AdGetV2DataPromotionType) Ptr() *AdGetV2DataPromotionType {
	return &v
}
