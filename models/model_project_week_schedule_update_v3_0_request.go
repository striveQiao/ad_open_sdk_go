/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectWeekScheduleUpdateV30Request struct for ProjectWeekScheduleUpdateV30Request
type ProjectWeekScheduleUpdateV30Request struct {
	// 广告主账户id
	AdvertiserId int64 `json:"advertiser_id"`
	// 批量修改投放时段，限制最多批量修改10个项目
	Data []*ProjectWeekScheduleUpdateV30RequestDataInner `json:"data"`
}
