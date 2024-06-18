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

// QianchuanEstimateEffectV10LiveScheduleType
type QianchuanEstimateEffectV10LiveScheduleType string

// List of qianchuan_estimate_effect_v1.0_live_schedule_type
const (
	SCHEDULE_FROM_NOW_QianchuanEstimateEffectV10LiveScheduleType        QianchuanEstimateEffectV10LiveScheduleType = "SCHEDULE_FROM_NOW"
	SCHEDULE_START_END_QianchuanEstimateEffectV10LiveScheduleType       QianchuanEstimateEffectV10LiveScheduleType = "SCHEDULE_START_END"
	SCHEDULE_TIME_FIXEDRANGE_QianchuanEstimateEffectV10LiveScheduleType QianchuanEstimateEffectV10LiveScheduleType = "SCHEDULE_TIME_FIXEDRANGE"
)

// All allowed values of QianchuanEstimateEffectV10LiveScheduleType enum
var AllowedQianchuanEstimateEffectV10LiveScheduleTypeEnumValues = []QianchuanEstimateEffectV10LiveScheduleType{
	"SCHEDULE_FROM_NOW",
	"SCHEDULE_START_END",
	"SCHEDULE_TIME_FIXEDRANGE",
}

// NewQianchuanEstimateEffectV10LiveScheduleTypeFromValue returns a pointer to a valid QianchuanEstimateEffectV10LiveScheduleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanEstimateEffectV10LiveScheduleTypeFromValue(v string) (*QianchuanEstimateEffectV10LiveScheduleType, error) {
	ev := QianchuanEstimateEffectV10LiveScheduleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanEstimateEffectV10LiveScheduleType: valid values are %v", v, AllowedQianchuanEstimateEffectV10LiveScheduleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanEstimateEffectV10LiveScheduleType) IsValid() bool {
	for _, existing := range AllowedQianchuanEstimateEffectV10LiveScheduleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_estimate_effect_v1.0_live_schedule_type value
func (v QianchuanEstimateEffectV10LiveScheduleType) Ptr() *QianchuanEstimateEffectV10LiveScheduleType {
	return &v
}
