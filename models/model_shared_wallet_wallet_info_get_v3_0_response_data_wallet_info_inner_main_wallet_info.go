/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletWalletInfoGetV30ResponseDataWalletInfoInnerMainWalletInfo 共享钱包(大钱包)信息
type SharedWalletWalletInfoGetV30ResponseDataWalletInfoInnerMainWalletInfo struct {
	// 公司ID
	CompanyId       *int64                                                                                `json:"company_id,omitempty"`
	MainSharedRange *SharedWalletWalletInfoGetV30ResponseDataWalletInfoInnerMainWalletInfoMainSharedRange `json:"main_shared_range,omitempty"`
}
