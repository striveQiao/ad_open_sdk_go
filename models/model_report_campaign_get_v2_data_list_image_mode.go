/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportCampaignGetV2DataListImageMode
type ReportCampaignGetV2DataListImageMode string

// List of report_campaign_get_v2_data_list_image_mode
const (
	CREATIVE_IMAGE_MODE_VIDEO_ReportCampaignGetV2DataListImageMode               ReportCampaignGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO_ReportCampaignGetV2DataListImageMode  ReportCampaignGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO"
	CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL_ReportCampaignGetV2DataListImageMode ReportCampaignGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL"
	SEARCH_AD_SMALL_IMAGE_ReportCampaignGetV2DataListImageMode                   ReportCampaignGetV2DataListImageMode = "SEARCH_AD_SMALL_IMAGE"
	CREATIVE_IMAGE_MODE_AWEME_LIVE_ReportCampaignGetV2DataListImageMode          ReportCampaignGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_AWEME_LIVE"
	TOUTIAO_SEARCH_AD_IMAGE_ReportCampaignGetV2DataListImageMode                 ReportCampaignGetV2DataListImageMode = "TOUTIAO_SEARCH_AD_IMAGE"
	CREATIVE_IMAGE_MODE_SMALL_ReportCampaignGetV2DataListImageMode               ReportCampaignGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_SMALL"
	CREATIVE_IMAGE_MODE_LARGE_ReportCampaignGetV2DataListImageMode               ReportCampaignGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_LARGE"
	CREATIVE_IMAGE_MODE_GROUP_ReportCampaignGetV2DataListImageMode               ReportCampaignGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_GROUP"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_ReportCampaignGetV2DataListImageMode        ReportCampaignGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH"
	MATERIAL_IMAGE_MODE_TITLE_ReportCampaignGetV2DataListImageMode               ReportCampaignGetV2DataListImageMode = "MATERIAL_IMAGE_MODE_TITLE"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_ReportCampaignGetV2DataListImageMode      ReportCampaignGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
	CREATIVE_IMAGE_MODE_DISPLAY_WINDOW_ReportCampaignGetV2DataListImageMode      ReportCampaignGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_DISPLAY_WINDOW"
	CREATIVE_IMAGE_MODE_DECORATION_COUPON_ReportCampaignGetV2DataListImageMode   ReportCampaignGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_DECORATION_COUPON"
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_ReportCampaignGetV2DataListImageMode      ReportCampaignGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
	CREATIVE_IMAGE_MODE_GIF_ReportCampaignGetV2DataListImageMode                 ReportCampaignGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_GIF"
	CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL_ReportCampaignGetV2DataListImageMode   ReportCampaignGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL"
)

// All allowed values of ReportCampaignGetV2DataListImageMode enum
var AllowedReportCampaignGetV2DataListImageModeEnumValues = []ReportCampaignGetV2DataListImageMode{
	"CREATIVE_IMAGE_MODE_VIDEO",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO",
	"CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL",
	"SEARCH_AD_SMALL_IMAGE",
	"CREATIVE_IMAGE_MODE_AWEME_LIVE",
	"TOUTIAO_SEARCH_AD_IMAGE",
	"CREATIVE_IMAGE_MODE_SMALL",
	"CREATIVE_IMAGE_MODE_LARGE",
	"CREATIVE_IMAGE_MODE_GROUP",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH",
	"MATERIAL_IMAGE_MODE_TITLE",
	"CREATIVE_IMAGE_MODE_VIDEO_VERTICAL",
	"CREATIVE_IMAGE_MODE_DISPLAY_WINDOW",
	"CREATIVE_IMAGE_MODE_DECORATION_COUPON",
	"CREATIVE_IMAGE_MODE_LARGE_VERTICAL",
	"CREATIVE_IMAGE_MODE_GIF",
	"CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL",
}

// NewReportCampaignGetV2DataListImageModeFromValue returns a pointer to a valid ReportCampaignGetV2DataListImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCampaignGetV2DataListImageModeFromValue(v string) (*ReportCampaignGetV2DataListImageMode, error) {
	ev := ReportCampaignGetV2DataListImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCampaignGetV2DataListImageMode: valid values are %v", v, AllowedReportCampaignGetV2DataListImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCampaignGetV2DataListImageMode) IsValid() bool {
	for _, existing := range AllowedReportCampaignGetV2DataListImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_campaign_get_v2_data_list_image_mode value
func (v ReportCampaignGetV2DataListImageMode) Ptr() *ReportCampaignGetV2DataListImageMode {
	return &v
}
