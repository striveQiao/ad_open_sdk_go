/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageCreateV2IosOsv
type AudiencePackageCreateV2IosOsv string

// List of audience_package_create_v2_ios_osv
const (
	Enum_0_0_AudiencePackageCreateV2IosOsv  AudiencePackageCreateV2IosOsv = "0.0"
	Enum_9_3_AudiencePackageCreateV2IosOsv  AudiencePackageCreateV2IosOsv = "9.3"
	Enum_13_2_AudiencePackageCreateV2IosOsv AudiencePackageCreateV2IosOsv = "13.2"
	Enum_4_1_AudiencePackageCreateV2IosOsv  AudiencePackageCreateV2IosOsv = "4.1"
	Enum_12_1_AudiencePackageCreateV2IosOsv AudiencePackageCreateV2IosOsv = "12.1"
	Enum_11_4_AudiencePackageCreateV2IosOsv AudiencePackageCreateV2IosOsv = "11.4"
	Enum_8_2_AudiencePackageCreateV2IosOsv  AudiencePackageCreateV2IosOsv = "8.2"
	Enum_5_0_AudiencePackageCreateV2IosOsv  AudiencePackageCreateV2IosOsv = "5.0"
	Enum_14_0_AudiencePackageCreateV2IosOsv AudiencePackageCreateV2IosOsv = "14.0"
	Enum_11_3_AudiencePackageCreateV2IosOsv AudiencePackageCreateV2IosOsv = "11.3"
	Enum_7_0_AudiencePackageCreateV2IosOsv  AudiencePackageCreateV2IosOsv = "7.0"
	Enum_13_3_AudiencePackageCreateV2IosOsv AudiencePackageCreateV2IosOsv = "13.3"
	Enum_13_4_AudiencePackageCreateV2IosOsv AudiencePackageCreateV2IosOsv = "13.4"
	Enum_12_4_AudiencePackageCreateV2IosOsv AudiencePackageCreateV2IosOsv = "12.4"
	Enum_4_3_AudiencePackageCreateV2IosOsv  AudiencePackageCreateV2IosOsv = "4.3"
	Enum_11_0_AudiencePackageCreateV2IosOsv AudiencePackageCreateV2IosOsv = "11.0"
	Enum_9_1_AudiencePackageCreateV2IosOsv  AudiencePackageCreateV2IosOsv = "9.1"
	Enum_13_1_AudiencePackageCreateV2IosOsv AudiencePackageCreateV2IosOsv = "13.1"
	Enum_4_0_AudiencePackageCreateV2IosOsv  AudiencePackageCreateV2IosOsv = "4.0"
	Enum_10_3_AudiencePackageCreateV2IosOsv AudiencePackageCreateV2IosOsv = "10.3"
	NONE_AudiencePackageCreateV2IosOsv      AudiencePackageCreateV2IosOsv = "NONE"
	Enum_13_0_AudiencePackageCreateV2IosOsv AudiencePackageCreateV2IosOsv = "13.0"
	Enum_5_1_AudiencePackageCreateV2IosOsv  AudiencePackageCreateV2IosOsv = "5.1"
	Enum_12_0_AudiencePackageCreateV2IosOsv AudiencePackageCreateV2IosOsv = "12.0"
	Enum_9_0_AudiencePackageCreateV2IosOsv  AudiencePackageCreateV2IosOsv = "9.0"
	Enum_8_0_AudiencePackageCreateV2IosOsv  AudiencePackageCreateV2IosOsv = "8.0"
	Enum_11_1_AudiencePackageCreateV2IosOsv AudiencePackageCreateV2IosOsv = "11.1"
	Enum_7_1_AudiencePackageCreateV2IosOsv  AudiencePackageCreateV2IosOsv = "7.1"
	Enum_9_2_AudiencePackageCreateV2IosOsv  AudiencePackageCreateV2IosOsv = "9.2"
	Enum_10_2_AudiencePackageCreateV2IosOsv AudiencePackageCreateV2IosOsv = "10.2"
	Enum_12_2_AudiencePackageCreateV2IosOsv AudiencePackageCreateV2IosOsv = "12.2"
	Enum_8_1_AudiencePackageCreateV2IosOsv  AudiencePackageCreateV2IosOsv = "8.1"
	Enum_12_3_AudiencePackageCreateV2IosOsv AudiencePackageCreateV2IosOsv = "12.3"
	Enum_6_0_AudiencePackageCreateV2IosOsv  AudiencePackageCreateV2IosOsv = "6.0"
	Enum_10_1_AudiencePackageCreateV2IosOsv AudiencePackageCreateV2IosOsv = "10.1"
	Enum_11_2_AudiencePackageCreateV2IosOsv AudiencePackageCreateV2IosOsv = "11.2"
	Enum_4_2_AudiencePackageCreateV2IosOsv  AudiencePackageCreateV2IosOsv = "4.2"
)

// All allowed values of AudiencePackageCreateV2IosOsv enum
var AllowedAudiencePackageCreateV2IosOsvEnumValues = []AudiencePackageCreateV2IosOsv{
	"0.0",
	"9.3",
	"13.2",
	"4.1",
	"12.1",
	"11.4",
	"8.2",
	"5.0",
	"14.0",
	"11.3",
	"7.0",
	"13.3",
	"13.4",
	"12.4",
	"4.3",
	"11.0",
	"9.1",
	"13.1",
	"4.0",
	"10.3",
	"NONE",
	"13.0",
	"5.1",
	"12.0",
	"9.0",
	"8.0",
	"11.1",
	"7.1",
	"9.2",
	"10.2",
	"12.2",
	"8.1",
	"12.3",
	"6.0",
	"10.1",
	"11.2",
	"4.2",
}

// NewAudiencePackageCreateV2IosOsvFromValue returns a pointer to a valid AudiencePackageCreateV2IosOsv
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageCreateV2IosOsvFromValue(v string) (*AudiencePackageCreateV2IosOsv, error) {
	ev := AudiencePackageCreateV2IosOsv(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageCreateV2IosOsv: valid values are %v", v, AllowedAudiencePackageCreateV2IosOsvEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageCreateV2IosOsv) IsValid() bool {
	for _, existing := range AllowedAudiencePackageCreateV2IosOsvEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_create_v2_ios_osv value
func (v AudiencePackageCreateV2IosOsv) Ptr() *AudiencePackageCreateV2IosOsv {
	return &v
}
