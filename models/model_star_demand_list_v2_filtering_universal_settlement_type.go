/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// StarDemandListV2FilteringUniversalSettlementType
type StarDemandListV2FilteringUniversalSettlementType string

// List of star_demand_list_v2_filtering_universal_settlement_type
const (
	STAR_SUPPORT_StarDemandListV2FilteringUniversalSettlementType StarDemandListV2FilteringUniversalSettlementType = "STAR_SUPPORT"
	EXCHANGE_StarDemandListV2FilteringUniversalSettlementType     StarDemandListV2FilteringUniversalSettlementType = "EXCHANGE"
	FLOW_SHARE_StarDemandListV2FilteringUniversalSettlementType   StarDemandListV2FilteringUniversalSettlementType = "FLOW_SHARE"
	CPA_StarDemandListV2FilteringUniversalSettlementType          StarDemandListV2FilteringUniversalSettlementType = "CPA"
	DOU_PLUS_StarDemandListV2FilteringUniversalSettlementType     StarDemandListV2FilteringUniversalSettlementType = "DOU_PLUS"
	MONEY_SHARE_StarDemandListV2FilteringUniversalSettlementType  StarDemandListV2FilteringUniversalSettlementType = "MONEY_SHARE"
	EFFECT_StarDemandListV2FilteringUniversalSettlementType       StarDemandListV2FilteringUniversalSettlementType = "EFFECT"
	CUSTOMIZE_StarDemandListV2FilteringUniversalSettlementType    StarDemandListV2FilteringUniversalSettlementType = "CUSTOMIZE"
	RANK_StarDemandListV2FilteringUniversalSettlementType         StarDemandListV2FilteringUniversalSettlementType = "RANK"
	FIXED_PRICE_StarDemandListV2FilteringUniversalSettlementType  StarDemandListV2FilteringUniversalSettlementType = "FIXED_PRICE"
	GIFT_StarDemandListV2FilteringUniversalSettlementType         StarDemandListV2FilteringUniversalSettlementType = "GIFT"
	CPM_StarDemandListV2FilteringUniversalSettlementType          StarDemandListV2FilteringUniversalSettlementType = "CPM"
	ALL_StarDemandListV2FilteringUniversalSettlementType          StarDemandListV2FilteringUniversalSettlementType = "ALL"
)

// All allowed values of StarDemandListV2FilteringUniversalSettlementType enum
var AllowedStarDemandListV2FilteringUniversalSettlementTypeEnumValues = []StarDemandListV2FilteringUniversalSettlementType{
	"STAR_SUPPORT",
	"EXCHANGE",
	"FLOW_SHARE",
	"CPA",
	"DOU_PLUS",
	"MONEY_SHARE",
	"EFFECT",
	"CUSTOMIZE",
	"RANK",
	"FIXED_PRICE",
	"GIFT",
	"CPM",
	"ALL",
}

// NewStarDemandListV2FilteringUniversalSettlementTypeFromValue returns a pointer to a valid StarDemandListV2FilteringUniversalSettlementType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarDemandListV2FilteringUniversalSettlementTypeFromValue(v string) (*StarDemandListV2FilteringUniversalSettlementType, error) {
	ev := StarDemandListV2FilteringUniversalSettlementType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarDemandListV2FilteringUniversalSettlementType: valid values are %v", v, AllowedStarDemandListV2FilteringUniversalSettlementTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarDemandListV2FilteringUniversalSettlementType) IsValid() bool {
	for _, existing := range AllowedStarDemandListV2FilteringUniversalSettlementTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_demand_list_v2_filtering_universal_settlement_type value
func (v StarDemandListV2FilteringUniversalSettlementType) Ptr() *StarDemandListV2FilteringUniversalSettlementType {
	return &v
}
