/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionAidGetV30ResponseDataPromotionMapDataInnerAdsInnerCreativesInner struct for PromotionAidGetV30ResponseDataPromotionMapDataInnerAdsInnerCreativesInner
type PromotionAidGetV30ResponseDataPromotionMapDataInnerAdsInnerCreativesInner struct {
	//
	CreativeId     *int64                                                                                   `json:"creative_id,omitempty"`
	ImageMaterials *PromotionAidGetV30ResponseDataPromotionMapDataInnerAdsInnerCreativesInnerImageMaterials `json:"image_materials,omitempty"`
	TitleMaterial  *PromotionAidGetV30ResponseDataPromotionMapDataInnerAdsInnerCreativesInnerTitleMaterial  `json:"title_material,omitempty"`
	VideoMaterial  *PromotionAidGetV30ResponseDataPromotionMapDataInnerAdsInnerCreativesInnerVideoMaterial  `json:"video_material,omitempty"`
}
