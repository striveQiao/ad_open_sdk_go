/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionCreateV30PromotionMaterialsDecorationMaterialImageMode
type PromotionCreateV30PromotionMaterialsDecorationMaterialImageMode string

// List of promotion_create_v3.0_promotion_materials_decoration_material_image_mode
const (
	CREATIVE_IMAGE_MODE_DECORATION_COUPON_PromotionCreateV30PromotionMaterialsDecorationMaterialImageMode PromotionCreateV30PromotionMaterialsDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_DECORATION_COUPON"
)

// All allowed values of PromotionCreateV30PromotionMaterialsDecorationMaterialImageMode enum
var AllowedPromotionCreateV30PromotionMaterialsDecorationMaterialImageModeEnumValues = []PromotionCreateV30PromotionMaterialsDecorationMaterialImageMode{
	"CREATIVE_IMAGE_MODE_DECORATION_COUPON",
}

// NewPromotionCreateV30PromotionMaterialsDecorationMaterialImageModeFromValue returns a pointer to a valid PromotionCreateV30PromotionMaterialsDecorationMaterialImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCreateV30PromotionMaterialsDecorationMaterialImageModeFromValue(v string) (*PromotionCreateV30PromotionMaterialsDecorationMaterialImageMode, error) {
	ev := PromotionCreateV30PromotionMaterialsDecorationMaterialImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCreateV30PromotionMaterialsDecorationMaterialImageMode: valid values are %v", v, AllowedPromotionCreateV30PromotionMaterialsDecorationMaterialImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCreateV30PromotionMaterialsDecorationMaterialImageMode) IsValid() bool {
	for _, existing := range AllowedPromotionCreateV30PromotionMaterialsDecorationMaterialImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_create_v3.0_promotion_materials_decoration_material_image_mode value
func (v PromotionCreateV30PromotionMaterialsDecorationMaterialImageMode) Ptr() *PromotionCreateV30PromotionMaterialsDecorationMaterialImageMode {
	return &v
}
