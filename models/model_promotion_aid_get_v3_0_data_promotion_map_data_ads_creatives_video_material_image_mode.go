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

// PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode
type PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode string

// List of promotion_aid_get_v3.0_data_promotion_map_data_ads_creatives_video_material_image_mode
const (
	CREATIVE_IMAGE_MODE_AWEME_LIVE_PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode                    PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode = "CREATIVE_IMAGE_MODE_AWEME_LIVE"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE_PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO_PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode        PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO"
	CREATIVE_IMAGE_MODE_DECORATION_COUPON_PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode             PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode = "CREATIVE_IMAGE_MODE_DECORATION_COUPON"
	CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE_PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode               PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode = "CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE"
	CREATIVE_IMAGE_MODE_DISPLAY_WINDOW_PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode                PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode = "CREATIVE_IMAGE_MODE_DISPLAY_WINDOW"
	CREATIVE_IMAGE_MODE_GIF_PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode                           PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode = "CREATIVE_IMAGE_MODE_GIF"
	CREATIVE_IMAGE_MODE_GROUP_PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode                         PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode = "CREATIVE_IMAGE_MODE_GROUP"
	CREATIVE_IMAGE_MODE_LARGE_PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode                         PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode = "CREATIVE_IMAGE_MODE_LARGE"
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode                PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
	CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL_PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode           PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL"
	CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL_PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode             PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL"
	CREATIVE_IMAGE_MODE_SMALL_PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode                         PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode = "CREATIVE_IMAGE_MODE_SMALL"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode                  PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO_PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode            PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO"
	CREATIVE_IMAGE_MODE_VIDEO_PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode                         PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode                PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
	INVALID_DATA_PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode                                      PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode = "INVALID_DATA"
	SEARCH_AD_SMALL_IMAGE_PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode                             PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode = "SEARCH_AD_SMALL_IMAGE"
	TOUTIAO_SEARCH_AD_IMAGE_PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode                           PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode = "TOUTIAO_SEARCH_AD_IMAGE"
)

// All allowed values of PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode enum
var AllowedPromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageModeEnumValues = []PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode{
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
	"INVALID_DATA",
	"SEARCH_AD_SMALL_IMAGE",
	"TOUTIAO_SEARCH_AD_IMAGE",
}

// NewPromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageModeFromValue returns a pointer to a valid PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageModeFromValue(v string) (*PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode, error) {
	ev := PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode: valid values are %v", v, AllowedPromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode) IsValid() bool {
	for _, existing := range AllowedPromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_aid_get_v3.0_data_promotion_map_data_ads_creatives_video_material_image_mode value
func (v PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode) Ptr() *PromotionAidGetV30DataPromotionMapDataAdsCreativesVideoMaterialImageMode {
	return &v
}
