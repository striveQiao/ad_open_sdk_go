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

// ToolsClueLifeGetV2DataListClueType
type ToolsClueLifeGetV2DataListClueType string

// List of tools_clue_life_get_v2_data_list_clue_type
const (
	FORM_ToolsClueLifeGetV2DataListClueType         ToolsClueLifeGetV2DataListClueType = "FORM"
	CONSULT_ToolsClueLifeGetV2DataListClueType      ToolsClueLifeGetV2DataListClueType = "CONSULT"
	SMARTPHONE_ToolsClueLifeGetV2DataListClueType   ToolsClueLifeGetV2DataListClueType = "SMARTPHONE"
	GROUP_BUYING_ToolsClueLifeGetV2DataListClueType ToolsClueLifeGetV2DataListClueType = "GROUP_BUYING"
)

// All allowed values of ToolsClueLifeGetV2DataListClueType enum
var AllowedToolsClueLifeGetV2DataListClueTypeEnumValues = []ToolsClueLifeGetV2DataListClueType{
	"FORM",
	"CONSULT",
	"SMARTPHONE",
	"GROUP_BUYING",
}

// NewToolsClueLifeGetV2DataListClueTypeFromValue returns a pointer to a valid ToolsClueLifeGetV2DataListClueType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsClueLifeGetV2DataListClueTypeFromValue(v string) (*ToolsClueLifeGetV2DataListClueType, error) {
	ev := ToolsClueLifeGetV2DataListClueType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsClueLifeGetV2DataListClueType: valid values are %v", v, AllowedToolsClueLifeGetV2DataListClueTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsClueLifeGetV2DataListClueType) IsValid() bool {
	for _, existing := range AllowedToolsClueLifeGetV2DataListClueTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_clue_life_get_v2_data_list_clue_type value
func (v ToolsClueLifeGetV2DataListClueType) Ptr() *ToolsClueLifeGetV2DataListClueType {
	return &v
}
