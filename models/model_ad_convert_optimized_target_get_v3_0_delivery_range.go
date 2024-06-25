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

// AdConvertOptimizedTargetGetV30DeliveryRange
type AdConvertOptimizedTargetGetV30DeliveryRange string

// List of ad_convert_optimized_target_get_v3.0_delivery_range
const (
	DEFAULT_AdConvertOptimizedTargetGetV30DeliveryRange AdConvertOptimizedTargetGetV30DeliveryRange = "DEFAULT"
	UNION_AdConvertOptimizedTargetGetV30DeliveryRange   AdConvertOptimizedTargetGetV30DeliveryRange = "UNION"
)

// All allowed values of AdConvertOptimizedTargetGetV30DeliveryRange enum
var AllowedAdConvertOptimizedTargetGetV30DeliveryRangeEnumValues = []AdConvertOptimizedTargetGetV30DeliveryRange{
	"DEFAULT",
	"UNION",
}

// NewAdConvertOptimizedTargetGetV30DeliveryRangeFromValue returns a pointer to a valid AdConvertOptimizedTargetGetV30DeliveryRange
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdConvertOptimizedTargetGetV30DeliveryRangeFromValue(v string) (*AdConvertOptimizedTargetGetV30DeliveryRange, error) {
	ev := AdConvertOptimizedTargetGetV30DeliveryRange(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdConvertOptimizedTargetGetV30DeliveryRange: valid values are %v", v, AllowedAdConvertOptimizedTargetGetV30DeliveryRangeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdConvertOptimizedTargetGetV30DeliveryRange) IsValid() bool {
	for _, existing := range AllowedAdConvertOptimizedTargetGetV30DeliveryRangeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_convert_optimized_target_get_v3.0_delivery_range value
func (v AdConvertOptimizedTargetGetV30DeliveryRange) Ptr() *AdConvertOptimizedTargetGetV30DeliveryRange {
	return &v
}
