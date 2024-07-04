/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsSiteCreateV2SiteType
type ToolsSiteCreateV2SiteType string

// List of tools_site_create_v2_site_type
const (
	STREAMING_ToolsSiteCreateV2SiteType     ToolsSiteCreateV2SiteType = "STREAMING"
	ECOMMERCEPAGE_ToolsSiteCreateV2SiteType ToolsSiteCreateV2SiteType = "ECOMMERCEPAGE"
)

// All allowed values of ToolsSiteCreateV2SiteType enum
var AllowedToolsSiteCreateV2SiteTypeEnumValues = []ToolsSiteCreateV2SiteType{
	"STREAMING",
	"ECOMMERCEPAGE",
}

// NewToolsSiteCreateV2SiteTypeFromValue returns a pointer to a valid ToolsSiteCreateV2SiteType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteCreateV2SiteTypeFromValue(v string) (*ToolsSiteCreateV2SiteType, error) {
	ev := ToolsSiteCreateV2SiteType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteCreateV2SiteType: valid values are %v", v, AllowedToolsSiteCreateV2SiteTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteCreateV2SiteType) IsValid() bool {
	for _, existing := range AllowedToolsSiteCreateV2SiteTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_create_v2_site_type value
func (v ToolsSiteCreateV2SiteType) Ptr() *ToolsSiteCreateV2SiteType {
	return &v
}
