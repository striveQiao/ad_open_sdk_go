/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// EventManagerOptimizedGoalGetV2V30MicroPromotionType
type EventManagerOptimizedGoalGetV2V30MicroPromotionType string

// List of event_manager_optimized_goal_get_v2_v3.0_micro_promotion_type
const (
	BYTE_APP_EventManagerOptimizedGoalGetV2V30MicroPromotionType    EventManagerOptimizedGoalGetV2V30MicroPromotionType = "BYTE_APP"
	BYTE_GAME_EventManagerOptimizedGoalGetV2V30MicroPromotionType   EventManagerOptimizedGoalGetV2V30MicroPromotionType = "BYTE_GAME"
	WECHAT_APP_EventManagerOptimizedGoalGetV2V30MicroPromotionType  EventManagerOptimizedGoalGetV2V30MicroPromotionType = "WECHAT_APP"
	WECHAT_GAME_EventManagerOptimizedGoalGetV2V30MicroPromotionType EventManagerOptimizedGoalGetV2V30MicroPromotionType = "WECHAT_GAME"
)

// All allowed values of EventManagerOptimizedGoalGetV2V30MicroPromotionType enum
var AllowedEventManagerOptimizedGoalGetV2V30MicroPromotionTypeEnumValues = []EventManagerOptimizedGoalGetV2V30MicroPromotionType{
	"BYTE_APP",
	"BYTE_GAME",
	"WECHAT_APP",
	"WECHAT_GAME",
}

// NewEventManagerOptimizedGoalGetV2V30MicroPromotionTypeFromValue returns a pointer to a valid EventManagerOptimizedGoalGetV2V30MicroPromotionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerOptimizedGoalGetV2V30MicroPromotionTypeFromValue(v string) (*EventManagerOptimizedGoalGetV2V30MicroPromotionType, error) {
	ev := EventManagerOptimizedGoalGetV2V30MicroPromotionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerOptimizedGoalGetV2V30MicroPromotionType: valid values are %v", v, AllowedEventManagerOptimizedGoalGetV2V30MicroPromotionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerOptimizedGoalGetV2V30MicroPromotionType) IsValid() bool {
	for _, existing := range AllowedEventManagerOptimizedGoalGetV2V30MicroPromotionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_optimized_goal_get_v2_v3.0_micro_promotion_type value
func (v EventManagerOptimizedGoalGetV2V30MicroPromotionType) Ptr() *EventManagerOptimizedGoalGetV2V30MicroPromotionType {
	return &v
}
