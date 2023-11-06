/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SubscribeAccountsListV30ResponseDataAdvertisersInner struct for SubscribeAccountsListV30ResponseDataAdvertisersInner
type SubscribeAccountsListV30ResponseDataAdvertisersInner struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	CoreUserId int64 `json:"core_user_id"`
	//
	Event *string `json:"event,omitempty"`
	//
	Reason *string                                        `json:"reason,omitempty"`
	Status *SubscribeAccountsListV30DataAdvertisersStatus `json:"status,omitempty"`
}
