/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoTemplateType
type PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoTemplateType string

// List of promotion_list_v3.0_data_list_promotion_materials_video_material_list_video_template_type
const (
	DPA_VIDEO_TEMPLATE_CUSTOM_PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoTemplateType     PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoTemplateType = "DPA_VIDEO_TEMPLATE_CUSTOM"
	DPA_VIDEO_TEMPLATE_DEPRECATED_PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoTemplateType PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoTemplateType = "DPA_VIDEO_TEMPLATE_DEPRECATED"
	DPA_VIDEO_TEMPLATE_SMART_PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoTemplateType      PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoTemplateType = "DPA_VIDEO_TEMPLATE_SMART"
)

// All allowed values of PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoTemplateType enum
var AllowedPromotionListV30DataListPromotionMaterialsVideoMaterialListVideoTemplateTypeEnumValues = []PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoTemplateType{
	"DPA_VIDEO_TEMPLATE_CUSTOM",
	"DPA_VIDEO_TEMPLATE_DEPRECATED",
	"DPA_VIDEO_TEMPLATE_SMART",
}

// NewPromotionListV30DataListPromotionMaterialsVideoMaterialListVideoTemplateTypeFromValue returns a pointer to a valid PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoTemplateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListPromotionMaterialsVideoMaterialListVideoTemplateTypeFromValue(v string) (*PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoTemplateType, error) {
	ev := PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoTemplateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoTemplateType: valid values are %v", v, AllowedPromotionListV30DataListPromotionMaterialsVideoMaterialListVideoTemplateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoTemplateType) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListPromotionMaterialsVideoMaterialListVideoTemplateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_promotion_materials_video_material_list_video_template_type value
func (v PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoTemplateType) Ptr() *PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoTemplateType {
	return &v
}
