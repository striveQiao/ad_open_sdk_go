/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupDetailV30ResponseDataDataAdInfoProductInfoInner struct for AdlabGroupDetailV30ResponseDataDataAdInfoProductInfoInner
type AdlabGroupDetailV30ResponseDataDataAdInfoProductInfoInner struct {
	// 投放条件id
	AssetId *int64 `json:"asset_id,omitempty"`
	// 商品id
	ProductId *int64 `json:"product_id,omitempty"`
	// 商品库id
	ProductPlatformId *int64 `json:"product_platform_id,omitempty"`
	// 附属商品信息
	Supplements []*AdlabGroupDetailV30ResponseDataDataAdInfoProductInfoInnerSupplementsInner `json:"supplements,omitempty"`
}
