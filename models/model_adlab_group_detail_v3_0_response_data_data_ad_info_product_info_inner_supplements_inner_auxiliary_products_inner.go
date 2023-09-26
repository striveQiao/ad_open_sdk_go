/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupDetailV30ResponseDataDataAdInfoProductInfoInnerSupplementsInnerAuxiliaryProductsInner struct for AdlabGroupDetailV30ResponseDataDataAdInfoProductInfoInnerSupplementsInnerAuxiliaryProductsInner
type AdlabGroupDetailV30ResponseDataDataAdInfoProductInfoInnerSupplementsInnerAuxiliaryProductsInner struct {
	// 附属商品的投放条件id
	AssetId *int64 `json:"asset_id,omitempty"`
	// 附属商品的商品id
	ProductId *string `json:"product_id,omitempty"`
	// 附属商品的商品库id
	ProductPlatformId *int64 `json:"product_platform_id,omitempty"`
}
