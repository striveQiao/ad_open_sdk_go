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

// BrandAdCreateV30DateQuantityStockIncreasePackageType
type BrandAdCreateV30DateQuantityStockIncreasePackageType string

// List of brand_ad_create_v3.0_date_quantity_stock_increase_package_type
const (
	NONE_BrandAdCreateV30DateQuantityStockIncreasePackageType BrandAdCreateV30DateQuantityStockIncreasePackageType = "NONE"
	PLUS_BrandAdCreateV30DateQuantityStockIncreasePackageType BrandAdCreateV30DateQuantityStockIncreasePackageType = "PLUS"
	PRO_BrandAdCreateV30DateQuantityStockIncreasePackageType  BrandAdCreateV30DateQuantityStockIncreasePackageType = "PRO"
)

// All allowed values of BrandAdCreateV30DateQuantityStockIncreasePackageType enum
var AllowedBrandAdCreateV30DateQuantityStockIncreasePackageTypeEnumValues = []BrandAdCreateV30DateQuantityStockIncreasePackageType{
	"NONE",
	"PLUS",
	"PRO",
}

// NewBrandAdCreateV30DateQuantityStockIncreasePackageTypeFromValue returns a pointer to a valid BrandAdCreateV30DateQuantityStockIncreasePackageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdCreateV30DateQuantityStockIncreasePackageTypeFromValue(v string) (*BrandAdCreateV30DateQuantityStockIncreasePackageType, error) {
	ev := BrandAdCreateV30DateQuantityStockIncreasePackageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdCreateV30DateQuantityStockIncreasePackageType: valid values are %v", v, AllowedBrandAdCreateV30DateQuantityStockIncreasePackageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdCreateV30DateQuantityStockIncreasePackageType) IsValid() bool {
	for _, existing := range AllowedBrandAdCreateV30DateQuantityStockIncreasePackageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_create_v3.0_date_quantity_stock_increase_package_type value
func (v BrandAdCreateV30DateQuantityStockIncreasePackageType) Ptr() *BrandAdCreateV30DateQuantityStockIncreasePackageType {
	return &v
}
