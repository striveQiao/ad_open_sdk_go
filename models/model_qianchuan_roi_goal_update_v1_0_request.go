/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanRoiGoalUpdateV10Request struct for QianchuanRoiGoalUpdateV10Request
type QianchuanRoiGoalUpdateV10Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	RoiGoalUpdates []*QianchuanRoiGoalUpdateV10RequestRoiGoalUpdatesInner `json:"roi_goal_updates"`
}
