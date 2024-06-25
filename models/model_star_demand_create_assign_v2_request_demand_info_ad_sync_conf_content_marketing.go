/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandCreateAssignV2RequestDemandInfoAdSyncConfContentMarketing 产出物推送品牌广告（内容热推/内容服务）配置
type StarDemandCreateAssignV2RequestDemandInfoAdSyncConfContentMarketing struct {
	// 是否用于视频素材投放（头像跳落地页）：2=是；1=否（用于原视频投放） 默认推送素材（为2，头像跳落地页）。内容服务为2不可改
	AdSync *int32 `json:"ad_sync,omitempty"`
	// 是否投放原生视频 0或不填 = 否；1 = 是
	AdSyncOrigin *int64 `json:"ad_sync_origin,omitempty"`
	// 广告主ID
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	// 是否自动投放 0或不填 = 否；1 = 是
	AutoSync *int64 `json:"auto_sync,omitempty"`
	// 自动推送素材包的平台：2-内容热推；3-内容服务 需要auto_sync=1，且不允许为空数组
	AutoSyncSpecifyMaterialPlatforms []int32 `json:"auto_sync_specify_material_platforms,omitempty"`
	// 推广产品链接
	ProductLink *string `json:"product_link,omitempty"`
	// 推广产品图片文件名（材料上传接口中使用到的文件名）
	ProductPics []string `json:"product_pics,omitempty"`
	// 投放原生视频的投放时间（单位：天） 大于0的整数
	SyncDuration *int64 `json:"sync_duration,omitempty"`
}
