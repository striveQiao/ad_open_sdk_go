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

// AdGetV2DataAudienceAc
type AdGetV2DataAudienceAc string

// List of ad_get_v2_data_audience_ac
const (
	Enum_5_G_AdGetV2DataAudienceAc AdGetV2DataAudienceAc = "5G"
	WIFI_AdGetV2DataAudienceAc     AdGetV2DataAudienceAc = "WIFI"
	Enum_3_G_AdGetV2DataAudienceAc AdGetV2DataAudienceAc = "3G"
	Enum_2_G_AdGetV2DataAudienceAc AdGetV2DataAudienceAc = "2G"
	Enum_4_G_AdGetV2DataAudienceAc AdGetV2DataAudienceAc = "4G"
)

// All allowed values of AdGetV2DataAudienceAc enum
var AllowedAdGetV2DataAudienceAcEnumValues = []AdGetV2DataAudienceAc{
	"5G",
	"WIFI",
	"3G",
	"2G",
	"4G",
}

// NewAdGetV2DataAudienceAcFromValue returns a pointer to a valid AdGetV2DataAudienceAc
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataAudienceAcFromValue(v string) (*AdGetV2DataAudienceAc, error) {
	ev := AdGetV2DataAudienceAc(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataAudienceAc: valid values are %v", v, AllowedAdGetV2DataAudienceAcEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataAudienceAc) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataAudienceAcEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_audience_ac value
func (v AdGetV2DataAudienceAc) Ptr() *AdGetV2DataAudienceAc {
	return &v
}
