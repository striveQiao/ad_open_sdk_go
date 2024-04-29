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

// CreativeDetailGetV30DataAdDataMiniProgramInfoType
type CreativeDetailGetV30DataAdDataMiniProgramInfoType string

// List of creative_detail_get_v3.0_data_ad_data_mini_program_info_type
const (
	BYTE_APP_CreativeDetailGetV30DataAdDataMiniProgramInfoType     CreativeDetailGetV30DataAdDataMiniProgramInfoType = "BYTE_APP"
	BYTE_GAME_CreativeDetailGetV30DataAdDataMiniProgramInfoType    CreativeDetailGetV30DataAdDataMiniProgramInfoType = "BYTE_GAME"
	SHELL_APP_CreativeDetailGetV30DataAdDataMiniProgramInfoType    CreativeDetailGetV30DataAdDataMiniProgramInfoType = "SHELL_APP"
	TEMPLATE_APP_CreativeDetailGetV30DataAdDataMiniProgramInfoType CreativeDetailGetV30DataAdDataMiniProgramInfoType = "TEMPLATE_APP"
)

// All allowed values of CreativeDetailGetV30DataAdDataMiniProgramInfoType enum
var AllowedCreativeDetailGetV30DataAdDataMiniProgramInfoTypeEnumValues = []CreativeDetailGetV30DataAdDataMiniProgramInfoType{
	"BYTE_APP",
	"BYTE_GAME",
	"SHELL_APP",
	"TEMPLATE_APP",
}

// NewCreativeDetailGetV30DataAdDataMiniProgramInfoTypeFromValue returns a pointer to a valid CreativeDetailGetV30DataAdDataMiniProgramInfoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeDetailGetV30DataAdDataMiniProgramInfoTypeFromValue(v string) (*CreativeDetailGetV30DataAdDataMiniProgramInfoType, error) {
	ev := CreativeDetailGetV30DataAdDataMiniProgramInfoType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeDetailGetV30DataAdDataMiniProgramInfoType: valid values are %v", v, AllowedCreativeDetailGetV30DataAdDataMiniProgramInfoTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeDetailGetV30DataAdDataMiniProgramInfoType) IsValid() bool {
	for _, existing := range AllowedCreativeDetailGetV30DataAdDataMiniProgramInfoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_detail_get_v3.0_data_ad_data_mini_program_info_type value
func (v CreativeDetailGetV30DataAdDataMiniProgramInfoType) Ptr() *CreativeDetailGetV30DataAdDataMiniProgramInfoType {
	return &v
}
