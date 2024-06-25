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

// QianchuanOrientationPackageGetV10DataListAutoExtendTargets
type QianchuanOrientationPackageGetV10DataListAutoExtendTargets string

// List of qianchuan_orientation_package_get_v1.0_data_list_auto_extend_targets
const (
	AGE_QianchuanOrientationPackageGetV10DataListAutoExtendTargets             QianchuanOrientationPackageGetV10DataListAutoExtendTargets = "AGE"
	REGION_QianchuanOrientationPackageGetV10DataListAutoExtendTargets          QianchuanOrientationPackageGetV10DataListAutoExtendTargets = "REGION"
	GENDER_QianchuanOrientationPackageGetV10DataListAutoExtendTargets          QianchuanOrientationPackageGetV10DataListAutoExtendTargets = "GENDER"
	CUSTOM_AUDIENCE_QianchuanOrientationPackageGetV10DataListAutoExtendTargets QianchuanOrientationPackageGetV10DataListAutoExtendTargets = "CUSTOM_AUDIENCE"
	INTEREST_ACTION_QianchuanOrientationPackageGetV10DataListAutoExtendTargets QianchuanOrientationPackageGetV10DataListAutoExtendTargets = "INTEREST_ACTION"
)

// All allowed values of QianchuanOrientationPackageGetV10DataListAutoExtendTargets enum
var AllowedQianchuanOrientationPackageGetV10DataListAutoExtendTargetsEnumValues = []QianchuanOrientationPackageGetV10DataListAutoExtendTargets{
	"AGE",
	"REGION",
	"GENDER",
	"CUSTOM_AUDIENCE",
	"INTEREST_ACTION",
}

// NewQianchuanOrientationPackageGetV10DataListAutoExtendTargetsFromValue returns a pointer to a valid QianchuanOrientationPackageGetV10DataListAutoExtendTargets
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanOrientationPackageGetV10DataListAutoExtendTargetsFromValue(v string) (*QianchuanOrientationPackageGetV10DataListAutoExtendTargets, error) {
	ev := QianchuanOrientationPackageGetV10DataListAutoExtendTargets(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanOrientationPackageGetV10DataListAutoExtendTargets: valid values are %v", v, AllowedQianchuanOrientationPackageGetV10DataListAutoExtendTargetsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanOrientationPackageGetV10DataListAutoExtendTargets) IsValid() bool {
	for _, existing := range AllowedQianchuanOrientationPackageGetV10DataListAutoExtendTargetsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_orientation_package_get_v1.0_data_list_auto_extend_targets value
func (v QianchuanOrientationPackageGetV10DataListAutoExtendTargets) Ptr() *QianchuanOrientationPackageGetV10DataListAutoExtendTargets {
	return &v
}
