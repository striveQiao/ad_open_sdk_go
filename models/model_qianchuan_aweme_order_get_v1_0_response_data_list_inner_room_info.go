/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeOrderGetV10ResponseDataListInnerRoomInfo 直播间信息
type QianchuanAwemeOrderGetV10ResponseDataListInnerRoomInfo struct {
	// 直播间封面
	RoomCover *string `json:"room_cover,omitempty"`
	// 直播间id
	RoomId     *int64                                               `json:"room_id,omitempty"`
	RoomStatus *QianchuanAwemeOrderGetV10DataListRoomInfoRoomStatus `json:"room_status,omitempty"`
	// 直播间标题
	RoomTitle *string `json:"room_title,omitempty"`
}
