/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdScheduleTimeUpdateV10Request struct for QianchuanAdScheduleTimeUpdateV10Request
type QianchuanAdScheduleTimeUpdateV10Request struct {
	//
	AdIds []int64 `json:"ad_ids"`
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	ScheduleTime *string `json:"schedule_time,omitempty"`
}
