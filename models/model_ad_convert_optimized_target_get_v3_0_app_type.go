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

// AdConvertOptimizedTargetGetV30AppType
type AdConvertOptimizedTargetGetV30AppType string

// List of ad_convert_optimized_target_get_v3.0_app_type
const (
	ANDROID_AdConvertOptimizedTargetGetV30AppType AdConvertOptimizedTargetGetV30AppType = "ANDROID"
	IOS_AdConvertOptimizedTargetGetV30AppType     AdConvertOptimizedTargetGetV30AppType = "IOS"
)

// All allowed values of AdConvertOptimizedTargetGetV30AppType enum
var AllowedAdConvertOptimizedTargetGetV30AppTypeEnumValues = []AdConvertOptimizedTargetGetV30AppType{
	"ANDROID",
	"IOS",
}

// NewAdConvertOptimizedTargetGetV30AppTypeFromValue returns a pointer to a valid AdConvertOptimizedTargetGetV30AppType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdConvertOptimizedTargetGetV30AppTypeFromValue(v string) (*AdConvertOptimizedTargetGetV30AppType, error) {
	ev := AdConvertOptimizedTargetGetV30AppType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdConvertOptimizedTargetGetV30AppType: valid values are %v", v, AllowedAdConvertOptimizedTargetGetV30AppTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdConvertOptimizedTargetGetV30AppType) IsValid() bool {
	for _, existing := range AllowedAdConvertOptimizedTargetGetV30AppTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_convert_optimized_target_get_v3.0_app_type value
func (v AdConvertOptimizedTargetGetV30AppType) Ptr() *AdConvertOptimizedTargetGetV30AppType {
	return &v
}
