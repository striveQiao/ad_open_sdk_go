/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletMainWalletGetV30ResponseDataMainWalletInfoUnallocatedBalanceEcpOnlyUnallocatedBalance 巨量千川业务线待分配余额(单位元)
type SharedWalletMainWalletGetV30ResponseDataMainWalletInfoUnallocatedBalanceEcpOnlyUnallocatedBalance struct {
	AvailableBalance   *SharedWalletMainWalletGetV30ResponseDataMainWalletInfoUnallocatedBalanceEcpOnlyUnallocatedBalanceAvailableBalance   `json:"available_balance,omitempty"`
	UnavailableBalance *SharedWalletMainWalletGetV30ResponseDataMainWalletInfoUnallocatedBalanceEcpOnlyUnallocatedBalanceUnavailableBalance `json:"unavailable_balance,omitempty"`
}
