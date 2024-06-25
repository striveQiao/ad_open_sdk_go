/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileUploadTaskCreateV2Request struct for FileUploadTaskCreateV2Request
type FileUploadTaskCreateV2Request struct {
	// 账户ID
	AccountId   int64                             `json:"account_id"`
	AccountType FileUploadTaskCreateV2AccountType `json:"account_type"`
	// 素材的文件名
	Filename string `json:"filename"`
	// 视频素材是否为AIGC生成
	IsAigc *bool `json:"is_aigc,omitempty"`
	// 上传视频场景是否是引导视频（此功能仅游戏行业白名单客户可用，如需使用请联系运营） 允许值：FALSE，TRUE
	IsGuideVideo *bool `json:"is_guide_video,omitempty"`
	// 标注是否允许生效授权保护（前提素材需要是自产！！） 目前只支持true   当account_type= AGTENT时，必填
	IsNeedAuth *bool `json:"is_need_auth,omitempty"`
	// 标签
	Labels []string `json:"labels,omitempty"`
	// 视频 url地址，如http://xxx.xxx
	VideoUrl string `json:"video_url"`
}
