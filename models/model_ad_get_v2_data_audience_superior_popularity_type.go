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

// AdGetV2DataAudienceSuperiorPopularityType
type AdGetV2DataAudienceSuperiorPopularityType string

// List of ad_get_v2_data_audience_superior_popularity_type
const (
	GAME_AdGetV2DataAudienceSuperiorPopularityType AdGetV2DataAudienceSuperiorPopularityType = "GAME"
	NONE_AdGetV2DataAudienceSuperiorPopularityType AdGetV2DataAudienceSuperiorPopularityType = "NONE"
)

// All allowed values of AdGetV2DataAudienceSuperiorPopularityType enum
var AllowedAdGetV2DataAudienceSuperiorPopularityTypeEnumValues = []AdGetV2DataAudienceSuperiorPopularityType{
	"GAME",
	"NONE",
}

// NewAdGetV2DataAudienceSuperiorPopularityTypeFromValue returns a pointer to a valid AdGetV2DataAudienceSuperiorPopularityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataAudienceSuperiorPopularityTypeFromValue(v string) (*AdGetV2DataAudienceSuperiorPopularityType, error) {
	ev := AdGetV2DataAudienceSuperiorPopularityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataAudienceSuperiorPopularityType: valid values are %v", v, AllowedAdGetV2DataAudienceSuperiorPopularityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataAudienceSuperiorPopularityType) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataAudienceSuperiorPopularityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_audience_superior_popularity_type value
func (v AdGetV2DataAudienceSuperiorPopularityType) Ptr() *AdGetV2DataAudienceSuperiorPopularityType {
	return &v
}
