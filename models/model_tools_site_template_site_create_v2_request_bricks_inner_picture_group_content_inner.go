/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateSiteCreateV2RequestBricksInnerPictureGroupContentInner struct for ToolsSiteTemplateSiteCreateV2RequestBricksInnerPictureGroupContentInner
type ToolsSiteTemplateSiteCreateV2RequestBricksInnerPictureGroupContentInner struct {
	LinkDto *ToolsSiteTemplateSiteCreateV2RequestBricksInnerPictureGroupContentInnerLinkDto `json:"link_dto,omitempty"`
	// 标签，用户自定义标注
	Tag *string `json:"tag,omitempty"`
	// 用户自行上传的图片url，当content不为空时，必填
	Url string `json:"url"`
}
