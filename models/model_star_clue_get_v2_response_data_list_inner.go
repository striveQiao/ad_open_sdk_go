/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarClueGetV2ResponseDataListInner struct for StarClueGetV2ResponseDataListInner
type StarClueGetV2ResponseDataListInner struct {
	//
	Address *string `json:"address,omitempty"`
	//
	AnchorName *string `json:"anchor_name,omitempty"`
	//
	AuthorName *string `json:"author_name,omitempty"`
	//
	AuthorUid *int64 `json:"author_uid,omitempty"`
	//
	CarSeries *string `json:"car_series,omitempty"`
	//
	CarType *string `json:"car_type,omitempty"`
	//
	City          *string                             `json:"city,omitempty"`
	ComponentType *StarClueGetV2DataListComponentType `json:"component_type,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	DemandId *int64 `json:"demand_id,omitempty"`
	//
	ItemId *int64 `json:"item_id,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	OrderId *int64 `json:"order_id,omitempty"`
	//
	Phone *string `json:"phone,omitempty"`
	//
	Province *string `json:"province,omitempty"`
	//
	StarId *int64 `json:"star_id,omitempty"`
}
