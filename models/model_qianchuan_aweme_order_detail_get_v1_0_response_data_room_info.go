/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeOrderDetailGetV10ResponseDataRoomInfo 直播间信息
type QianchuanAwemeOrderDetailGetV10ResponseDataRoomInfo struct {
	// 直播间封面
	RoomCover *string `json:"room_cover,omitempty"`
	// 直播间id
	RoomId     *int64                                                 `json:"room_id,omitempty"`
	RoomStatus *QianchuanAwemeOrderDetailGetV10DataRoomInfoRoomStatus `json:"room_status,omitempty"`
	// 直播间标题
	RoomTitle *string `json:"room_title,omitempty"`
}
