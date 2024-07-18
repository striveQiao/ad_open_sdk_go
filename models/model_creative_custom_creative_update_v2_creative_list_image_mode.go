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

// CreativeCustomCreativeUpdateV2CreativeListImageMode
type CreativeCustomCreativeUpdateV2CreativeListImageMode string

// List of creative_custom_creative_update_v2_creative_list_image_mode
const (
	CREATIVE_IMAGE_MODE_VIDEO_CreativeCustomCreativeUpdateV2CreativeListImageMode                         CreativeCustomCreativeUpdateV2CreativeListImageMode = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE_CreativeCustomCreativeUpdateV2CreativeListImageMode               CreativeCustomCreativeUpdateV2CreativeListImageMode = "CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE"
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_CreativeCustomCreativeUpdateV2CreativeListImageMode                CreativeCustomCreativeUpdateV2CreativeListImageMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
	CREATIVE_IMAGE_MODE_GIF_CreativeCustomCreativeUpdateV2CreativeListImageMode                           CreativeCustomCreativeUpdateV2CreativeListImageMode = "CREATIVE_IMAGE_MODE_GIF"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_CreativeCustomCreativeUpdateV2CreativeListImageMode                  CreativeCustomCreativeUpdateV2CreativeListImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_CreativeCustomCreativeUpdateV2CreativeListImageMode                CreativeCustomCreativeUpdateV2CreativeListImageMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
	CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL_CreativeCustomCreativeUpdateV2CreativeListImageMode             CreativeCustomCreativeUpdateV2CreativeListImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL"
	CREATIVE_IMAGE_MODE_DECORATION_COUPON_CreativeCustomCreativeUpdateV2CreativeListImageMode             CreativeCustomCreativeUpdateV2CreativeListImageMode = "CREATIVE_IMAGE_MODE_DECORATION_COUPON"
	CREATIVE_IMAGE_MODE_SMALL_CreativeCustomCreativeUpdateV2CreativeListImageMode                         CreativeCustomCreativeUpdateV2CreativeListImageMode = "CREATIVE_IMAGE_MODE_SMALL"
	CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL_CreativeCustomCreativeUpdateV2CreativeListImageMode           CreativeCustomCreativeUpdateV2CreativeListImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE_CreativeCustomCreativeUpdateV2CreativeListImageMode CreativeCustomCreativeUpdateV2CreativeListImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO_CreativeCustomCreativeUpdateV2CreativeListImageMode        CreativeCustomCreativeUpdateV2CreativeListImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO"
	CREATIVE_IMAGE_MODE_AWEME_LIVE_CreativeCustomCreativeUpdateV2CreativeListImageMode                    CreativeCustomCreativeUpdateV2CreativeListImageMode = "CREATIVE_IMAGE_MODE_AWEME_LIVE"
	TOUTIAO_SEARCH_AD_IMAGE_CreativeCustomCreativeUpdateV2CreativeListImageMode                           CreativeCustomCreativeUpdateV2CreativeListImageMode = "TOUTIAO_SEARCH_AD_IMAGE"
	CREATIVE_IMAGE_MODE_GROUP_CreativeCustomCreativeUpdateV2CreativeListImageMode                         CreativeCustomCreativeUpdateV2CreativeListImageMode = "CREATIVE_IMAGE_MODE_GROUP"
	CREATIVE_IMAGE_MODE_DISPLAY_WINDOW_CreativeCustomCreativeUpdateV2CreativeListImageMode                CreativeCustomCreativeUpdateV2CreativeListImageMode = "CREATIVE_IMAGE_MODE_DISPLAY_WINDOW"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO_CreativeCustomCreativeUpdateV2CreativeListImageMode            CreativeCustomCreativeUpdateV2CreativeListImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO"
	CREATIVE_IMAGE_MODE_LARGE_CreativeCustomCreativeUpdateV2CreativeListImageMode                         CreativeCustomCreativeUpdateV2CreativeListImageMode = "CREATIVE_IMAGE_MODE_LARGE"
	SEARCH_AD_SMALL_IMAGE_CreativeCustomCreativeUpdateV2CreativeListImageMode                             CreativeCustomCreativeUpdateV2CreativeListImageMode = "SEARCH_AD_SMALL_IMAGE"
)

// All allowed values of CreativeCustomCreativeUpdateV2CreativeListImageMode enum
var AllowedCreativeCustomCreativeUpdateV2CreativeListImageModeEnumValues = []CreativeCustomCreativeUpdateV2CreativeListImageMode{
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

// NewCreativeCustomCreativeUpdateV2CreativeListImageModeFromValue returns a pointer to a valid CreativeCustomCreativeUpdateV2CreativeListImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeCustomCreativeUpdateV2CreativeListImageModeFromValue(v string) (*CreativeCustomCreativeUpdateV2CreativeListImageMode, error) {
	ev := CreativeCustomCreativeUpdateV2CreativeListImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeCustomCreativeUpdateV2CreativeListImageMode: valid values are %v", v, AllowedCreativeCustomCreativeUpdateV2CreativeListImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeCustomCreativeUpdateV2CreativeListImageMode) IsValid() bool {
	for _, existing := range AllowedCreativeCustomCreativeUpdateV2CreativeListImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_custom_creative_update_v2_creative_list_image_mode value
func (v CreativeCustomCreativeUpdateV2CreativeListImageMode) Ptr() *CreativeCustomCreativeUpdateV2CreativeListImageMode {
	return &v
}
