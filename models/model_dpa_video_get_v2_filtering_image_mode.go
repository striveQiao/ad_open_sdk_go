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

// DpaVideoGetV2FilteringImageMode
type DpaVideoGetV2FilteringImageMode string

// List of dpa_video_get_v2_filtering_image_mode
const (
	CREATIVE_IMAGE_MODE_VIDEO_DpaVideoGetV2FilteringImageMode                         DpaVideoGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE_DpaVideoGetV2FilteringImageMode               DpaVideoGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE"
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_DpaVideoGetV2FilteringImageMode                DpaVideoGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
	CREATIVE_IMAGE_MODE_GIF_DpaVideoGetV2FilteringImageMode                           DpaVideoGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_GIF"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_DpaVideoGetV2FilteringImageMode                  DpaVideoGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_DpaVideoGetV2FilteringImageMode                DpaVideoGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
	CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL_DpaVideoGetV2FilteringImageMode             DpaVideoGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL"
	CREATIVE_IMAGE_MODE_DECORATION_COUPON_DpaVideoGetV2FilteringImageMode             DpaVideoGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_DECORATION_COUPON"
	CREATIVE_IMAGE_MODE_SMALL_DpaVideoGetV2FilteringImageMode                         DpaVideoGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_SMALL"
	CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL_DpaVideoGetV2FilteringImageMode           DpaVideoGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE_DpaVideoGetV2FilteringImageMode DpaVideoGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO_DpaVideoGetV2FilteringImageMode        DpaVideoGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO"
	CREATIVE_IMAGE_MODE_AWEME_LIVE_DpaVideoGetV2FilteringImageMode                    DpaVideoGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_AWEME_LIVE"
	TOUTIAO_SEARCH_AD_IMAGE_DpaVideoGetV2FilteringImageMode                           DpaVideoGetV2FilteringImageMode = "TOUTIAO_SEARCH_AD_IMAGE"
	CREATIVE_IMAGE_MODE_GROUP_DpaVideoGetV2FilteringImageMode                         DpaVideoGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_GROUP"
	CREATIVE_IMAGE_MODE_DISPLAY_WINDOW_DpaVideoGetV2FilteringImageMode                DpaVideoGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_DISPLAY_WINDOW"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO_DpaVideoGetV2FilteringImageMode            DpaVideoGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO"
	CREATIVE_IMAGE_MODE_LARGE_DpaVideoGetV2FilteringImageMode                         DpaVideoGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_LARGE"
	SEARCH_AD_SMALL_IMAGE_DpaVideoGetV2FilteringImageMode                             DpaVideoGetV2FilteringImageMode = "SEARCH_AD_SMALL_IMAGE"
)

// All allowed values of DpaVideoGetV2FilteringImageMode enum
var AllowedDpaVideoGetV2FilteringImageModeEnumValues = []DpaVideoGetV2FilteringImageMode{
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

// NewDpaVideoGetV2FilteringImageModeFromValue returns a pointer to a valid DpaVideoGetV2FilteringImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaVideoGetV2FilteringImageModeFromValue(v string) (*DpaVideoGetV2FilteringImageMode, error) {
	ev := DpaVideoGetV2FilteringImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaVideoGetV2FilteringImageMode: valid values are %v", v, AllowedDpaVideoGetV2FilteringImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaVideoGetV2FilteringImageMode) IsValid() bool {
	for _, existing := range AllowedDpaVideoGetV2FilteringImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_video_get_v2_filtering_image_mode value
func (v DpaVideoGetV2FilteringImageMode) Ptr() *DpaVideoGetV2FilteringImageMode {
	return &v
}
