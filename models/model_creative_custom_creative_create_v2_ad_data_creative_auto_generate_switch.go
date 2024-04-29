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

// CreativeCustomCreativeCreateV2AdDataCreativeAutoGenerateSwitch
type CreativeCustomCreativeCreateV2AdDataCreativeAutoGenerateSwitch int64

// List of creative_custom_creative_create_v2_ad_data_creative_auto_generate_switch
const (
	Enum_0_CreativeCustomCreativeCreateV2AdDataCreativeAutoGenerateSwitch CreativeCustomCreativeCreateV2AdDataCreativeAutoGenerateSwitch = 0
	Enum_1_CreativeCustomCreativeCreateV2AdDataCreativeAutoGenerateSwitch CreativeCustomCreativeCreateV2AdDataCreativeAutoGenerateSwitch = 1
)

// All allowed values of CreativeCustomCreativeCreateV2AdDataCreativeAutoGenerateSwitch enum
var AllowedCreativeCustomCreativeCreateV2AdDataCreativeAutoGenerateSwitchEnumValues = []CreativeCustomCreativeCreateV2AdDataCreativeAutoGenerateSwitch{
	0,
	1,
}

// NewCreativeCustomCreativeCreateV2AdDataCreativeAutoGenerateSwitchFromValue returns a pointer to a valid CreativeCustomCreativeCreateV2AdDataCreativeAutoGenerateSwitch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeCustomCreativeCreateV2AdDataCreativeAutoGenerateSwitchFromValue(v int64) (*CreativeCustomCreativeCreateV2AdDataCreativeAutoGenerateSwitch, error) {
	ev := CreativeCustomCreativeCreateV2AdDataCreativeAutoGenerateSwitch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeCustomCreativeCreateV2AdDataCreativeAutoGenerateSwitch: valid values are %v", v, AllowedCreativeCustomCreativeCreateV2AdDataCreativeAutoGenerateSwitchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeCustomCreativeCreateV2AdDataCreativeAutoGenerateSwitch) IsValid() bool {
	for _, existing := range AllowedCreativeCustomCreativeCreateV2AdDataCreativeAutoGenerateSwitchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_custom_creative_create_v2_ad_data_creative_auto_generate_switch value
func (v CreativeCustomCreativeCreateV2AdDataCreativeAutoGenerateSwitch) Ptr() *CreativeCustomCreativeCreateV2AdDataCreativeAutoGenerateSwitch {
	return &v
}
