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

// CreativeCustomCreativeUpdateV2AdDataSupplementsSupplementType
type CreativeCustomCreativeUpdateV2AdDataSupplementsSupplementType string

// List of creative_custom_creative_update_v2_ad_data_supplements_supplement_type
const (
	NORMAL_CreativeCustomCreativeUpdateV2AdDataSupplementsSupplementType     CreativeCustomCreativeUpdateV2AdDataSupplementsSupplementType = "NORMAL"
	CLOUD_GAME_CreativeCustomCreativeUpdateV2AdDataSupplementsSupplementType CreativeCustomCreativeUpdateV2AdDataSupplementsSupplementType = "CLOUD_GAME"
)

// All allowed values of CreativeCustomCreativeUpdateV2AdDataSupplementsSupplementType enum
var AllowedCreativeCustomCreativeUpdateV2AdDataSupplementsSupplementTypeEnumValues = []CreativeCustomCreativeUpdateV2AdDataSupplementsSupplementType{
	"NORMAL",
	"CLOUD_GAME",
}

// NewCreativeCustomCreativeUpdateV2AdDataSupplementsSupplementTypeFromValue returns a pointer to a valid CreativeCustomCreativeUpdateV2AdDataSupplementsSupplementType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeCustomCreativeUpdateV2AdDataSupplementsSupplementTypeFromValue(v string) (*CreativeCustomCreativeUpdateV2AdDataSupplementsSupplementType, error) {
	ev := CreativeCustomCreativeUpdateV2AdDataSupplementsSupplementType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeCustomCreativeUpdateV2AdDataSupplementsSupplementType: valid values are %v", v, AllowedCreativeCustomCreativeUpdateV2AdDataSupplementsSupplementTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeCustomCreativeUpdateV2AdDataSupplementsSupplementType) IsValid() bool {
	for _, existing := range AllowedCreativeCustomCreativeUpdateV2AdDataSupplementsSupplementTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_custom_creative_update_v2_ad_data_supplements_supplement_type value
func (v CreativeCustomCreativeUpdateV2AdDataSupplementsSupplementType) Ptr() *CreativeCustomCreativeUpdateV2AdDataSupplementsSupplementType {
	return &v
}
