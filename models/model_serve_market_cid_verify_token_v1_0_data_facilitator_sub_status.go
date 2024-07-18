/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ServeMarketCidVerifyTokenV10DataFacilitatorSubStatus
type ServeMarketCidVerifyTokenV10DataFacilitatorSubStatus int64

// List of serve_market_cid_verify_token_v1.0_data_facilitator_sub_status
const (
	Enum_10_ServeMarketCidVerifyTokenV10DataFacilitatorSubStatus  ServeMarketCidVerifyTokenV10DataFacilitatorSubStatus = 10
	Enum_30_ServeMarketCidVerifyTokenV10DataFacilitatorSubStatus  ServeMarketCidVerifyTokenV10DataFacilitatorSubStatus = 30
	Enum_40_ServeMarketCidVerifyTokenV10DataFacilitatorSubStatus  ServeMarketCidVerifyTokenV10DataFacilitatorSubStatus = 40
	Enum_45_ServeMarketCidVerifyTokenV10DataFacilitatorSubStatus  ServeMarketCidVerifyTokenV10DataFacilitatorSubStatus = 45
	Enum_50_ServeMarketCidVerifyTokenV10DataFacilitatorSubStatus  ServeMarketCidVerifyTokenV10DataFacilitatorSubStatus = 50
	Enum_60_ServeMarketCidVerifyTokenV10DataFacilitatorSubStatus  ServeMarketCidVerifyTokenV10DataFacilitatorSubStatus = 60
	Enum_100_ServeMarketCidVerifyTokenV10DataFacilitatorSubStatus ServeMarketCidVerifyTokenV10DataFacilitatorSubStatus = 100
)

// All allowed values of ServeMarketCidVerifyTokenV10DataFacilitatorSubStatus enum
var AllowedServeMarketCidVerifyTokenV10DataFacilitatorSubStatusEnumValues = []ServeMarketCidVerifyTokenV10DataFacilitatorSubStatus{
	10,
	30,
	40,
	45,
	50,
	60,
	100,
}

// NewServeMarketCidVerifyTokenV10DataFacilitatorSubStatusFromValue returns a pointer to a valid ServeMarketCidVerifyTokenV10DataFacilitatorSubStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServeMarketCidVerifyTokenV10DataFacilitatorSubStatusFromValue(v int64) (*ServeMarketCidVerifyTokenV10DataFacilitatorSubStatus, error) {
	ev := ServeMarketCidVerifyTokenV10DataFacilitatorSubStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServeMarketCidVerifyTokenV10DataFacilitatorSubStatus: valid values are %v", v, AllowedServeMarketCidVerifyTokenV10DataFacilitatorSubStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServeMarketCidVerifyTokenV10DataFacilitatorSubStatus) IsValid() bool {
	for _, existing := range AllowedServeMarketCidVerifyTokenV10DataFacilitatorSubStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to serve_market_cid_verify_token_v1.0_data_facilitator_sub_status value
func (v ServeMarketCidVerifyTokenV10DataFacilitatorSubStatus) Ptr() *ServeMarketCidVerifyTokenV10DataFacilitatorSubStatus {
	return &v
}
