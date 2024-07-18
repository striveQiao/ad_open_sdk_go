/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorCreateV30RequestAnchorInfo
type NativeAnchorCreateV30RequestAnchorInfo struct {
	AnchorType         *NativeAnchorCreateV30AnchorInfoAnchorType                `json:"anchor_type,omitempty"`
	AppEcommerceAnchor *NativeAnchorCreateV30RequestAnchorInfoAppEcommerceAnchor `json:"app_ecommerce_anchor,omitempty"`
	GameAnchor         *NativeAnchorCreateV30RequestAnchorInfoGameAnchor         `json:"game_anchor,omitempty"`
	NetServiceAnchor   *NativeAnchorCreateV30RequestAnchorInfoNetServiceAnchor   `json:"net_service_anchor,omitempty"`
	PrivateChat        *NativeAnchorCreateV30RequestAnchorInfoPrivateChat        `json:"private_chat,omitempty"`
	ShoppingCartAnchor *NativeAnchorCreateV30RequestAnchorInfoShoppingCartAnchor `json:"shopping_cart_anchor,omitempty"`
	// 锚点工具名称（内部管理展示）
	ToolTitle *string `json:"tool_title,omitempty"`
}
