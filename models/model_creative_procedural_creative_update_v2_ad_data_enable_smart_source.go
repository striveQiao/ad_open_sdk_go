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

// CreativeProceduralCreativeUpdateV2AdDataEnableSmartSource
type CreativeProceduralCreativeUpdateV2AdDataEnableSmartSource string

// List of creative_procedural_creative_update_v2_ad_data_enable_smart_source
const (
	ON_CreativeProceduralCreativeUpdateV2AdDataEnableSmartSource  CreativeProceduralCreativeUpdateV2AdDataEnableSmartSource = "ON"
	OFF_CreativeProceduralCreativeUpdateV2AdDataEnableSmartSource CreativeProceduralCreativeUpdateV2AdDataEnableSmartSource = "OFF"
)

// All allowed values of CreativeProceduralCreativeUpdateV2AdDataEnableSmartSource enum
var AllowedCreativeProceduralCreativeUpdateV2AdDataEnableSmartSourceEnumValues = []CreativeProceduralCreativeUpdateV2AdDataEnableSmartSource{
	"ON",
	"OFF",
}

// NewCreativeProceduralCreativeUpdateV2AdDataEnableSmartSourceFromValue returns a pointer to a valid CreativeProceduralCreativeUpdateV2AdDataEnableSmartSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeProceduralCreativeUpdateV2AdDataEnableSmartSourceFromValue(v string) (*CreativeProceduralCreativeUpdateV2AdDataEnableSmartSource, error) {
	ev := CreativeProceduralCreativeUpdateV2AdDataEnableSmartSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeProceduralCreativeUpdateV2AdDataEnableSmartSource: valid values are %v", v, AllowedCreativeProceduralCreativeUpdateV2AdDataEnableSmartSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeProceduralCreativeUpdateV2AdDataEnableSmartSource) IsValid() bool {
	for _, existing := range AllowedCreativeProceduralCreativeUpdateV2AdDataEnableSmartSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_procedural_creative_update_v2_ad_data_enable_smart_source value
func (v CreativeProceduralCreativeUpdateV2AdDataEnableSmartSource) Ptr() *CreativeProceduralCreativeUpdateV2AdDataEnableSmartSource {
	return &v
}
