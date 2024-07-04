/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectUpdateV30RequestDeliverySetting
type ProjectUpdateV30RequestDeliverySetting struct {
	//
	Bid *float64 `json:"bid,omitempty"`
	//
	Budget     *float64                                   `json:"budget,omitempty"`
	BudgetMode *ProjectUpdateV30DeliverySettingBudgetMode `json:"budget_mode,omitempty"`
	//
	CpaBid *float64 `json:"cpa_bid,omitempty"`
	//
	DeepCpabid *float64 `json:"deep_cpabid,omitempty"`
	//
	EndTime           *string                                           `json:"end_time,omitempty"`
	FilterNightSwitch *ProjectUpdateV30DeliverySettingFilterNightSwitch `json:"filter_night_switch,omitempty"`
	//
	FirstRoiGoal *float64 `json:"first_roi_goal,omitempty"`
	//
	RoiGoal *float64 `json:"roi_goal,omitempty"`
	//
	ScheduleTime           *string                                                `json:"schedule_time,omitempty"`
	ScheduleType           *ProjectUpdateV30DeliverySettingScheduleType           `json:"schedule_type,omitempty"`
	SearchContinueDelivery *ProjectUpdateV30DeliverySettingSearchContinueDelivery `json:"search_continue_delivery,omitempty"`
	//
	ShopMultiRoiGoals []*ProjectUpdateV30RequestDeliverySettingShopMultiRoiGoalsInner `json:"shop_multi_roi_goals,omitempty"`
}
