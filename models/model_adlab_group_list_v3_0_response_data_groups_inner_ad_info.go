/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupListV30ResponseDataGroupsInnerAdInfo 广告计划维度信息
type AdlabGroupListV30ResponseDataGroupsInnerAdInfo struct {
	AdTarget *AdlabGroupListV30DataGroupsAdInfoAdTarget `json:"ad_target,omitempty"`
	AppType  *AdlabGroupListV30DataGroupsAdInfoAppType  `json:"app_type,omitempty"`
	// 资产id
	AssetsIds []int64                                                 `json:"assets_ids,omitempty"`
	Audience  *AdlabGroupListV30ResponseDataGroupsInnerAdInfoAudience `json:"audience,omitempty"`
	// 自定义转化目标
	ConvertId *int64 `json:"convert_id,omitempty"`
	// 目标转化出价
	CpaBid      *float64                                      `json:"cpa_bid,omitempty"`
	DeepBidType *AdlabGroupListV30DataGroupsAdInfoDeepBidType `json:"deep_bid_type,omitempty"`
	// 深度优化出价
	DeepCpaBid         *float64                                                     `json:"deep_cpa_bid,omitempty"`
	DeepExternalAction *AdlabGroupListV30DataGroupsAdInfoDeepExternalAction         `json:"deep_external_action,omitempty"`
	DeliveryRange      *AdlabGroupListV30ResponseDataGroupsInnerAdInfoDeliveryRange `json:"delivery_range,omitempty"`
	DownloadMode       *AdlabGroupListV30DataGroupsAdInfoDownloadMode               `json:"download_mode,omitempty"`
	DownloadType       *AdlabGroupListV30DataGroupsAdInfoDownloadType               `json:"download_type,omitempty"`
	// 下载链接
	DownloadUrl *string `json:"download_url,omitempty"`
	// 投放结束时间
	EndTime        *string                                          `json:"end_time,omitempty"`
	ExternalAction *AdlabGroupListV30DataGroupsAdInfoExternalAction `json:"external_action,omitempty"`
	// 直达链接
	OpenUrl *string `json:"open_url,omitempty"`
	// 应用包名
	Package *string `json:"package,omitempty"`
	// 深度转化ROI系数
	RoiGoal      *float64                                       `json:"roi_goal,omitempty"`
	ScheduleType *AdlabGroupListV30DataGroupsAdInfoScheduleType `json:"schedule_type,omitempty"`
	// 投放起始时间
	StartTime       *string                                                        `json:"start_time,omitempty"`
	TrackUrlSetting *AdlabGroupListV30ResponseDataGroupsInnerAdInfoTrackUrlSetting `json:"track_url_setting,omitempty"`
	// 投放时段
	WeekSchedule []*[]int64 `json:"week_schedule,omitempty"`
}
