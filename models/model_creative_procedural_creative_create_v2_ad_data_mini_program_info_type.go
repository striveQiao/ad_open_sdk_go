/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeProceduralCreativeCreateV2AdDataMiniProgramInfoType
type CreativeProceduralCreativeCreateV2AdDataMiniProgramInfoType string

// List of creative_procedural_creative_create_v2_ad_data_mini_program_info_type
const (
	SHELL_APP_CreativeProceduralCreativeCreateV2AdDataMiniProgramInfoType    CreativeProceduralCreativeCreateV2AdDataMiniProgramInfoType = "SHELL_APP"
	TEMPLATE_APP_CreativeProceduralCreativeCreateV2AdDataMiniProgramInfoType CreativeProceduralCreativeCreateV2AdDataMiniProgramInfoType = "TEMPLATE_APP"
	BYTE_GAME_CreativeProceduralCreativeCreateV2AdDataMiniProgramInfoType    CreativeProceduralCreativeCreateV2AdDataMiniProgramInfoType = "BYTE_GAME"
	BYTE_APP_CreativeProceduralCreativeCreateV2AdDataMiniProgramInfoType     CreativeProceduralCreativeCreateV2AdDataMiniProgramInfoType = "BYTE_APP"
)

// All allowed values of CreativeProceduralCreativeCreateV2AdDataMiniProgramInfoType enum
var AllowedCreativeProceduralCreativeCreateV2AdDataMiniProgramInfoTypeEnumValues = []CreativeProceduralCreativeCreateV2AdDataMiniProgramInfoType{
	"SHELL_APP",
	"TEMPLATE_APP",
	"BYTE_GAME",
	"BYTE_APP",
}

// NewCreativeProceduralCreativeCreateV2AdDataMiniProgramInfoTypeFromValue returns a pointer to a valid CreativeProceduralCreativeCreateV2AdDataMiniProgramInfoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeProceduralCreativeCreateV2AdDataMiniProgramInfoTypeFromValue(v string) (*CreativeProceduralCreativeCreateV2AdDataMiniProgramInfoType, error) {
	ev := CreativeProceduralCreativeCreateV2AdDataMiniProgramInfoType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeProceduralCreativeCreateV2AdDataMiniProgramInfoType: valid values are %v", v, AllowedCreativeProceduralCreativeCreateV2AdDataMiniProgramInfoTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeProceduralCreativeCreateV2AdDataMiniProgramInfoType) IsValid() bool {
	for _, existing := range AllowedCreativeProceduralCreativeCreateV2AdDataMiniProgramInfoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_procedural_creative_create_v2_ad_data_mini_program_info_type value
func (v CreativeProceduralCreativeCreateV2AdDataMiniProgramInfoType) Ptr() *CreativeProceduralCreativeCreateV2AdDataMiniProgramInfoType {
	return &v
}
