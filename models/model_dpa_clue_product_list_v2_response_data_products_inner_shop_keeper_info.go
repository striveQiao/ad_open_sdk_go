/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaClueProductListV2ResponseDataProductsInnerShopKeeperInfo 商家信息
type DpaClueProductListV2ResponseDataProductsInnerShopKeeperInfo struct {
	// 商家地址
	Address *string `json:"address,omitempty"`
	// 商家ID
	ShopKeeperId *string `json:"shop_keeper_id,omitempty"`
	// 商家logo
	ShopKeeperLogo *string `json:"shop_keeper_logo,omitempty"`
	// 商家名称
	ShopKeeperName *string `json:"shop_keeper_name,omitempty"`
	// 商家落地页
	ShopKeeperUrl *string `json:"shop_keeper_url,omitempty"`
	// 商家安卓落地页
	ShopKeeperUrlAndroidApp *string `json:"shop_keeper_url_android_app,omitempty"`
	// 商家iOS落地页
	ShopKeeperUrlIosApp *string `json:"shop_keeper_url_ios_app,omitempty"`
	// 商家H5落地页
	ShopKeeperUrlMobile *string `json:"shop_keeper_url_mobile,omitempty"`
	// 商家落地页universal link
	ShopKeeperUrlUniversalLink *string `json:"shop_keeper_url_universal_link,omitempty"`
}
