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

// PromotionCreateV30AutoExtendTraffic
type PromotionCreateV30AutoExtendTraffic string

// List of promotion_create_v3.0_auto_extend_traffic
const (
	OFF_PromotionCreateV30AutoExtendTraffic PromotionCreateV30AutoExtendTraffic = "OFF"
	ON_PromotionCreateV30AutoExtendTraffic  PromotionCreateV30AutoExtendTraffic = "ON"
)

// All allowed values of PromotionCreateV30AutoExtendTraffic enum
var AllowedPromotionCreateV30AutoExtendTrafficEnumValues = []PromotionCreateV30AutoExtendTraffic{
	"OFF",
	"ON",
}

// NewPromotionCreateV30AutoExtendTrafficFromValue returns a pointer to a valid PromotionCreateV30AutoExtendTraffic
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCreateV30AutoExtendTrafficFromValue(v string) (*PromotionCreateV30AutoExtendTraffic, error) {
	ev := PromotionCreateV30AutoExtendTraffic(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCreateV30AutoExtendTraffic: valid values are %v", v, AllowedPromotionCreateV30AutoExtendTrafficEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCreateV30AutoExtendTraffic) IsValid() bool {
	for _, existing := range AllowedPromotionCreateV30AutoExtendTrafficEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_create_v3.0_auto_extend_traffic value
func (v PromotionCreateV30AutoExtendTraffic) Ptr() *PromotionCreateV30AutoExtendTraffic {
	return &v
}
