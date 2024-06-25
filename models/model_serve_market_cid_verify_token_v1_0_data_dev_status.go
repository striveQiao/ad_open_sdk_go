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

// ServeMarketCidVerifyTokenV10DataDevStatus
type ServeMarketCidVerifyTokenV10DataDevStatus int64

// List of serve_market_cid_verify_token_v1.0_data_dev_status
const (
	Enum_1_ServeMarketCidVerifyTokenV10DataDevStatus ServeMarketCidVerifyTokenV10DataDevStatus = 1
	Enum_2_ServeMarketCidVerifyTokenV10DataDevStatus ServeMarketCidVerifyTokenV10DataDevStatus = 2
	Enum_3_ServeMarketCidVerifyTokenV10DataDevStatus ServeMarketCidVerifyTokenV10DataDevStatus = 3
	Enum_4_ServeMarketCidVerifyTokenV10DataDevStatus ServeMarketCidVerifyTokenV10DataDevStatus = 4
	Enum_5_ServeMarketCidVerifyTokenV10DataDevStatus ServeMarketCidVerifyTokenV10DataDevStatus = 5
	Enum_6_ServeMarketCidVerifyTokenV10DataDevStatus ServeMarketCidVerifyTokenV10DataDevStatus = 6
	Enum_7_ServeMarketCidVerifyTokenV10DataDevStatus ServeMarketCidVerifyTokenV10DataDevStatus = 7
	Enum_8_ServeMarketCidVerifyTokenV10DataDevStatus ServeMarketCidVerifyTokenV10DataDevStatus = 8
)

// All allowed values of ServeMarketCidVerifyTokenV10DataDevStatus enum
var AllowedServeMarketCidVerifyTokenV10DataDevStatusEnumValues = []ServeMarketCidVerifyTokenV10DataDevStatus{
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
}

// NewServeMarketCidVerifyTokenV10DataDevStatusFromValue returns a pointer to a valid ServeMarketCidVerifyTokenV10DataDevStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServeMarketCidVerifyTokenV10DataDevStatusFromValue(v int64) (*ServeMarketCidVerifyTokenV10DataDevStatus, error) {
	ev := ServeMarketCidVerifyTokenV10DataDevStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServeMarketCidVerifyTokenV10DataDevStatus: valid values are %v", v, AllowedServeMarketCidVerifyTokenV10DataDevStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServeMarketCidVerifyTokenV10DataDevStatus) IsValid() bool {
	for _, existing := range AllowedServeMarketCidVerifyTokenV10DataDevStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to serve_market_cid_verify_token_v1.0_data_dev_status value
func (v ServeMarketCidVerifyTokenV10DataDevStatus) Ptr() *ServeMarketCidVerifyTokenV10DataDevStatus {
	return &v
}
