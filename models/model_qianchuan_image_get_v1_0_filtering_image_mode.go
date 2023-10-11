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

// QianchuanImageGetV10FilteringImageMode
type QianchuanImageGetV10FilteringImageMode string

// List of qianchuan_image_get_v1.0_filtering_image_mode
const (
	LARGE_QianchuanImageGetV10FilteringImageMode          QianchuanImageGetV10FilteringImageMode = "LARGE"
	LARGE_VERTICAL_QianchuanImageGetV10FilteringImageMode QianchuanImageGetV10FilteringImageMode = "LARGE_VERTICAL"
	SMALL_QianchuanImageGetV10FilteringImageMode          QianchuanImageGetV10FilteringImageMode = "SMALL"
	SQUARE_QianchuanImageGetV10FilteringImageMode         QianchuanImageGetV10FilteringImageMode = "SQUARE"
	UNION_SPLASH_QianchuanImageGetV10FilteringImageMode   QianchuanImageGetV10FilteringImageMode = "UNION_SPLASH"
)

// All allowed values of QianchuanImageGetV10FilteringImageMode enum
var AllowedQianchuanImageGetV10FilteringImageModeEnumValues = []QianchuanImageGetV10FilteringImageMode{
	"LARGE",
	"LARGE_VERTICAL",
	"SMALL",
	"SQUARE",
	"UNION_SPLASH",
}

// NewQianchuanImageGetV10FilteringImageModeFromValue returns a pointer to a valid QianchuanImageGetV10FilteringImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanImageGetV10FilteringImageModeFromValue(v string) (*QianchuanImageGetV10FilteringImageMode, error) {
	ev := QianchuanImageGetV10FilteringImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanImageGetV10FilteringImageMode: valid values are %v", v, AllowedQianchuanImageGetV10FilteringImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanImageGetV10FilteringImageMode) IsValid() bool {
	for _, existing := range AllowedQianchuanImageGetV10FilteringImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_image_get_v1.0_filtering_image_mode value
func (v QianchuanImageGetV10FilteringImageMode) Ptr() *QianchuanImageGetV10FilteringImageMode {
	return &v
}
