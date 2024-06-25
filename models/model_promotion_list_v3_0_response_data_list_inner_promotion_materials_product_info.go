/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionListV30ResponseDataListInnerPromotionMaterialsProductInfo
type PromotionListV30ResponseDataListInnerPromotionMaterialsProductInfo struct {
	//
	ImageIds []string `json:"image_ids,omitempty"`
	//
	ProductImageFields []string                                                               `json:"product_image_fields,omitempty"`
	ProductImageType   *PromotionListV30DataListPromotionMaterialsProductInfoProductImageType `json:"product_image_type,omitempty"`
	//
	ProductNameFields []string                                                              `json:"product_name_fields,omitempty"`
	ProductNameType   *PromotionListV30DataListPromotionMaterialsProductInfoProductNameType `json:"product_name_type,omitempty"`
	//
	ProductSellingPointFields []string                                                                      `json:"product_selling_point_fields,omitempty"`
	ProductSellingPointType   *PromotionListV30DataListPromotionMaterialsProductInfoProductSellingPointType `json:"product_selling_point_type,omitempty"`
	//
	SellingPoints []string `json:"selling_points,omitempty"`
	//
	Titles []string `json:"titles,omitempty"`
}
