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

// QianchuanTodayLiveRoomProductListGetV10DataListExplainStatus
type QianchuanTodayLiveRoomProductListGetV10DataListExplainStatus string

// List of qianchuan_today_live_room_product_list_get_v1.0_data_list_explain_status
const (
	BEINGEXPLAIN_QianchuanTodayLiveRoomProductListGetV10DataListExplainStatus QianchuanTodayLiveRoomProductListGetV10DataListExplainStatus = "BEINGEXPLAIN"
	HASEXPLAIN_QianchuanTodayLiveRoomProductListGetV10DataListExplainStatus   QianchuanTodayLiveRoomProductListGetV10DataListExplainStatus = "HASEXPLAIN"
	UNEXPLAIN_QianchuanTodayLiveRoomProductListGetV10DataListExplainStatus    QianchuanTodayLiveRoomProductListGetV10DataListExplainStatus = "UNEXPLAIN"
)

// All allowed values of QianchuanTodayLiveRoomProductListGetV10DataListExplainStatus enum
var AllowedQianchuanTodayLiveRoomProductListGetV10DataListExplainStatusEnumValues = []QianchuanTodayLiveRoomProductListGetV10DataListExplainStatus{
	"BEINGEXPLAIN",
	"HASEXPLAIN",
	"UNEXPLAIN",
}

// NewQianchuanTodayLiveRoomProductListGetV10DataListExplainStatusFromValue returns a pointer to a valid QianchuanTodayLiveRoomProductListGetV10DataListExplainStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanTodayLiveRoomProductListGetV10DataListExplainStatusFromValue(v string) (*QianchuanTodayLiveRoomProductListGetV10DataListExplainStatus, error) {
	ev := QianchuanTodayLiveRoomProductListGetV10DataListExplainStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanTodayLiveRoomProductListGetV10DataListExplainStatus: valid values are %v", v, AllowedQianchuanTodayLiveRoomProductListGetV10DataListExplainStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanTodayLiveRoomProductListGetV10DataListExplainStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanTodayLiveRoomProductListGetV10DataListExplainStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_today_live_room_product_list_get_v1.0_data_list_explain_status value
func (v QianchuanTodayLiveRoomProductListGetV10DataListExplainStatus) Ptr() *QianchuanTodayLiveRoomProductListGetV10DataListExplainStatus {
	return &v
}
