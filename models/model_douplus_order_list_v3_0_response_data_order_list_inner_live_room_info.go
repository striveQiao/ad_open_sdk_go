/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DouplusOrderListV30ResponseDataOrderListInnerLiveRoomInfo
type DouplusOrderListV30ResponseDataOrderListInnerLiveRoomInfo struct {
	//
	RoomCover []string `json:"room_cover,omitempty"`
	//
	RoomId     *int64                                                  `json:"room_id,omitempty"`
	RoomStatus *DouplusOrderListV30DataOrderListLiveRoomInfoRoomStatus `json:"room_status,omitempty"`
	//
	RoomTitle *string `json:"room_title,omitempty"`
}
