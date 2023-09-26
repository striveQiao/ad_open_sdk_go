/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsEventAssetsGetV2ResponseDataQuickAppInner struct for ToolsEventAssetsGetV2ResponseDataQuickAppInner
type ToolsEventAssetsGetV2ResponseDataQuickAppInner struct {
	// 快应用资产ID
	AssetId *int64 `json:"asset_id,omitempty"`
	// 快应用名称
	AssetName *string `json:"asset_name,omitempty"`
	// 快应用包名
	PackageName *string `json:"package_name,omitempty"`
	// 快应用ID
	QuickAppId *int64 `json:"quick_app_id,omitempty"`
}
