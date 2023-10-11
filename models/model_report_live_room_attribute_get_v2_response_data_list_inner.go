/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportLiveRoomAttributeGetV2ResponseDataListInner struct for ReportLiveRoomAttributeGetV2ResponseDataListInner
type ReportLiveRoomAttributeGetV2ResponseDataListInner struct {
	//
	AnchorAvatar *string `json:"anchor_avatar,omitempty"`
	//
	AnchorId *int64 `json:"anchor_id,omitempty"`
	//
	AnchorNick *string `json:"anchor_nick,omitempty"`
	//
	RoomCover *string `json:"room_cover,omitempty"`
	//
	RoomCreateTime **string `json:"room_create_time,omitempty"`
	//
	RoomFinishTime **string `json:"room_finish_time,omitempty"`
	//
	RoomId *int64 `json:"room_id,omitempty"`
	//
	RoomQrcode *string                                         `json:"room_qrcode,omitempty"`
	RoomStatus *ReportLiveRoomAttributeGetV2DataListRoomStatus `json:"room_status,omitempty"`
	//
	RoomTitle *string `json:"room_title,omitempty"`
}
