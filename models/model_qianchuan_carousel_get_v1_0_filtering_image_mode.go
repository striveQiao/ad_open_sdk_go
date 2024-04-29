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

// QianchuanCarouselGetV10FilteringImageMode
type QianchuanCarouselGetV10FilteringImageMode string

// List of qianchuan_carousel_get_v1.0_filtering_image_mode
const (
	CAROUSEL_QianchuanCarouselGetV10FilteringImageMode QianchuanCarouselGetV10FilteringImageMode = "CAROUSEL"
)

// All allowed values of QianchuanCarouselGetV10FilteringImageMode enum
var AllowedQianchuanCarouselGetV10FilteringImageModeEnumValues = []QianchuanCarouselGetV10FilteringImageMode{
	"CAROUSEL",
}

// NewQianchuanCarouselGetV10FilteringImageModeFromValue returns a pointer to a valid QianchuanCarouselGetV10FilteringImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCarouselGetV10FilteringImageModeFromValue(v string) (*QianchuanCarouselGetV10FilteringImageMode, error) {
	ev := QianchuanCarouselGetV10FilteringImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCarouselGetV10FilteringImageMode: valid values are %v", v, AllowedQianchuanCarouselGetV10FilteringImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCarouselGetV10FilteringImageMode) IsValid() bool {
	for _, existing := range AllowedQianchuanCarouselGetV10FilteringImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_carousel_get_v1.0_filtering_image_mode value
func (v QianchuanCarouselGetV10FilteringImageMode) Ptr() *QianchuanCarouselGetV10FilteringImageMode {
	return &v
}
