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

// NativeAnchorGetV30FilteringLandingType
type NativeAnchorGetV30FilteringLandingType string

// List of native_anchor_get_v3.0_filtering_landing_type
const (
	APP_NativeAnchorGetV30FilteringLandingType            NativeAnchorGetV30FilteringLandingType = "APP"
	ARTICLE_NativeAnchorGetV30FilteringLandingType        NativeAnchorGetV30FilteringLandingType = "ARTICLE"
	AWEME_NativeAnchorGetV30FilteringLandingType          NativeAnchorGetV30FilteringLandingType = "AWEME"
	BRAND_EXTERNAL_NativeAnchorGetV30FilteringLandingType NativeAnchorGetV30FilteringLandingType = "BRAND_EXTERNAL"
	DPA_NativeAnchorGetV30FilteringLandingType            NativeAnchorGetV30FilteringLandingType = "DPA"
	GOODS_NativeAnchorGetV30FilteringLandingType          NativeAnchorGetV30FilteringLandingType = "GOODS"
	LINK_NativeAnchorGetV30FilteringLandingType           NativeAnchorGetV30FilteringLandingType = "LINK"
	LIVE_NativeAnchorGetV30FilteringLandingType           NativeAnchorGetV30FilteringLandingType = "LIVE"
	MICRO_GAME_NativeAnchorGetV30FilteringLandingType     NativeAnchorGetV30FilteringLandingType = "MICRO_GAME"
	QUICK_APP_NativeAnchorGetV30FilteringLandingType      NativeAnchorGetV30FilteringLandingType = "QUICK_APP"
	SHOP_NativeAnchorGetV30FilteringLandingType           NativeAnchorGetV30FilteringLandingType = "SHOP"
	STORE_NativeAnchorGetV30FilteringLandingType          NativeAnchorGetV30FilteringLandingType = "STORE"
)

// All allowed values of NativeAnchorGetV30FilteringLandingType enum
var AllowedNativeAnchorGetV30FilteringLandingTypeEnumValues = []NativeAnchorGetV30FilteringLandingType{
	"APP",
	"ARTICLE",
	"AWEME",
	"BRAND_EXTERNAL",
	"DPA",
	"GOODS",
	"LINK",
	"LIVE",
	"MICRO_GAME",
	"QUICK_APP",
	"SHOP",
	"STORE",
}

// NewNativeAnchorGetV30FilteringLandingTypeFromValue returns a pointer to a valid NativeAnchorGetV30FilteringLandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorGetV30FilteringLandingTypeFromValue(v string) (*NativeAnchorGetV30FilteringLandingType, error) {
	ev := NativeAnchorGetV30FilteringLandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorGetV30FilteringLandingType: valid values are %v", v, AllowedNativeAnchorGetV30FilteringLandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorGetV30FilteringLandingType) IsValid() bool {
	for _, existing := range AllowedNativeAnchorGetV30FilteringLandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_get_v3.0_filtering_landing_type value
func (v NativeAnchorGetV30FilteringLandingType) Ptr() *NativeAnchorGetV30FilteringLandingType {
	return &v
}
