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

// ClueFormUpdateV2ExtendInfoSignUpConfigSignType
type ClueFormUpdateV2ExtendInfoSignUpConfigSignType string

// List of clue_form_update_v2_extend_info_sign_up_config_sign_type
const (
	SIGN_TYPE_SCROLL_BAR_ClueFormUpdateV2ExtendInfoSignUpConfigSignType  ClueFormUpdateV2ExtendInfoSignUpConfigSignType = "SIGN_TYPE_SCROLL_BAR"
	SIGN_TYPE_SCROLL_WALL_ClueFormUpdateV2ExtendInfoSignUpConfigSignType ClueFormUpdateV2ExtendInfoSignUpConfigSignType = "SIGN_TYPE_SCROLL_WALL"
)

// All allowed values of ClueFormUpdateV2ExtendInfoSignUpConfigSignType enum
var AllowedClueFormUpdateV2ExtendInfoSignUpConfigSignTypeEnumValues = []ClueFormUpdateV2ExtendInfoSignUpConfigSignType{
	"SIGN_TYPE_SCROLL_BAR",
	"SIGN_TYPE_SCROLL_WALL",
}

// NewClueFormUpdateV2ExtendInfoSignUpConfigSignTypeFromValue returns a pointer to a valid ClueFormUpdateV2ExtendInfoSignUpConfigSignType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueFormUpdateV2ExtendInfoSignUpConfigSignTypeFromValue(v string) (*ClueFormUpdateV2ExtendInfoSignUpConfigSignType, error) {
	ev := ClueFormUpdateV2ExtendInfoSignUpConfigSignType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueFormUpdateV2ExtendInfoSignUpConfigSignType: valid values are %v", v, AllowedClueFormUpdateV2ExtendInfoSignUpConfigSignTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueFormUpdateV2ExtendInfoSignUpConfigSignType) IsValid() bool {
	for _, existing := range AllowedClueFormUpdateV2ExtendInfoSignUpConfigSignTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_form_update_v2_extend_info_sign_up_config_sign_type value
func (v ClueFormUpdateV2ExtendInfoSignUpConfigSignType) Ptr() *ClueFormUpdateV2ExtendInfoSignUpConfigSignType {
	return &v
}
