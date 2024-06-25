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

// CreativeProceduralCreativeCreateV2AdDataSupplementsSupplementType
type CreativeProceduralCreativeCreateV2AdDataSupplementsSupplementType string

// List of creative_procedural_creative_create_v2_ad_data_supplements_supplement_type
const (
	NORMAL_CreativeProceduralCreativeCreateV2AdDataSupplementsSupplementType     CreativeProceduralCreativeCreateV2AdDataSupplementsSupplementType = "NORMAL"
	CLOUD_GAME_CreativeProceduralCreativeCreateV2AdDataSupplementsSupplementType CreativeProceduralCreativeCreateV2AdDataSupplementsSupplementType = "CLOUD_GAME"
)

// All allowed values of CreativeProceduralCreativeCreateV2AdDataSupplementsSupplementType enum
var AllowedCreativeProceduralCreativeCreateV2AdDataSupplementsSupplementTypeEnumValues = []CreativeProceduralCreativeCreateV2AdDataSupplementsSupplementType{
	"NORMAL",
	"CLOUD_GAME",
}

// NewCreativeProceduralCreativeCreateV2AdDataSupplementsSupplementTypeFromValue returns a pointer to a valid CreativeProceduralCreativeCreateV2AdDataSupplementsSupplementType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeProceduralCreativeCreateV2AdDataSupplementsSupplementTypeFromValue(v string) (*CreativeProceduralCreativeCreateV2AdDataSupplementsSupplementType, error) {
	ev := CreativeProceduralCreativeCreateV2AdDataSupplementsSupplementType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeProceduralCreativeCreateV2AdDataSupplementsSupplementType: valid values are %v", v, AllowedCreativeProceduralCreativeCreateV2AdDataSupplementsSupplementTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeProceduralCreativeCreateV2AdDataSupplementsSupplementType) IsValid() bool {
	for _, existing := range AllowedCreativeProceduralCreativeCreateV2AdDataSupplementsSupplementTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_procedural_creative_create_v2_ad_data_supplements_supplement_type value
func (v CreativeProceduralCreativeCreateV2AdDataSupplementsSupplementType) Ptr() *CreativeProceduralCreativeCreateV2AdDataSupplementsSupplementType {
	return &v
}
