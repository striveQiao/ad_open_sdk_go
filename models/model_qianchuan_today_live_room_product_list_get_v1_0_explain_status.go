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

// QianchuanTodayLiveRoomProductListGetV10ExplainStatus
type QianchuanTodayLiveRoomProductListGetV10ExplainStatus string

// List of qianchuan_today_live_room_product_list_get_v1.0_explain_status
const (
	ALL_QianchuanTodayLiveRoomProductListGetV10ExplainStatus        QianchuanTodayLiveRoomProductListGetV10ExplainStatus = "ALL"
	HASEXPLAIN_QianchuanTodayLiveRoomProductListGetV10ExplainStatus QianchuanTodayLiveRoomProductListGetV10ExplainStatus = "HASEXPLAIN"
	UNEXPLAIN_QianchuanTodayLiveRoomProductListGetV10ExplainStatus  QianchuanTodayLiveRoomProductListGetV10ExplainStatus = "UNEXPLAIN"
)

// All allowed values of QianchuanTodayLiveRoomProductListGetV10ExplainStatus enum
var AllowedQianchuanTodayLiveRoomProductListGetV10ExplainStatusEnumValues = []QianchuanTodayLiveRoomProductListGetV10ExplainStatus{
	"ALL",
	"HASEXPLAIN",
	"UNEXPLAIN",
}

// NewQianchuanTodayLiveRoomProductListGetV10ExplainStatusFromValue returns a pointer to a valid QianchuanTodayLiveRoomProductListGetV10ExplainStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanTodayLiveRoomProductListGetV10ExplainStatusFromValue(v string) (*QianchuanTodayLiveRoomProductListGetV10ExplainStatus, error) {
	ev := QianchuanTodayLiveRoomProductListGetV10ExplainStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanTodayLiveRoomProductListGetV10ExplainStatus: valid values are %v", v, AllowedQianchuanTodayLiveRoomProductListGetV10ExplainStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanTodayLiveRoomProductListGetV10ExplainStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanTodayLiveRoomProductListGetV10ExplainStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_today_live_room_product_list_get_v1.0_explain_status value
func (v QianchuanTodayLiveRoomProductListGetV10ExplainStatus) Ptr() *QianchuanTodayLiveRoomProductListGetV10ExplainStatus {
	return &v
}
