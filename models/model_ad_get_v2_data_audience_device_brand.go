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

// AdGetV2DataAudienceDeviceBrand
type AdGetV2DataAudienceDeviceBrand string

// List of ad_get_v2_data_audience_device_brand
const (
	HISENSE_AdGetV2DataAudienceDeviceBrand     AdGetV2DataAudienceDeviceBrand = "HISENSE"
	PEPPER_AdGetV2DataAudienceDeviceBrand      AdGetV2DataAudienceDeviceBrand = "PEPPER"
	LG_AdGetV2DataAudienceDeviceBrand          AdGetV2DataAudienceDeviceBrand = "LG"
	MEIZU_AdGetV2DataAudienceDeviceBrand       AdGetV2DataAudienceDeviceBrand = "MEIZU"
	QIKU_AdGetV2DataAudienceDeviceBrand        AdGetV2DataAudienceDeviceBrand = "QIKU"
	ZTE_AdGetV2DataAudienceDeviceBrand         AdGetV2DataAudienceDeviceBrand = "ZTE"
	SAMSUNG_AdGetV2DataAudienceDeviceBrand     AdGetV2DataAudienceDeviceBrand = "SAMSUNG"
	HTC_AdGetV2DataAudienceDeviceBrand         AdGetV2DataAudienceDeviceBrand = "HTC"
	XIAOMI_AdGetV2DataAudienceDeviceBrand      AdGetV2DataAudienceDeviceBrand = "XIAOMI"
	CHINAMOBILE_AdGetV2DataAudienceDeviceBrand AdGetV2DataAudienceDeviceBrand = "CHINAMOBILE"
	Enum_360_AdGetV2DataAudienceDeviceBrand    AdGetV2DataAudienceDeviceBrand = "360"
	VIVO_AdGetV2DataAudienceDeviceBrand        AdGetV2DataAudienceDeviceBrand = "VIVO"
	NUBIA_AdGetV2DataAudienceDeviceBrand       AdGetV2DataAudienceDeviceBrand = "NUBIA"
	SMARTISAN_AdGetV2DataAudienceDeviceBrand   AdGetV2DataAudienceDeviceBrand = "SMARTISAN"
	GIONEE_AdGetV2DataAudienceDeviceBrand      AdGetV2DataAudienceDeviceBrand = "GIONEE"
	MOTO_AdGetV2DataAudienceDeviceBrand        AdGetV2DataAudienceDeviceBrand = "MOTO"
	NOKIA_AdGetV2DataAudienceDeviceBrand       AdGetV2DataAudienceDeviceBrand = "NOKIA"
	OPPO_AdGetV2DataAudienceDeviceBrand        AdGetV2DataAudienceDeviceBrand = "OPPO"
	LETV_AdGetV2DataAudienceDeviceBrand        AdGetV2DataAudienceDeviceBrand = "LETV"
	HUAWEI_AdGetV2DataAudienceDeviceBrand      AdGetV2DataAudienceDeviceBrand = "HUAWEI"
	LENOVO_AdGetV2DataAudienceDeviceBrand      AdGetV2DataAudienceDeviceBrand = "LENOVO"
	ONEPLUS_AdGetV2DataAudienceDeviceBrand     AdGetV2DataAudienceDeviceBrand = "ONEPLUS"
	HONOR_AdGetV2DataAudienceDeviceBrand       AdGetV2DataAudienceDeviceBrand = "HONOR"
	GOOGLE_AdGetV2DataAudienceDeviceBrand      AdGetV2DataAudienceDeviceBrand = "GOOGLE"
	TCL_AdGetV2DataAudienceDeviceBrand         AdGetV2DataAudienceDeviceBrand = "TCL"
	SONY_AdGetV2DataAudienceDeviceBrand        AdGetV2DataAudienceDeviceBrand = "SONY"
	COOLPAD_AdGetV2DataAudienceDeviceBrand     AdGetV2DataAudienceDeviceBrand = "COOLPAD"
	APPLE_AdGetV2DataAudienceDeviceBrand       AdGetV2DataAudienceDeviceBrand = "APPLE"
)

// All allowed values of AdGetV2DataAudienceDeviceBrand enum
var AllowedAdGetV2DataAudienceDeviceBrandEnumValues = []AdGetV2DataAudienceDeviceBrand{
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
