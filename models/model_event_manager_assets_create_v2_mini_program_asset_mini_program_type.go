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

// EventManagerAssetsCreateV2MiniProgramAssetMiniProgramType
type EventManagerAssetsCreateV2MiniProgramAssetMiniProgramType string

// List of event_manager_assets_create_v2_mini_program_asset_mini_program_type
const (
	BYTE_APP_EventManagerAssetsCreateV2MiniProgramAssetMiniProgramType  EventManagerAssetsCreateV2MiniProgramAssetMiniProgramType = "BYTE_APP"
	BYTE_GAME_EventManagerAssetsCreateV2MiniProgramAssetMiniProgramType EventManagerAssetsCreateV2MiniProgramAssetMiniProgramType = "BYTE_GAME"
)

// All allowed values of EventManagerAssetsCreateV2MiniProgramAssetMiniProgramType enum
var AllowedEventManagerAssetsCreateV2MiniProgramAssetMiniProgramTypeEnumValues = []EventManagerAssetsCreateV2MiniProgramAssetMiniProgramType{
	"BYTE_APP",
	"BYTE_GAME",
}

// NewEventManagerAssetsCreateV2MiniProgramAssetMiniProgramTypeFromValue returns a pointer to a valid EventManagerAssetsCreateV2MiniProgramAssetMiniProgramType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerAssetsCreateV2MiniProgramAssetMiniProgramTypeFromValue(v string) (*EventManagerAssetsCreateV2MiniProgramAssetMiniProgramType, error) {
	ev := EventManagerAssetsCreateV2MiniProgramAssetMiniProgramType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerAssetsCreateV2MiniProgramAssetMiniProgramType: valid values are %v", v, AllowedEventManagerAssetsCreateV2MiniProgramAssetMiniProgramTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerAssetsCreateV2MiniProgramAssetMiniProgramType) IsValid() bool {
	for _, existing := range AllowedEventManagerAssetsCreateV2MiniProgramAssetMiniProgramTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_assets_create_v2_mini_program_asset_mini_program_type value
func (v EventManagerAssetsCreateV2MiniProgramAssetMiniProgramType) Ptr() *EventManagerAssetsCreateV2MiniProgramAssetMiniProgramType {
	return &v
}
