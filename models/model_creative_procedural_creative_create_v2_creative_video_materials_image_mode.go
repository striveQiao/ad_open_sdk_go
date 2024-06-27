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

// CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode
type CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode string

// List of creative_procedural_creative_create_v2_creative_video_materials_image_mode
const (
	CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE_CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode               CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE"
	CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL_CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode             CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL"
	TOUTIAO_SEARCH_AD_IMAGE_CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode                           CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode = "TOUTIAO_SEARCH_AD_IMAGE"
	CREATIVE_IMAGE_MODE_GIF_CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode                           CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_GIF"
	CREATIVE_IMAGE_MODE_DECORATION_COUPON_CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode             CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_DECORATION_COUPON"
	CREATIVE_IMAGE_MODE_GROUP_CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode                         CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_GROUP"
	CREATIVE_IMAGE_MODE_LARGE_CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode                         CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_LARGE"
	CREATIVE_IMAGE_MODE_VIDEO_CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode                         CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode                CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode                CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO_CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode        CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO"
	SEARCH_AD_SMALL_IMAGE_CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode                             CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode = "SEARCH_AD_SMALL_IMAGE"
	CREATIVE_IMAGE_MODE_DISPLAY_WINDOW_CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode                CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_DISPLAY_WINDOW"
	CREATIVE_IMAGE_MODE_AWEME_LIVE_CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode                    CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_AWEME_LIVE"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO_CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode            CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO"
	CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL_CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode           CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL"
	CREATIVE_IMAGE_MODE_SMALL_CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode                         CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_SMALL"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode                  CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE_CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE"
)

// All allowed values of CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode enum
var AllowedCreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageModeEnumValues = []CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode{
	"CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE",
	"CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL",
	"TOUTIAO_SEARCH_AD_IMAGE",
	"CREATIVE_IMAGE_MODE_GIF",
	"CREATIVE_IMAGE_MODE_DECORATION_COUPON",
	"CREATIVE_IMAGE_MODE_GROUP",
	"CREATIVE_IMAGE_MODE_LARGE",
	"CREATIVE_IMAGE_MODE_VIDEO",
	"CREATIVE_IMAGE_MODE_VIDEO_VERTICAL",
	"CREATIVE_IMAGE_MODE_LARGE_VERTICAL",
	"CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO",
	"SEARCH_AD_SMALL_IMAGE",
	"CREATIVE_IMAGE_MODE_DISPLAY_WINDOW",
	"CREATIVE_IMAGE_MODE_AWEME_LIVE",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO",
	"CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL",
	"CREATIVE_IMAGE_MODE_SMALL",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH",
	"CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE",
}

// NewCreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageModeFromValue returns a pointer to a valid CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageModeFromValue(v string) (*CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode, error) {
	ev := CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode: valid values are %v", v, AllowedCreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode) IsValid() bool {
	for _, existing := range AllowedCreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_procedural_creative_create_v2_creative_video_materials_image_mode value
func (v CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode) Ptr() *CreativeProceduralCreativeCreateV2CreativeVideoMaterialsImageMode {
	return &v
}
