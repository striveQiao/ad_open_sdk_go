/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorCreateV30RequestAnchorInfoGameAnchorGamePackageListInner struct for NativeAnchorCreateV30RequestAnchorInfoGameAnchorGamePackageListInner
type NativeAnchorCreateV30RequestAnchorInfoGameAnchorGamePackageListInner struct {
	// 安卓礼包码，字符限制0～20
	AndroidPackageCode *string `json:"android_package_code,omitempty"`
	// 游戏礼包名称，字符限制0～14
	GamePackageName *string `json:"game_package_name,omitempty"`
	// 礼包内的礼品配置，数量限制0～8
	Gift []*NativeAnchorCreateV30RequestAnchorInfoGameAnchorGamePackageListInnerGiftInner `json:"gift,omitempty"`
	// 礼包使用结束期限，格式 yyyy.MM.dd
	GiftEndDate *string `json:"gift_end_date,omitempty"`
	// 礼包使用规则，字符限制0～30
	GiftRegulation *string `json:"gift_regulation,omitempty"`
	// 礼包使用开始期限，格式 yyyy.MM.dd
	GiftStartDate *string `json:"gift_start_date,omitempty"`
	// ios礼包码，字符限制0～20
	IosPackageCode *string `json:"ios_package_code,omitempty"`
}
