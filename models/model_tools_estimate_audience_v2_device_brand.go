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

// ToolsEstimateAudienceV2DeviceBrand
type ToolsEstimateAudienceV2DeviceBrand string

// List of tools_estimate_audience_v2_device_brand
const (
	HISENSE_ToolsEstimateAudienceV2DeviceBrand     ToolsEstimateAudienceV2DeviceBrand = "HISENSE"
	PEPPER_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "PEPPER"
	LG_ToolsEstimateAudienceV2DeviceBrand          ToolsEstimateAudienceV2DeviceBrand = "LG"
	MEIZU_ToolsEstimateAudienceV2DeviceBrand       ToolsEstimateAudienceV2DeviceBrand = "MEIZU"
	QIKU_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "QIKU"
	ZTE_ToolsEstimateAudienceV2DeviceBrand         ToolsEstimateAudienceV2DeviceBrand = "ZTE"
	SAMSUNG_ToolsEstimateAudienceV2DeviceBrand     ToolsEstimateAudienceV2DeviceBrand = "SAMSUNG"
	HTC_ToolsEstimateAudienceV2DeviceBrand         ToolsEstimateAudienceV2DeviceBrand = "HTC"
	XIAOMI_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "XIAOMI"
	CHINAMOBILE_ToolsEstimateAudienceV2DeviceBrand ToolsEstimateAudienceV2DeviceBrand = "CHINAMOBILE"
	Enum_360_ToolsEstimateAudienceV2DeviceBrand    ToolsEstimateAudienceV2DeviceBrand = "360"
	VIVO_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "VIVO"
	NUBIA_ToolsEstimateAudienceV2DeviceBrand       ToolsEstimateAudienceV2DeviceBrand = "NUBIA"
	SMARTISAN_ToolsEstimateAudienceV2DeviceBrand   ToolsEstimateAudienceV2DeviceBrand = "SMARTISAN"
	GIONEE_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "GIONEE"
	MOTO_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "MOTO"
	NOKIA_ToolsEstimateAudienceV2DeviceBrand       ToolsEstimateAudienceV2DeviceBrand = "NOKIA"
	OPPO_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "OPPO"
	LETV_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "LETV"
	HUAWEI_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "HUAWEI"
	LENOVO_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "LENOVO"
	ONEPLUS_ToolsEstimateAudienceV2DeviceBrand     ToolsEstimateAudienceV2DeviceBrand = "ONEPLUS"
	HONOR_ToolsEstimateAudienceV2DeviceBrand       ToolsEstimateAudienceV2DeviceBrand = "HONOR"
	GOOGLE_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "GOOGLE"
	TCL_ToolsEstimateAudienceV2DeviceBrand         ToolsEstimateAudienceV2DeviceBrand = "TCL"
	SONY_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "SONY"
	COOLPAD_ToolsEstimateAudienceV2DeviceBrand     ToolsEstimateAudienceV2DeviceBrand = "COOLPAD"
	APPLE_ToolsEstimateAudienceV2DeviceBrand       ToolsEstimateAudienceV2DeviceBrand = "APPLE"
)

// All allowed values of ToolsEstimateAudienceV2DeviceBrand enum
var AllowedToolsEstimateAudienceV2DeviceBrandEnumValues = []ToolsEstimateAudienceV2DeviceBrand{
	"HISENSE",
	"PEPPER",
	"LG",
	"MEIZU",
	"QIKU",
	"ZTE",
	"SAMSUNG",
	"HTC",
	"XIAOMI",
	"CHINAMOBILE",
	"360",
	"VIVO",
	"NUBIA",
	"SMARTISAN",
	"GIONEE",
	"MOTO",
	"NOKIA",
	"OPPO",
	"LETV",
	"HUAWEI",
	"LENOVO",
	"ONEPLUS",
	"HONOR",
	"GOOGLE",
	"TCL",
	"SONY",
	"COOLPAD",
	"APPLE",
}

// NewToolsEstimateAudienceV2DeviceBrandFromValue returns a pointer to a valid ToolsEstimateAudienceV2DeviceBrand
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEstimateAudienceV2DeviceBrandFromValue(v string) (*ToolsEstimateAudienceV2DeviceBrand, error) {
	ev := ToolsEstimateAudienceV2DeviceBrand(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEstimateAudienceV2DeviceBrand: valid values are %v", v, AllowedToolsEstimateAudienceV2DeviceBrandEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEstimateAudienceV2DeviceBrand) IsValid() bool {
	for _, existing := range AllowedToolsEstimateAudienceV2DeviceBrandEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_estimate_audience_v2_device_brand value
func (v ToolsEstimateAudienceV2DeviceBrand) Ptr() *ToolsEstimateAudienceV2DeviceBrand {
	return &v
}
