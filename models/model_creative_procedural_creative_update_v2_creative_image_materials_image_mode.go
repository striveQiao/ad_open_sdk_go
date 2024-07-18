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

// CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode
type CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode string

// List of creative_procedural_creative_update_v2_creative_image_materials_image_mode
const (
	CREATIVE_IMAGE_MODE_VIDEO_CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode                         CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE_CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode               CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE"
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode                CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
	CREATIVE_IMAGE_MODE_GIF_CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode                           CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_GIF"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode                  CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode                CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
	CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL_CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode             CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL"
	CREATIVE_IMAGE_MODE_DECORATION_COUPON_CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode             CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_DECORATION_COUPON"
	CREATIVE_IMAGE_MODE_SMALL_CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode                         CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_SMALL"
	CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL_CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode           CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE_CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO_CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode        CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO"
	CREATIVE_IMAGE_MODE_AWEME_LIVE_CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode                    CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_AWEME_LIVE"
	TOUTIAO_SEARCH_AD_IMAGE_CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode                           CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode = "TOUTIAO_SEARCH_AD_IMAGE"
	CREATIVE_IMAGE_MODE_GROUP_CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode                         CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_GROUP"
	CREATIVE_IMAGE_MODE_DISPLAY_WINDOW_CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode                CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_DISPLAY_WINDOW"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO_CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode            CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO"
	CREATIVE_IMAGE_MODE_LARGE_CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode                         CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_LARGE"
	SEARCH_AD_SMALL_IMAGE_CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode                             CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode = "SEARCH_AD_SMALL_IMAGE"
)

// All allowed values of CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode enum
var AllowedCreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageModeEnumValues = []CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode{
	"CREATIVE_IMAGE_MODE_VIDEO",
	"CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE",
	"CREATIVE_IMAGE_MODE_LARGE_VERTICAL",
	"CREATIVE_IMAGE_MODE_GIF",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH",
	"CREATIVE_IMAGE_MODE_VIDEO_VERTICAL",
	"CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL",
	"CREATIVE_IMAGE_MODE_DECORATION_COUPON",
	"CREATIVE_IMAGE_MODE_SMALL",
	"CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL",
	"CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE",
	"CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO",
	"CREATIVE_IMAGE_MODE_AWEME_LIVE",
	"TOUTIAO_SEARCH_AD_IMAGE",
	"CREATIVE_IMAGE_MODE_GROUP",
	"CREATIVE_IMAGE_MODE_DISPLAY_WINDOW",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO",
	"CREATIVE_IMAGE_MODE_LARGE",
	"SEARCH_AD_SMALL_IMAGE",
}

// NewCreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageModeFromValue returns a pointer to a valid CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageModeFromValue(v string) (*CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode, error) {
	ev := CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode: valid values are %v", v, AllowedCreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode) IsValid() bool {
	for _, existing := range AllowedCreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_procedural_creative_update_v2_creative_image_materials_image_mode value
func (v CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode) Ptr() *CreativeProceduralCreativeUpdateV2CreativeImageMaterialsImageMode {
	return &v
}
