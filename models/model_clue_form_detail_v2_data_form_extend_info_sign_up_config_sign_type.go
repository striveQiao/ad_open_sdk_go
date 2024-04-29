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

// ClueFormDetailV2DataFormExtendInfoSignUpConfigSignType
type ClueFormDetailV2DataFormExtendInfoSignUpConfigSignType string

// List of clue_form_detail_v2_data_form_extend_info_sign_up_config_sign_type
const (
	SIGN_TYPE_SCROLL_BAR_ClueFormDetailV2DataFormExtendInfoSignUpConfigSignType  ClueFormDetailV2DataFormExtendInfoSignUpConfigSignType = "SIGN_TYPE_SCROLL_BAR"
	SIGN_TYPE_SCROLL_WALL_ClueFormDetailV2DataFormExtendInfoSignUpConfigSignType ClueFormDetailV2DataFormExtendInfoSignUpConfigSignType = "SIGN_TYPE_SCROLL_WALL"
)

// All allowed values of ClueFormDetailV2DataFormExtendInfoSignUpConfigSignType enum
var AllowedClueFormDetailV2DataFormExtendInfoSignUpConfigSignTypeEnumValues = []ClueFormDetailV2DataFormExtendInfoSignUpConfigSignType{
	"SIGN_TYPE_SCROLL_BAR",
	"SIGN_TYPE_SCROLL_WALL",
}

// NewClueFormDetailV2DataFormExtendInfoSignUpConfigSignTypeFromValue returns a pointer to a valid ClueFormDetailV2DataFormExtendInfoSignUpConfigSignType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueFormDetailV2DataFormExtendInfoSignUpConfigSignTypeFromValue(v string) (*ClueFormDetailV2DataFormExtendInfoSignUpConfigSignType, error) {
	ev := ClueFormDetailV2DataFormExtendInfoSignUpConfigSignType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueFormDetailV2DataFormExtendInfoSignUpConfigSignType: valid values are %v", v, AllowedClueFormDetailV2DataFormExtendInfoSignUpConfigSignTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueFormDetailV2DataFormExtendInfoSignUpConfigSignType) IsValid() bool {
	for _, existing := range AllowedClueFormDetailV2DataFormExtendInfoSignUpConfigSignTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_form_detail_v2_data_form_extend_info_sign_up_config_sign_type value
func (v ClueFormDetailV2DataFormExtendInfoSignUpConfigSignType) Ptr() *ClueFormDetailV2DataFormExtendInfoSignUpConfigSignType {
	return &v
}
