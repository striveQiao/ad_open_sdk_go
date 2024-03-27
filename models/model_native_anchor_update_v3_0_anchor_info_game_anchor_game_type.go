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

// NativeAnchorUpdateV30AnchorInfoGameAnchorGameType
type NativeAnchorUpdateV30AnchorInfoGameAnchorGameType string

// List of native_anchor_update_v3.0_anchor_info_game_anchor_game_type
const (
	GENERAL_NativeAnchorUpdateV30AnchorInfoGameAnchorGameType    NativeAnchorUpdateV30AnchorInfoGameAnchorGameType = "GENERAL"
	MICRO_GAME_NativeAnchorUpdateV30AnchorInfoGameAnchorGameType NativeAnchorUpdateV30AnchorInfoGameAnchorGameType = "MICRO_GAME"
)

// All allowed values of NativeAnchorUpdateV30AnchorInfoGameAnchorGameType enum
var AllowedNativeAnchorUpdateV30AnchorInfoGameAnchorGameTypeEnumValues = []NativeAnchorUpdateV30AnchorInfoGameAnchorGameType{
	"GENERAL",
	"MICRO_GAME",
}

// NewNativeAnchorUpdateV30AnchorInfoGameAnchorGameTypeFromValue returns a pointer to a valid NativeAnchorUpdateV30AnchorInfoGameAnchorGameType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorUpdateV30AnchorInfoGameAnchorGameTypeFromValue(v string) (*NativeAnchorUpdateV30AnchorInfoGameAnchorGameType, error) {
	ev := NativeAnchorUpdateV30AnchorInfoGameAnchorGameType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorUpdateV30AnchorInfoGameAnchorGameType: valid values are %v", v, AllowedNativeAnchorUpdateV30AnchorInfoGameAnchorGameTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorUpdateV30AnchorInfoGameAnchorGameType) IsValid() bool {
	for _, existing := range AllowedNativeAnchorUpdateV30AnchorInfoGameAnchorGameTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_update_v3.0_anchor_info_game_anchor_game_type value
func (v NativeAnchorUpdateV30AnchorInfoGameAnchorGameType) Ptr() *NativeAnchorUpdateV30AnchorInfoGameAnchorGameType {
	return &v
}
