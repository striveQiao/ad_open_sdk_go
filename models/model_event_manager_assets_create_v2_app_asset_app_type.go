/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// EventManagerAssetsCreateV2AppAssetAppType
type EventManagerAssetsCreateV2AppAssetAppType string

// List of event_manager_assets_create_v2_app_asset_app_type
const (
	ANDROID_EventManagerAssetsCreateV2AppAssetAppType EventManagerAssetsCreateV2AppAssetAppType = "Android"
	IOS_EventManagerAssetsCreateV2AppAssetAppType     EventManagerAssetsCreateV2AppAssetAppType = "IOS"
)

// All allowed values of EventManagerAssetsCreateV2AppAssetAppType enum
var AllowedEventManagerAssetsCreateV2AppAssetAppTypeEnumValues = []EventManagerAssetsCreateV2AppAssetAppType{
	"Android",
	"IOS",
}

// NewEventManagerAssetsCreateV2AppAssetAppTypeFromValue returns a pointer to a valid EventManagerAssetsCreateV2AppAssetAppType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerAssetsCreateV2AppAssetAppTypeFromValue(v string) (*EventManagerAssetsCreateV2AppAssetAppType, error) {
	ev := EventManagerAssetsCreateV2AppAssetAppType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerAssetsCreateV2AppAssetAppType: valid values are %v", v, AllowedEventManagerAssetsCreateV2AppAssetAppTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerAssetsCreateV2AppAssetAppType) IsValid() bool {
	for _, existing := range AllowedEventManagerAssetsCreateV2AppAssetAppTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_assets_create_v2_app_asset_app_type value
func (v EventManagerAssetsCreateV2AppAssetAppType) Ptr() *EventManagerAssetsCreateV2AppAssetAppType {
	return &v
}
