/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdStatusUpdateV10Request struct for QianchuanAdStatusUpdateV10Request
type QianchuanAdStatusUpdateV10Request struct {
	//
	AdIds []int64 `json:"ad_ids"`
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Budget *float64 `json:"budget,omitempty"`
	//
	OptStatus interface{} `json:"opt_status"`
	//
	ReviveBudget *float64 `json:"revive_budget,omitempty"`
	//
	ScheduleFixedRange *int64 `json:"schedule_fixed_range,omitempty"`
}
