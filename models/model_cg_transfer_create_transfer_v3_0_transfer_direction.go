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

// CgTransferCreateTransferV30TransferDirection
type CgTransferCreateTransferV30TransferDirection string

// List of cg_transfer_create_transfer_v3.0_transfer_direction
const (
	TRANSFER_IN_CgTransferCreateTransferV30TransferDirection  CgTransferCreateTransferV30TransferDirection = "TRANSFER_IN"
	TRANSFER_OUT_CgTransferCreateTransferV30TransferDirection CgTransferCreateTransferV30TransferDirection = "TRANSFER_OUT"
)

// All allowed values of CgTransferCreateTransferV30TransferDirection enum
var AllowedCgTransferCreateTransferV30TransferDirectionEnumValues = []CgTransferCreateTransferV30TransferDirection{
	"TRANSFER_IN",
	"TRANSFER_OUT",
}

// NewCgTransferCreateTransferV30TransferDirectionFromValue returns a pointer to a valid CgTransferCreateTransferV30TransferDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCgTransferCreateTransferV30TransferDirectionFromValue(v string) (*CgTransferCreateTransferV30TransferDirection, error) {
	ev := CgTransferCreateTransferV30TransferDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CgTransferCreateTransferV30TransferDirection: valid values are %v", v, AllowedCgTransferCreateTransferV30TransferDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CgTransferCreateTransferV30TransferDirection) IsValid() bool {
	for _, existing := range AllowedCgTransferCreateTransferV30TransferDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cg_transfer_create_transfer_v3.0_transfer_direction value
func (v CgTransferCreateTransferV30TransferDirection) Ptr() *CgTransferCreateTransferV30TransferDirection {
	return &v
}
