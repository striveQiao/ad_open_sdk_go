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

// ToolsBidSuggestV2Carrier
type ToolsBidSuggestV2Carrier string

// List of tools_bid_suggest_v2_carrier
const (
	UNICOM_ToolsBidSuggestV2Carrier ToolsBidSuggestV2Carrier = "UNICOM"
	MOBILE_ToolsBidSuggestV2Carrier ToolsBidSuggestV2Carrier = "MOBILE"
	TELCOM_ToolsBidSuggestV2Carrier ToolsBidSuggestV2Carrier = "TELCOM"
)

// All allowed values of ToolsBidSuggestV2Carrier enum
var AllowedToolsBidSuggestV2CarrierEnumValues = []ToolsBidSuggestV2Carrier{
	"UNICOM",
	"MOBILE",
	"TELCOM",
}

// NewToolsBidSuggestV2CarrierFromValue returns a pointer to a valid ToolsBidSuggestV2Carrier
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2CarrierFromValue(v string) (*ToolsBidSuggestV2Carrier, error) {
	ev := ToolsBidSuggestV2Carrier(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2Carrier: valid values are %v", v, AllowedToolsBidSuggestV2CarrierEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2Carrier) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2CarrierEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_carrier value
func (v ToolsBidSuggestV2Carrier) Ptr() *ToolsBidSuggestV2Carrier {
	return &v
}
