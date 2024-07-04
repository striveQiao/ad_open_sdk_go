/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// BrandCreativeCreateV30CreativeDisplayMode
type BrandCreativeCreateV30CreativeDisplayMode string

// List of brand_creative_create_v3.0_creative_display_mode
const (
	CTR_BrandCreativeCreateV30CreativeDisplayMode    BrandCreativeCreateV30CreativeDisplayMode = "CTR"
	RANDOM_BrandCreativeCreateV30CreativeDisplayMode BrandCreativeCreateV30CreativeDisplayMode = "RANDOM"
)

// All allowed values of BrandCreativeCreateV30CreativeDisplayMode enum
var AllowedBrandCreativeCreateV30CreativeDisplayModeEnumValues = []BrandCreativeCreateV30CreativeDisplayMode{
	"CTR",
	"RANDOM",
}

// NewBrandCreativeCreateV30CreativeDisplayModeFromValue returns a pointer to a valid BrandCreativeCreateV30CreativeDisplayMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandCreativeCreateV30CreativeDisplayModeFromValue(v string) (*BrandCreativeCreateV30CreativeDisplayMode, error) {
	ev := BrandCreativeCreateV30CreativeDisplayMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandCreativeCreateV30CreativeDisplayMode: valid values are %v", v, AllowedBrandCreativeCreateV30CreativeDisplayModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandCreativeCreateV30CreativeDisplayMode) IsValid() bool {
	for _, existing := range AllowedBrandCreativeCreateV30CreativeDisplayModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_creative_create_v3.0_creative_display_mode value
func (v BrandCreativeCreateV30CreativeDisplayMode) Ptr() *BrandCreativeCreateV30CreativeDisplayMode {
	return &v
}
