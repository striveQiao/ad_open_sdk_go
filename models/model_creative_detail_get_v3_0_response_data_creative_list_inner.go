/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeDetailGetV30ResponseDataCreativeListInner struct for CreativeDetailGetV30ResponseDataCreativeListInner
type CreativeDetailGetV30ResponseDataCreativeListInner struct {
	// 摘要素材
	AbstractMaterials []*CreativeDetailGetV30ResponseDataCreativeListInnerAbstractMaterialsInner `json:"abstract_materials,omitempty"`
	// 创意id
	CreativeId         *int64                                                               `json:"creative_id,omitempty"`
	DecorationMaterial *CreativeDetailGetV30ResponseDataCreativeListInnerDecorationMaterial `json:"decoration_material,omitempty"`
	// 创意图片素材
	ImageMaterials      []*CreativeDetailGetV30ResponseDataCreativeListInnerImageMaterialsInner `json:"image_materials,omitempty"`
	ImageMode           *CreativeDetailGetV30DataCreativeListImageMode                          `json:"image_mode,omitempty"`
	InteractiveMaterial *CreativeDetailGetV30ResponseDataCreativeListInnerInteractiveMaterial   `json:"interactive_material,omitempty"`
	PlayableMaterial    *CreativeDetailGetV30ResponseDataCreativeListInnerPlayableMaterial      `json:"playable_material,omitempty"`
	TitleMaterial       *CreativeDetailGetV30ResponseDataCreativeListInnerTitleMaterial         `json:"title_material,omitempty"`
	VideoMaterial       *CreativeDetailGetV30ResponseDataCreativeListInnerVideoMaterial         `json:"video_material,omitempty"`
}
