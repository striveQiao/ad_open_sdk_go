/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsEventAllAssetsDetailV2ResponseDataAssetListInner struct for ToolsEventAllAssetsDetailV2ResponseDataAssetListInner
type ToolsEventAllAssetsDetailV2ResponseDataAssetListInner struct {
	// 应用ID，当asset_type = APP 时返回
	AppId *int64 `json:"app_id,omitempty"`
	// 应用名称，当asset_type = APP 时返回
	AppName *string                                          `json:"app_name,omitempty"`
	AppType *ToolsEventAllAssetsDetailV2DataAssetListAppType `json:"app_type,omitempty"`
	// 资产id，所有资产均会返回
	AssetId *int64 `json:"asset_id,omitempty"`
	// 资产名称，所有资产均会返回
	AssetName *string                                            `json:"asset_name,omitempty"`
	AssetType *ToolsEventAllAssetsDetailV2DataAssetListAssetType `json:"asset_type,omitempty"`
	// 应用下载链接，当asset_type = APP 时返回
	DownloadUrl *string `json:"download_url,omitempty"`
	// 小程序appid，当asset_type = MINI_PROGRAME 时返回
	MicroAppId *string `json:"micro_app_id,omitempty"`
	// 小程序资产id，当asset_type = MINI_PROGRAME 时返回
	MicroAppInstanceId *int64 `json:"micro_app_instance_id,omitempty"`
	// 应用包id，当asset_type = APP 时返回，当调用API创建资产时，Android应用必填，建议获取保存
	PackageId *string `json:"package_id,omitempty"`
	// 应用包名，当asset_type = APP 时返回
	PackageName *string `json:"package_name,omitempty"`
	// 快应用id，当asset_type = QUICK_APP 时返回
	QuickAppId *int64 `json:"quick_app_id,omitempty"`
	// 快应用包名，当asset_type = QUICK_APP 时返回
	QuickAppPackageName *string `json:"quick_app_package_name,omitempty"`
	// 橙子落地页站点id，当asset_type = TETRIS_EXTERNAL 时返回，表示该资产是基于哪一个橙子落地页站点创建的，在巨量广告-事件资产页面，橙子落地页显示为该站点id
	SiteId *int64 `json:"site_id,omitempty"`
}
