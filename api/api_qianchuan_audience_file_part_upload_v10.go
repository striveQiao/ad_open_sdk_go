/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
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

// QianchuanAudienceFilePartUploadV10ApiService QianchuanAudienceFilePartUploadV10Api service
type QianchuanAudienceFilePartUploadV10ApiService service

type ApiOpenApiV10QianchuanAudienceFilePartUploadPostRequest struct {
	ctx          context.Context
	ApiService   *QianchuanAudienceFilePartUploadV10ApiService
	advertiserId *int64
	file         *FormFileInfo
	isFinished   *int32
	partNum      *int64
	fileKey      *string
}

func (r *ApiOpenApiV10QianchuanAudienceFilePartUploadPostRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanAudienceFilePartUploadPostRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanAudienceFilePartUploadPostRequest) File(file *FormFileInfo) *ApiOpenApiV10QianchuanAudienceFilePartUploadPostRequest {
	r.file = file
	return r
}

func (r *ApiOpenApiV10QianchuanAudienceFilePartUploadPostRequest) IsFinished(isFinished int32) *ApiOpenApiV10QianchuanAudienceFilePartUploadPostRequest {
	r.isFinished = &isFinished
	return r
}

func (r *ApiOpenApiV10QianchuanAudienceFilePartUploadPostRequest) PartNum(partNum int64) *ApiOpenApiV10QianchuanAudienceFilePartUploadPostRequest {
	r.partNum = &partNum
	return r
}

func (r *ApiOpenApiV10QianchuanAudienceFilePartUploadPostRequest) FileKey(fileKey string) *ApiOpenApiV10QianchuanAudienceFilePartUploadPostRequest {
	r.fileKey = &fileKey
	return r
}

func (r *ApiOpenApiV10QianchuanAudienceFilePartUploadPostRequest) Execute() (*QianchuanAudienceFilePartUploadV10Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV10QianchuanAudienceFilePartUploadPostRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanAudienceFilePartUploadPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanAudienceFilePartUploadPostRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanAudienceFilePartUploadPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanAudienceFilePartUploadPost Method for OpenApiV10QianchuanAudienceFilePartUploadPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanAudienceFilePartUploadPostRequest
*/
func (a *QianchuanAudienceFilePartUploadV10ApiService) Post(ctx context.Context) *ApiOpenApiV10QianchuanAudienceFilePartUploadPostRequest {
	return &ApiOpenApiV10QianchuanAudienceFilePartUploadPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanAudienceFilePartUploadV10Response
func (a *QianchuanAudienceFilePartUploadV10ApiService) postExecute(r *ApiOpenApiV10QianchuanAudienceFilePartUploadPostRequest) (*QianchuanAudienceFilePartUploadV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanAudienceFilePartUploadV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/audience_file/part_upload/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.file == nil {
		return localVarReturnValue, nil, ReportError("file is required and must be specified")
	}
	if r.isFinished == nil {
		return localVarReturnValue, nil, ReportError("isFinished is required and must be specified")
	}
	if *r.isFinished < 0 {
		return localVarReturnValue, nil, ReportError("isFinished must be greater than 0")
	}
	if *r.isFinished > 1 {
		return localVarReturnValue, nil, ReportError("isFinished must be less than 1")
	}
	if r.partNum == nil {
		return localVarReturnValue, nil, ReportError("partNum is required and must be specified")
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
	if r.file != nil {
		formFiles["file"] = r.file
	}
	if r.fileKey != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "file_key", r.fileKey)
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "is_finished", r.isFinished)
	parameterAddToHeaderOrQuery(localVarFormParams, "part_num", r.partNum)
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
