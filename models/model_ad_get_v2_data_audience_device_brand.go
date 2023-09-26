/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdGetV2DataAudienceDeviceBrand
type AdGetV2DataAudienceDeviceBrand string

// List of ad_get_v2_data_audience_device_brand
const (
	APPLE_AdGetV2DataAudienceDeviceBrand       AdGetV2DataAudienceDeviceBrand = "APPLE"
	PEPPER_AdGetV2DataAudienceDeviceBrand      AdGetV2DataAudienceDeviceBrand = "PEPPER"
	QIKU_AdGetV2DataAudienceDeviceBrand        AdGetV2DataAudienceDeviceBrand = "QIKU"
	NUBIA_AdGetV2DataAudienceDeviceBrand       AdGetV2DataAudienceDeviceBrand = "NUBIA"
	OPPO_AdGetV2DataAudienceDeviceBrand        AdGetV2DataAudienceDeviceBrand = "OPPO"
	GOOGLE_AdGetV2DataAudienceDeviceBrand      AdGetV2DataAudienceDeviceBrand = "GOOGLE"
	HISENSE_AdGetV2DataAudienceDeviceBrand     AdGetV2DataAudienceDeviceBrand = "HISENSE"
	COOLPAD_AdGetV2DataAudienceDeviceBrand     AdGetV2DataAudienceDeviceBrand = "COOLPAD"
	SAMSUNG_AdGetV2DataAudienceDeviceBrand     AdGetV2DataAudienceDeviceBrand = "SAMSUNG"
	SMARTISAN_AdGetV2DataAudienceDeviceBrand   AdGetV2DataAudienceDeviceBrand = "SMARTISAN"
	ONEPLUS_AdGetV2DataAudienceDeviceBrand     AdGetV2DataAudienceDeviceBrand = "ONEPLUS"
	ZTE_AdGetV2DataAudienceDeviceBrand         AdGetV2DataAudienceDeviceBrand = "ZTE"
	HTC_AdGetV2DataAudienceDeviceBrand         AdGetV2DataAudienceDeviceBrand = "HTC"
	LG_AdGetV2DataAudienceDeviceBrand          AdGetV2DataAudienceDeviceBrand = "LG"
	XIAOMI_AdGetV2DataAudienceDeviceBrand      AdGetV2DataAudienceDeviceBrand = "XIAOMI"
	MEIZU_AdGetV2DataAudienceDeviceBrand       AdGetV2DataAudienceDeviceBrand = "MEIZU"
	HONOR_AdGetV2DataAudienceDeviceBrand       AdGetV2DataAudienceDeviceBrand = "HONOR"
	LETV_AdGetV2DataAudienceDeviceBrand        AdGetV2DataAudienceDeviceBrand = "LETV"
	TCL_AdGetV2DataAudienceDeviceBrand         AdGetV2DataAudienceDeviceBrand = "TCL"
	HUAWEI_AdGetV2DataAudienceDeviceBrand      AdGetV2DataAudienceDeviceBrand = "HUAWEI"
	SONY_AdGetV2DataAudienceDeviceBrand        AdGetV2DataAudienceDeviceBrand = "SONY"
	Enum_360_AdGetV2DataAudienceDeviceBrand    AdGetV2DataAudienceDeviceBrand = "360"
	CHINAMOBILE_AdGetV2DataAudienceDeviceBrand AdGetV2DataAudienceDeviceBrand = "CHINAMOBILE"
	GIONEE_AdGetV2DataAudienceDeviceBrand      AdGetV2DataAudienceDeviceBrand = "GIONEE"
	VIVO_AdGetV2DataAudienceDeviceBrand        AdGetV2DataAudienceDeviceBrand = "VIVO"
	NOKIA_AdGetV2DataAudienceDeviceBrand       AdGetV2DataAudienceDeviceBrand = "NOKIA"
	MOTO_AdGetV2DataAudienceDeviceBrand        AdGetV2DataAudienceDeviceBrand = "MOTO"
	LENOVO_AdGetV2DataAudienceDeviceBrand      AdGetV2DataAudienceDeviceBrand = "LENOVO"
)

// All allowed values of AdGetV2DataAudienceDeviceBrand enum
var AllowedAdGetV2DataAudienceDeviceBrandEnumValues = []AdGetV2DataAudienceDeviceBrand{
	"APPLE",
	"PEPPER",
	"QIKU",
	"NUBIA",
	"OPPO",
	"GOOGLE",
	"HISENSE",
	"COOLPAD",
	"SAMSUNG",
	"SMARTISAN",
	"ONEPLUS",
	"ZTE",
	"HTC",
	"LG",
	"XIAOMI",
	"MEIZU",
	"HONOR",
	"LETV",
	"TCL",
	"HUAWEI",
	"SONY",
	"360",
	"CHINAMOBILE",
	"GIONEE",
	"VIVO",
	"NOKIA",
	"MOTO",
	"LENOVO",
}

// NewAdGetV2DataAudienceDeviceBrandFromValue returns a pointer to a valid AdGetV2DataAudienceDeviceBrand
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataAudienceDeviceBrandFromValue(v string) (*AdGetV2DataAudienceDeviceBrand, error) {
	ev := AdGetV2DataAudienceDeviceBrand(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataAudienceDeviceBrand: valid values are %v", v, AllowedAdGetV2DataAudienceDeviceBrandEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataAudienceDeviceBrand) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataAudienceDeviceBrandEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_audience_device_brand value
func (v AdGetV2DataAudienceDeviceBrand) Ptr() *AdGetV2DataAudienceDeviceBrand {
	return &v
}
