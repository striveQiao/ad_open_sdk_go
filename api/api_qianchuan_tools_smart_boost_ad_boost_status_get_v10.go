/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
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

// QianchuanToolsSmartBoostAdBoostStatusGetV10ApiService QianchuanToolsSmartBoostAdBoostStatusGetV10Api service
type QianchuanToolsSmartBoostAdBoostStatusGetV10ApiService service

type ApiOpenApiV10QianchuanToolsSmartBoostAdBoostStatusGetGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanToolsSmartBoostAdBoostStatusGetV10ApiService
	advertiserId *int64
	adIds        *[]int64
}

// 千川广告主账户ID。
func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostStatusGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostStatusGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 广告计划id，限制1-10个。
func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostStatusGetGetRequest) AdIds(adIds []int64) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostStatusGetGetRequest {
	r.adIds = &adIds
	return r
}

func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostStatusGetGetRequest) Execute() (*QianchuanToolsSmartBoostAdBoostStatusGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostStatusGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostStatusGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostStatusGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostStatusGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanToolsSmartBoostAdBoostStatusGetGet Method for OpenApiV10QianchuanToolsSmartBoostAdBoostStatusGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanToolsSmartBoostAdBoostStatusGetGetRequest
*/
func (a *QianchuanToolsSmartBoostAdBoostStatusGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostStatusGetGetRequest {
	return &ApiOpenApiV10QianchuanToolsSmartBoostAdBoostStatusGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanToolsSmartBoostAdBoostStatusGetV10Response
func (a *QianchuanToolsSmartBoostAdBoostStatusGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostStatusGetGetRequest) (*QianchuanToolsSmartBoostAdBoostStatusGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *QianchuanToolsSmartBoostAdBoostStatusGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/tools/smart_boost/ad_boost/status/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
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
