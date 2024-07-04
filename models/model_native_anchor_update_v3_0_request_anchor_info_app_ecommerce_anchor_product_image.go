/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorUpdateV30RequestAnchorInfoAppEcommerceAnchorProductImage 商品大图，推荐比例 1：1
type NativeAnchorUpdateV30RequestAnchorInfoAppEcommerceAnchorProductImage struct {
	// 图片高度
	Height *int64 `json:"height,omitempty"`
	// 图片image_id
	Uri *string `json:"uri,omitempty"`
	// 图片宽度
	Width *int64 `json:"width,omitempty"`
}
