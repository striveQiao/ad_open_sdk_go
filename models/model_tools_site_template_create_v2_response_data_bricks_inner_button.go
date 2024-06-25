/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateCreateV2ResponseDataBricksInnerButton 按钮组件描述
type ToolsSiteTemplateCreateV2ResponseDataBricksInnerButton struct {
	AppointEvent *ToolsSiteTemplateCreateV2ResponseDataBricksInnerButtonAppointEvent `json:"appoint_event,omitempty"`
	// linkEvent事件自定义图片链接
	BgImageUrl    *string                                                              `json:"bg_image_url,omitempty"`
	DownloadEvent *ToolsSiteTemplateCreateV2ResponseDataBricksInnerButtonDownloadEvent `json:"download_event,omitempty"`
	EventType     ToolsSiteTemplateCreateV2DataBricksButtonEventType                   `json:"event_type"`
	LinkEvent     *ToolsSiteTemplateCreateV2ResponseDataBricksInnerButtonLinkEvent     `json:"link_event,omitempty"`
	PhoneEvent    *ToolsSiteTemplateCreateV2ResponseDataBricksInnerButtonPhoneEvent    `json:"phone_event,omitempty"`
}
