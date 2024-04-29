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

// CreativeTemplateDetailGetV2TemplateType
type CreativeTemplateDetailGetV2TemplateType string

// List of creative_template_detail_get_v2_template_type
const (
	HORIZONTAL_TO_VERTICAL_DYNAMIC_BACK_VIDEO_CreativeTemplateDetailGetV2TemplateType CreativeTemplateDetailGetV2TemplateType = "HORIZONTAL_TO_VERTICAL_DYNAMIC_BACK_VIDEO"
	VERTICAL_TO_HORIZONTAL_DYNAMIC_BACK_VIDEO_CreativeTemplateDetailGetV2TemplateType CreativeTemplateDetailGetV2TemplateType = "VERTICAL_TO_HORIZONTAL_DYNAMIC_BACK_VIDEO"
	VERTICAL_TO_HORIZONTAL_STATIC_BACK_VIDEO_CreativeTemplateDetailGetV2TemplateType  CreativeTemplateDetailGetV2TemplateType = "VERTICAL_TO_HORIZONTAL_STATIC_BACK_VIDEO"
)

// All allowed values of CreativeTemplateDetailGetV2TemplateType enum
var AllowedCreativeTemplateDetailGetV2TemplateTypeEnumValues = []CreativeTemplateDetailGetV2TemplateType{
	"HORIZONTAL_TO_VERTICAL_DYNAMIC_BACK_VIDEO",
	"VERTICAL_TO_HORIZONTAL_DYNAMIC_BACK_VIDEO",
	"VERTICAL_TO_HORIZONTAL_STATIC_BACK_VIDEO",
}

// NewCreativeTemplateDetailGetV2TemplateTypeFromValue returns a pointer to a valid CreativeTemplateDetailGetV2TemplateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeTemplateDetailGetV2TemplateTypeFromValue(v string) (*CreativeTemplateDetailGetV2TemplateType, error) {
	ev := CreativeTemplateDetailGetV2TemplateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeTemplateDetailGetV2TemplateType: valid values are %v", v, AllowedCreativeTemplateDetailGetV2TemplateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeTemplateDetailGetV2TemplateType) IsValid() bool {
	for _, existing := range AllowedCreativeTemplateDetailGetV2TemplateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_template_detail_get_v2_template_type value
func (v CreativeTemplateDetailGetV2TemplateType) Ptr() *CreativeTemplateDetailGetV2TemplateType {
	return &v
}
