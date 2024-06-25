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

// QianchuanOrientationPackageGetV10DataListAwemeFanBehaviorsDays
type QianchuanOrientationPackageGetV10DataListAwemeFanBehaviorsDays string

// List of qianchuan_orientation_package_get_v1.0_data_list_aweme_fan_behaviors_days
const (
	DAYS_15_QianchuanOrientationPackageGetV10DataListAwemeFanBehaviorsDays QianchuanOrientationPackageGetV10DataListAwemeFanBehaviorsDays = "DAYS_15"
	DAYS_30_QianchuanOrientationPackageGetV10DataListAwemeFanBehaviorsDays QianchuanOrientationPackageGetV10DataListAwemeFanBehaviorsDays = "DAYS_30"
	DAYS_60_QianchuanOrientationPackageGetV10DataListAwemeFanBehaviorsDays QianchuanOrientationPackageGetV10DataListAwemeFanBehaviorsDays = "DAYS_60"
)

// All allowed values of QianchuanOrientationPackageGetV10DataListAwemeFanBehaviorsDays enum
var AllowedQianchuanOrientationPackageGetV10DataListAwemeFanBehaviorsDaysEnumValues = []QianchuanOrientationPackageGetV10DataListAwemeFanBehaviorsDays{
	"DAYS_15",
	"DAYS_30",
	"DAYS_60",
}

// NewQianchuanOrientationPackageGetV10DataListAwemeFanBehaviorsDaysFromValue returns a pointer to a valid QianchuanOrientationPackageGetV10DataListAwemeFanBehaviorsDays
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanOrientationPackageGetV10DataListAwemeFanBehaviorsDaysFromValue(v string) (*QianchuanOrientationPackageGetV10DataListAwemeFanBehaviorsDays, error) {
	ev := QianchuanOrientationPackageGetV10DataListAwemeFanBehaviorsDays(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanOrientationPackageGetV10DataListAwemeFanBehaviorsDays: valid values are %v", v, AllowedQianchuanOrientationPackageGetV10DataListAwemeFanBehaviorsDaysEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanOrientationPackageGetV10DataListAwemeFanBehaviorsDays) IsValid() bool {
	for _, existing := range AllowedQianchuanOrientationPackageGetV10DataListAwemeFanBehaviorsDaysEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_orientation_package_get_v1.0_data_list_aweme_fan_behaviors_days value
func (v QianchuanOrientationPackageGetV10DataListAwemeFanBehaviorsDays) Ptr() *QianchuanOrientationPackageGetV10DataListAwemeFanBehaviorsDays {
	return &v
}
