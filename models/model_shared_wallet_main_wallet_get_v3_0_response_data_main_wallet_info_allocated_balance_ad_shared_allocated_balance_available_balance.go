/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletMainWalletGetV30ResponseDataMainWalletInfoAllocatedBalanceAdSharedAllocatedBalanceAvailableBalance 可用余额(单位元)
type SharedWalletMainWalletGetV30ResponseDataMainWalletInfoAllocatedBalanceAdSharedAllocatedBalanceAvailableBalance struct {
	// 预付竞价可用余额(单位元)
	PrepayBiddingBalance *float64 `json:"prepay_bidding_balance,omitempty"`
	// 预付品牌可用余额(单位元)
	PrepayBrandBalance *float64 `json:"prepay_brand_balance,omitempty"`
	// 预付通用可用余额(单位元)
	PrepayGeneralBalance *float64 `json:"prepay_general_balance,omitempty"`
}
