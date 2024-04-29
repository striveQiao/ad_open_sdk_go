/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
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

// CreativeAutoGenerateConfigGetV2ApiService CreativeAutoGenerateConfigGetV2Api service
type CreativeAutoGenerateConfigGetV2ApiService service

type ApiOpenApi2CreativeAutoGenerateConfigGetGetRequest struct {
	ctx          context.Context
	ApiService   *CreativeAutoGenerateConfigGetV2ApiService
	advertiserId *int64
	adId         *int64
}

// 广告主ID
func (r *ApiOpenApi2CreativeAutoGenerateConfigGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2CreativeAutoGenerateConfigGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 广告计划ID
func (r *ApiOpenApi2CreativeAutoGenerateConfigGetGetRequest) AdId(adId int64) *ApiOpenApi2CreativeAutoGenerateConfigGetGetRequest {
	r.adId = &adId
	return r
}

func (r *ApiOpenApi2CreativeAutoGenerateConfigGetGetRequest) Execute() (*CreativeAutoGenerateConfigGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2CreativeAutoGenerateConfigGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2CreativeAutoGenerateConfigGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2CreativeAutoGenerateConfigGetGetRequest) WithLog(enable bool) *ApiOpenApi2CreativeAutoGenerateConfigGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2CreativeAutoGenerateConfigGetGet Method for OpenApi2CreativeAutoGenerateConfigGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2CreativeAutoGenerateConfigGetGetRequest
*/
func (a *CreativeAutoGenerateConfigGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2CreativeAutoGenerateConfigGetGetRequest {
	return &ApiOpenApi2CreativeAutoGenerateConfigGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreativeAutoGenerateConfigGetV2Response
func (a *CreativeAutoGenerateConfigGetV2ApiService) getExecute(r *ApiOpenApi2CreativeAutoGenerateConfigGetGetRequest) (*CreativeAutoGenerateConfigGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *CreativeAutoGenerateConfigGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/creative/auto_generate_config/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.adId == nil {
		return localVarReturnValue, nil, ReportError("adId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "ad_id", r.adId)
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
