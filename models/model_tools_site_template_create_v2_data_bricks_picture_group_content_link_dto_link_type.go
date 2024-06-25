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

// ToolsSiteTemplateCreateV2DataBricksPictureGroupContentLinkDtoLinkType
type ToolsSiteTemplateCreateV2DataBricksPictureGroupContentLinkDtoLinkType string

// List of tools_site_template_create_v2_data_bricks_picture_group_content_link_dto_link_type
const (
	QUICK_APP_ToolsSiteTemplateCreateV2DataBricksPictureGroupContentLinkDtoLinkType ToolsSiteTemplateCreateV2DataBricksPictureGroupContentLinkDtoLinkType = "QUICK_APP"
	SCHEME_ToolsSiteTemplateCreateV2DataBricksPictureGroupContentLinkDtoLinkType    ToolsSiteTemplateCreateV2DataBricksPictureGroupContentLinkDtoLinkType = "SCHEME"
	URL_ToolsSiteTemplateCreateV2DataBricksPictureGroupContentLinkDtoLinkType       ToolsSiteTemplateCreateV2DataBricksPictureGroupContentLinkDtoLinkType = "URL"
)

// All allowed values of ToolsSiteTemplateCreateV2DataBricksPictureGroupContentLinkDtoLinkType enum
var AllowedToolsSiteTemplateCreateV2DataBricksPictureGroupContentLinkDtoLinkTypeEnumValues = []ToolsSiteTemplateCreateV2DataBricksPictureGroupContentLinkDtoLinkType{
	"QUICK_APP",
	"SCHEME",
	"URL",
}

// NewToolsSiteTemplateCreateV2DataBricksPictureGroupContentLinkDtoLinkTypeFromValue returns a pointer to a valid ToolsSiteTemplateCreateV2DataBricksPictureGroupContentLinkDtoLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteTemplateCreateV2DataBricksPictureGroupContentLinkDtoLinkTypeFromValue(v string) (*ToolsSiteTemplateCreateV2DataBricksPictureGroupContentLinkDtoLinkType, error) {
	ev := ToolsSiteTemplateCreateV2DataBricksPictureGroupContentLinkDtoLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteTemplateCreateV2DataBricksPictureGroupContentLinkDtoLinkType: valid values are %v", v, AllowedToolsSiteTemplateCreateV2DataBricksPictureGroupContentLinkDtoLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteTemplateCreateV2DataBricksPictureGroupContentLinkDtoLinkType) IsValid() bool {
	for _, existing := range AllowedToolsSiteTemplateCreateV2DataBricksPictureGroupContentLinkDtoLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_template_create_v2_data_bricks_picture_group_content_link_dto_link_type value
func (v ToolsSiteTemplateCreateV2DataBricksPictureGroupContentLinkDtoLinkType) Ptr() *ToolsSiteTemplateCreateV2DataBricksPictureGroupContentLinkDtoLinkType {
	return &v
}
