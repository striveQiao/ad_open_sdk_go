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

// ToolsSiteTemplateCreateV2DataBricksButtonEventType
type ToolsSiteTemplateCreateV2DataBricksButtonEventType string

// List of tools_site_template_create_v2_data_bricks_button_event_type
const (
	APPOINT_EVENT_ToolsSiteTemplateCreateV2DataBricksButtonEventType   ToolsSiteTemplateCreateV2DataBricksButtonEventType = "APPOINT_EVENT"
	DOWNLOAD_EVENT_ToolsSiteTemplateCreateV2DataBricksButtonEventType  ToolsSiteTemplateCreateV2DataBricksButtonEventType = "DOWNLOAD_EVENT"
	LINK_EVENT_ToolsSiteTemplateCreateV2DataBricksButtonEventType      ToolsSiteTemplateCreateV2DataBricksButtonEventType = "LINK_EVENT"
	TELEPHONE_EVENT_ToolsSiteTemplateCreateV2DataBricksButtonEventType ToolsSiteTemplateCreateV2DataBricksButtonEventType = "TELEPHONE_EVENT"
)

// All allowed values of ToolsSiteTemplateCreateV2DataBricksButtonEventType enum
var AllowedToolsSiteTemplateCreateV2DataBricksButtonEventTypeEnumValues = []ToolsSiteTemplateCreateV2DataBricksButtonEventType{
	"APPOINT_EVENT",
	"DOWNLOAD_EVENT",
	"LINK_EVENT",
	"TELEPHONE_EVENT",
}

// NewToolsSiteTemplateCreateV2DataBricksButtonEventTypeFromValue returns a pointer to a valid ToolsSiteTemplateCreateV2DataBricksButtonEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteTemplateCreateV2DataBricksButtonEventTypeFromValue(v string) (*ToolsSiteTemplateCreateV2DataBricksButtonEventType, error) {
	ev := ToolsSiteTemplateCreateV2DataBricksButtonEventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteTemplateCreateV2DataBricksButtonEventType: valid values are %v", v, AllowedToolsSiteTemplateCreateV2DataBricksButtonEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteTemplateCreateV2DataBricksButtonEventType) IsValid() bool {
	for _, existing := range AllowedToolsSiteTemplateCreateV2DataBricksButtonEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_template_create_v2_data_bricks_button_event_type value
func (v ToolsSiteTemplateCreateV2DataBricksButtonEventType) Ptr() *ToolsSiteTemplateCreateV2DataBricksButtonEventType {
	return &v
}
