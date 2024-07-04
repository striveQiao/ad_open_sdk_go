/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeCustomCreativeUpdateV2AdDataDynamicCreativeSwitch
type CreativeCustomCreativeUpdateV2AdDataDynamicCreativeSwitch string

// List of creative_custom_creative_update_v2_ad_data_dynamic_creative_switch
const (
	DYNAMIC_CREATIVE_SUBLINK_CreativeCustomCreativeUpdateV2AdDataDynamicCreativeSwitch  CreativeCustomCreativeUpdateV2AdDataDynamicCreativeSwitch = "DYNAMIC_CREATIVE_SUBLINK"
	DYNAMIC_CREATIVE_ON_CreativeCustomCreativeUpdateV2AdDataDynamicCreativeSwitch       CreativeCustomCreativeUpdateV2AdDataDynamicCreativeSwitch = "DYNAMIC_CREATIVE_ON"
	DYNAMIC_CREATIVE_TITLE_CreativeCustomCreativeUpdateV2AdDataDynamicCreativeSwitch    CreativeCustomCreativeUpdateV2AdDataDynamicCreativeSwitch = "DYNAMIC_CREATIVE_TITLE"
	DYNAMIC_CREATIVE_ABSTRACT_CreativeCustomCreativeUpdateV2AdDataDynamicCreativeSwitch CreativeCustomCreativeUpdateV2AdDataDynamicCreativeSwitch = "DYNAMIC_CREATIVE_ABSTRACT"
)

// All allowed values of CreativeCustomCreativeUpdateV2AdDataDynamicCreativeSwitch enum
var AllowedCreativeCustomCreativeUpdateV2AdDataDynamicCreativeSwitchEnumValues = []CreativeCustomCreativeUpdateV2AdDataDynamicCreativeSwitch{
	"DYNAMIC_CREATIVE_SUBLINK",
	"DYNAMIC_CREATIVE_ON",
	"DYNAMIC_CREATIVE_TITLE",
	"DYNAMIC_CREATIVE_ABSTRACT",
}

// NewCreativeCustomCreativeUpdateV2AdDataDynamicCreativeSwitchFromValue returns a pointer to a valid CreativeCustomCreativeUpdateV2AdDataDynamicCreativeSwitch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeCustomCreativeUpdateV2AdDataDynamicCreativeSwitchFromValue(v string) (*CreativeCustomCreativeUpdateV2AdDataDynamicCreativeSwitch, error) {
	ev := CreativeCustomCreativeUpdateV2AdDataDynamicCreativeSwitch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeCustomCreativeUpdateV2AdDataDynamicCreativeSwitch: valid values are %v", v, AllowedCreativeCustomCreativeUpdateV2AdDataDynamicCreativeSwitchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeCustomCreativeUpdateV2AdDataDynamicCreativeSwitch) IsValid() bool {
	for _, existing := range AllowedCreativeCustomCreativeUpdateV2AdDataDynamicCreativeSwitchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_custom_creative_update_v2_ad_data_dynamic_creative_switch value
func (v CreativeCustomCreativeUpdateV2AdDataDynamicCreativeSwitch) Ptr() *CreativeCustomCreativeUpdateV2AdDataDynamicCreativeSwitch {
	return &v
}
