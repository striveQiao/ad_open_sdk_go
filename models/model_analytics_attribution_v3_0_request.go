/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AnalyticsAttributionV30Request struct for AnalyticsAttributionV30Request
type AnalyticsAttributionV30Request struct {
	// 开发者申请的应用APP_ID
	AppId int64 `json:"app_id"`
	// 当前场景只支持convert，勿传其他值
	AttributeLabel *string `json:"attribute_label,omitempty"`
	// 当前场景只支持4，勿传其他值
	BizType *int64                                 `json:"biz_type,omitempty"`
	Context *AnalyticsAttributionV30RequestContext `json:"context,omitempty"`
	// 事件类型
	EventType  *string                                   `json:"event_type,omitempty"`
	Properties *AnalyticsAttributionV30RequestProperties `json:"properties,omitempty"`
	// 广告主标识，自定义值，用于标识数据来源，例如：tb
	Source *string `json:"source,omitempty"`
	// 订单时间；这里timestamp指订单发生的时间，而不是数据上报时间，必须上报秒级时间戳；且回传超过1天的下单会返回错误
	Timestamp *int64 `json:"timestamp,omitempty"`
}
