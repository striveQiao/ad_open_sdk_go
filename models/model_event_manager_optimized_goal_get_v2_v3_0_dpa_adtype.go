/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// EventManagerOptimizedGoalGetV2V30DpaAdtype
type EventManagerOptimizedGoalGetV2V30DpaAdtype string

// List of event_manager_optimized_goal_get_v2_v3.0_dpa_adtype
const (
	DPA_APP_EventManagerOptimizedGoalGetV2V30DpaAdtype  EventManagerOptimizedGoalGetV2V30DpaAdtype = "DPA_APP"
	DPA_LINK_EventManagerOptimizedGoalGetV2V30DpaAdtype EventManagerOptimizedGoalGetV2V30DpaAdtype = "DPA_LINK"
)

// All allowed values of EventManagerOptimizedGoalGetV2V30DpaAdtype enum
var AllowedEventManagerOptimizedGoalGetV2V30DpaAdtypeEnumValues = []EventManagerOptimizedGoalGetV2V30DpaAdtype{
	"DPA_APP",
	"DPA_LINK",
}

// NewEventManagerOptimizedGoalGetV2V30DpaAdtypeFromValue returns a pointer to a valid EventManagerOptimizedGoalGetV2V30DpaAdtype
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerOptimizedGoalGetV2V30DpaAdtypeFromValue(v string) (*EventManagerOptimizedGoalGetV2V30DpaAdtype, error) {
	ev := EventManagerOptimizedGoalGetV2V30DpaAdtype(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerOptimizedGoalGetV2V30DpaAdtype: valid values are %v", v, AllowedEventManagerOptimizedGoalGetV2V30DpaAdtypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerOptimizedGoalGetV2V30DpaAdtype) IsValid() bool {
	for _, existing := range AllowedEventManagerOptimizedGoalGetV2V30DpaAdtypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_optimized_goal_get_v2_v3.0_dpa_adtype value
func (v EventManagerOptimizedGoalGetV2V30DpaAdtype) Ptr() *EventManagerOptimizedGoalGetV2V30DpaAdtype {
	return &v
}
