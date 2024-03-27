/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementAndroidBasicPackagePublishV2Request struct for ToolsAppManagementAndroidBasicPackagePublishV2Request
type ToolsAppManagementAndroidBasicPackagePublishV2Request struct {
	// 账户id指可以接的账号体系如广告主id、巨量纵横组织id等
	AccountId   int64                                                     `json:"account_id"`
	AccountType ToolsAppManagementAndroidBasicPackagePublishV2AccountType `json:"account_type"`
	// 应用包id
	PackageId string `json:"package_id"`
}
