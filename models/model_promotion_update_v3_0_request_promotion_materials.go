/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionUpdateV30RequestPromotionMaterials
type PromotionUpdateV30RequestPromotionMaterials struct {
	//
	AdvancedDcSettings []*PromotionUpdateV30PromotionMaterialsAdvancedDcSettings `json:"advanced_dc_settings,omitempty"`
	//
	AnchorMaterialList        []*PromotionUpdateV30RequestPromotionMaterialsAnchorMaterialListInner `json:"anchor_material_list,omitempty"`
	BlueFlowMaterialRecommend *PromotionUpdateV30PromotionMaterialsBlueFlowMaterialRecommend        `json:"blue_flow_material_recommend,omitempty"`
	//
	CallToActionButtons []string `json:"call_to_action_buttons,omitempty"`
	//
	CarouselMaterialList []*PromotionUpdateV30RequestPromotionMaterialsCarouselMaterialListInner `json:"carousel_material_list,omitempty"`
	//
	ComponentMaterialList []*PromotionUpdateV30RequestPromotionMaterialsComponentMaterialListInner `json:"component_material_list,omitempty"`
	DecorationMaterial    *PromotionUpdateV30RequestPromotionMaterialsDecorationMaterial           `json:"decoration_material,omitempty"`
	DynamicCreativeSwitch *PromotionUpdateV30PromotionMaterialsDynamicCreativeSwitch               `json:"dynamic_creative_switch,omitempty"`
	//
	ExternalUrlField *string `json:"external_url_field,omitempty"`
	//
	ExternalUrlMaterialList []string `json:"external_url_material_list,omitempty"`
	//
	ExternalUrlParams *string `json:"external_url_params,omitempty"`
	//
	ImageMaterialList     []*PromotionUpdateV30RequestPromotionMaterialsImageMaterialListInner `json:"image_material_list,omitempty"`
	IntelligentGeneration *PromotionUpdateV30PromotionMaterialsIntelligentGeneration           `json:"intelligent_generation,omitempty"`
	MiniProgramInfo       *PromotionUpdateV30RequestPromotionMaterialsMiniProgramInfo          `json:"mini_program_info,omitempty"`
	//
	OpenUrl *string `json:"open_url,omitempty"`
	//
	OpenUrlField *string `json:"open_url_field,omitempty"`
	//
	OpenUrlParams *string `json:"open_url_params,omitempty"`
	//
	OpenUrls []string `json:"open_urls,omitempty"`
	//
	PlayableUrlMaterialList []string                                                `json:"playable_url_material_list,omitempty"`
	ProductInfo             *PromotionUpdateV30RequestPromotionMaterialsProductInfo `json:"product_info,omitempty"`
	//
	TextAbstractList []*PromotionUpdateV30RequestPromotionMaterialsTextAbstractListInner `json:"text_abstract_list,omitempty"`
	//
	TitleMaterialList []*PromotionUpdateV30RequestPromotionMaterialsTitleMaterialListInner `json:"title_material_list,omitempty"`
	//
	Ulink *string `json:"ulink,omitempty"`
	//
	VideoMaterialList []*PromotionUpdateV30RequestPromotionMaterialsVideoMaterialListInner `json:"video_material_list,omitempty"`
	//
	WebUrlMaterialList []string `json:"web_url_material_list,omitempty"`
}
