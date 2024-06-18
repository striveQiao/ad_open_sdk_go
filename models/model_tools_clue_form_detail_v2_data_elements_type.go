/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsClueFormDetailV2DataElementsType
type ToolsClueFormDetailV2DataElementsType string

// List of tools_clue_form_detail_v2_data_elements_type
const (
	CALCULATOR_ToolsClueFormDetailV2DataElementsType ToolsClueFormDetailV2DataElementsType = "CALCULATOR"
	CHECKBOX_ToolsClueFormDetailV2DataElementsType   ToolsClueFormDetailV2DataElementsType = "CHECKBOX"
	CITY_ToolsClueFormDetailV2DataElementsType       ToolsClueFormDetailV2DataElementsType = "CITY"
	DATE_ToolsClueFormDetailV2DataElementsType       ToolsClueFormDetailV2DataElementsType = "DATE"
	EMAIL_ToolsClueFormDetailV2DataElementsType      ToolsClueFormDetailV2DataElementsType = "EMAIL"
	NAME_ToolsClueFormDetailV2DataElementsType       ToolsClueFormDetailV2DataElementsType = "NAME"
	NUMBER_ToolsClueFormDetailV2DataElementsType     ToolsClueFormDetailV2DataElementsType = "NUMBER"
	RADIO_ToolsClueFormDetailV2DataElementsType      ToolsClueFormDetailV2DataElementsType = "RADIO"
	SELECT_ToolsClueFormDetailV2DataElementsType     ToolsClueFormDetailV2DataElementsType = "SELECT"
	SEX_ToolsClueFormDetailV2DataElementsType        ToolsClueFormDetailV2DataElementsType = "SEX"
	TELEPHONE_ToolsClueFormDetailV2DataElementsType  ToolsClueFormDetailV2DataElementsType = "TELEPHONE"
	TEXT_ToolsClueFormDetailV2DataElementsType       ToolsClueFormDetailV2DataElementsType = "TEXT"
	TEXTAREA_ToolsClueFormDetailV2DataElementsType   ToolsClueFormDetailV2DataElementsType = "TEXTAREA"
)

// All allowed values of ToolsClueFormDetailV2DataElementsType enum
var AllowedToolsClueFormDetailV2DataElementsTypeEnumValues = []ToolsClueFormDetailV2DataElementsType{
	"CALCULATOR",
	"CHECKBOX",
	"CITY",
	"DATE",
	"EMAIL",
	"NAME",
	"NUMBER",
	"RADIO",
	"SELECT",
	"SEX",
	"TELEPHONE",
	"TEXT",
	"TEXTAREA",
}

// NewToolsClueFormDetailV2DataElementsTypeFromValue returns a pointer to a valid ToolsClueFormDetailV2DataElementsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsClueFormDetailV2DataElementsTypeFromValue(v string) (*ToolsClueFormDetailV2DataElementsType, error) {
	ev := ToolsClueFormDetailV2DataElementsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsClueFormDetailV2DataElementsType: valid values are %v", v, AllowedToolsClueFormDetailV2DataElementsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsClueFormDetailV2DataElementsType) IsValid() bool {
	for _, existing := range AllowedToolsClueFormDetailV2DataElementsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_clue_form_detail_v2_data_elements_type value
func (v ToolsClueFormDetailV2DataElementsType) Ptr() *ToolsClueFormDetailV2DataElementsType {
	return &v
}
