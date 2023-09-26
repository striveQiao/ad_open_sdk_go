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

// AudiencePackageUpdateV2DeviceBrand
type AudiencePackageUpdateV2DeviceBrand string

// List of audience_package_update_v2_device_brand
const (
	APPLE_AudiencePackageUpdateV2DeviceBrand       AudiencePackageUpdateV2DeviceBrand = "APPLE"
	PEPPER_AudiencePackageUpdateV2DeviceBrand      AudiencePackageUpdateV2DeviceBrand = "PEPPER"
	QIKU_AudiencePackageUpdateV2DeviceBrand        AudiencePackageUpdateV2DeviceBrand = "QIKU"
	NUBIA_AudiencePackageUpdateV2DeviceBrand       AudiencePackageUpdateV2DeviceBrand = "NUBIA"
	OPPO_AudiencePackageUpdateV2DeviceBrand        AudiencePackageUpdateV2DeviceBrand = "OPPO"
	GOOGLE_AudiencePackageUpdateV2DeviceBrand      AudiencePackageUpdateV2DeviceBrand = "GOOGLE"
	HISENSE_AudiencePackageUpdateV2DeviceBrand     AudiencePackageUpdateV2DeviceBrand = "HISENSE"
	COOLPAD_AudiencePackageUpdateV2DeviceBrand     AudiencePackageUpdateV2DeviceBrand = "COOLPAD"
	SAMSUNG_AudiencePackageUpdateV2DeviceBrand     AudiencePackageUpdateV2DeviceBrand = "SAMSUNG"
	SMARTISAN_AudiencePackageUpdateV2DeviceBrand   AudiencePackageUpdateV2DeviceBrand = "SMARTISAN"
	ONEPLUS_AudiencePackageUpdateV2DeviceBrand     AudiencePackageUpdateV2DeviceBrand = "ONEPLUS"
	ZTE_AudiencePackageUpdateV2DeviceBrand         AudiencePackageUpdateV2DeviceBrand = "ZTE"
	HTC_AudiencePackageUpdateV2DeviceBrand         AudiencePackageUpdateV2DeviceBrand = "HTC"
	LG_AudiencePackageUpdateV2DeviceBrand          AudiencePackageUpdateV2DeviceBrand = "LG"
	XIAOMI_AudiencePackageUpdateV2DeviceBrand      AudiencePackageUpdateV2DeviceBrand = "XIAOMI"
	MEIZU_AudiencePackageUpdateV2DeviceBrand       AudiencePackageUpdateV2DeviceBrand = "MEIZU"
	HONOR_AudiencePackageUpdateV2DeviceBrand       AudiencePackageUpdateV2DeviceBrand = "HONOR"
	LETV_AudiencePackageUpdateV2DeviceBrand        AudiencePackageUpdateV2DeviceBrand = "LETV"
	TCL_AudiencePackageUpdateV2DeviceBrand         AudiencePackageUpdateV2DeviceBrand = "TCL"
	HUAWEI_AudiencePackageUpdateV2DeviceBrand      AudiencePackageUpdateV2DeviceBrand = "HUAWEI"
	SONY_AudiencePackageUpdateV2DeviceBrand        AudiencePackageUpdateV2DeviceBrand = "SONY"
	Enum_360_AudiencePackageUpdateV2DeviceBrand    AudiencePackageUpdateV2DeviceBrand = "360"
	CHINAMOBILE_AudiencePackageUpdateV2DeviceBrand AudiencePackageUpdateV2DeviceBrand = "CHINAMOBILE"
	GIONEE_AudiencePackageUpdateV2DeviceBrand      AudiencePackageUpdateV2DeviceBrand = "GIONEE"
	VIVO_AudiencePackageUpdateV2DeviceBrand        AudiencePackageUpdateV2DeviceBrand = "VIVO"
	NOKIA_AudiencePackageUpdateV2DeviceBrand       AudiencePackageUpdateV2DeviceBrand = "NOKIA"
	MOTO_AudiencePackageUpdateV2DeviceBrand        AudiencePackageUpdateV2DeviceBrand = "MOTO"
	LENOVO_AudiencePackageUpdateV2DeviceBrand      AudiencePackageUpdateV2DeviceBrand = "LENOVO"
)

// All allowed values of AudiencePackageUpdateV2DeviceBrand enum
var AllowedAudiencePackageUpdateV2DeviceBrandEnumValues = []AudiencePackageUpdateV2DeviceBrand{
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

// NewAudiencePackageUpdateV2DeviceBrandFromValue returns a pointer to a valid AudiencePackageUpdateV2DeviceBrand
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageUpdateV2DeviceBrandFromValue(v string) (*AudiencePackageUpdateV2DeviceBrand, error) {
	ev := AudiencePackageUpdateV2DeviceBrand(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageUpdateV2DeviceBrand: valid values are %v", v, AllowedAudiencePackageUpdateV2DeviceBrandEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageUpdateV2DeviceBrand) IsValid() bool {
	for _, existing := range AllowedAudiencePackageUpdateV2DeviceBrandEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_update_v2_device_brand value
func (v AudiencePackageUpdateV2DeviceBrand) Ptr() *AudiencePackageUpdateV2DeviceBrand {
	return &v
}
