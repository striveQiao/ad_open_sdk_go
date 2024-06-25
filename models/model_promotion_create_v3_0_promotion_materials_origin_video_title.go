/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionCreateV30PromotionMaterialsOriginVideoTitle
type PromotionCreateV30PromotionMaterialsOriginVideoTitle string

// List of promotion_create_v3.0_promotion_materials_origin_video_title
const (
	OFF_PromotionCreateV30PromotionMaterialsOriginVideoTitle PromotionCreateV30PromotionMaterialsOriginVideoTitle = "OFF"
	ON_PromotionCreateV30PromotionMaterialsOriginVideoTitle  PromotionCreateV30PromotionMaterialsOriginVideoTitle = "ON"
)

// All allowed values of PromotionCreateV30PromotionMaterialsOriginVideoTitle enum
var AllowedPromotionCreateV30PromotionMaterialsOriginVideoTitleEnumValues = []PromotionCreateV30PromotionMaterialsOriginVideoTitle{
	"OFF",
	"ON",
}

// NewPromotionCreateV30PromotionMaterialsOriginVideoTitleFromValue returns a pointer to a valid PromotionCreateV30PromotionMaterialsOriginVideoTitle
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCreateV30PromotionMaterialsOriginVideoTitleFromValue(v string) (*PromotionCreateV30PromotionMaterialsOriginVideoTitle, error) {
	ev := PromotionCreateV30PromotionMaterialsOriginVideoTitle(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCreateV30PromotionMaterialsOriginVideoTitle: valid values are %v", v, AllowedPromotionCreateV30PromotionMaterialsOriginVideoTitleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCreateV30PromotionMaterialsOriginVideoTitle) IsValid() bool {
	for _, existing := range AllowedPromotionCreateV30PromotionMaterialsOriginVideoTitleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_create_v3.0_promotion_materials_origin_video_title value
func (v PromotionCreateV30PromotionMaterialsOriginVideoTitle) Ptr() *PromotionCreateV30PromotionMaterialsOriginVideoTitle {
	return &v
}
