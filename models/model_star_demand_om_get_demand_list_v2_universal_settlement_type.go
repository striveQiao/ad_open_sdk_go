/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// StarDemandOmGetDemandListV2UniversalSettlementType
type StarDemandOmGetDemandListV2UniversalSettlementType int64

// List of star_demand_om_get_demand_list_v2_universal_settlement_type
const (
	Enum_16_StarDemandOmGetDemandListV2UniversalSettlementType StarDemandOmGetDemandListV2UniversalSettlementType = 16
	Enum_19_StarDemandOmGetDemandListV2UniversalSettlementType StarDemandOmGetDemandListV2UniversalSettlementType = 19
)

// All allowed values of StarDemandOmGetDemandListV2UniversalSettlementType enum
var AllowedStarDemandOmGetDemandListV2UniversalSettlementTypeEnumValues = []StarDemandOmGetDemandListV2UniversalSettlementType{
	16,
	19,
}

// NewStarDemandOmGetDemandListV2UniversalSettlementTypeFromValue returns a pointer to a valid StarDemandOmGetDemandListV2UniversalSettlementType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarDemandOmGetDemandListV2UniversalSettlementTypeFromValue(v int64) (*StarDemandOmGetDemandListV2UniversalSettlementType, error) {
	ev := StarDemandOmGetDemandListV2UniversalSettlementType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarDemandOmGetDemandListV2UniversalSettlementType: valid values are %v", v, AllowedStarDemandOmGetDemandListV2UniversalSettlementTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarDemandOmGetDemandListV2UniversalSettlementType) IsValid() bool {
	for _, existing := range AllowedStarDemandOmGetDemandListV2UniversalSettlementTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_demand_om_get_demand_list_v2_universal_settlement_type value
func (v StarDemandOmGetDemandListV2UniversalSettlementType) Ptr() *StarDemandOmGetDemandListV2UniversalSettlementType {
	return &v
}
