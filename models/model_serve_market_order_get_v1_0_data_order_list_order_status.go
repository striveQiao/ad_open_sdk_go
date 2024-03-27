/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ServeMarketOrderGetV10DataOrderListOrderStatus
type ServeMarketOrderGetV10DataOrderListOrderStatus string

// List of serve_market_order_get_v1.0_data_order_list_order_status
const (
	UNPAID_ServeMarketOrderGetV10DataOrderListOrderStatus            ServeMarketOrderGetV10DataOrderListOrderStatus = "UNPAID"
	CANCELED_ServeMarketOrderGetV10DataOrderListOrderStatus          ServeMarketOrderGetV10DataOrderListOrderStatus = "CANCELED"
	OVERTIME_CANCELED_ServeMarketOrderGetV10DataOrderListOrderStatus ServeMarketOrderGetV10DataOrderListOrderStatus = "OVERTIME_CANCELED"
	SERVING_ServeMarketOrderGetV10DataOrderListOrderStatus           ServeMarketOrderGetV10DataOrderListOrderStatus = "SERVING"
	REFUND_APPLY_ServeMarketOrderGetV10DataOrderListOrderStatus      ServeMarketOrderGetV10DataOrderListOrderStatus = "REFUND_APPLY"
	REFUND_REJECT_ServeMarketOrderGetV10DataOrderListOrderStatus     ServeMarketOrderGetV10DataOrderListOrderStatus = "REFUND_REJECT"
	REFUND_APPROVED_ServeMarketOrderGetV10DataOrderListOrderStatus   ServeMarketOrderGetV10DataOrderListOrderStatus = "REFUND_APPROVED"
	REFUND_SUCCESS_ServeMarketOrderGetV10DataOrderListOrderStatus    ServeMarketOrderGetV10DataOrderListOrderStatus = "REFUND_SUCCESS"
	DELIVERED_ServeMarketOrderGetV10DataOrderListOrderStatus         ServeMarketOrderGetV10DataOrderListOrderStatus = "DELIVERED"
	FINISHED_ServeMarketOrderGetV10DataOrderListOrderStatus          ServeMarketOrderGetV10DataOrderListOrderStatus = "FINISHED"
	OVERTIME_FINISHED_ServeMarketOrderGetV10DataOrderListOrderStatus ServeMarketOrderGetV10DataOrderListOrderStatus = "OVERTIME_FINISHED"
	DELIVERY_REJECT_ServeMarketOrderGetV10DataOrderListOrderStatus   ServeMarketOrderGetV10DataOrderListOrderStatus = "DELIVERY_REJECT"
)

// All allowed values of ServeMarketOrderGetV10DataOrderListOrderStatus enum
var AllowedServeMarketOrderGetV10DataOrderListOrderStatusEnumValues = []ServeMarketOrderGetV10DataOrderListOrderStatus{
	"UNPAID",
	"CANCELED",
	"OVERTIME_CANCELED",
	"SERVING",
	"REFUND_APPLY",
	"REFUND_REJECT",
	"REFUND_APPROVED",
	"REFUND_SUCCESS",
	"DELIVERED",
	"FINISHED",
	"OVERTIME_FINISHED",
	"DELIVERY_REJECT",
}

// NewServeMarketOrderGetV10DataOrderListOrderStatusFromValue returns a pointer to a valid ServeMarketOrderGetV10DataOrderListOrderStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServeMarketOrderGetV10DataOrderListOrderStatusFromValue(v string) (*ServeMarketOrderGetV10DataOrderListOrderStatus, error) {
	ev := ServeMarketOrderGetV10DataOrderListOrderStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServeMarketOrderGetV10DataOrderListOrderStatus: valid values are %v", v, AllowedServeMarketOrderGetV10DataOrderListOrderStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServeMarketOrderGetV10DataOrderListOrderStatus) IsValid() bool {
	for _, existing := range AllowedServeMarketOrderGetV10DataOrderListOrderStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to serve_market_order_get_v1.0_data_order_list_order_status value
func (v ServeMarketOrderGetV10DataOrderListOrderStatus) Ptr() *ServeMarketOrderGetV10DataOrderListOrderStatus {
	return &v
}
