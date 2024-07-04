/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
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

// ToolsCreativeWordSelectV2ApiService ToolsCreativeWordSelectV2Api service
type ToolsCreativeWordSelectV2ApiService service

type ApiOpenApi2ToolsCreativeWordSelectGetRequest struct {
	ctx             context.Context
	ApiService      *ToolsCreativeWordSelectV2ApiService
	advertiserId    *int64
	creativeWordIds *[]int64
}

// 广告主ID
func (r *ApiOpenApi2ToolsCreativeWordSelectGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsCreativeWordSelectGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 创意词包id列表，如不填默认返回所有创意词包
func (r *ApiOpenApi2ToolsCreativeWordSelectGetRequest) CreativeWordIds(creativeWordIds []int64) *ApiOpenApi2ToolsCreativeWordSelectGetRequest {
	r.creativeWordIds = &creativeWordIds
	return r
}

func (r *ApiOpenApi2ToolsCreativeWordSelectGetRequest) Execute() (*ToolsCreativeWordSelectV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsCreativeWordSelectGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsCreativeWordSelectGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsCreativeWordSelectGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsCreativeWordSelectGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsCreativeWordSelectGet Method for OpenApi2ToolsCreativeWordSelectGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsCreativeWordSelectGetRequest
*/
func (a *ToolsCreativeWordSelectV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsCreativeWordSelectGetRequest {
	return &ApiOpenApi2ToolsCreativeWordSelectGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsCreativeWordSelectV2Response
func (a *ToolsCreativeWordSelectV2ApiService) getExecute(r *ApiOpenApi2ToolsCreativeWordSelectGetRequest) (*ToolsCreativeWordSelectV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsCreativeWordSelectV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/creative_word/select/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.creativeWordIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "creative_word_ids", r.creativeWordIds)
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
