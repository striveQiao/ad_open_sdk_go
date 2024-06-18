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

// PromotionListV30FilteringHasCarryMaterial
type PromotionListV30FilteringHasCarryMaterial string

// List of promotion_list_v3.0_filtering_has_carry_material
const (
	TRUE_PromotionListV30FilteringHasCarryMaterial  PromotionListV30FilteringHasCarryMaterial = "TRUE"
	FALSE_PromotionListV30FilteringHasCarryMaterial PromotionListV30FilteringHasCarryMaterial = "FALSE"
)

// All allowed values of PromotionListV30FilteringHasCarryMaterial enum
var AllowedPromotionListV30FilteringHasCarryMaterialEnumValues = []PromotionListV30FilteringHasCarryMaterial{
	"TRUE",
	"FALSE",
}

// NewPromotionListV30FilteringHasCarryMaterialFromValue returns a pointer to a valid PromotionListV30FilteringHasCarryMaterial
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30FilteringHasCarryMaterialFromValue(v string) (*PromotionListV30FilteringHasCarryMaterial, error) {
	ev := PromotionListV30FilteringHasCarryMaterial(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30FilteringHasCarryMaterial: valid values are %v", v, AllowedPromotionListV30FilteringHasCarryMaterialEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30FilteringHasCarryMaterial) IsValid() bool {
	for _, existing := range AllowedPromotionListV30FilteringHasCarryMaterialEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_filtering_has_carry_material value
func (v PromotionListV30FilteringHasCarryMaterial) Ptr() *PromotionListV30FilteringHasCarryMaterial {
	return &v
}
