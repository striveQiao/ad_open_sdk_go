/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandAdGetV30ResponseDataAdsInnerAdStockInfosInner struct for BrandAdGetV30ResponseDataAdsInnerAdStockInfosInner
type BrandAdGetV30ResponseDataAdsInnerAdStockInfosInner struct {
	AdForm               BrandAdGetV30DataAdsAdStockInfosAdForm                `json:"ad_form"`
	AppOrigin            BrandAdGetV30DataAdsAdStockInfosAppOrigin             `json:"app_origin"`
	SuperiorCreativeType *BrandAdGetV30DataAdsAdStockInfosSuperiorCreativeType `json:"superior_creative_type,omitempty"`
}
