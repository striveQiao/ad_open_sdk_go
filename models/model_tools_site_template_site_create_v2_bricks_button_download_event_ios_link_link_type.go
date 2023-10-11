/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventIosLinkLinkType
type ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventIosLinkLinkType string

// List of tools_site_template_site_create_v2_bricks_button_download_event_ios_link_link_type
const (
	QUICK_APP_ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventIosLinkLinkType ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventIosLinkLinkType = "QUICK_APP"
	SCHEME_ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventIosLinkLinkType    ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventIosLinkLinkType = "SCHEME"
	URL_ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventIosLinkLinkType       ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventIosLinkLinkType = "URL"
)

// All allowed values of ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventIosLinkLinkType enum
var AllowedToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventIosLinkLinkTypeEnumValues = []ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventIosLinkLinkType{
	"QUICK_APP",
	"SCHEME",
	"URL",
}

// NewToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventIosLinkLinkTypeFromValue returns a pointer to a valid ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventIosLinkLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventIosLinkLinkTypeFromValue(v string) (*ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventIosLinkLinkType, error) {
	ev := ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventIosLinkLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventIosLinkLinkType: valid values are %v", v, AllowedToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventIosLinkLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventIosLinkLinkType) IsValid() bool {
	for _, existing := range AllowedToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventIosLinkLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_template_site_create_v2_bricks_button_download_event_ios_link_link_type value
func (v ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventIosLinkLinkType) Ptr() *ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventIosLinkLinkType {
	return &v
}
