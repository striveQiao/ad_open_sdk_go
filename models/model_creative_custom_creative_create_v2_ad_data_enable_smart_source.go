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

// CreativeCustomCreativeCreateV2AdDataEnableSmartSource
type CreativeCustomCreativeCreateV2AdDataEnableSmartSource string

// List of creative_custom_creative_create_v2_ad_data_enable_smart_source
const (
	OFF_CreativeCustomCreativeCreateV2AdDataEnableSmartSource CreativeCustomCreativeCreateV2AdDataEnableSmartSource = "OFF"
	ON_CreativeCustomCreativeCreateV2AdDataEnableSmartSource  CreativeCustomCreativeCreateV2AdDataEnableSmartSource = "ON"
)

// All allowed values of CreativeCustomCreativeCreateV2AdDataEnableSmartSource enum
var AllowedCreativeCustomCreativeCreateV2AdDataEnableSmartSourceEnumValues = []CreativeCustomCreativeCreateV2AdDataEnableSmartSource{
	"OFF",
	"ON",
}

// NewCreativeCustomCreativeCreateV2AdDataEnableSmartSourceFromValue returns a pointer to a valid CreativeCustomCreativeCreateV2AdDataEnableSmartSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeCustomCreativeCreateV2AdDataEnableSmartSourceFromValue(v string) (*CreativeCustomCreativeCreateV2AdDataEnableSmartSource, error) {
	ev := CreativeCustomCreativeCreateV2AdDataEnableSmartSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeCustomCreativeCreateV2AdDataEnableSmartSource: valid values are %v", v, AllowedCreativeCustomCreativeCreateV2AdDataEnableSmartSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeCustomCreativeCreateV2AdDataEnableSmartSource) IsValid() bool {
	for _, existing := range AllowedCreativeCustomCreativeCreateV2AdDataEnableSmartSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_custom_creative_create_v2_ad_data_enable_smart_source value
func (v CreativeCustomCreativeCreateV2AdDataEnableSmartSource) Ptr() *CreativeCustomCreativeCreateV2AdDataEnableSmartSource {
	return &v
}
