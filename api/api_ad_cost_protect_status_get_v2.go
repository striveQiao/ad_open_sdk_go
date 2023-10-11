/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
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

// AdCostProtectStatusGetV2ApiService AdCostProtectStatusGetV2Api service
type AdCostProtectStatusGetV2ApiService service

type ApiOpenApi2AdCostProtectStatusGetGetRequest struct {
	ctx          context.Context
	ApiService   *AdCostProtectStatusGetV2ApiService
	advertiserId *int64
	adIds        *[]int64
}

// 广告主id
func (r *ApiOpenApi2AdCostProtectStatusGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2AdCostProtectStatusGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 广告计划id，每次最多传入50个
func (r *ApiOpenApi2AdCostProtectStatusGetGetRequest) AdIds(adIds []int64) *ApiOpenApi2AdCostProtectStatusGetGetRequest {
	r.adIds = &adIds
	return r
}

func (r *ApiOpenApi2AdCostProtectStatusGetGetRequest) Execute() (*AdCostProtectStatusGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2AdCostProtectStatusGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2AdCostProtectStatusGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2AdCostProtectStatusGetGetRequest) WithLog(enable bool) *ApiOpenApi2AdCostProtectStatusGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2AdCostProtectStatusGetGet Method for OpenApi2AdCostProtectStatusGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2AdCostProtectStatusGetGetRequest
*/
func (a *AdCostProtectStatusGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2AdCostProtectStatusGetGetRequest {
	return &ApiOpenApi2AdCostProtectStatusGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AdCostProtectStatusGetV2Response
func (a *AdCostProtectStatusGetV2ApiService) getExecute(r *ApiOpenApi2AdCostProtectStatusGetGetRequest) (*AdCostProtectStatusGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *AdCostProtectStatusGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/ad/cost_protect_status/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 1 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 1")
	}
	if r.adIds == nil {
		return localVarReturnValue, nil, ReportError("adIds is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "ad_ids", r.adIds)
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
