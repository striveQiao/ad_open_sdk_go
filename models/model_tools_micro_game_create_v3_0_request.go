/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsMicroGameCreateV30Request struct for ToolsMicroGameCreateV30Request
type ToolsMicroGameCreateV30Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	AppId    string                                 `json:"app_id"`
	GameLink ToolsMicroGameCreateV30RequestGameLink `json:"game_link"`
	//
	Remark string `json:"remark"`
}
