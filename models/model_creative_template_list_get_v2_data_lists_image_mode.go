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

// CreativeTemplateListGetV2DataListsImageMode
type CreativeTemplateListGetV2DataListsImageMode string

// List of creative_template_list_get_v2_data_lists_image_mode
const (
	VIDEO_HORIZONTAL_CreativeTemplateListGetV2DataListsImageMode CreativeTemplateListGetV2DataListsImageMode = "VIDEO_HORIZONTAL"
	VIDEO_VERTICAL_CreativeTemplateListGetV2DataListsImageMode   CreativeTemplateListGetV2DataListsImageMode = "VIDEO_VERTICAL"
)

// All allowed values of CreativeTemplateListGetV2DataListsImageMode enum
var AllowedCreativeTemplateListGetV2DataListsImageModeEnumValues = []CreativeTemplateListGetV2DataListsImageMode{
	"VIDEO_HORIZONTAL",
	"VIDEO_VERTICAL",
}

// NewCreativeTemplateListGetV2DataListsImageModeFromValue returns a pointer to a valid CreativeTemplateListGetV2DataListsImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeTemplateListGetV2DataListsImageModeFromValue(v string) (*CreativeTemplateListGetV2DataListsImageMode, error) {
	ev := CreativeTemplateListGetV2DataListsImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeTemplateListGetV2DataListsImageMode: valid values are %v", v, AllowedCreativeTemplateListGetV2DataListsImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeTemplateListGetV2DataListsImageMode) IsValid() bool {
	for _, existing := range AllowedCreativeTemplateListGetV2DataListsImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_template_list_get_v2_data_lists_image_mode value
func (v CreativeTemplateListGetV2DataListsImageMode) Ptr() *CreativeTemplateListGetV2DataListsImageMode {
	return &v
}
