/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerPicture 图片组件描述
type ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerPicture struct {
	LinkDto *ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerPictureLinkDto `json:"link_dto,omitempty"`
	// 标签，用户自定义标注
	Tag *string `json:"tag,omitempty"`
	// 用户自行上传的图片url，当`picture`不为空时，有值
	Url string `json:"url"`
}
