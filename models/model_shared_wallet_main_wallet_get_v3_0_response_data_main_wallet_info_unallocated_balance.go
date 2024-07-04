/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletMainWalletGetV30ResponseDataMainWalletInfoUnallocatedBalance 钱包待分配余额(单位元)
type SharedWalletMainWalletGetV30ResponseDataMainWalletInfoUnallocatedBalance struct {
	AdOnlyUnallocatedBalance    *SharedWalletMainWalletGetV30ResponseDataMainWalletInfoUnallocatedBalanceAdOnlyUnallocatedBalance    `json:"ad_only_unallocated_balance,omitempty"`
	AdSharedUnallocatedBalance  *SharedWalletMainWalletGetV30ResponseDataMainWalletInfoUnallocatedBalanceAdSharedUnallocatedBalance  `json:"ad_shared_unallocated_balance,omitempty"`
	EcpOnlyUnallocatedBalance   *SharedWalletMainWalletGetV30ResponseDataMainWalletInfoUnallocatedBalanceEcpOnlyUnallocatedBalance   `json:"ecp_only_unallocated_balance,omitempty"`
	LocalOnlyUnallocatedBalance *SharedWalletMainWalletGetV30ResponseDataMainWalletInfoUnallocatedBalanceLocalOnlyUnallocatedBalance `json:"local_only_unallocated_balance,omitempty"`
}
