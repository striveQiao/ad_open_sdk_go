/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletMainWalletGetV30ResponseDataMainWalletInfoAllocatedBalanceAdOnlyAllocatedBalanceAvailableBalance 可用余额(单位元)
type SharedWalletMainWalletGetV30ResponseDataMainWalletInfoAllocatedBalanceAdOnlyAllocatedBalanceAvailableBalance struct {
	// 授信竞价可用余额(单位元)
	CreditBiddingBalance *float64 `json:"credit_bidding_balance,omitempty"`
	// 授信品牌可用余额(单位元)
	CreditBrandBalance *float64 `json:"credit_brand_balance,omitempty"`
	// 授信通用可用余额(单位元)
	CreditGeneralBalance *float64 `json:"credit_general_balance,omitempty"`
	// 预付竞价可用余额(单位元)
	PrepayBiddingBalance *float64 `json:"prepay_bidding_balance,omitempty"`
	// 预付品牌可用余额(单位元)
	PrepayBrandBalance *float64 `json:"prepay_brand_balance,omitempty"`
	// 预付通用可用余额(单位元)
	PrepayGeneralBalance *float64 `json:"prepay_general_balance,omitempty"`
}
