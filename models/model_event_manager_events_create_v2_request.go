/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerEventsCreateV2Request struct for EventManagerEventsCreateV2Request
type EventManagerEventsCreateV2Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 资产ID
	AssetId int64 `json:"asset_id"`
	// 事件ID
	EventId int64 `json:"event_id"`
	// 事件回传方式列表，允许值：`JSSDK` JS埋码 、`EXTERNAL_API` API回传 、`XPATH` XPath圈选
	TrackTypes []*EventManagerEventsCreateV2TrackTypes `json:"track_types"`
}
