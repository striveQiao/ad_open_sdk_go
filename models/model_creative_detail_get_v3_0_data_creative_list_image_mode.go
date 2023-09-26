/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeDetailGetV30DataCreativeListImageMode
type CreativeDetailGetV30DataCreativeListImageMode string

// List of creative_detail_get_v3.0_data_creative_list_image_mode
const (
	CREATIVE_IMAGE_MODE_AWEME_LIVE_CreativeDetailGetV30DataCreativeListImageMode                    CreativeDetailGetV30DataCreativeListImageMode = "CREATIVE_IMAGE_MODE_AWEME_LIVE"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE_CreativeDetailGetV30DataCreativeListImageMode CreativeDetailGetV30DataCreativeListImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO_CreativeDetailGetV30DataCreativeListImageMode        CreativeDetailGetV30DataCreativeListImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO"
	CREATIVE_IMAGE_MODE_DECORATION_COUPON_CreativeDetailGetV30DataCreativeListImageMode             CreativeDetailGetV30DataCreativeListImageMode = "CREATIVE_IMAGE_MODE_DECORATION_COUPON"
	CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE_CreativeDetailGetV30DataCreativeListImageMode               CreativeDetailGetV30DataCreativeListImageMode = "CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE"
	CREATIVE_IMAGE_MODE_DISPLAY_WINDOW_CreativeDetailGetV30DataCreativeListImageMode                CreativeDetailGetV30DataCreativeListImageMode = "CREATIVE_IMAGE_MODE_DISPLAY_WINDOW"
	CREATIVE_IMAGE_MODE_GIF_CreativeDetailGetV30DataCreativeListImageMode                           CreativeDetailGetV30DataCreativeListImageMode = "CREATIVE_IMAGE_MODE_GIF"
	CREATIVE_IMAGE_MODE_GROUP_CreativeDetailGetV30DataCreativeListImageMode                         CreativeDetailGetV30DataCreativeListImageMode = "CREATIVE_IMAGE_MODE_GROUP"
	CREATIVE_IMAGE_MODE_LARGE_CreativeDetailGetV30DataCreativeListImageMode                         CreativeDetailGetV30DataCreativeListImageMode = "CREATIVE_IMAGE_MODE_LARGE"
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_CreativeDetailGetV30DataCreativeListImageMode                CreativeDetailGetV30DataCreativeListImageMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
	CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL_CreativeDetailGetV30DataCreativeListImageMode           CreativeDetailGetV30DataCreativeListImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL"
	CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL_CreativeDetailGetV30DataCreativeListImageMode             CreativeDetailGetV30DataCreativeListImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL"
	CREATIVE_IMAGE_MODE_SMALL_CreativeDetailGetV30DataCreativeListImageMode                         CreativeDetailGetV30DataCreativeListImageMode = "CREATIVE_IMAGE_MODE_SMALL"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_CreativeDetailGetV30DataCreativeListImageMode                  CreativeDetailGetV30DataCreativeListImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO_CreativeDetailGetV30DataCreativeListImageMode            CreativeDetailGetV30DataCreativeListImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO"
	CREATIVE_IMAGE_MODE_VIDEO_CreativeDetailGetV30DataCreativeListImageMode                         CreativeDetailGetV30DataCreativeListImageMode = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_CreativeDetailGetV30DataCreativeListImageMode                CreativeDetailGetV30DataCreativeListImageMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
	SEARCH_AD_SMALL_IMAGE_CreativeDetailGetV30DataCreativeListImageMode                             CreativeDetailGetV30DataCreativeListImageMode = "SEARCH_AD_SMALL_IMAGE"
	TOUTIAO_SEARCH_AD_IMAGE_CreativeDetailGetV30DataCreativeListImageMode                           CreativeDetailGetV30DataCreativeListImageMode = "TOUTIAO_SEARCH_AD_IMAGE"
)

// All allowed values of CreativeDetailGetV30DataCreativeListImageMode enum
var AllowedCreativeDetailGetV30DataCreativeListImageModeEnumValues = []CreativeDetailGetV30DataCreativeListImageMode{
	"CREATIVE_IMAGE_MODE_AWEME_LIVE",
	"CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE",
	"CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO",
	"CREATIVE_IMAGE_MODE_DECORATION_COUPON",
	"CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE",
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

// NewCreativeDetailGetV30DataCreativeListImageModeFromValue returns a pointer to a valid CreativeDetailGetV30DataCreativeListImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeDetailGetV30DataCreativeListImageModeFromValue(v string) (*CreativeDetailGetV30DataCreativeListImageMode, error) {
	ev := CreativeDetailGetV30DataCreativeListImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeDetailGetV30DataCreativeListImageMode: valid values are %v", v, AllowedCreativeDetailGetV30DataCreativeListImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeDetailGetV30DataCreativeListImageMode) IsValid() bool {
	for _, existing := range AllowedCreativeDetailGetV30DataCreativeListImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_detail_get_v3.0_data_creative_list_image_mode value
func (v CreativeDetailGetV30DataCreativeListImageMode) Ptr() *CreativeDetailGetV30DataCreativeListImageMode {
	return &v
}
