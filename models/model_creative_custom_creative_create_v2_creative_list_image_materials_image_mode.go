/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode
type CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode string

// List of creative_custom_creative_create_v2_creative_list_image_materials_image_mode
const (
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO_CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode        CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO"
	SEARCH_AD_SMALL_IMAGE_CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode                             CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode = "SEARCH_AD_SMALL_IMAGE"
	TOUTIAO_SEARCH_AD_IMAGE_CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode                           CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode = "TOUTIAO_SEARCH_AD_IMAGE"
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode                CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
	CREATIVE_IMAGE_MODE_GIF_CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode                           CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_GIF"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO_CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode            CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO"
	CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL_CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode             CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL"
	CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE_CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode               CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE"
	CREATIVE_IMAGE_MODE_DISPLAY_WINDOW_CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode                CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_DISPLAY_WINDOW"
	CREATIVE_IMAGE_MODE_AWEME_LIVE_CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode                    CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_AWEME_LIVE"
	CREATIVE_IMAGE_MODE_VIDEO_CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode                         CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_GROUP_CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode                         CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_GROUP"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE_CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE"
	CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL_CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode           CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode                CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
	CREATIVE_IMAGE_MODE_DECORATION_COUPON_CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode             CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_DECORATION_COUPON"
	CREATIVE_IMAGE_MODE_LARGE_CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode                         CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_LARGE"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode                  CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH"
	CREATIVE_IMAGE_MODE_SMALL_CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode                         CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_SMALL"
)

// All allowed values of CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode enum
var AllowedCreativeCustomCreativeCreateV2CreativeListImageMaterialsImageModeEnumValues = []CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode{
	"CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO",
	"SEARCH_AD_SMALL_IMAGE",
	"TOUTIAO_SEARCH_AD_IMAGE",
	"CREATIVE_IMAGE_MODE_LARGE_VERTICAL",
	"CREATIVE_IMAGE_MODE_GIF",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO",
	"CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL",
	"CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE",
	"CREATIVE_IMAGE_MODE_DISPLAY_WINDOW",
	"CREATIVE_IMAGE_MODE_AWEME_LIVE",
	"CREATIVE_IMAGE_MODE_VIDEO",
	"CREATIVE_IMAGE_MODE_GROUP",
	"CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE",
	"CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL",
	"CREATIVE_IMAGE_MODE_VIDEO_VERTICAL",
	"CREATIVE_IMAGE_MODE_DECORATION_COUPON",
	"CREATIVE_IMAGE_MODE_LARGE",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH",
	"CREATIVE_IMAGE_MODE_SMALL",
}

// NewCreativeCustomCreativeCreateV2CreativeListImageMaterialsImageModeFromValue returns a pointer to a valid CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeCustomCreativeCreateV2CreativeListImageMaterialsImageModeFromValue(v string) (*CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode, error) {
	ev := CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode: valid values are %v", v, AllowedCreativeCustomCreativeCreateV2CreativeListImageMaterialsImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode) IsValid() bool {
	for _, existing := range AllowedCreativeCustomCreativeCreateV2CreativeListImageMaterialsImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_custom_creative_create_v2_creative_list_image_materials_image_mode value
func (v CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode) Ptr() *CreativeCustomCreativeCreateV2CreativeListImageMaterialsImageMode {
	return &v
}
