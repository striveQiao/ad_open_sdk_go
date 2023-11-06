/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateSiteCreateV2RequestBricksInnerButtonAppointEventLink 链接信息
type ToolsSiteTemplateSiteCreateV2RequestBricksInnerButtonAppointEventLink struct {
	LinkType ToolsSiteTemplateSiteCreateV2BricksButtonAppointEventLinkLinkType `json:"link_type"`
	// 快应用地址，当`link_type`为QUICK_APP是，必填
	QuickApp *string `json:"quick_app,omitempty"`
	// scheme地址，当`link_type`为SCHEME时，必填
	Scheme *string `json:"scheme,omitempty"`
	// 链接地址，当`link_type`为URL时，必填
	Url *string `json:"url,omitempty"`
}
