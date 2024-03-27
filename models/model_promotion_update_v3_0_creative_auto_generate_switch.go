/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionUpdateV30CreativeAutoGenerateSwitch
type PromotionUpdateV30CreativeAutoGenerateSwitch string

// List of promotion_update_v3.0_creative_auto_generate_switch
const (
	OFF_PromotionUpdateV30CreativeAutoGenerateSwitch PromotionUpdateV30CreativeAutoGenerateSwitch = "OFF"
	ON_PromotionUpdateV30CreativeAutoGenerateSwitch  PromotionUpdateV30CreativeAutoGenerateSwitch = "ON"
)

// All allowed values of PromotionUpdateV30CreativeAutoGenerateSwitch enum
var AllowedPromotionUpdateV30CreativeAutoGenerateSwitchEnumValues = []PromotionUpdateV30CreativeAutoGenerateSwitch{
	"OFF",
	"ON",
}

// NewPromotionUpdateV30CreativeAutoGenerateSwitchFromValue returns a pointer to a valid PromotionUpdateV30CreativeAutoGenerateSwitch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionUpdateV30CreativeAutoGenerateSwitchFromValue(v string) (*PromotionUpdateV30CreativeAutoGenerateSwitch, error) {
	ev := PromotionUpdateV30CreativeAutoGenerateSwitch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionUpdateV30CreativeAutoGenerateSwitch: valid values are %v", v, AllowedPromotionUpdateV30CreativeAutoGenerateSwitchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionUpdateV30CreativeAutoGenerateSwitch) IsValid() bool {
	for _, existing := range AllowedPromotionUpdateV30CreativeAutoGenerateSwitchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_update_v3.0_creative_auto_generate_switch value
func (v PromotionUpdateV30CreativeAutoGenerateSwitch) Ptr() *PromotionUpdateV30CreativeAutoGenerateSwitch {
	return &v
}
