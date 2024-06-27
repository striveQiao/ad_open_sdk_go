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

// ClueSmartphoneGetV2DataListValidateType
type ClueSmartphoneGetV2DataListValidateType string

// List of clue_smartphone_get_v2_data_list_validate_type
const (
	VALIDITY_PRIORITY_ClueSmartphoneGetV2DataListValidateType ClueSmartphoneGetV2DataListValidateType = "VALIDITY_PRIORITY"
	NONE_VERIFICATION_ClueSmartphoneGetV2DataListValidateType ClueSmartphoneGetV2DataListValidateType = "NONE_VERIFICATION"
	AUTO_VERIFICATION_ClueSmartphoneGetV2DataListValidateType ClueSmartphoneGetV2DataListValidateType = "AUTO_VERIFICATION"
	ALL_VERIFICATION_ClueSmartphoneGetV2DataListValidateType  ClueSmartphoneGetV2DataListValidateType = "ALL_VERIFICATION"
	CLUE_PRIORITY_ClueSmartphoneGetV2DataListValidateType     ClueSmartphoneGetV2DataListValidateType = "CLUE_PRIORITY"
)

// All allowed values of ClueSmartphoneGetV2DataListValidateType enum
var AllowedClueSmartphoneGetV2DataListValidateTypeEnumValues = []ClueSmartphoneGetV2DataListValidateType{
	"VALIDITY_PRIORITY",
	"NONE_VERIFICATION",
	"AUTO_VERIFICATION",
	"ALL_VERIFICATION",
	"CLUE_PRIORITY",
}

// NewClueSmartphoneGetV2DataListValidateTypeFromValue returns a pointer to a valid ClueSmartphoneGetV2DataListValidateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueSmartphoneGetV2DataListValidateTypeFromValue(v string) (*ClueSmartphoneGetV2DataListValidateType, error) {
	ev := ClueSmartphoneGetV2DataListValidateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueSmartphoneGetV2DataListValidateType: valid values are %v", v, AllowedClueSmartphoneGetV2DataListValidateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueSmartphoneGetV2DataListValidateType) IsValid() bool {
	for _, existing := range AllowedClueSmartphoneGetV2DataListValidateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_smartphone_get_v2_data_list_validate_type value
func (v ClueSmartphoneGetV2DataListValidateType) Ptr() *ClueSmartphoneGetV2DataListValidateType {
	return &v
}
