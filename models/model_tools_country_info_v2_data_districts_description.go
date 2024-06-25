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

// ToolsCountryInfoV2DataDistrictsDescription
type ToolsCountryInfoV2DataDistrictsDescription string

// List of tools_country_info_v2_data_districts_description
const (
	COUNTRY_ToolsCountryInfoV2DataDistrictsDescription ToolsCountryInfoV2DataDistrictsDescription = "COUNTRY"
	STATE_ToolsCountryInfoV2DataDistrictsDescription   ToolsCountryInfoV2DataDistrictsDescription = "STATE"
	REGION_ToolsCountryInfoV2DataDistrictsDescription  ToolsCountryInfoV2DataDistrictsDescription = "REGION"
)

// All allowed values of ToolsCountryInfoV2DataDistrictsDescription enum
var AllowedToolsCountryInfoV2DataDistrictsDescriptionEnumValues = []ToolsCountryInfoV2DataDistrictsDescription{
	"COUNTRY",
	"STATE",
	"REGION",
}

// NewToolsCountryInfoV2DataDistrictsDescriptionFromValue returns a pointer to a valid ToolsCountryInfoV2DataDistrictsDescription
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsCountryInfoV2DataDistrictsDescriptionFromValue(v string) (*ToolsCountryInfoV2DataDistrictsDescription, error) {
	ev := ToolsCountryInfoV2DataDistrictsDescription(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsCountryInfoV2DataDistrictsDescription: valid values are %v", v, AllowedToolsCountryInfoV2DataDistrictsDescriptionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsCountryInfoV2DataDistrictsDescription) IsValid() bool {
	for _, existing := range AllowedToolsCountryInfoV2DataDistrictsDescriptionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_country_info_v2_data_districts_description value
func (v ToolsCountryInfoV2DataDistrictsDescription) Ptr() *ToolsCountryInfoV2DataDistrictsDescription {
	return &v
}
