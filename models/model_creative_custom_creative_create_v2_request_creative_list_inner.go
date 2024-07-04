/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeCustomCreativeCreateV2RequestCreativeListInner struct for CreativeCustomCreativeCreateV2RequestCreativeListInner
type CreativeCustomCreativeCreateV2RequestCreativeListInner struct {
	//
	AbstractMaterials []*CreativeCustomCreativeCreateV2RequestCreativeListInnerAbstractMaterialsInner `json:"abstract_materials,omitempty"`
	//
	ComponentMaterials []*CreativeCustomCreativeCreateV2RequestCreativeListInnerComponentMaterialsInner `json:"component_materials,omitempty"`
	//
	CreativeId         *int64                                                                    `json:"creative_id,omitempty"`
	DecorationMaterial *CreativeCustomCreativeCreateV2RequestCreativeListInnerDecorationMaterial `json:"decoration_material,omitempty"`
	DerivePosterCid    *CreativeCustomCreativeCreateV2CreativeListDerivePosterCid                `json:"derive_poster_cid,omitempty"`
	//
	ImageMaterials      []*CreativeCustomCreativeCreateV2RequestCreativeListInnerImageMaterialsInner `json:"image_materials,omitempty"`
	ImageMode           CreativeCustomCreativeCreateV2CreativeListImageMode                          `json:"image_mode"`
	InteractiveMaterial *CreativeCustomCreativeCreateV2RequestCreativeListInnerInteractiveMaterial   `json:"interactive_material,omitempty"`
	PlayableMaterial    *CreativeCustomCreativeCreateV2RequestCreativeListInnerPlayableMaterial      `json:"playable_material,omitempty"`
	SubTitleMaterial    *CreativeCustomCreativeCreateV2RequestCreativeListInnerSubTitleMaterial      `json:"sub_title_material,omitempty"`
	//
	ThirdPartyId  *string                                                              `json:"third_party_id,omitempty"`
	TitleMaterial *CreativeCustomCreativeCreateV2RequestCreativeListInnerTitleMaterial `json:"title_material,omitempty"`
	VideoMaterial *CreativeCustomCreativeCreateV2RequestCreativeListInnerVideoMaterial `json:"video_material,omitempty"`
}
