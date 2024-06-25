/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarChallengeInfoV2ResponseDataTaskInfoListInnerDemandInfoAdSyncConfOceanEngine
type StarChallengeInfoV2ResponseDataTaskInfoListInnerDemandInfoAdSyncConfOceanEngine struct {
	//
	AdSyncOrigin *int64 `json:"ad_sync_origin,omitempty"`
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	AutoSync *int64 `json:"auto_sync,omitempty"`
	//
	DouPlusUid *int64 `json:"dou_plus_uid,omitempty"`
	//
	ProductLink *string `json:"product_link,omitempty"`
	//
	ProductPics []string `json:"product_pics,omitempty"`
	//
	SyncDuration *int64 `json:"sync_duration,omitempty"`
}
