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

// ToolsSiteTemplateSiteCreateV2BricksButtonLinkEventLinkLinkType
type ToolsSiteTemplateSiteCreateV2BricksButtonLinkEventLinkLinkType string

// List of tools_site_template_site_create_v2_bricks_button_link_event_link_link_type
const (
	QUICK_APP_ToolsSiteTemplateSiteCreateV2BricksButtonLinkEventLinkLinkType ToolsSiteTemplateSiteCreateV2BricksButtonLinkEventLinkLinkType = "QUICK_APP"
	SCHEME_ToolsSiteTemplateSiteCreateV2BricksButtonLinkEventLinkLinkType    ToolsSiteTemplateSiteCreateV2BricksButtonLinkEventLinkLinkType = "SCHEME"
	URL_ToolsSiteTemplateSiteCreateV2BricksButtonLinkEventLinkLinkType       ToolsSiteTemplateSiteCreateV2BricksButtonLinkEventLinkLinkType = "URL"
)

// All allowed values of ToolsSiteTemplateSiteCreateV2BricksButtonLinkEventLinkLinkType enum
var AllowedToolsSiteTemplateSiteCreateV2BricksButtonLinkEventLinkLinkTypeEnumValues = []ToolsSiteTemplateSiteCreateV2BricksButtonLinkEventLinkLinkType{
	"QUICK_APP",
	"SCHEME",
	"URL",
}

// NewToolsSiteTemplateSiteCreateV2BricksButtonLinkEventLinkLinkTypeFromValue returns a pointer to a valid ToolsSiteTemplateSiteCreateV2BricksButtonLinkEventLinkLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteTemplateSiteCreateV2BricksButtonLinkEventLinkLinkTypeFromValue(v string) (*ToolsSiteTemplateSiteCreateV2BricksButtonLinkEventLinkLinkType, error) {
	ev := ToolsSiteTemplateSiteCreateV2BricksButtonLinkEventLinkLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteTemplateSiteCreateV2BricksButtonLinkEventLinkLinkType: valid values are %v", v, AllowedToolsSiteTemplateSiteCreateV2BricksButtonLinkEventLinkLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteTemplateSiteCreateV2BricksButtonLinkEventLinkLinkType) IsValid() bool {
	for _, existing := range AllowedToolsSiteTemplateSiteCreateV2BricksButtonLinkEventLinkLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_template_site_create_v2_bricks_button_link_event_link_link_type value
func (v ToolsSiteTemplateSiteCreateV2BricksButtonLinkEventLinkLinkType) Ptr() *ToolsSiteTemplateSiteCreateV2BricksButtonLinkEventLinkLinkType {
	return &v
}
