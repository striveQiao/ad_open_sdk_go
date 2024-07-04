/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdCreateV10RequestDeliverySetting
type QianchuanAdCreateV10RequestDeliverySetting struct {
	//
	Budget     float64                                       `json:"budget"`
	BudgetMode QianchuanAdCreateV10DeliverySettingBudgetMode `json:"budget_mode"`
	//
	CpaBid             *float64                                               `json:"cpa_bid,omitempty"`
	DeepBidType        *QianchuanAdCreateV10DeliverySettingDeepBidType        `json:"deep_bid_type,omitempty"`
	DeepExternalAction *QianchuanAdCreateV10DeliverySettingDeepExternalAction `json:"deep_external_action,omitempty"`
	//
	EndTime        *string                                           `json:"end_time,omitempty"`
	ExternalAction QianchuanAdCreateV10DeliverySettingExternalAction `json:"external_action"`
	// 抢首屏ROI目标降低幅度
	GrabFirstScreenBidAdjustRate *int64 `json:"grab_first_screen_bid_adjust_rate,omitempty"`
	// 抢首屏开关
	GrabFirstScreenSwitch *bool                                                `json:"grab_first_screen_switch,omitempty"`
	LiveScheduleType      *QianchuanAdCreateV10DeliverySettingLiveScheduleType `json:"live_schedule_type,omitempty"`
	//
	ProductNewOpen *bool                                        `json:"product_new_open,omitempty"`
	QcpxMode       *QianchuanAdCreateV10DeliverySettingQcpxMode `json:"qcpx_mode,omitempty"`
	//
	RoiGoal *float64 `json:"roi_goal,omitempty"`
	//
	ScheduleFixedRange *int64 `json:"schedule_fixed_range,omitempty"`
	//
	ScheduleTime *string                                         `json:"schedule_time,omitempty"`
	SmartBidType QianchuanAdCreateV10DeliverySettingSmartBidType `json:"smart_bid_type"`
	//
	StartTime         *string                                               `json:"start_time,omitempty"`
	VideoScheduleType *QianchuanAdCreateV10DeliverySettingVideoScheduleType `json:"video_schedule_type,omitempty"`
}
