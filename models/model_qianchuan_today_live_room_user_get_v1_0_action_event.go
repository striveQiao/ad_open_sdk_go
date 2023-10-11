/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanTodayLiveRoomUserGetV10ActionEvent
type QianchuanTodayLiveRoomUserGetV10ActionEvent string

// List of qianchuan_today_live_room_user_get_v1.0_action_event
const (
	ENTER_QianchuanTodayLiveRoomUserGetV10ActionEvent QianchuanTodayLiveRoomUserGetV10ActionEvent = "ENTER"
	PAY_QianchuanTodayLiveRoomUserGetV10ActionEvent   QianchuanTodayLiveRoomUserGetV10ActionEvent = "PAY"
)

// All allowed values of QianchuanTodayLiveRoomUserGetV10ActionEvent enum
var AllowedQianchuanTodayLiveRoomUserGetV10ActionEventEnumValues = []QianchuanTodayLiveRoomUserGetV10ActionEvent{
	"ENTER",
	"PAY",
}

// NewQianchuanTodayLiveRoomUserGetV10ActionEventFromValue returns a pointer to a valid QianchuanTodayLiveRoomUserGetV10ActionEvent
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanTodayLiveRoomUserGetV10ActionEventFromValue(v string) (*QianchuanTodayLiveRoomUserGetV10ActionEvent, error) {
	ev := QianchuanTodayLiveRoomUserGetV10ActionEvent(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanTodayLiveRoomUserGetV10ActionEvent: valid values are %v", v, AllowedQianchuanTodayLiveRoomUserGetV10ActionEventEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanTodayLiveRoomUserGetV10ActionEvent) IsValid() bool {
	for _, existing := range AllowedQianchuanTodayLiveRoomUserGetV10ActionEventEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_today_live_room_user_get_v1.0_action_event value
func (v QianchuanTodayLiveRoomUserGetV10ActionEvent) Ptr() *QianchuanTodayLiveRoomUserGetV10ActionEvent {
	return &v
}
