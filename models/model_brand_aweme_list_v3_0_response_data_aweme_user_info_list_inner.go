/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandAwemeListV30ResponseDataAwemeUserInfoListInner struct for BrandAwemeListV30ResponseDataAwemeUserInfoListInner
type BrandAwemeListV30ResponseDataAwemeUserInfoListInner struct {
	// 广告主ID
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	// 抖音号
	AdvertiserName *string                                           `json:"advertiser_name,omitempty"`
	AppName        *BrandAwemeListV30DataAwemeUserInfoListAppName    `json:"app_name,omitempty"`
	AuthStatus     *BrandAwemeListV30DataAwemeUserInfoListAuthStatus `json:"auth_status,omitempty"`
	AuthType       *BrandAwemeListV30DataAwemeUserInfoListAuthType   `json:"auth_type,omitempty"`
	// 抖音号ID
	AwemeId *string `json:"aweme_id,omitempty"`
	// 抖音号名称
	AwemeName *string `json:"aweme_name,omitempty"`
	// 绑定时间
	BindTime *string `json:"bind_time,omitempty"`
	// 剩余有效天数
	LeftValidDays   *int64                                                 `json:"left_valid_days,omitempty"`
	OperatePlatform *BrandAwemeListV30DataAwemeUserInfoListOperatePlatform `json:"operate_platform,omitempty"`
}
