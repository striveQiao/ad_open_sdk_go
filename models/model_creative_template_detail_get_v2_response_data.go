/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeTemplateDetailGetV2ResponseData
type CreativeTemplateDetailGetV2ResponseData struct {
	// 模板填充的图片内容列表
	TemplateImgSchema []*CreativeTemplateDetailGetV2ResponseDataTemplateImgSchemaInner `json:"template_img_schema,omitempty"`
	TemplateIndex     *CreativeTemplateDetailGetV2DataTemplateIndex                    `json:"template_index,omitempty"`
	TemplatePos       *CreativeTemplateDetailGetV2DataTemplatePos                      `json:"template_pos,omitempty"`
	// 当模板与素材的位置关系为自定义时，该值表示素材出现在模板中基于左上角[x,y]的坐标，单位为像素
	TemplatePosCustom []int64 `json:"template_pos_custom,omitempty"`
	// 模板填充的文本内容列表
	TemplateTextSchema []*CreativeTemplateDetailGetV2ResponseDataTemplateTextSchemaInner `json:"template_text_schema,omitempty"`
}
