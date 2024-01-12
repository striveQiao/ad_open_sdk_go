/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaAssetsDetailReadV2ResponseDataListInner struct for DpaAssetsDetailReadV2ResponseDataListInner
type DpaAssetsDetailReadV2ResponseDataListInner struct {
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	AssetCreateTime **string `json:"asset_create_time,omitempty"`
	//
	AssetId *int64 `json:"asset_id,omitempty"`
	//
	AssetModifyTime **string                                  `json:"asset_modify_time,omitempty"`
	AssetType       *DpaAssetsDetailReadV2DataListAssetType   `json:"asset_type,omitempty"`
	AuditStatus     *DpaAssetsDetailReadV2DataListAuditStatus `json:"audit_status,omitempty"`
	//
	PlatformId *int64 `json:"platform_id,omitempty"`
	//
	ProductId *int64                               `json:"product_id,omitempty"`
	Source    *DpaAssetsDetailReadV2DataListSource `json:"source,omitempty"`
	Status    *DpaAssetsDetailReadV2DataListStatus `json:"status,omitempty"`
}
