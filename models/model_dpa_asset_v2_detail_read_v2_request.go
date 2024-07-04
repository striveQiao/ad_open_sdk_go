/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaAssetV2DetailReadV2Request struct for DpaAssetV2DetailReadV2Request
type DpaAssetV2DetailReadV2Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 投放条件ID, 最多允许传入100个
	AssetIds []int64 `json:"asset_ids"`
	// 投放条件对应的线索商品ID, 最多允许传入100个
	UniqueProductIds []int64 `json:"unique_product_ids"`
}
