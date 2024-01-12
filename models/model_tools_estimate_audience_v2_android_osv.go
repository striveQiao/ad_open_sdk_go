/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsEstimateAudienceV2AndroidOsv
type ToolsEstimateAudienceV2AndroidOsv string

// List of tools_estimate_audience_v2_android_osv
const (
	Enum_2_2_ToolsEstimateAudienceV2AndroidOsv  ToolsEstimateAudienceV2AndroidOsv = "2.2"
	Enum_3_2_ToolsEstimateAudienceV2AndroidOsv  ToolsEstimateAudienceV2AndroidOsv = "3.2"
	Enum_4_5_ToolsEstimateAudienceV2AndroidOsv  ToolsEstimateAudienceV2AndroidOsv = "4.5"
	Enum_2_3_ToolsEstimateAudienceV2AndroidOsv  ToolsEstimateAudienceV2AndroidOsv = "2.3"
	Enum_11_0_ToolsEstimateAudienceV2AndroidOsv ToolsEstimateAudienceV2AndroidOsv = "11.0"
	Enum_3_0_ToolsEstimateAudienceV2AndroidOsv  ToolsEstimateAudienceV2AndroidOsv = "3.0"
	Enum_5_1_ToolsEstimateAudienceV2AndroidOsv  ToolsEstimateAudienceV2AndroidOsv = "5.1"
	Enum_0_0_ToolsEstimateAudienceV2AndroidOsv  ToolsEstimateAudienceV2AndroidOsv = "0.0"
	Enum_7_0_ToolsEstimateAudienceV2AndroidOsv  ToolsEstimateAudienceV2AndroidOsv = "7.0"
	Enum_9_0_ToolsEstimateAudienceV2AndroidOsv  ToolsEstimateAudienceV2AndroidOsv = "9.0"
	Enum_10_1_ToolsEstimateAudienceV2AndroidOsv ToolsEstimateAudienceV2AndroidOsv = "10.1"
	Enum_10_0_ToolsEstimateAudienceV2AndroidOsv ToolsEstimateAudienceV2AndroidOsv = "10.0"
	Enum_5_0_ToolsEstimateAudienceV2AndroidOsv  ToolsEstimateAudienceV2AndroidOsv = "5.0"
	Enum_7_1_ToolsEstimateAudienceV2AndroidOsv  ToolsEstimateAudienceV2AndroidOsv = "7.1"
	Enum_4_0_ToolsEstimateAudienceV2AndroidOsv  ToolsEstimateAudienceV2AndroidOsv = "4.0"
	Enum_4_4_ToolsEstimateAudienceV2AndroidOsv  ToolsEstimateAudienceV2AndroidOsv = "4.4"
	Enum_6_0_ToolsEstimateAudienceV2AndroidOsv  ToolsEstimateAudienceV2AndroidOsv = "6.0"
	NONE_ToolsEstimateAudienceV2AndroidOsv      ToolsEstimateAudienceV2AndroidOsv = "NONE"
	Enum_4_2_ToolsEstimateAudienceV2AndroidOsv  ToolsEstimateAudienceV2AndroidOsv = "4.2"
	Enum_2_1_ToolsEstimateAudienceV2AndroidOsv  ToolsEstimateAudienceV2AndroidOsv = "2.1"
	Enum_2_0_ToolsEstimateAudienceV2AndroidOsv  ToolsEstimateAudienceV2AndroidOsv = "2.0"
	Enum_4_1_ToolsEstimateAudienceV2AndroidOsv  ToolsEstimateAudienceV2AndroidOsv = "4.1"
	Enum_4_3_ToolsEstimateAudienceV2AndroidOsv  ToolsEstimateAudienceV2AndroidOsv = "4.3"
	Enum_8_1_ToolsEstimateAudienceV2AndroidOsv  ToolsEstimateAudienceV2AndroidOsv = "8.1"
	Enum_3_1_ToolsEstimateAudienceV2AndroidOsv  ToolsEstimateAudienceV2AndroidOsv = "3.1"
	Enum_8_0_ToolsEstimateAudienceV2AndroidOsv  ToolsEstimateAudienceV2AndroidOsv = "8.0"
)

// All allowed values of ToolsEstimateAudienceV2AndroidOsv enum
var AllowedToolsEstimateAudienceV2AndroidOsvEnumValues = []ToolsEstimateAudienceV2AndroidOsv{
	"2.2",
	"3.2",
	"4.5",
	"2.3",
	"11.0",
	"3.0",
	"5.1",
	"0.0",
	"7.0",
	"9.0",
	"10.1",
	"10.0",
	"5.0",
	"7.1",
	"4.0",
	"4.4",
	"6.0",
	"NONE",
	"4.2",
	"2.1",
	"2.0",
	"4.1",
	"4.3",
	"8.1",
	"3.1",
	"8.0",
}

// NewToolsEstimateAudienceV2AndroidOsvFromValue returns a pointer to a valid ToolsEstimateAudienceV2AndroidOsv
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEstimateAudienceV2AndroidOsvFromValue(v string) (*ToolsEstimateAudienceV2AndroidOsv, error) {
	ev := ToolsEstimateAudienceV2AndroidOsv(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEstimateAudienceV2AndroidOsv: valid values are %v", v, AllowedToolsEstimateAudienceV2AndroidOsvEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEstimateAudienceV2AndroidOsv) IsValid() bool {
	for _, existing := range AllowedToolsEstimateAudienceV2AndroidOsvEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_estimate_audience_v2_android_osv value
func (v ToolsEstimateAudienceV2AndroidOsv) Ptr() *ToolsEstimateAudienceV2AndroidOsv {
	return &v
}
