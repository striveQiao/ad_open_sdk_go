/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// FileAudioAdV2ApiService FileAudioAdV2Api service
type FileAudioAdV2ApiService service

type ApiOpenApi2FileAudioAdPostRequest struct {
	ctx            context.Context
	ApiService     *FileAudioAdV2ApiService
	advertiserId   *int64
	uploadType     *FileAudioAdV2UploadType
	audioFile      *FormFileInfo
	audioSignature *string
	audioUrl       *string
}

// 广告主ID
func (r *ApiOpenApi2FileAudioAdPostRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2FileAudioAdPostRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2FileAudioAdPostRequest) UploadType(uploadType FileAudioAdV2UploadType) *ApiOpenApi2FileAudioAdPostRequest {
	r.uploadType = &uploadType
	return r
}

func (r *ApiOpenApi2FileAudioAdPostRequest) AudioFile(audioFile *FormFileInfo) *ApiOpenApi2FileAudioAdPostRequest {
	r.audioFile = audioFile
	return r
}

// 图片的md5值(用于服务端校验) upload_type为UPLOAD_BY_FILE
func (r *ApiOpenApi2FileAudioAdPostRequest) AudioSignature(audioSignature string) *ApiOpenApi2FileAudioAdPostRequest {
	r.audioSignature = &audioSignature
	return r
}

func (r *ApiOpenApi2FileAudioAdPostRequest) AudioUrl(audioUrl string) *ApiOpenApi2FileAudioAdPostRequest {
	r.audioUrl = &audioUrl
	return r
}

func (r *ApiOpenApi2FileAudioAdPostRequest) Execute() (*FileAudioAdV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2FileAudioAdPostRequest) AccessToken(accessToken string) *ApiOpenApi2FileAudioAdPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2FileAudioAdPostRequest) WithLog(enable bool) *ApiOpenApi2FileAudioAdPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2FileAudioAdPost Method for OpenApi2FileAudioAdPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2FileAudioAdPostRequest
*/
func (a *FileAudioAdV2ApiService) Post(ctx context.Context) *ApiOpenApi2FileAudioAdPostRequest {
	return &ApiOpenApi2FileAudioAdPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FileAudioAdV2Response
func (a *FileAudioAdV2ApiService) postExecute(r *ApiOpenApi2FileAudioAdPostRequest) (*FileAudioAdV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *FileAudioAdV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/file/audio/ad/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.uploadType == nil {
		return localVarReturnValue, nil, ReportError("uploadType is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "advertiser_id", r.advertiserId)
	if r.audioFile != nil {
		formFiles["audio_file"] = r.audioFile
	}
	if r.audioSignature != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "audio_signature", r.audioSignature)
	}
	if r.audioUrl != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "audio_url", r.audioUrl)
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "upload_type", r.uploadType)
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
