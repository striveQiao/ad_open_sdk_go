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

// EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes
type EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes string

// List of event_manager_available_events_get_v2_data_event_configs_track_types
const (
	JSSDK_EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes             EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes = "JSSDK"
	XPATH_EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes             EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes = "XPATH"
	APPLICATION_API_EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes   EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes = "APPLICATION_API"
	EXTERNAL_API_EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes      EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes = "EXTERNAL_API"
	APPLICATION_SDK_EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes   EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes = "APPLICATION_SDK"
	MINI_PROGRAME_SDK_EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes = "MINI_PROGRAME_SDK"
	MINI_PROGRAME_API_EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes = "MINI_PROGRAME_API"
	OFFLINE_CALLBACK_EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes  EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes = "OFFLINE_CALLBACK"
	OFFLINE_API_EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes       EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes = "OFFLINE_API"
	APPLOG_EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes            EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes = "APPLOG"
	TETRIS_EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes            EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes = "TETRIS"
	TETRIS_FE_EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes         EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes = "TETRIS_FE"
	FLY_FISH_EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes          EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes = "FLY_FISH"
	E_CHAT_EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes            EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes = "E_CHAT"
	QUICK_APP_API_EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes     EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes = "QUICK_APP_API"
	CLUE_COMPONENT_EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes    EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes = "CLUE_COMPONENT"
)

// All allowed values of EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes enum
var AllowedEventManagerAvailableEventsGetV2DataEventConfigsTrackTypesEnumValues = []EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes{
	"JSSDK",
	"XPATH",
	"APPLICATION_API",
	"EXTERNAL_API",
	"APPLICATION_SDK",
	"MINI_PROGRAME_SDK",
	"MINI_PROGRAME_API",
	"OFFLINE_CALLBACK",
	"OFFLINE_API",
	"APPLOG",
	"TETRIS",
	"TETRIS_FE",
	"FLY_FISH",
	"E_CHAT",
	"QUICK_APP_API",
	"CLUE_COMPONENT",
}

// NewEventManagerAvailableEventsGetV2DataEventConfigsTrackTypesFromValue returns a pointer to a valid EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerAvailableEventsGetV2DataEventConfigsTrackTypesFromValue(v string) (*EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes, error) {
	ev := EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes: valid values are %v", v, AllowedEventManagerAvailableEventsGetV2DataEventConfigsTrackTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes) IsValid() bool {
	for _, existing := range AllowedEventManagerAvailableEventsGetV2DataEventConfigsTrackTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_available_events_get_v2_data_event_configs_track_types value
func (v EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes) Ptr() *EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes {
	return &v
}
