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

// CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode
type CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode string

// List of creative_procedural_creative_create_v2_creative_image_materials_image_mode
const (
	CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL_CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode             CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL"
	SEARCH_AD_SMALL_IMAGE_CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode                             CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode = "SEARCH_AD_SMALL_IMAGE"
	CREATIVE_IMAGE_MODE_SMALL_CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode                         CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_SMALL"
	TOUTIAO_SEARCH_AD_IMAGE_CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode                           CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode = "TOUTIAO_SEARCH_AD_IMAGE"
	CREATIVE_IMAGE_MODE_GIF_CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode                           CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_GIF"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO_CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode        CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO_CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode            CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO"
	CREATIVE_IMAGE_MODE_AWEME_LIVE_CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode                    CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_AWEME_LIVE"
	CREATIVE_IMAGE_MODE_GROUP_CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode                         CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_GROUP"
	CREATIVE_IMAGE_MODE_DISPLAY_WINDOW_CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode                CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_DISPLAY_WINDOW"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode                  CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH"
	CREATIVE_IMAGE_MODE_DECORATION_COUPON_CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode             CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_DECORATION_COUPON"
	CREATIVE_IMAGE_MODE_VIDEO_CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode                         CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_LARGE_CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode                         CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_LARGE"
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode                CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE_CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode                CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
	CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE_CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode               CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE"
	CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL_CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode           CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL"
)

// All allowed values of CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode enum
var AllowedCreativeProceduralCreativeCreateV2CreativeImageMaterialsImageModeEnumValues = []CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode{
	"CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL",
	"SEARCH_AD_SMALL_IMAGE",
	"CREATIVE_IMAGE_MODE_SMALL",
	"TOUTIAO_SEARCH_AD_IMAGE",
	"CREATIVE_IMAGE_MODE_GIF",
	"CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO",
	"CREATIVE_IMAGE_MODE_AWEME_LIVE",
	"CREATIVE_IMAGE_MODE_GROUP",
	"CREATIVE_IMAGE_MODE_DISPLAY_WINDOW",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH",
	"CREATIVE_IMAGE_MODE_DECORATION_COUPON",
	"CREATIVE_IMAGE_MODE_VIDEO",
	"CREATIVE_IMAGE_MODE_LARGE",
	"CREATIVE_IMAGE_MODE_LARGE_VERTICAL",
	"CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE",
	"CREATIVE_IMAGE_MODE_VIDEO_VERTICAL",
	"CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE",
	"CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL",
}

// NewCreativeProceduralCreativeCreateV2CreativeImageMaterialsImageModeFromValue returns a pointer to a valid CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeProceduralCreativeCreateV2CreativeImageMaterialsImageModeFromValue(v string) (*CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode, error) {
	ev := CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode: valid values are %v", v, AllowedCreativeProceduralCreativeCreateV2CreativeImageMaterialsImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode) IsValid() bool {
	for _, existing := range AllowedCreativeProceduralCreativeCreateV2CreativeImageMaterialsImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_procedural_creative_create_v2_creative_image_materials_image_mode value
func (v CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode) Ptr() *CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode {
	return &v
}
