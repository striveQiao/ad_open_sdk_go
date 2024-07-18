/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorUpdateV30RequestAnchorInfo
type NativeAnchorUpdateV30RequestAnchorInfo struct {
	// 待更新的锚点id
	AnchorId           string                                                    `json:"anchor_id"`
	AnchorType         NativeAnchorUpdateV30AnchorInfoAnchorType                 `json:"anchor_type"`
	AppEcommerceAnchor *NativeAnchorUpdateV30RequestAnchorInfoAppEcommerceAnchor `json:"app_ecommerce_anchor,omitempty"`
	GameAnchor         *NativeAnchorUpdateV30RequestAnchorInfoGameAnchor         `json:"game_anchor,omitempty"`
	NetServiceAnchor   *NativeAnchorUpdateV30RequestAnchorInfoNetServiceAnchor   `json:"net_service_anchor,omitempty"`
	PrivateChat        *NativeAnchorUpdateV30RequestAnchorInfoPrivateChat        `json:"private_chat,omitempty"`
	ShoppingCartAnchor *NativeAnchorUpdateV30RequestAnchorInfoShoppingCartAnchor `json:"shopping_cart_anchor,omitempty"`
	// 锚点工具名称（内部管理展示）
	ToolTitle string `json:"tool_title"`
}
