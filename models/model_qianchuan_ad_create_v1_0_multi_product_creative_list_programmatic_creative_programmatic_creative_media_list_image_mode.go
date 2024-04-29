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

// QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode
type QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode string

// List of qianchuan_ad_create_v1.0_multi_product_creative_list_programmatic_creative_programmatic_creative_media_list_image_mode
const (
	AWEME_LIVE_ROOM_QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode = "AWEME_LIVE_ROOM"
	CAROUSEL_QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode        QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode = "CAROUSEL"
	LARGE_QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode           QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode = "LARGE"
	LARGE_VERTICAL_QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode  QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode = "LARGE_VERTICAL"
	SMALL_QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode           QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode = "SMALL"
	SQUARE_QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode          QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode = "SQUARE"
	UNION_SPLASH_QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode    QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode = "UNION_SPLASH"
	VIDEO_LARGE_QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode     QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode = "VIDEO_LARGE"
	VIDEO_VERTICAL_QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode  QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode = "VIDEO_VERTICAL"
)

// All allowed values of QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode enum
var AllowedQianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageModeEnumValues = []QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode{
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

// NewQianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageModeFromValue returns a pointer to a valid QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageModeFromValue(v string) (*QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode, error) {
	ev := QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode: valid values are %v", v, AllowedQianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_multi_product_creative_list_programmatic_creative_programmatic_creative_media_list_image_mode value
func (v QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode) Ptr() *QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode {
	return &v
}
