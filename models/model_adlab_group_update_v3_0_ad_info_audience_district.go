/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupUpdateV30AdInfoAudienceDistrict
type AdlabGroupUpdateV30AdInfoAudienceDistrict string

// List of adlab_group_update_v3.0_ad_info_audience_district
const (
	BUSINESS_DISTRICT_AdlabGroupUpdateV30AdInfoAudienceDistrict AdlabGroupUpdateV30AdInfoAudienceDistrict = "BUSINESS_DISTRICT"
	NONE_AdlabGroupUpdateV30AdInfoAudienceDistrict              AdlabGroupUpdateV30AdInfoAudienceDistrict = "NONE"
	REGION_AdlabGroupUpdateV30AdInfoAudienceDistrict            AdlabGroupUpdateV30AdInfoAudienceDistrict = "REGION"
)

// All allowed values of AdlabGroupUpdateV30AdInfoAudienceDistrict enum
var AllowedAdlabGroupUpdateV30AdInfoAudienceDistrictEnumValues = []AdlabGroupUpdateV30AdInfoAudienceDistrict{
	"BUSINESS_DISTRICT",
	"NONE",
	"REGION",
}

// NewAdlabGroupUpdateV30AdInfoAudienceDistrictFromValue returns a pointer to a valid AdlabGroupUpdateV30AdInfoAudienceDistrict
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupUpdateV30AdInfoAudienceDistrictFromValue(v string) (*AdlabGroupUpdateV30AdInfoAudienceDistrict, error) {
	ev := AdlabGroupUpdateV30AdInfoAudienceDistrict(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupUpdateV30AdInfoAudienceDistrict: valid values are %v", v, AllowedAdlabGroupUpdateV30AdInfoAudienceDistrictEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupUpdateV30AdInfoAudienceDistrict) IsValid() bool {
	for _, existing := range AllowedAdlabGroupUpdateV30AdInfoAudienceDistrictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_update_v3.0_ad_info_audience_district value
func (v AdlabGroupUpdateV30AdInfoAudienceDistrict) Ptr() *AdlabGroupUpdateV30AdInfoAudienceDistrict {
	return &v
}
