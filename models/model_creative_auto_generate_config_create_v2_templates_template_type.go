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

// CreativeAutoGenerateConfigCreateV2TemplatesTemplateType
type CreativeAutoGenerateConfigCreateV2TemplatesTemplateType string

// List of creative_auto_generate_config_create_v2_templates_template_type
const (
	HORIZONTAL_TO_VERTICAL_DYNAMIC_BACK_VIDEO_CreativeAutoGenerateConfigCreateV2TemplatesTemplateType CreativeAutoGenerateConfigCreateV2TemplatesTemplateType = "HORIZONTAL_TO_VERTICAL_DYNAMIC_BACK_VIDEO"
	VERTICAL_TO_HORIZONTAL_DYNAMIC_BACK_VIDEO_CreativeAutoGenerateConfigCreateV2TemplatesTemplateType CreativeAutoGenerateConfigCreateV2TemplatesTemplateType = "VERTICAL_TO_HORIZONTAL_DYNAMIC_BACK_VIDEO"
	VERTICAL_TO_HORIZONTAL_STATIC_BACK_VIDEO_CreativeAutoGenerateConfigCreateV2TemplatesTemplateType  CreativeAutoGenerateConfigCreateV2TemplatesTemplateType = "VERTICAL_TO_HORIZONTAL_STATIC_BACK_VIDEO"
)

// All allowed values of CreativeAutoGenerateConfigCreateV2TemplatesTemplateType enum
var AllowedCreativeAutoGenerateConfigCreateV2TemplatesTemplateTypeEnumValues = []CreativeAutoGenerateConfigCreateV2TemplatesTemplateType{
	"HORIZONTAL_TO_VERTICAL_DYNAMIC_BACK_VIDEO",
	"VERTICAL_TO_HORIZONTAL_DYNAMIC_BACK_VIDEO",
	"VERTICAL_TO_HORIZONTAL_STATIC_BACK_VIDEO",
}

// NewCreativeAutoGenerateConfigCreateV2TemplatesTemplateTypeFromValue returns a pointer to a valid CreativeAutoGenerateConfigCreateV2TemplatesTemplateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeAutoGenerateConfigCreateV2TemplatesTemplateTypeFromValue(v string) (*CreativeAutoGenerateConfigCreateV2TemplatesTemplateType, error) {
	ev := CreativeAutoGenerateConfigCreateV2TemplatesTemplateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeAutoGenerateConfigCreateV2TemplatesTemplateType: valid values are %v", v, AllowedCreativeAutoGenerateConfigCreateV2TemplatesTemplateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeAutoGenerateConfigCreateV2TemplatesTemplateType) IsValid() bool {
	for _, existing := range AllowedCreativeAutoGenerateConfigCreateV2TemplatesTemplateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_auto_generate_config_create_v2_templates_template_type value
func (v CreativeAutoGenerateConfigCreateV2TemplatesTemplateType) Ptr() *CreativeAutoGenerateConfigCreateV2TemplatesTemplateType {
	return &v
}
