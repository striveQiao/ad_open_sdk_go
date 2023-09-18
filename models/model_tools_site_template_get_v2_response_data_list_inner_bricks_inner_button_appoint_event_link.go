/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerButtonAppointEventLink 链接信息
type ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerButtonAppointEventLink struct {
	LinkType ToolsSiteTemplateGetV2DataListBricksButtonAppointEventLinkLinkType `json:"link_type"`
	// 快应用地址，当`link_type`等于`QUICK_APP`时，有值
	QuickApp *string `json:"quick_app,omitempty"`
	// scheme地址，当`link_type`等于`SCHEME`时，有值
	Scheme *string `json:"scheme,omitempty"`
	// 链接地址，当`link_type`等于`URL`时，有值
	Url *string `json:"url,omitempty"`
}
