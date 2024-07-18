/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerEventConfigsGetV2ResponseDataEventConfigsInner struct for EventManagerEventConfigsGetV2ResponseDataEventConfigsInner
type EventManagerEventConfigsGetV2ResponseDataEventConfigsInner struct {
	AttributionConfiguration *EventManagerEventConfigsGetV2ResponseDataEventConfigsInnerAttributionConfiguration `json:"attribution_configuration,omitempty"`
	// 事件创建时间
	CreateTime      *string                                                       `json:"create_time,omitempty"`
	DebuggingStatus *EventManagerEventConfigsGetV2DataEventConfigsDebuggingStatus `json:"debugging_status,omitempty"`
	// 事件中文名称
	EventCnName *string `json:"event_cn_name,omitempty"`
	// 事件ID
	EventId *int64 `json:"event_id,omitempty"`
	// 事件类型
	EventType *string `json:"event_type,omitempty"`
	// 事件的附加属性
	Properties []*EventManagerEventConfigsGetV2ResponseDataEventConfigsInnerPropertiesInner `json:"properties,omitempty"`
	// 事件回传方式列表，允许值:落地页支持:`JSSDK` JS埋码 、`EXTERNAL_API` API回传、`XPATH` XPath圈选应用支持：`APPLICATION_API` 应用API、`APPLICATION_SDK` 应用SDK、快应用支持：`QUICK_APP_API` 快应用API
	TrackTypes []*EventManagerEventConfigsGetV2DataEventConfigsTrackTypes `json:"track_types,omitempty"`
}
