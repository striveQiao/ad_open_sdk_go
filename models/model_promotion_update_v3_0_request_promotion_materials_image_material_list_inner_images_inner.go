/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionUpdateV30RequestPromotionMaterialsImageMaterialListInnerImagesInner struct for PromotionUpdateV30RequestPromotionMaterialsImageMaterialListInnerImagesInner
type PromotionUpdateV30RequestPromotionMaterialsImageMaterialListInnerImagesInner struct {
	//
	ImageId *string `json:"image_id,omitempty"`
	//
	TemplateDataList []*PromotionCreateV30RequestPromotionMaterialsImageMaterialListInnerImagesInnerTemplateDataListInner `json:"template_data_list,omitempty"`
	//
	TemplateId *int64 `json:"template_id,omitempty"`
}
