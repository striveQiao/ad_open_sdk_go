/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupUpdateV30AdInfoAudienceLocationType
type AdlabGroupUpdateV30AdInfoAudienceLocationType string

// List of adlab_group_update_v3.0_ad_info_audience_location_type
const (
	ALL_AdlabGroupUpdateV30AdInfoAudienceLocationType     AdlabGroupUpdateV30AdInfoAudienceLocationType = "ALL"
	CURRENT_AdlabGroupUpdateV30AdInfoAudienceLocationType AdlabGroupUpdateV30AdInfoAudienceLocationType = "CURRENT"
	HOME_AdlabGroupUpdateV30AdInfoAudienceLocationType    AdlabGroupUpdateV30AdInfoAudienceLocationType = "HOME"
	TRAVEL_AdlabGroupUpdateV30AdInfoAudienceLocationType  AdlabGroupUpdateV30AdInfoAudienceLocationType = "TRAVEL"
)

// All allowed values of AdlabGroupUpdateV30AdInfoAudienceLocationType enum
var AllowedAdlabGroupUpdateV30AdInfoAudienceLocationTypeEnumValues = []AdlabGroupUpdateV30AdInfoAudienceLocationType{
	"ALL",
	"CURRENT",
	"HOME",
	"TRAVEL",
}

// NewAdlabGroupUpdateV30AdInfoAudienceLocationTypeFromValue returns a pointer to a valid AdlabGroupUpdateV30AdInfoAudienceLocationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupUpdateV30AdInfoAudienceLocationTypeFromValue(v string) (*AdlabGroupUpdateV30AdInfoAudienceLocationType, error) {
	ev := AdlabGroupUpdateV30AdInfoAudienceLocationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupUpdateV30AdInfoAudienceLocationType: valid values are %v", v, AllowedAdlabGroupUpdateV30AdInfoAudienceLocationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupUpdateV30AdInfoAudienceLocationType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupUpdateV30AdInfoAudienceLocationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_update_v3.0_ad_info_audience_location_type value
func (v AdlabGroupUpdateV30AdInfoAudienceLocationType) Ptr() *AdlabGroupUpdateV30AdInfoAudienceLocationType {
	return &v
}
