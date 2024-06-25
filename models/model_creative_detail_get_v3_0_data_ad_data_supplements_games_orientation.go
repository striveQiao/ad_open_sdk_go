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

// CreativeDetailGetV30DataAdDataSupplementsGamesOrientation
type CreativeDetailGetV30DataAdDataSupplementsGamesOrientation string

// List of creative_detail_get_v3.0_data_ad_data_supplements_games_orientation
const (
	VERTICAL_CreativeDetailGetV30DataAdDataSupplementsGamesOrientation   CreativeDetailGetV30DataAdDataSupplementsGamesOrientation = "VERTICAL"
	HORIZONTAL_CreativeDetailGetV30DataAdDataSupplementsGamesOrientation CreativeDetailGetV30DataAdDataSupplementsGamesOrientation = "HORIZONTAL"
)

// All allowed values of CreativeDetailGetV30DataAdDataSupplementsGamesOrientation enum
var AllowedCreativeDetailGetV30DataAdDataSupplementsGamesOrientationEnumValues = []CreativeDetailGetV30DataAdDataSupplementsGamesOrientation{
	"VERTICAL",
	"HORIZONTAL",
}

// NewCreativeDetailGetV30DataAdDataSupplementsGamesOrientationFromValue returns a pointer to a valid CreativeDetailGetV30DataAdDataSupplementsGamesOrientation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeDetailGetV30DataAdDataSupplementsGamesOrientationFromValue(v string) (*CreativeDetailGetV30DataAdDataSupplementsGamesOrientation, error) {
	ev := CreativeDetailGetV30DataAdDataSupplementsGamesOrientation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeDetailGetV30DataAdDataSupplementsGamesOrientation: valid values are %v", v, AllowedCreativeDetailGetV30DataAdDataSupplementsGamesOrientationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeDetailGetV30DataAdDataSupplementsGamesOrientation) IsValid() bool {
	for _, existing := range AllowedCreativeDetailGetV30DataAdDataSupplementsGamesOrientationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_detail_get_v3.0_data_ad_data_supplements_games_orientation value
func (v CreativeDetailGetV30DataAdDataSupplementsGamesOrientation) Ptr() *CreativeDetailGetV30DataAdDataSupplementsGamesOrientation {
	return &v
}
