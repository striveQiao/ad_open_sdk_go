/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// EventManagerOptimizedGoalGetV2V30LandingType
type EventManagerOptimizedGoalGetV2V30LandingType string

// List of event_manager_optimized_goal_get_v2_v3.0_landing_type
const (
	APP_EventManagerOptimizedGoalGetV2V30LandingType           EventManagerOptimizedGoalGetV2V30LandingType = "APP"
	DPA_EventManagerOptimizedGoalGetV2V30LandingType           EventManagerOptimizedGoalGetV2V30LandingType = "DPA"
	LINK_EventManagerOptimizedGoalGetV2V30LandingType          EventManagerOptimizedGoalGetV2V30LandingType = "LINK"
	MICRO_GAME_EventManagerOptimizedGoalGetV2V30LandingType    EventManagerOptimizedGoalGetV2V30LandingType = "MICRO_GAME"
	NATIVE_ACTION_EventManagerOptimizedGoalGetV2V30LandingType EventManagerOptimizedGoalGetV2V30LandingType = "NATIVE_ACTION"
	QUICK_APP_EventManagerOptimizedGoalGetV2V30LandingType     EventManagerOptimizedGoalGetV2V30LandingType = "QUICK_APP"
	SHOP_EventManagerOptimizedGoalGetV2V30LandingType          EventManagerOptimizedGoalGetV2V30LandingType = "SHOP"
)

// All allowed values of EventManagerOptimizedGoalGetV2V30LandingType enum
var AllowedEventManagerOptimizedGoalGetV2V30LandingTypeEnumValues = []EventManagerOptimizedGoalGetV2V30LandingType{
	"APP",
	"DPA",
	"LINK",
	"MICRO_GAME",
	"NATIVE_ACTION",
	"QUICK_APP",
	"SHOP",
}

// NewEventManagerOptimizedGoalGetV2V30LandingTypeFromValue returns a pointer to a valid EventManagerOptimizedGoalGetV2V30LandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerOptimizedGoalGetV2V30LandingTypeFromValue(v string) (*EventManagerOptimizedGoalGetV2V30LandingType, error) {
	ev := EventManagerOptimizedGoalGetV2V30LandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerOptimizedGoalGetV2V30LandingType: valid values are %v", v, AllowedEventManagerOptimizedGoalGetV2V30LandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerOptimizedGoalGetV2V30LandingType) IsValid() bool {
	for _, existing := range AllowedEventManagerOptimizedGoalGetV2V30LandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_optimized_goal_get_v2_v3.0_landing_type value
func (v EventManagerOptimizedGoalGetV2V30LandingType) Ptr() *EventManagerOptimizedGoalGetV2V30LandingType {
	return &v
}
