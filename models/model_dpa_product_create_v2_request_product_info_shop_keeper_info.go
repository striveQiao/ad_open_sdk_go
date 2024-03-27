/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaProductCreateV2RequestProductInfoShopKeeperInfo
type DpaProductCreateV2RequestProductInfoShopKeeperInfo struct {
	//
	Address *string `json:"address,omitempty"`
	//
	ShopKeeperId *string `json:"shop_keeper_id,omitempty"`
	//
	ShopKeeperName *string `json:"shop_keeper_name,omitempty"`
	//
	ShopKeeperUrl *string `json:"shop_keeper_url,omitempty"`
	//
	ShopKeeperUrlAndroidApp *string `json:"shop_keeper_url_android_app,omitempty"`
	//
	ShopKeeperUrlIosApp *string `json:"shop_keeper_url_ios_app,omitempty"`
	//
	ShopKeeperUrlMobile *string `json:"shop_keeper_url_mobile,omitempty"`
	//
	ShopKeeperUrlUniversalLink *string `json:"shop_keeper_url_universal_link,omitempty"`
}
