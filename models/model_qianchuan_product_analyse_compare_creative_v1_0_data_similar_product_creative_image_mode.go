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

// QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageMode
type QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageMode string

// List of qianchuan_product_analyse_compare_creative_v1.0_data_similar_product_creative_image_mode
const (
	AWEME_LIVE_ROOM_QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageMode QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageMode = "AWEME_LIVE_ROOM"
	LARGE_QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageMode           QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageMode = "LARGE"
	LARGE_VERTICAL_QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageMode  QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageMode = "LARGE_VERTICAL"
	SMALL_QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageMode           QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageMode = "SMALL"
	SQUARE_QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageMode          QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageMode = "SQUARE"
	UNION_SPLASH_QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageMode    QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageMode = "UNION_SPLASH"
	VIDEO_LARGE_QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageMode     QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageMode = "VIDEO_LARGE"
	VIDEO_VERTICAL_QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageMode  QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageMode = "VIDEO_VERTICAL"
)

// All allowed values of QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageMode enum
var AllowedQianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageModeEnumValues = []QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageMode{
	"AWEME_LIVE_ROOM",
	"LARGE",
	"LARGE_VERTICAL",
	"SMALL",
	"SQUARE",
	"UNION_SPLASH",
	"VIDEO_LARGE",
	"VIDEO_VERTICAL",
}

// NewQianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageModeFromValue returns a pointer to a valid QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageModeFromValue(v string) (*QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageMode, error) {
	ev := QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageMode: valid values are %v", v, AllowedQianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageMode) IsValid() bool {
	for _, existing := range AllowedQianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_product_analyse_compare_creative_v1.0_data_similar_product_creative_image_mode value
func (v QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageMode) Ptr() *QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageMode {
	return &v
}
