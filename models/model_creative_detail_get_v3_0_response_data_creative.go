/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeDetailGetV30ResponseDataCreative 程序化素材信息
type CreativeDetailGetV30ResponseDataCreative struct {
	// 摘要素材
	AbstractMaterials []*CreativeDetailGetV30ResponseDataCreativeAbstractMaterialsInner `json:"abstract_materials,omitempty"`
	// 组件信息
	ComponentMaterials []*CreativeDetailGetV30ResponseDataCreativeComponentMaterialsInner `json:"component_materials,omitempty"`
	DecorationMaterial *CreativeDetailGetV30ResponseDataCreativeDecorationMaterial        `json:"decoration_material,omitempty"`
	//
	ImageMaterials   []*CreativeDetailGetV30ResponseDataCreativeImageMaterialsInner `json:"image_materials,omitempty"`
	SubTitleMaterial *CreativeDetailGetV30ResponseDataCreativeSubTitleMaterial      `json:"sub_title_material,omitempty"`
	// 创意标题素材
	TitleMaterials []*CreativeDetailGetV30ResponseDataCreativeTitleMaterialsInner `json:"title_materials,omitempty"`
	// 视频素材信息
	VideoMaterials []*CreativeDetailGetV30ResponseDataCreativeVideoMaterialsInner `json:"video_materials,omitempty"`
}
