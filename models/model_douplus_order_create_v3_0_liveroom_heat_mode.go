/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DouplusOrderCreateV30LiveroomHeatMode
type DouplusOrderCreateV30LiveroomHeatMode string

// List of douplus_order_create_v3.0_liveroom_heat_mode
const (
	BY_ROOM_DouplusOrderCreateV30LiveroomHeatMode  DouplusOrderCreateV30LiveroomHeatMode = "BY_ROOM"
	BY_VIDEO_DouplusOrderCreateV30LiveroomHeatMode DouplusOrderCreateV30LiveroomHeatMode = "BY_VIDEO"
)

// All allowed values of DouplusOrderCreateV30LiveroomHeatMode enum
var AllowedDouplusOrderCreateV30LiveroomHeatModeEnumValues = []DouplusOrderCreateV30LiveroomHeatMode{
	"BY_ROOM",
	"BY_VIDEO",
}

// NewDouplusOrderCreateV30LiveroomHeatModeFromValue returns a pointer to a valid DouplusOrderCreateV30LiveroomHeatMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDouplusOrderCreateV30LiveroomHeatModeFromValue(v string) (*DouplusOrderCreateV30LiveroomHeatMode, error) {
	ev := DouplusOrderCreateV30LiveroomHeatMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DouplusOrderCreateV30LiveroomHeatMode: valid values are %v", v, AllowedDouplusOrderCreateV30LiveroomHeatModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DouplusOrderCreateV30LiveroomHeatMode) IsValid() bool {
	for _, existing := range AllowedDouplusOrderCreateV30LiveroomHeatModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to douplus_order_create_v3.0_liveroom_heat_mode value
func (v DouplusOrderCreateV30LiveroomHeatMode) Ptr() *DouplusOrderCreateV30LiveroomHeatMode {
	return &v
}
