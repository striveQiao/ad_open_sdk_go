/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletMainWalletGetV30ResponseDataMainWalletInfoUnallocatedBalanceAdOnlyUnallocatedBalance 巨量广告业务线待分配余额(单位元)
type SharedWalletMainWalletGetV30ResponseDataMainWalletInfoUnallocatedBalanceAdOnlyUnallocatedBalance struct {
	AvailableBalance   *SharedWalletMainWalletGetV30ResponseDataMainWalletInfoUnallocatedBalanceAdOnlyUnallocatedBalanceAvailableBalance   `json:"available_balance,omitempty"`
	UnavailableBalance *SharedWalletMainWalletGetV30ResponseDataMainWalletInfoUnallocatedBalanceAdOnlyUnallocatedBalanceUnavailableBalance `json:"unavailable_balance,omitempty"`
}
