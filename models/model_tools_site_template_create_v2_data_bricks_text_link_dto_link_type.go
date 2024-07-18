/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsSiteTemplateCreateV2DataBricksTextLinkDtoLinkType
type ToolsSiteTemplateCreateV2DataBricksTextLinkDtoLinkType string

// List of tools_site_template_create_v2_data_bricks_text_link_dto_link_type
const (
	QUICK_APP_ToolsSiteTemplateCreateV2DataBricksTextLinkDtoLinkType ToolsSiteTemplateCreateV2DataBricksTextLinkDtoLinkType = "QUICK_APP"
	SCHEME_ToolsSiteTemplateCreateV2DataBricksTextLinkDtoLinkType    ToolsSiteTemplateCreateV2DataBricksTextLinkDtoLinkType = "SCHEME"
	URL_ToolsSiteTemplateCreateV2DataBricksTextLinkDtoLinkType       ToolsSiteTemplateCreateV2DataBricksTextLinkDtoLinkType = "URL"
)

// All allowed values of ToolsSiteTemplateCreateV2DataBricksTextLinkDtoLinkType enum
var AllowedToolsSiteTemplateCreateV2DataBricksTextLinkDtoLinkTypeEnumValues = []ToolsSiteTemplateCreateV2DataBricksTextLinkDtoLinkType{
	"QUICK_APP",
	"SCHEME",
	"URL",
}

// NewToolsSiteTemplateCreateV2DataBricksTextLinkDtoLinkTypeFromValue returns a pointer to a valid ToolsSiteTemplateCreateV2DataBricksTextLinkDtoLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteTemplateCreateV2DataBricksTextLinkDtoLinkTypeFromValue(v string) (*ToolsSiteTemplateCreateV2DataBricksTextLinkDtoLinkType, error) {
	ev := ToolsSiteTemplateCreateV2DataBricksTextLinkDtoLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteTemplateCreateV2DataBricksTextLinkDtoLinkType: valid values are %v", v, AllowedToolsSiteTemplateCreateV2DataBricksTextLinkDtoLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteTemplateCreateV2DataBricksTextLinkDtoLinkType) IsValid() bool {
	for _, existing := range AllowedToolsSiteTemplateCreateV2DataBricksTextLinkDtoLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_template_create_v2_data_bricks_text_link_dto_link_type value
func (v ToolsSiteTemplateCreateV2DataBricksTextLinkDtoLinkType) Ptr() *ToolsSiteTemplateCreateV2DataBricksTextLinkDtoLinkType {
	return &v
}
