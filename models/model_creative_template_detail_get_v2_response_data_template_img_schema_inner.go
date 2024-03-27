/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeTemplateDetailGetV2ResponseDataTemplateImgSchemaInner struct for CreativeTemplateDetailGetV2ResponseDataTemplateImgSchemaInner
type CreativeTemplateDetailGetV2ResponseDataTemplateImgSchemaInner struct {
	// 图片建议的大于等于的高度
	ImgHeight *int64 `json:"img_height,omitempty"`
	// 图片建议的大于等于的宽度
	ImgWidth *int64 `json:"img_width,omitempty"`
	// 填充内容的键
	Key *string `json:"key,omitempty"`
	// 填充内容的名称，e.g. 背景图
	Name *string `json:"name,omitempty"`
	// 默认填充内容的值（图片类型的填充值为图片url）
	Value *string `json:"value,omitempty"`
}
