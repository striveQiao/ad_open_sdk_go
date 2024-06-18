/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanImageGetV10DataListImageMode
type QianchuanImageGetV10DataListImageMode string

// List of qianchuan_image_get_v1.0_data_list_image_mode
const (
	LARGE_QianchuanImageGetV10DataListImageMode          QianchuanImageGetV10DataListImageMode = "LARGE"
	LARGE_VERTICAL_QianchuanImageGetV10DataListImageMode QianchuanImageGetV10DataListImageMode = "LARGE_VERTICAL"
	SMALL_QianchuanImageGetV10DataListImageMode          QianchuanImageGetV10DataListImageMode = "SMALL"
	SQUARE_QianchuanImageGetV10DataListImageMode         QianchuanImageGetV10DataListImageMode = "SQUARE"
	UNION_SPLASH_QianchuanImageGetV10DataListImageMode   QianchuanImageGetV10DataListImageMode = "UNION_SPLASH"
	VIDEO_LARGE_QianchuanImageGetV10DataListImageMode    QianchuanImageGetV10DataListImageMode = "VIDEO_LARGE"
	VIDEO_VERTICAL_QianchuanImageGetV10DataListImageMode QianchuanImageGetV10DataListImageMode = "VIDEO_VERTICAL"
)

// All allowed values of QianchuanImageGetV10DataListImageMode enum
var AllowedQianchuanImageGetV10DataListImageModeEnumValues = []QianchuanImageGetV10DataListImageMode{
	"LARGE",
	"LARGE_VERTICAL",
	"SMALL",
	"SQUARE",
	"UNION_SPLASH",
	"VIDEO_LARGE",
	"VIDEO_VERTICAL",
}

// NewQianchuanImageGetV10DataListImageModeFromValue returns a pointer to a valid QianchuanImageGetV10DataListImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanImageGetV10DataListImageModeFromValue(v string) (*QianchuanImageGetV10DataListImageMode, error) {
	ev := QianchuanImageGetV10DataListImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanImageGetV10DataListImageMode: valid values are %v", v, AllowedQianchuanImageGetV10DataListImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanImageGetV10DataListImageMode) IsValid() bool {
	for _, existing := range AllowedQianchuanImageGetV10DataListImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_image_get_v1.0_data_list_image_mode value
func (v QianchuanImageGetV10DataListImageMode) Ptr() *QianchuanImageGetV10DataListImageMode {
	return &v
}
