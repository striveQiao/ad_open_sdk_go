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

// AdGetV2DataIntelligentFlowSwitch
type AdGetV2DataIntelligentFlowSwitch string

// List of ad_get_v2_data_intelligent_flow_switch
const (
	ON_AdGetV2DataIntelligentFlowSwitch  AdGetV2DataIntelligentFlowSwitch = "ON"
	OFF_AdGetV2DataIntelligentFlowSwitch AdGetV2DataIntelligentFlowSwitch = "OFF"
)

// All allowed values of AdGetV2DataIntelligentFlowSwitch enum
var AllowedAdGetV2DataIntelligentFlowSwitchEnumValues = []AdGetV2DataIntelligentFlowSwitch{
	"ON",
	"OFF",
}

// NewAdGetV2DataIntelligentFlowSwitchFromValue returns a pointer to a valid AdGetV2DataIntelligentFlowSwitch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataIntelligentFlowSwitchFromValue(v string) (*AdGetV2DataIntelligentFlowSwitch, error) {
	ev := AdGetV2DataIntelligentFlowSwitch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataIntelligentFlowSwitch: valid values are %v", v, AllowedAdGetV2DataIntelligentFlowSwitchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataIntelligentFlowSwitch) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataIntelligentFlowSwitchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_intelligent_flow_switch value
func (v AdGetV2DataIntelligentFlowSwitch) Ptr() *AdGetV2DataIntelligentFlowSwitch {
	return &v
}
