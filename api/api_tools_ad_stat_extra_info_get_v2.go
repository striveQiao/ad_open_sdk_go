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

// ToolsAdStatExtraInfoGetV2ApiService ToolsAdStatExtraInfoGetV2Api service
type ToolsAdStatExtraInfoGetV2ApiService service

type ApiOpenApi2ToolsAdStatExtraInfoGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsAdStatExtraInfoGetV2ApiService
	adIds        *[]int64
	advertiserId *int64
}

func (r *ApiOpenApi2ToolsAdStatExtraInfoGetGetRequest) AdIds(adIds []int64) *ApiOpenApi2ToolsAdStatExtraInfoGetGetRequest {
	r.adIds = &adIds
	return r
}

func (r *ApiOpenApi2ToolsAdStatExtraInfoGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsAdStatExtraInfoGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsAdStatExtraInfoGetGetRequest) Execute() (*ToolsAdStatExtraInfoGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsAdStatExtraInfoGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAdStatExtraInfoGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAdStatExtraInfoGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsAdStatExtraInfoGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAdStatExtraInfoGetGet Method for OpenApi2ToolsAdStatExtraInfoGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsAdStatExtraInfoGetGetRequest
*/
func (a *ToolsAdStatExtraInfoGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsAdStatExtraInfoGetGetRequest {
	return &ApiOpenApi2ToolsAdStatExtraInfoGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAdStatExtraInfoGetV2Response
func (a *ToolsAdStatExtraInfoGetV2ApiService) getExecute(r *ApiOpenApi2ToolsAdStatExtraInfoGetGetRequest) (*ToolsAdStatExtraInfoGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsAdStatExtraInfoGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/ad_stat_extra_info/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.adIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ad_ids", r.adIds)
	}
	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
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
