/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode
type PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode string

// List of promotion_list_v3.0_data_list_promotion_materials_image_material_list_image_mode
const (
	CREATIVE_IMAGE_MODE_AWEME_LIVE_PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode                    PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode = "CREATIVE_IMAGE_MODE_AWEME_LIVE"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE_PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO_PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode        PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO"
	CREATIVE_IMAGE_MODE_DISPLAY_WINDOW_PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode                PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode = "CREATIVE_IMAGE_MODE_DISPLAY_WINDOW"
	CREATIVE_IMAGE_MODE_GIF_PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode                           PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode = "CREATIVE_IMAGE_MODE_GIF"
	CREATIVE_IMAGE_MODE_GROUP_PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode                         PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode = "CREATIVE_IMAGE_MODE_GROUP"
	CREATIVE_IMAGE_MODE_LARGE_PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode                         PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode = "CREATIVE_IMAGE_MODE_LARGE"
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode                PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
	CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL_PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode           PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL"
	CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL_PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode             PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL"
	CREATIVE_IMAGE_MODE_SMALL_PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode                         PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode = "CREATIVE_IMAGE_MODE_SMALL"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode                  PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO_PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode            PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO"
	CREATIVE_IMAGE_MODE_VIDEO_PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode                         PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode                PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
	SEARCH_AD_SMALL_IMAGE_PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode                             PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode = "SEARCH_AD_SMALL_IMAGE"
	TOUTIAO_SEARCH_AD_IMAGE_PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode                           PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode = "TOUTIAO_SEARCH_AD_IMAGE"
)

// All allowed values of PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode enum
var AllowedPromotionListV30DataListPromotionMaterialsImageMaterialListImageModeEnumValues = []PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode{
	"CREATIVE_IMAGE_MODE_AWEME_LIVE",
	"CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE",
	"CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO",
	"CREATIVE_IMAGE_MODE_DISPLAY_WINDOW",
	"CREATIVE_IMAGE_MODE_GIF",
	"CREATIVE_IMAGE_MODE_GROUP",
	"CREATIVE_IMAGE_MODE_LARGE",
	"CREATIVE_IMAGE_MODE_LARGE_VERTICAL",
	"CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL",
	"CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL",
	"CREATIVE_IMAGE_MODE_SMALL",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO",
	"CREATIVE_IMAGE_MODE_VIDEO",
	"CREATIVE_IMAGE_MODE_VIDEO_VERTICAL",
	"SEARCH_AD_SMALL_IMAGE",
	"TOUTIAO_SEARCH_AD_IMAGE",
}

// NewPromotionListV30DataListPromotionMaterialsImageMaterialListImageModeFromValue returns a pointer to a valid PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListPromotionMaterialsImageMaterialListImageModeFromValue(v string) (*PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode, error) {
	ev := PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode: valid values are %v", v, AllowedPromotionListV30DataListPromotionMaterialsImageMaterialListImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListPromotionMaterialsImageMaterialListImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_promotion_materials_image_material_list_image_mode value
func (v PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode) Ptr() *PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode {
	return &v
}
