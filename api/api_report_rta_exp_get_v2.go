/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
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

// ReportRtaExpGetV2ApiService ReportRtaExpGetV2Api service
type ReportRtaExpGetV2ApiService service

type ApiOpenApi2ReportRtaExpGetGetRequest struct {
	ctx          context.Context
	ApiService   *ReportRtaExpGetV2ApiService
	advertiserId *int64
	rtaId        *int64
	startDate    *string
	endDate      *string
	strategy     *int64
}

// 广告主ID
func (r *ApiOpenApi2ReportRtaExpGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ReportRtaExpGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// RTA ID
func (r *ApiOpenApi2ReportRtaExpGetGetRequest) RtaId(rtaId int64) *ApiOpenApi2ReportRtaExpGetGetRequest {
	r.rtaId = &rtaId
	return r
}

// 开始时间，格式YYYYMMDD，示例：20220701
func (r *ApiOpenApi2ReportRtaExpGetGetRequest) StartDate(startDate string) *ApiOpenApi2ReportRtaExpGetGetRequest {
	r.startDate = &startDate
	return r
}

// 结束时间，格式YYYYMMDD，示例：20220701
func (r *ApiOpenApi2ReportRtaExpGetGetRequest) EndDate(endDate string) *ApiOpenApi2ReportRtaExpGetGetRequest {
	r.endDate = &endDate
	return r
}

// 联合实验策略。允许值：0 代表基线策略 ，传入1、2、3、4、5、6、7、8、9
func (r *ApiOpenApi2ReportRtaExpGetGetRequest) Strategy(strategy int64) *ApiOpenApi2ReportRtaExpGetGetRequest {
	r.strategy = &strategy
	return r
}

func (r *ApiOpenApi2ReportRtaExpGetGetRequest) Execute() (*ReportRtaExpGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ReportRtaExpGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ReportRtaExpGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ReportRtaExpGetGetRequest) WithLog(enable bool) *ApiOpenApi2ReportRtaExpGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ReportRtaExpGetGet Method for OpenApi2ReportRtaExpGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ReportRtaExpGetGetRequest
*/
func (a *ReportRtaExpGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ReportRtaExpGetGetRequest {
	return &ApiOpenApi2ReportRtaExpGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReportRtaExpGetV2Response
func (a *ReportRtaExpGetV2ApiService) getExecute(r *ApiOpenApi2ReportRtaExpGetGetRequest) (*ReportRtaExpGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ReportRtaExpGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/report/rta_exp/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.rtaId == nil {
		return localVarReturnValue, nil, ReportError("rtaId is required and must be specified")
	}
	if r.startDate == nil {
		return localVarReturnValue, nil, ReportError("startDate is required and must be specified")
	}
	if r.endDate == nil {
		return localVarReturnValue, nil, ReportError("endDate is required and must be specified")
	}
	if r.strategy == nil {
		return localVarReturnValue, nil, ReportError("strategy is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "rta_id", r.rtaId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "strategy", r.strategy)
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
