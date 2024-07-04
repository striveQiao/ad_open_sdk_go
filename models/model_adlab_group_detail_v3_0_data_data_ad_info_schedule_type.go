/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupDetailV30DataDataAdInfoScheduleType
type AdlabGroupDetailV30DataDataAdInfoScheduleType string

// List of adlab_group_detail_v3.0_data_data_ad_info_schedule_type
const (
	SCHEDULE_FROM_NOW_AdlabGroupDetailV30DataDataAdInfoScheduleType  AdlabGroupDetailV30DataDataAdInfoScheduleType = "SCHEDULE_FROM_NOW"
	SCHEDULE_START_END_AdlabGroupDetailV30DataDataAdInfoScheduleType AdlabGroupDetailV30DataDataAdInfoScheduleType = "SCHEDULE_START_END"
)

// All allowed values of AdlabGroupDetailV30DataDataAdInfoScheduleType enum
var AllowedAdlabGroupDetailV30DataDataAdInfoScheduleTypeEnumValues = []AdlabGroupDetailV30DataDataAdInfoScheduleType{
	"SCHEDULE_FROM_NOW",
	"SCHEDULE_START_END",
}

// NewAdlabGroupDetailV30DataDataAdInfoScheduleTypeFromValue returns a pointer to a valid AdlabGroupDetailV30DataDataAdInfoScheduleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupDetailV30DataDataAdInfoScheduleTypeFromValue(v string) (*AdlabGroupDetailV30DataDataAdInfoScheduleType, error) {
	ev := AdlabGroupDetailV30DataDataAdInfoScheduleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupDetailV30DataDataAdInfoScheduleType: valid values are %v", v, AllowedAdlabGroupDetailV30DataDataAdInfoScheduleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupDetailV30DataDataAdInfoScheduleType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupDetailV30DataDataAdInfoScheduleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_detail_v3.0_data_data_ad_info_schedule_type value
func (v AdlabGroupDetailV30DataDataAdInfoScheduleType) Ptr() *AdlabGroupDetailV30DataDataAdInfoScheduleType {
	return &v
}
