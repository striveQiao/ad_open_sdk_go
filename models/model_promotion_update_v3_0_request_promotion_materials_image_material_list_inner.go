/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionUpdateV30RequestPromotionMaterialsImageMaterialListInner struct for PromotionUpdateV30RequestPromotionMaterialsImageMaterialListInner
type PromotionUpdateV30RequestPromotionMaterialsImageMaterialListInner struct {
	ImageMode PromotionUpdateV30PromotionMaterialsImageMaterialListImageMode `json:"image_mode"`
	//
	Images []*PromotionUpdateV30RequestPromotionMaterialsImageMaterialListInnerImagesInner `json:"images"`
}
