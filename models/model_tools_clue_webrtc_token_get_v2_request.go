/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueWebrtcTokenGetV2Request struct for ToolsClueWebrtcTokenGetV2Request
type ToolsClueWebrtcTokenGetV2Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 客服ID。平台下客服ID不得重复。
	UserId int64 `json:"user_id"`
}
