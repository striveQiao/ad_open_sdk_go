/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsEstimateAudienceV2Carrier
type ToolsEstimateAudienceV2Carrier string

// List of tools_estimate_audience_v2_carrier
const (
	TELCOM_ToolsEstimateAudienceV2Carrier ToolsEstimateAudienceV2Carrier = "TELCOM"
	UNICOM_ToolsEstimateAudienceV2Carrier ToolsEstimateAudienceV2Carrier = "UNICOM"
	MOBILE_ToolsEstimateAudienceV2Carrier ToolsEstimateAudienceV2Carrier = "MOBILE"
)

// All allowed values of ToolsEstimateAudienceV2Carrier enum
var AllowedToolsEstimateAudienceV2CarrierEnumValues = []ToolsEstimateAudienceV2Carrier{
	"TELCOM",
	"UNICOM",
	"MOBILE",
}

// NewToolsEstimateAudienceV2CarrierFromValue returns a pointer to a valid ToolsEstimateAudienceV2Carrier
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEstimateAudienceV2CarrierFromValue(v string) (*ToolsEstimateAudienceV2Carrier, error) {
	ev := ToolsEstimateAudienceV2Carrier(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEstimateAudienceV2Carrier: valid values are %v", v, AllowedToolsEstimateAudienceV2CarrierEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEstimateAudienceV2Carrier) IsValid() bool {
	for _, existing := range AllowedToolsEstimateAudienceV2CarrierEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_estimate_audience_v2_carrier value
func (v ToolsEstimateAudienceV2Carrier) Ptr() *ToolsEstimateAudienceV2Carrier {
	return &v
}
