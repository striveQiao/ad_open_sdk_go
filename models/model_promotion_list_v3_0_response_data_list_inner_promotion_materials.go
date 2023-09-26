/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionListV30ResponseDataListInnerPromotionMaterials
type PromotionListV30ResponseDataListInnerPromotionMaterials struct {
	//
	AnchorMaterialList []*PromotionListV30ResponseDataListInnerPromotionMaterialsAnchorMaterialListInner `json:"anchor_material_list,omitempty"`
	//
	CallToActionButtons []string `json:"call_to_action_buttons,omitempty"`
	//
	CarouselMaterialList []*PromotionListV30ResponseDataListInnerPromotionMaterialsCarouselMaterialListInner `json:"carousel_material_list,omitempty"`
	//
	ComponentMaterialList []*PromotionListV30ResponseDataListInnerPromotionMaterialsComponentMaterialListInner `json:"component_material_list,omitempty"`
	DecorationMaterial    *PromotionListV30ResponseDataListInnerPromotionMaterialsDecorationMaterial           `json:"decoration_material,omitempty"`
	DynamicCreativeSwitch *PromotionListV30DataListPromotionMaterialsDynamicCreativeSwitch                     `json:"dynamic_creative_switch,omitempty"`
	//
	ExternalUrlField *string `json:"external_url_field,omitempty"`
	//
	ExternalUrlMaterialList []string `json:"external_url_material_list,omitempty"`
	//
	ExternalUrlParams *string `json:"external_url_params,omitempty"`
	//
	ImageMaterialList     []*PromotionListV30ResponseDataListInnerPromotionMaterialsImageMaterialListInner `json:"image_material_list,omitempty"`
	IntelligentGeneration *PromotionListV30DataListPromotionMaterialsIntelligentGeneration                 `json:"intelligent_generation,omitempty"`
	MiniProgramInfo       *PromotionListV30ResponseDataListInnerPromotionMaterialsMiniProgramInfo          `json:"mini_program_info,omitempty"`
	//
	OpenUrl *string `json:"open_url,omitempty"`
	//
	OpenUrlField *string `json:"open_url_field,omitempty"`
	//
	OpenUrlParams *string                                                `json:"open_url_params,omitempty"`
	OpenUrlType   *PromotionListV30DataListPromotionMaterialsOpenUrlType `json:"open_url_type,omitempty"`
	//
	OpenUrls   []string                                              `json:"open_urls,omitempty"`
	ParamsType *PromotionListV30DataListPromotionMaterialsParamsType `json:"params_type,omitempty"`
	//
	PlayableUrlMaterialList []string                                                            `json:"playable_url_material_list,omitempty"`
	ProductInfo             *PromotionListV30ResponseDataListInnerPromotionMaterialsProductInfo `json:"product_info,omitempty"`
	//
	TextAbstractList []*PromotionListV30ResponseDataListInnerPromotionMaterialsTextAbstractListInner `json:"text_abstract_list,omitempty"`
	//
	TitleMaterialList []*PromotionListV30ResponseDataListInnerPromotionMaterialsTitleMaterialListInner `json:"title_material_list,omitempty"`
	//
	Ulink *string `json:"ulink,omitempty"`
	//
	UlinkUrl *string `json:"ulink_url,omitempty"`
	//
	VideoMaterialList []*PromotionListV30ResponseDataListInnerPromotionMaterialsVideoMaterialListInner `json:"video_material_list,omitempty"`
	//
	WebUrlMaterialList []string `json:"web_url_material_list,omitempty"`
}
