/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPlayableUploadV2Request struct for ToolsPlayableUploadV2Request
type ToolsPlayableUploadV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	// 试玩素材文件 【格式说明】： 1.格式为zip，大小不超过3M 2.已接入穿山甲JS-SDK，并且调用方式为window.openAppStore(); 3.主html文件需命名为index，且位于一级目录中 4.试玩广告播放方向字段应存储于config.json文件中，位于一级目录中，取值为0,1,2 5.文件名称仅支持大小写字母、数字、减号和下划线 6.素材中不允许使用 mraid.js 格式 7.素材不允许通过外部网络加载动态素材 8.素材中不允许包含JS 重定向 9.素材不允许发出 HTTP 请求 10.素材中不允许存在内容为空的文件
	PlayablePackage *FormFileInfo `json:"playable_package"`
}
