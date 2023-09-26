/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeAutoGenerateConfigCreateV2SelectTemplateMode
type CreativeAutoGenerateConfigCreateV2SelectTemplateMode string

// List of creative_auto_generate_config_create_v2_select_template_mode
const (
	CUSTOM_SELECT_TEMPLATE_CreativeAutoGenerateConfigCreateV2SelectTemplateMode CreativeAutoGenerateConfigCreateV2SelectTemplateMode = "CUSTOM_SELECT_TEMPLATE"
	SYSTEM_SELECT_TEMPLATE_CreativeAutoGenerateConfigCreateV2SelectTemplateMode CreativeAutoGenerateConfigCreateV2SelectTemplateMode = "SYSTEM_SELECT_TEMPLATE"
)

// All allowed values of CreativeAutoGenerateConfigCreateV2SelectTemplateMode enum
var AllowedCreativeAutoGenerateConfigCreateV2SelectTemplateModeEnumValues = []CreativeAutoGenerateConfigCreateV2SelectTemplateMode{
	"CUSTOM_SELECT_TEMPLATE",
	"SYSTEM_SELECT_TEMPLATE",
}

// NewCreativeAutoGenerateConfigCreateV2SelectTemplateModeFromValue returns a pointer to a valid CreativeAutoGenerateConfigCreateV2SelectTemplateMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeAutoGenerateConfigCreateV2SelectTemplateModeFromValue(v string) (*CreativeAutoGenerateConfigCreateV2SelectTemplateMode, error) {
	ev := CreativeAutoGenerateConfigCreateV2SelectTemplateMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeAutoGenerateConfigCreateV2SelectTemplateMode: valid values are %v", v, AllowedCreativeAutoGenerateConfigCreateV2SelectTemplateModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeAutoGenerateConfigCreateV2SelectTemplateMode) IsValid() bool {
	for _, existing := range AllowedCreativeAutoGenerateConfigCreateV2SelectTemplateModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_auto_generate_config_create_v2_select_template_mode value
func (v CreativeAutoGenerateConfigCreateV2SelectTemplateMode) Ptr() *CreativeAutoGenerateConfigCreateV2SelectTemplateMode {
	return &v
}
