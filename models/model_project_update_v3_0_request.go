/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectUpdateV30Request struct for ProjectUpdateV30Request
type ProjectUpdateV30Request struct {
	//
	AdvertiserId    int64                                   `json:"advertiser_id"`
	Audience        *ProjectUpdateV30RequestAudience        `json:"audience,omitempty"`
	AudienceExtend  *ProjectUpdateV30AudienceExtend         `json:"audience_extend,omitempty"`
	DeliverySetting *ProjectUpdateV30RequestDeliverySetting `json:"delivery_setting,omitempty"`
	DownloadMode    *ProjectUpdateV30DownloadMode           `json:"download_mode,omitempty"`
	//
	DpaCategories []int64 `json:"dpa_categories,omitempty"`
	//
	Keywords []*ProjectUpdateV30RequestKeywordsInner `json:"keywords,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	OpenUrl *string `json:"open_url,omitempty"`
	//
	OpenUrlField *string `json:"open_url_field,omitempty"`
	//
	OpenUrlParams *string `json:"open_url_params,omitempty"`
	//
	ProductTarget []*ProjectUpdateV30RequestProductTargetInner `json:"product_target,omitempty"`
	//
	ProjectId int64 `json:"project_id"`
	//
	SearchBidRatio  *float32                                `json:"search_bid_ratio,omitempty"`
	TrackUrlSetting *ProjectUpdateV30RequestTrackUrlSetting `json:"track_url_setting,omitempty"`
	//
	UlinkUrl *string `json:"ulink_url,omitempty"`
}
