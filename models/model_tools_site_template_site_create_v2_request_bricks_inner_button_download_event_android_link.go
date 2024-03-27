/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateSiteCreateV2RequestBricksInnerButtonDownloadEventAndroidLink android链接信息
type ToolsSiteTemplateSiteCreateV2RequestBricksInnerButtonDownloadEventAndroidLink struct {
	// 应用描述，为了展示效果，推荐12个中文字符长度
	Description *string                                                                   `json:"description,omitempty"`
	LinkType    ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventAndroidLinkLinkType `json:"link_type"`
	// 快应用地址，当`link_type`为QUICK_APP是，必填
	QuickApp *string `json:"quick_app,omitempty"`
	// 链接地址，当`link_type`为URL时，必填，
	Url *string `json:"url,omitempty"`
}
