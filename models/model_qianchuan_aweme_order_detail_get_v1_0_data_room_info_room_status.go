/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeOrderDetailGetV10DataRoomInfoRoomStatus
type QianchuanAwemeOrderDetailGetV10DataRoomInfoRoomStatus string

// List of qianchuan_aweme_order_detail_get_v1.0_data_room_info_room_status
const (
	FINISH_QianchuanAwemeOrderDetailGetV10DataRoomInfoRoomStatus  QianchuanAwemeOrderDetailGetV10DataRoomInfoRoomStatus = "FINISH"
	LIVING_QianchuanAwemeOrderDetailGetV10DataRoomInfoRoomStatus  QianchuanAwemeOrderDetailGetV10DataRoomInfoRoomStatus = "LIVING"
	PAUSE_QianchuanAwemeOrderDetailGetV10DataRoomInfoRoomStatus   QianchuanAwemeOrderDetailGetV10DataRoomInfoRoomStatus = "PAUSE"
	PREPARE_QianchuanAwemeOrderDetailGetV10DataRoomInfoRoomStatus QianchuanAwemeOrderDetailGetV10DataRoomInfoRoomStatus = "PREPARE"
	UNKNOW_QianchuanAwemeOrderDetailGetV10DataRoomInfoRoomStatus  QianchuanAwemeOrderDetailGetV10DataRoomInfoRoomStatus = "UNKNOW"
)

// All allowed values of QianchuanAwemeOrderDetailGetV10DataRoomInfoRoomStatus enum
var AllowedQianchuanAwemeOrderDetailGetV10DataRoomInfoRoomStatusEnumValues = []QianchuanAwemeOrderDetailGetV10DataRoomInfoRoomStatus{
	"FINISH",
	"LIVING",
	"PAUSE",
	"PREPARE",
	"UNKNOW",
}

// NewQianchuanAwemeOrderDetailGetV10DataRoomInfoRoomStatusFromValue returns a pointer to a valid QianchuanAwemeOrderDetailGetV10DataRoomInfoRoomStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderDetailGetV10DataRoomInfoRoomStatusFromValue(v string) (*QianchuanAwemeOrderDetailGetV10DataRoomInfoRoomStatus, error) {
	ev := QianchuanAwemeOrderDetailGetV10DataRoomInfoRoomStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderDetailGetV10DataRoomInfoRoomStatus: valid values are %v", v, AllowedQianchuanAwemeOrderDetailGetV10DataRoomInfoRoomStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderDetailGetV10DataRoomInfoRoomStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderDetailGetV10DataRoomInfoRoomStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_detail_get_v1.0_data_room_info_room_status value
func (v QianchuanAwemeOrderDetailGetV10DataRoomInfoRoomStatus) Ptr() *QianchuanAwemeOrderDetailGetV10DataRoomInfoRoomStatus {
	return &v
}
