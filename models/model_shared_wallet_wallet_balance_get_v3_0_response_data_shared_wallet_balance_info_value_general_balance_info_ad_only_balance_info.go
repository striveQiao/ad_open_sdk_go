/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletWalletBalanceGetV30ResponseDataSharedWalletBalanceInfoValueGeneralBalanceInfoAdOnlyBalanceInfo 巨量广告业务线余额信息
type SharedWalletWalletBalanceGetV30ResponseDataSharedWalletBalanceInfoValueGeneralBalanceInfoAdOnlyBalanceInfo struct {
	// 授信竞价可用余额(单位元)
	CreditBiddingValidBalance *float64 `json:"credit_bidding_valid_balance,omitempty"`
	// 授信品牌可用余额(单位元)
	CreditBrandValidBalance *float64 `json:"credit_brand_valid_balance,omitempty"`
	// 授信通用可用余额(单位元)
	CreditGeneralValidBalance *float64 `json:"credit_general_valid_balance,omitempty"`
	// 预付竞价可用余额(单位元)
	PrepayBiddingValidBalance *float64 `json:"prepay_bidding_valid_balance,omitempty"`
	// 预付品牌可用余额(单位元)
	PrepayBrandValidBalance *float64 `json:"prepay_brand_valid_balance,omitempty"`
	// 预付通用可用余额(单位元)
	PrepayGeneralValidBalance *float64 `json:"prepay_general_valid_balance,omitempty"`
	// 总余额(单位元)
	TotalBalance *float64 `json:"total_balance,omitempty"`
	// 总可用余额(单位元)
	TotalValidBalance *float64 `json:"total_valid_balance,omitempty"`
}
