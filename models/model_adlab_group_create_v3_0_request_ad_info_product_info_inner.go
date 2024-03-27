/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupCreateV30RequestAdInfoProductInfoInner struct for AdlabGroupCreateV30RequestAdInfoProductInfoInner
type AdlabGroupCreateV30RequestAdInfoProductInfoInner struct {
	// 投放条件id
	AssetId *int64 `json:"asset_id,omitempty"`
	// 商品id
	ProductId *int64 `json:"product_id,omitempty"`
	// 商品库id
	ProductPlatformId *int64 `json:"product_platform_id,omitempty"`
	// 附属商品信息，当推广目的为LINK销售线索收集时可选
	Supplements []*AdlabGroupCreateV30RequestAdInfoProductInfoInnerSupplementsInner `json:"supplements,omitempty"`
}
