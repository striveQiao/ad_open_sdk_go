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

// EventManagerDeepBidTypeGetV30DeliveryMode
type EventManagerDeepBidTypeGetV30DeliveryMode string

// List of event_manager_deep_bid_type_get_v3.0_delivery_mode
const (
	MANUAL_EventManagerDeepBidTypeGetV30DeliveryMode     EventManagerDeepBidTypeGetV30DeliveryMode = "MANUAL"
	PROCEDURAL_EventManagerDeepBidTypeGetV30DeliveryMode EventManagerDeepBidTypeGetV30DeliveryMode = "PROCEDURAL"
)

// All allowed values of EventManagerDeepBidTypeGetV30DeliveryMode enum
var AllowedEventManagerDeepBidTypeGetV30DeliveryModeEnumValues = []EventManagerDeepBidTypeGetV30DeliveryMode{
	"MANUAL",
	"PROCEDURAL",
}

// NewEventManagerDeepBidTypeGetV30DeliveryModeFromValue returns a pointer to a valid EventManagerDeepBidTypeGetV30DeliveryMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerDeepBidTypeGetV30DeliveryModeFromValue(v string) (*EventManagerDeepBidTypeGetV30DeliveryMode, error) {
	ev := EventManagerDeepBidTypeGetV30DeliveryMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerDeepBidTypeGetV30DeliveryMode: valid values are %v", v, AllowedEventManagerDeepBidTypeGetV30DeliveryModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerDeepBidTypeGetV30DeliveryMode) IsValid() bool {
	for _, existing := range AllowedEventManagerDeepBidTypeGetV30DeliveryModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_deep_bid_type_get_v3.0_delivery_mode value
func (v EventManagerDeepBidTypeGetV30DeliveryMode) Ptr() *EventManagerDeepBidTypeGetV30DeliveryMode {
	return &v
}
