/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionListV30ResponseDataListInnerPromotionMaterialsDecorationMaterial 家装卡券素材
type PromotionListV30ResponseDataListInnerPromotionMaterialsDecorationMaterial struct {
	// 卡券活动id
	ActivityId *string                                                                `json:"activity_id,omitempty"`
	ImageMode  *PromotionListV30DataListPromotionMaterialsDecorationMaterialImageMode `json:"image_mode,omitempty"`
	//
	MaterialId *int64 `json:"material_id,omitempty"`
}
