/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StardeliveryTaskDetailV30ResponseDataStarTaskAssetInfo 事件资产&监测链接信息
type StardeliveryTaskDetailV30ResponseDataStarTaskAssetInfo struct {
	// 任务使用的安卓应用点击监测链接，返回的是新建任务时选择的监测链接组ID下的点击监测链接
	AndroidActionTrackUrl *string `json:"android_action_track_url,omitempty"`
	// 安卓应用资产id
	AndroidAssetId *int64 `json:"android_asset_id,omitempty"`
	// 任务使用的iOS应用点击监测链接，返回的是新建任务时选择的监测链接组ID下的点击监测链接
	IosActionTrackUrl *string `json:"ios_action_track_url,omitempty"`
	// iOS应用资产id
	IosAssetId *int64 `json:"ios_asset_id,omitempty"`
}
