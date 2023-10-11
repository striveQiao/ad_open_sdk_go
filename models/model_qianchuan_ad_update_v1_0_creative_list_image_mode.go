/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdUpdateV10CreativeListImageMode
type QianchuanAdUpdateV10CreativeListImageMode string

// List of qianchuan_ad_update_v1.0_creative_list_image_mode
const (
	AWEME_LIVE_ROOM_QianchuanAdUpdateV10CreativeListImageMode QianchuanAdUpdateV10CreativeListImageMode = "AWEME_LIVE_ROOM"
	CAROUSEL_QianchuanAdUpdateV10CreativeListImageMode        QianchuanAdUpdateV10CreativeListImageMode = "CAROUSEL"
	LARGE_QianchuanAdUpdateV10CreativeListImageMode           QianchuanAdUpdateV10CreativeListImageMode = "LARGE"
	LARGE_VERTICAL_QianchuanAdUpdateV10CreativeListImageMode  QianchuanAdUpdateV10CreativeListImageMode = "LARGE_VERTICAL"
	SMALL_QianchuanAdUpdateV10CreativeListImageMode           QianchuanAdUpdateV10CreativeListImageMode = "SMALL"
	SQUARE_QianchuanAdUpdateV10CreativeListImageMode          QianchuanAdUpdateV10CreativeListImageMode = "SQUARE"
	UNION_SPLASH_QianchuanAdUpdateV10CreativeListImageMode    QianchuanAdUpdateV10CreativeListImageMode = "UNION_SPLASH"
	VIDEO_LARGE_QianchuanAdUpdateV10CreativeListImageMode     QianchuanAdUpdateV10CreativeListImageMode = "VIDEO_LARGE"
	VIDEO_VERTICAL_QianchuanAdUpdateV10CreativeListImageMode  QianchuanAdUpdateV10CreativeListImageMode = "VIDEO_VERTICAL"
)

// All allowed values of QianchuanAdUpdateV10CreativeListImageMode enum
var AllowedQianchuanAdUpdateV10CreativeListImageModeEnumValues = []QianchuanAdUpdateV10CreativeListImageMode{
	"AWEME_LIVE_ROOM",
	"CAROUSEL",
	"LARGE",
	"LARGE_VERTICAL",
	"SMALL",
	"SQUARE",
	"UNION_SPLASH",
	"VIDEO_LARGE",
	"VIDEO_VERTICAL",
}

// NewQianchuanAdUpdateV10CreativeListImageModeFromValue returns a pointer to a valid QianchuanAdUpdateV10CreativeListImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10CreativeListImageModeFromValue(v string) (*QianchuanAdUpdateV10CreativeListImageMode, error) {
	ev := QianchuanAdUpdateV10CreativeListImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10CreativeListImageMode: valid values are %v", v, AllowedQianchuanAdUpdateV10CreativeListImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10CreativeListImageMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10CreativeListImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_creative_list_image_mode value
func (v QianchuanAdUpdateV10CreativeListImageMode) Ptr() *QianchuanAdUpdateV10CreativeListImageMode {
	return &v
}
