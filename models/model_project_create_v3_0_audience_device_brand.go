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

// ProjectCreateV30AudienceDeviceBrand
type ProjectCreateV30AudienceDeviceBrand string

// List of project_create_v3.0_audience_device_brand
const (
	Enum_360_ProjectCreateV30AudienceDeviceBrand    ProjectCreateV30AudienceDeviceBrand = "360"
	APPLE_ProjectCreateV30AudienceDeviceBrand       ProjectCreateV30AudienceDeviceBrand = "APPLE"
	CHINAMOBILE_ProjectCreateV30AudienceDeviceBrand ProjectCreateV30AudienceDeviceBrand = "CHINAMOBILE"
	COOLPAD_ProjectCreateV30AudienceDeviceBrand     ProjectCreateV30AudienceDeviceBrand = "COOLPAD"
	GIONEE_ProjectCreateV30AudienceDeviceBrand      ProjectCreateV30AudienceDeviceBrand = "GIONEE"
	GOOGLE_ProjectCreateV30AudienceDeviceBrand      ProjectCreateV30AudienceDeviceBrand = "GOOGLE"
	HISENSE_ProjectCreateV30AudienceDeviceBrand     ProjectCreateV30AudienceDeviceBrand = "HISENSE"
	HONOR_ProjectCreateV30AudienceDeviceBrand       ProjectCreateV30AudienceDeviceBrand = "HONOR"
	HTC_ProjectCreateV30AudienceDeviceBrand         ProjectCreateV30AudienceDeviceBrand = "HTC"
	HUAWEI_ProjectCreateV30AudienceDeviceBrand      ProjectCreateV30AudienceDeviceBrand = "HUAWEI"
	LENOVO_ProjectCreateV30AudienceDeviceBrand      ProjectCreateV30AudienceDeviceBrand = "LENOVO"
	LETV_ProjectCreateV30AudienceDeviceBrand        ProjectCreateV30AudienceDeviceBrand = "LETV"
	LG_ProjectCreateV30AudienceDeviceBrand          ProjectCreateV30AudienceDeviceBrand = "LG"
	MEIZU_ProjectCreateV30AudienceDeviceBrand       ProjectCreateV30AudienceDeviceBrand = "MEIZU"
	MOTO_ProjectCreateV30AudienceDeviceBrand        ProjectCreateV30AudienceDeviceBrand = "MOTO"
	NOKIA_ProjectCreateV30AudienceDeviceBrand       ProjectCreateV30AudienceDeviceBrand = "NOKIA"
	NUBIA_ProjectCreateV30AudienceDeviceBrand       ProjectCreateV30AudienceDeviceBrand = "NUBIA"
	ONEPLUS_ProjectCreateV30AudienceDeviceBrand     ProjectCreateV30AudienceDeviceBrand = "ONEPLUS"
	OPPO_ProjectCreateV30AudienceDeviceBrand        ProjectCreateV30AudienceDeviceBrand = "OPPO"
	PEPPER_ProjectCreateV30AudienceDeviceBrand      ProjectCreateV30AudienceDeviceBrand = "PEPPER"
	QIKU_ProjectCreateV30AudienceDeviceBrand        ProjectCreateV30AudienceDeviceBrand = "QIKU"
	SAMSUNG_ProjectCreateV30AudienceDeviceBrand     ProjectCreateV30AudienceDeviceBrand = "SAMSUNG"
	SMARTISAN_ProjectCreateV30AudienceDeviceBrand   ProjectCreateV30AudienceDeviceBrand = "SMARTISAN"
	SONY_ProjectCreateV30AudienceDeviceBrand        ProjectCreateV30AudienceDeviceBrand = "SONY"
	TCL_ProjectCreateV30AudienceDeviceBrand         ProjectCreateV30AudienceDeviceBrand = "TCL"
	VIVO_ProjectCreateV30AudienceDeviceBrand        ProjectCreateV30AudienceDeviceBrand = "VIVO"
	XIAOMI_ProjectCreateV30AudienceDeviceBrand      ProjectCreateV30AudienceDeviceBrand = "XIAOMI"
	ZTE_ProjectCreateV30AudienceDeviceBrand         ProjectCreateV30AudienceDeviceBrand = "ZTE"
)

// All allowed values of ProjectCreateV30AudienceDeviceBrand enum
var AllowedProjectCreateV30AudienceDeviceBrandEnumValues = []ProjectCreateV30AudienceDeviceBrand{
	"360",
	"APPLE",
	"CHINAMOBILE",
	"COOLPAD",
	"GIONEE",
	"GOOGLE",
	"HISENSE",
	"HONOR",
	"HTC",
	"HUAWEI",
	"LENOVO",
	"LETV",
	"LG",
	"MEIZU",
	"MOTO",
	"NOKIA",
	"NUBIA",
	"ONEPLUS",
	"OPPO",
	"PEPPER",
	"QIKU",
	"SAMSUNG",
	"SMARTISAN",
	"SONY",
	"TCL",
	"VIVO",
	"XIAOMI",
	"ZTE",
}

// NewProjectCreateV30AudienceDeviceBrandFromValue returns a pointer to a valid ProjectCreateV30AudienceDeviceBrand
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AudienceDeviceBrandFromValue(v string) (*ProjectCreateV30AudienceDeviceBrand, error) {
	ev := ProjectCreateV30AudienceDeviceBrand(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AudienceDeviceBrand: valid values are %v", v, AllowedProjectCreateV30AudienceDeviceBrandEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AudienceDeviceBrand) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AudienceDeviceBrandEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_audience_device_brand value
func (v ProjectCreateV30AudienceDeviceBrand) Ptr() *ProjectCreateV30AudienceDeviceBrand {
	return &v
}
