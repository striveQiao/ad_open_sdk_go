/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueFormDetailV2DataFormElementsElementType
type ClueFormDetailV2DataFormElementsElementType string

// List of clue_form_detail_v2_data_form_elements_element_type
const (
	CHECKBOX_ClueFormDetailV2DataFormElementsElementType   ClueFormDetailV2DataFormElementsElementType = "CHECKBOX"
	SEX_ClueFormDetailV2DataFormElementsElementType        ClueFormDetailV2DataFormElementsElementType = "SEX"
	CITY_ClueFormDetailV2DataFormElementsElementType       ClueFormDetailV2DataFormElementsElementType = "CITY"
	TEXTAREA_ClueFormDetailV2DataFormElementsElementType   ClueFormDetailV2DataFormElementsElementType = "TEXTAREA"
	NUMBER_ClueFormDetailV2DataFormElementsElementType     ClueFormDetailV2DataFormElementsElementType = "NUMBER"
	DATE_ClueFormDetailV2DataFormElementsElementType       ClueFormDetailV2DataFormElementsElementType = "DATE"
	RADIO_ClueFormDetailV2DataFormElementsElementType      ClueFormDetailV2DataFormElementsElementType = "RADIO"
	TELEPHONE_ClueFormDetailV2DataFormElementsElementType  ClueFormDetailV2DataFormElementsElementType = "TELEPHONE"
	TEXT_ClueFormDetailV2DataFormElementsElementType       ClueFormDetailV2DataFormElementsElementType = "TEXT"
	CALCULATOR_ClueFormDetailV2DataFormElementsElementType ClueFormDetailV2DataFormElementsElementType = "CALCULATOR"
	EMAIL_ClueFormDetailV2DataFormElementsElementType      ClueFormDetailV2DataFormElementsElementType = "EMAIL"
	SELECT_ClueFormDetailV2DataFormElementsElementType     ClueFormDetailV2DataFormElementsElementType = "SELECT"
	NAME_ClueFormDetailV2DataFormElementsElementType       ClueFormDetailV2DataFormElementsElementType = "NAME"
)

// All allowed values of ClueFormDetailV2DataFormElementsElementType enum
var AllowedClueFormDetailV2DataFormElementsElementTypeEnumValues = []ClueFormDetailV2DataFormElementsElementType{
	"CHECKBOX",
	"SEX",
	"CITY",
	"TEXTAREA",
	"NUMBER",
	"DATE",
	"RADIO",
	"TELEPHONE",
	"TEXT",
	"CALCULATOR",
	"EMAIL",
	"SELECT",
	"NAME",
}

// NewClueFormDetailV2DataFormElementsElementTypeFromValue returns a pointer to a valid ClueFormDetailV2DataFormElementsElementType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueFormDetailV2DataFormElementsElementTypeFromValue(v string) (*ClueFormDetailV2DataFormElementsElementType, error) {
	ev := ClueFormDetailV2DataFormElementsElementType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueFormDetailV2DataFormElementsElementType: valid values are %v", v, AllowedClueFormDetailV2DataFormElementsElementTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueFormDetailV2DataFormElementsElementType) IsValid() bool {
	for _, existing := range AllowedClueFormDetailV2DataFormElementsElementTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_form_detail_v2_data_form_elements_element_type value
func (v ClueFormDetailV2DataFormElementsElementType) Ptr() *ClueFormDetailV2DataFormElementsElementType {
	return &v
}
