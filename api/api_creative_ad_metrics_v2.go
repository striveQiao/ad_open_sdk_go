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

// CreativeAdMetricsV2ApiService CreativeAdMetricsV2Api service
type CreativeAdMetricsV2ApiService service

type ApiOpenApi2CreativeAdMetricsGetRequest struct {
	ctx          context.Context
	ApiService   *CreativeAdMetricsV2ApiService
	advertiserId *int64
	adIds        *[]int64
	startDate    *string
	endDate      *string
}

// 广告主 ID
func (r *ApiOpenApi2CreativeAdMetricsGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2CreativeAdMetricsGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 广告计划 ID
func (r *ApiOpenApi2CreativeAdMetricsGetRequest) AdIds(adIds []int64) *ApiOpenApi2CreativeAdMetricsGetRequest {
	r.adIds = &adIds
	return r
}

// 统计开始日期，格式 yyyy-mm-dd
func (r *ApiOpenApi2CreativeAdMetricsGetRequest) StartDate(startDate string) *ApiOpenApi2CreativeAdMetricsGetRequest {
	r.startDate = &startDate
	return r
}

// 统计结束日期，格式 yyyy-mm-dd
func (r *ApiOpenApi2CreativeAdMetricsGetRequest) EndDate(endDate string) *ApiOpenApi2CreativeAdMetricsGetRequest {
	r.endDate = &endDate
	return r
}

func (r *ApiOpenApi2CreativeAdMetricsGetRequest) Execute() (*CreativeAdMetricsV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2CreativeAdMetricsGetRequest) AccessToken(accessToken string) *ApiOpenApi2CreativeAdMetricsGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2CreativeAdMetricsGetRequest) WithLog(enable bool) *ApiOpenApi2CreativeAdMetricsGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2CreativeAdMetricsGet Method for OpenApi2CreativeAdMetricsGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2CreativeAdMetricsGetRequest
*/
func (a *CreativeAdMetricsV2ApiService) Get(ctx context.Context) *ApiOpenApi2CreativeAdMetricsGetRequest {
	return &ApiOpenApi2CreativeAdMetricsGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreativeAdMetricsV2Response
func (a *CreativeAdMetricsV2ApiService) getExecute(r *ApiOpenApi2CreativeAdMetricsGetRequest) (*CreativeAdMetricsV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *CreativeAdMetricsV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/creative/ad/metrics/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.adIds == nil {
		return localVarReturnValue, nil, ReportError("adIds is required and must be specified")
	}
	if r.startDate == nil {
		return localVarReturnValue, nil, ReportError("startDate is required and must be specified")
	}
	if r.endDate == nil {
		return localVarReturnValue, nil, ReportError("endDate is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "ad_ids", r.adIds)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate)
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
