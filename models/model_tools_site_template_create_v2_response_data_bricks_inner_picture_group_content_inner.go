/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateCreateV2ResponseDataBricksInnerPictureGroupContentInner struct for ToolsSiteTemplateCreateV2ResponseDataBricksInnerPictureGroupContentInner
type ToolsSiteTemplateCreateV2ResponseDataBricksInnerPictureGroupContentInner struct {
	LinkDto *ToolsSiteTemplateCreateV2ResponseDataBricksInnerPictureGroupContentInnerLinkDto `json:"link_dto,omitempty"`
	// 标签，用户自定义标注
	Tag *string `json:"tag,omitempty"`
	// 用户自行上传的图片url，当content不为空时，有值
	Url string `json:"url"`
}
