/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandListV2ResponseDataListInner struct for StarDemandListV2ResponseDataListInner
type StarDemandListV2ResponseDataListInner struct {
	ComponentType *StarDemandListV2DataListComponentType `json:"component_type,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	DemandId *int64 `json:"demand_id,omitempty"`
	//
	DemandName              *string                                          `json:"demand_name,omitempty"`
	UniversalSettlementType *StarDemandListV2DataListUniversalSettlementType `json:"universal_settlement_type,omitempty"`
}
