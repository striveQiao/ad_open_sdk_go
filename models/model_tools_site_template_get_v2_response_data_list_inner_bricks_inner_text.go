/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerText 文本组件描述
type ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerText struct {
	// 文本内容，长度至少为1
	Content string                                                             `json:"content"`
	LinkDto *ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerTextLinkDto `json:"link_dto,omitempty"`
}
