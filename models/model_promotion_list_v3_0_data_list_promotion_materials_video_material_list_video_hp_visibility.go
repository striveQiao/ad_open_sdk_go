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

// PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoHpVisibility
type PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoHpVisibility string

// List of promotion_list_v3.0_data_list_promotion_materials_video_material_list_video_hp_visibility
const (
	ALWAYS_VISIBLE_PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoHpVisibility         PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoHpVisibility = "ALWAYS_VISIBLE"
	HIDE_AFTER_END_DATE_PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoHpVisibility    PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoHpVisibility = "HIDE_AFTER_END_DATE"
	HIDE_AFTER_NO_PLAYBACK_PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoHpVisibility PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoHpVisibility = "HIDE_AFTER_NO_PLAYBACK"
	HIDE_VIDEO_ON_HP_PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoHpVisibility       PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoHpVisibility = "HIDE_VIDEO_ON_HP"
)

// All allowed values of PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoHpVisibility enum
var AllowedPromotionListV30DataListPromotionMaterialsVideoMaterialListVideoHpVisibilityEnumValues = []PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoHpVisibility{
	"ALWAYS_VISIBLE",
	"HIDE_AFTER_END_DATE",
	"HIDE_AFTER_NO_PLAYBACK",
	"HIDE_VIDEO_ON_HP",
}

// NewPromotionListV30DataListPromotionMaterialsVideoMaterialListVideoHpVisibilityFromValue returns a pointer to a valid PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoHpVisibility
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListPromotionMaterialsVideoMaterialListVideoHpVisibilityFromValue(v string) (*PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoHpVisibility, error) {
	ev := PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoHpVisibility(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoHpVisibility: valid values are %v", v, AllowedPromotionListV30DataListPromotionMaterialsVideoMaterialListVideoHpVisibilityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoHpVisibility) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListPromotionMaterialsVideoMaterialListVideoHpVisibilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_promotion_materials_video_material_list_video_hp_visibility value
func (v PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoHpVisibility) Ptr() *PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoHpVisibility {
	return &v
}
