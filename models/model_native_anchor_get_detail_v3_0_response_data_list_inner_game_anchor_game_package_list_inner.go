/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorGetDetailV30ResponseDataListInnerGameAnchorGamePackageListInner struct for NativeAnchorGetDetailV30ResponseDataListInnerGameAnchorGamePackageListInner
type NativeAnchorGetDetailV30ResponseDataListInnerGameAnchorGamePackageListInner struct {
	//
	AndroidPackageCode *string `json:"android_package_code,omitempty"`
	//
	GameGiftRegulation *string `json:"game_gift_regulation,omitempty"`
	//
	GamePackageName string `json:"game_package_name"`
	//
	Gift []*NativeAnchorGetDetailV30ResponseDataListInnerGameAnchorGamePackageListInnerGiftInner `json:"gift,omitempty"`
	//
	GiftEndDate *string `json:"gift_end_date,omitempty"`
	//
	GiftStartDate *string `json:"gift_start_date,omitempty"`
	//
	IosPackageCode *string `json:"ios_package_code,omitempty"`
}
