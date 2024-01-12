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

// EventManagerAssetsCreateV2AppAssetAppCreateType
type EventManagerAssetsCreateV2AppAssetAppCreateType string

// List of event_manager_assets_create_v2_app_asset_app_create_type
const (
	NORMAL_EventManagerAssetsCreateV2AppAssetAppCreateType EventManagerAssetsCreateV2AppAssetAppCreateType = "NORMAL"
	UG_EventManagerAssetsCreateV2AppAssetAppCreateType     EventManagerAssetsCreateV2AppAssetAppCreateType = "UG"
)

// All allowed values of EventManagerAssetsCreateV2AppAssetAppCreateType enum
var AllowedEventManagerAssetsCreateV2AppAssetAppCreateTypeEnumValues = []EventManagerAssetsCreateV2AppAssetAppCreateType{
	"NORMAL",
	"UG",
}

// NewEventManagerAssetsCreateV2AppAssetAppCreateTypeFromValue returns a pointer to a valid EventManagerAssetsCreateV2AppAssetAppCreateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerAssetsCreateV2AppAssetAppCreateTypeFromValue(v string) (*EventManagerAssetsCreateV2AppAssetAppCreateType, error) {
	ev := EventManagerAssetsCreateV2AppAssetAppCreateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerAssetsCreateV2AppAssetAppCreateType: valid values are %v", v, AllowedEventManagerAssetsCreateV2AppAssetAppCreateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerAssetsCreateV2AppAssetAppCreateType) IsValid() bool {
	for _, existing := range AllowedEventManagerAssetsCreateV2AppAssetAppCreateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_assets_create_v2_app_asset_app_create_type value
func (v EventManagerAssetsCreateV2AppAssetAppCreateType) Ptr() *EventManagerAssetsCreateV2AppAssetAppCreateType {
	return &v
}
