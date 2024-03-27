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

// PromotionUpdateV30PromotionMaterialsImageMaterialListImageMode
type PromotionUpdateV30PromotionMaterialsImageMaterialListImageMode string

// List of promotion_update_v3.0_promotion_materials_image_material_list_image_mode
const (
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE_PromotionUpdateV30PromotionMaterialsImageMaterialListImageMode PromotionUpdateV30PromotionMaterialsImageMaterialListImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE"
	CREATIVE_IMAGE_MODE_LARGE_PromotionUpdateV30PromotionMaterialsImageMaterialListImageMode                         PromotionUpdateV30PromotionMaterialsImageMaterialListImageMode = "CREATIVE_IMAGE_MODE_LARGE"
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_PromotionUpdateV30PromotionMaterialsImageMaterialListImageMode                PromotionUpdateV30PromotionMaterialsImageMaterialListImageMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
	CREATIVE_IMAGE_MODE_SMALL_PromotionUpdateV30PromotionMaterialsImageMaterialListImageMode                         PromotionUpdateV30PromotionMaterialsImageMaterialListImageMode = "CREATIVE_IMAGE_MODE_SMALL"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_PromotionUpdateV30PromotionMaterialsImageMaterialListImageMode                  PromotionUpdateV30PromotionMaterialsImageMaterialListImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH"
	SEARCH_AD_SMALL_IMAGE_PromotionUpdateV30PromotionMaterialsImageMaterialListImageMode                             PromotionUpdateV30PromotionMaterialsImageMaterialListImageMode = "SEARCH_AD_SMALL_IMAGE"
	TOUTIAO_SEARCH_AD_IMAGE_PromotionUpdateV30PromotionMaterialsImageMaterialListImageMode                           PromotionUpdateV30PromotionMaterialsImageMaterialListImageMode = "TOUTIAO_SEARCH_AD_IMAGE"
)

// All allowed values of PromotionUpdateV30PromotionMaterialsImageMaterialListImageMode enum
var AllowedPromotionUpdateV30PromotionMaterialsImageMaterialListImageModeEnumValues = []PromotionUpdateV30PromotionMaterialsImageMaterialListImageMode{
	"CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE",
	"CREATIVE_IMAGE_MODE_LARGE",
	"CREATIVE_IMAGE_MODE_LARGE_VERTICAL",
	"CREATIVE_IMAGE_MODE_SMALL",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH",
	"SEARCH_AD_SMALL_IMAGE",
	"TOUTIAO_SEARCH_AD_IMAGE",
}

// NewPromotionUpdateV30PromotionMaterialsImageMaterialListImageModeFromValue returns a pointer to a valid PromotionUpdateV30PromotionMaterialsImageMaterialListImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionUpdateV30PromotionMaterialsImageMaterialListImageModeFromValue(v string) (*PromotionUpdateV30PromotionMaterialsImageMaterialListImageMode, error) {
	ev := PromotionUpdateV30PromotionMaterialsImageMaterialListImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionUpdateV30PromotionMaterialsImageMaterialListImageMode: valid values are %v", v, AllowedPromotionUpdateV30PromotionMaterialsImageMaterialListImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionUpdateV30PromotionMaterialsImageMaterialListImageMode) IsValid() bool {
	for _, existing := range AllowedPromotionUpdateV30PromotionMaterialsImageMaterialListImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_update_v3.0_promotion_materials_image_material_list_image_mode value
func (v PromotionUpdateV30PromotionMaterialsImageMaterialListImageMode) Ptr() *PromotionUpdateV30PromotionMaterialsImageMaterialListImageMode {
	return &v
}
