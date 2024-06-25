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

// ToolsBidSuggestV2AndroidOsv
type ToolsBidSuggestV2AndroidOsv string

// List of tools_bid_suggest_v2_android_osv
const (
	Enum_5_0_ToolsBidSuggestV2AndroidOsv  ToolsBidSuggestV2AndroidOsv = "5.0"
	Enum_3_0_ToolsBidSuggestV2AndroidOsv  ToolsBidSuggestV2AndroidOsv = "3.0"
	Enum_2_0_ToolsBidSuggestV2AndroidOsv  ToolsBidSuggestV2AndroidOsv = "2.0"
	Enum_3_1_ToolsBidSuggestV2AndroidOsv  ToolsBidSuggestV2AndroidOsv = "3.1"
	Enum_8_0_ToolsBidSuggestV2AndroidOsv  ToolsBidSuggestV2AndroidOsv = "8.0"
	Enum_10_1_ToolsBidSuggestV2AndroidOsv ToolsBidSuggestV2AndroidOsv = "10.1"
	Enum_2_3_ToolsBidSuggestV2AndroidOsv  ToolsBidSuggestV2AndroidOsv = "2.3"
	Enum_7_0_ToolsBidSuggestV2AndroidOsv  ToolsBidSuggestV2AndroidOsv = "7.0"
	Enum_6_0_ToolsBidSuggestV2AndroidOsv  ToolsBidSuggestV2AndroidOsv = "6.0"
	Enum_10_0_ToolsBidSuggestV2AndroidOsv ToolsBidSuggestV2AndroidOsv = "10.0"
	Enum_9_0_ToolsBidSuggestV2AndroidOsv  ToolsBidSuggestV2AndroidOsv = "9.0"
	Enum_2_1_ToolsBidSuggestV2AndroidOsv  ToolsBidSuggestV2AndroidOsv = "2.1"
	Enum_7_1_ToolsBidSuggestV2AndroidOsv  ToolsBidSuggestV2AndroidOsv = "7.1"
	Enum_5_1_ToolsBidSuggestV2AndroidOsv  ToolsBidSuggestV2AndroidOsv = "5.1"
	Enum_4_0_ToolsBidSuggestV2AndroidOsv  ToolsBidSuggestV2AndroidOsv = "4.0"
	NONE_ToolsBidSuggestV2AndroidOsv      ToolsBidSuggestV2AndroidOsv = "NONE"
	Enum_0_0_ToolsBidSuggestV2AndroidOsv  ToolsBidSuggestV2AndroidOsv = "0.0"
	Enum_4_2_ToolsBidSuggestV2AndroidOsv  ToolsBidSuggestV2AndroidOsv = "4.2"
	Enum_4_5_ToolsBidSuggestV2AndroidOsv  ToolsBidSuggestV2AndroidOsv = "4.5"
	Enum_3_2_ToolsBidSuggestV2AndroidOsv  ToolsBidSuggestV2AndroidOsv = "3.2"
	Enum_4_4_ToolsBidSuggestV2AndroidOsv  ToolsBidSuggestV2AndroidOsv = "4.4"
	Enum_11_0_ToolsBidSuggestV2AndroidOsv ToolsBidSuggestV2AndroidOsv = "11.0"
	Enum_4_3_ToolsBidSuggestV2AndroidOsv  ToolsBidSuggestV2AndroidOsv = "4.3"
	Enum_2_2_ToolsBidSuggestV2AndroidOsv  ToolsBidSuggestV2AndroidOsv = "2.2"
	Enum_8_1_ToolsBidSuggestV2AndroidOsv  ToolsBidSuggestV2AndroidOsv = "8.1"
	Enum_4_1_ToolsBidSuggestV2AndroidOsv  ToolsBidSuggestV2AndroidOsv = "4.1"
)

// All allowed values of ToolsBidSuggestV2AndroidOsv enum
var AllowedToolsBidSuggestV2AndroidOsvEnumValues = []ToolsBidSuggestV2AndroidOsv{
	"5.0",
	"3.0",
	"2.0",
	"3.1",
	"8.0",
	"10.1",
	"2.3",
	"7.0",
	"6.0",
	"10.0",
	"9.0",
	"2.1",
	"7.1",
	"5.1",
	"4.0",
	"NONE",
	"0.0",
	"4.2",
	"4.5",
	"3.2",
	"4.4",
	"11.0",
	"4.3",
	"2.2",
	"8.1",
	"4.1",
}

// NewToolsBidSuggestV2AndroidOsvFromValue returns a pointer to a valid ToolsBidSuggestV2AndroidOsv
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2AndroidOsvFromValue(v string) (*ToolsBidSuggestV2AndroidOsv, error) {
	ev := ToolsBidSuggestV2AndroidOsv(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2AndroidOsv: valid values are %v", v, AllowedToolsBidSuggestV2AndroidOsvEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2AndroidOsv) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2AndroidOsvEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_android_osv value
func (v ToolsBidSuggestV2AndroidOsv) Ptr() *ToolsBidSuggestV2AndroidOsv {
	return &v
}
