/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeTemplateListGetV2ResponseDataListsInnerTemplateTagsInner struct for CreativeTemplateListGetV2ResponseDataListsInnerTemplateTagsInner
type CreativeTemplateListGetV2ResponseDataListsInnerTemplateTagsInner struct {
	// 模板标签ID
	TemplateTagId *string `json:"template_tag_id,omitempty"`
	// 模板标签名称，e.g. \"价格优惠\"、\"节日氛围\"、\"玩法吸引力\"、\"利益点露出\"
	TemplateTagName *string `json:"template_tag_name,omitempty"`
}
