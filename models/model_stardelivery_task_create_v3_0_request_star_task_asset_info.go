/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StardeliveryTaskCreateV30RequestStarTaskAssetInfo 事件管理资产&监测链接组信息
type StardeliveryTaskCreateV30RequestStarTaskAssetInfo struct {
	// 安卓应用资产id，可通过「获取已创建资产列表」接口查询应用资产id，必须与star_task_android_track_url_group_id 同时传入，否则报错
	AndroidAssetId *int64 `json:"android_asset_id,omitempty"`
	// 安卓应用监测链接组ID，您可调用「获取事件资产下的监测链接组」接口查询资产下已创建的监测链接组ID(track_url_group_id)
	AndroidTrackUrlGroupId *int64 `json:"android_track_url_group_id,omitempty"`
	// iOS应用资产id，可通过「获取已创建资产列表」接口查询应用资产id，必须与star_task_ios_track_url_group_id 同时传入，否则报错
	IosAssetId *int64 `json:"ios_asset_id,omitempty"`
	// iOS应用监测链接组ID，您可调用「获取事件资产下的监测链接组」接口查询资产下已创建的监测链接组ID(track_url_group_id)
	IosTrackUrlGroupId *int64 `json:"ios_track_url_group_id,omitempty"`
}
