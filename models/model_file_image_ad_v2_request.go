/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileImageAdV2Request struct for FileImageAdV2Request
type FileImageAdV2Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 素材的文件名，可自定义素材名，不传则默认取文件名，最长255个字符 注：若同一素材已进行上传，重新上传不会改名
	Filename *string `json:"filename,omitempty"`
	// 图片文件 upload_type为UPLOAD_BY_FILE必填 格式: jpg、jpeg、png、bmp、gif, 大小1.5M内
	ImageFile **FormFileInfo `json:"image_file,omitempty"`
	// 图片的md5值(用于服务端校验) upload_type为UPLOAD_BY_FILE
	ImageSignature *string `json:"image_signature,omitempty"`
	// 图片url地址，如http://xxx.xxx upload_type为UPLOAD_BY_URL必填
	ImageUrl *string `json:"image_url,omitempty"`
	// 图片素材是否为AIGC生成
	IsAigc     *bool                    `json:"is_aigc,omitempty"`
	UploadType *FileImageAdV2UploadType `json:"upload_type,omitempty"`
}
