/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandCreateAssignV2RequestDemandInfoAdSyncConf 广告推送配置
type StarDemandCreateAssignV2RequestDemandInfoAdSyncConf struct {
	ContentMarketing *StarDemandCreateAssignV2RequestDemandInfoAdSyncConfContentMarketing `json:"content_marketing,omitempty"`
	DouPlus          *StarDemandCreateAssignV2RequestDemandInfoAdSyncConfDouPlus          `json:"dou_plus,omitempty"`
	EcomShop         *StarDemandCreateAssignV2RequestDemandInfoAdSyncConfEcomShop         `json:"ecom_shop,omitempty"`
	LocalPromotion   *StarDemandCreateAssignV2RequestDemandInfoAdSyncConfLocalPromotion   `json:"local_promotion,omitempty"`
	OceanEngine      *StarDemandCreateAssignV2RequestDemandInfoAdSyncConfOceanEngine      `json:"ocean_engine,omitempty"`
	Qianchuan        *StarDemandCreateAssignV2RequestDemandInfoAdSyncConfQianchuan        `json:"qianchuan,omitempty"`
}
