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

// PromotionUpdateV30PromotionMaterialsDynamicCreativeSwitch
type PromotionUpdateV30PromotionMaterialsDynamicCreativeSwitch string

// List of promotion_update_v3.0_promotion_materials_dynamic_creative_switch
const (
	OFF_PromotionUpdateV30PromotionMaterialsDynamicCreativeSwitch PromotionUpdateV30PromotionMaterialsDynamicCreativeSwitch = "OFF"
	ON_PromotionUpdateV30PromotionMaterialsDynamicCreativeSwitch  PromotionUpdateV30PromotionMaterialsDynamicCreativeSwitch = "ON"
)

// All allowed values of PromotionUpdateV30PromotionMaterialsDynamicCreativeSwitch enum
var AllowedPromotionUpdateV30PromotionMaterialsDynamicCreativeSwitchEnumValues = []PromotionUpdateV30PromotionMaterialsDynamicCreativeSwitch{
	"OFF",
	"ON",
}

// NewPromotionUpdateV30PromotionMaterialsDynamicCreativeSwitchFromValue returns a pointer to a valid PromotionUpdateV30PromotionMaterialsDynamicCreativeSwitch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionUpdateV30PromotionMaterialsDynamicCreativeSwitchFromValue(v string) (*PromotionUpdateV30PromotionMaterialsDynamicCreativeSwitch, error) {
	ev := PromotionUpdateV30PromotionMaterialsDynamicCreativeSwitch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionUpdateV30PromotionMaterialsDynamicCreativeSwitch: valid values are %v", v, AllowedPromotionUpdateV30PromotionMaterialsDynamicCreativeSwitchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionUpdateV30PromotionMaterialsDynamicCreativeSwitch) IsValid() bool {
	for _, existing := range AllowedPromotionUpdateV30PromotionMaterialsDynamicCreativeSwitchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_update_v3.0_promotion_materials_dynamic_creative_switch value
func (v PromotionUpdateV30PromotionMaterialsDynamicCreativeSwitch) Ptr() *PromotionUpdateV30PromotionMaterialsDynamicCreativeSwitch {
	return &v
}
