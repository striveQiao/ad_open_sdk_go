/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerPictureGroupContentInner struct for ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerPictureGroupContentInner
type ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerPictureGroupContentInner struct {
	LinkDto *ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerPictureGroupContentInnerLinkDto `json:"link_dto,omitempty"`
	// 标签
	Tag *string `json:"tag,omitempty"`
	// 用户自行上传的图片url，当content不为空时，有值
	Url string `json:"url"`
}
