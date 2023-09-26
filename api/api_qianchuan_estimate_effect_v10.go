/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
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

// QianchuanEstimateEffectV10ApiService QianchuanEstimateEffectV10Api service
type QianchuanEstimateEffectV10ApiService service

type ApiOpenApiV10QianchuanEstimateEffectGetRequest struct {
	ctx                context.Context
	ApiService         *QianchuanEstimateEffectV10ApiService
	advertiserId       *int64
	awemeId            *int64
	externalAction     *QianchuanEstimateEffectV10ExternalAction
	budgetMode         *QianchuanEstimateEffectV10BudgetMode
	budget             *float64
	liveScheduleType   *QianchuanEstimateEffectV10LiveScheduleType
	deepExternalAction *QianchuanEstimateEffectV10DeepExternalAction
	deepBidType        *QianchuanEstimateEffectV10DeepBidType
	startTime          *string
	endTime            *string
	scheduleTime       *string
	scheduleFixedRange *int64
}

func (r *ApiOpenApiV10QianchuanEstimateEffectGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanEstimateEffectGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanEstimateEffectGetRequest) AwemeId(awemeId int64) *ApiOpenApiV10QianchuanEstimateEffectGetRequest {
	r.awemeId = &awemeId
	return r
}

func (r *ApiOpenApiV10QianchuanEstimateEffectGetRequest) ExternalAction(externalAction QianchuanEstimateEffectV10ExternalAction) *ApiOpenApiV10QianchuanEstimateEffectGetRequest {
	r.externalAction = &externalAction
	return r
}

func (r *ApiOpenApiV10QianchuanEstimateEffectGetRequest) BudgetMode(budgetMode QianchuanEstimateEffectV10BudgetMode) *ApiOpenApiV10QianchuanEstimateEffectGetRequest {
	r.budgetMode = &budgetMode
	return r
}

func (r *ApiOpenApiV10QianchuanEstimateEffectGetRequest) Budget(budget float64) *ApiOpenApiV10QianchuanEstimateEffectGetRequest {
	r.budget = &budget
	return r
}

func (r *ApiOpenApiV10QianchuanEstimateEffectGetRequest) LiveScheduleType(liveScheduleType QianchuanEstimateEffectV10LiveScheduleType) *ApiOpenApiV10QianchuanEstimateEffectGetRequest {
	r.liveScheduleType = &liveScheduleType
	return r
}

func (r *ApiOpenApiV10QianchuanEstimateEffectGetRequest) DeepExternalAction(deepExternalAction QianchuanEstimateEffectV10DeepExternalAction) *ApiOpenApiV10QianchuanEstimateEffectGetRequest {
	r.deepExternalAction = &deepExternalAction
	return r
}

func (r *ApiOpenApiV10QianchuanEstimateEffectGetRequest) DeepBidType(deepBidType QianchuanEstimateEffectV10DeepBidType) *ApiOpenApiV10QianchuanEstimateEffectGetRequest {
	r.deepBidType = &deepBidType
	return r
}

func (r *ApiOpenApiV10QianchuanEstimateEffectGetRequest) StartTime(startTime string) *ApiOpenApiV10QianchuanEstimateEffectGetRequest {
	r.startTime = &startTime
	return r
}

func (r *ApiOpenApiV10QianchuanEstimateEffectGetRequest) EndTime(endTime string) *ApiOpenApiV10QianchuanEstimateEffectGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApiV10QianchuanEstimateEffectGetRequest) ScheduleTime(scheduleTime string) *ApiOpenApiV10QianchuanEstimateEffectGetRequest {
	r.scheduleTime = &scheduleTime
	return r
}

func (r *ApiOpenApiV10QianchuanEstimateEffectGetRequest) ScheduleFixedRange(scheduleFixedRange int64) *ApiOpenApiV10QianchuanEstimateEffectGetRequest {
	r.scheduleFixedRange = &scheduleFixedRange
	return r
}

func (r *ApiOpenApiV10QianchuanEstimateEffectGetRequest) Execute() (*QianchuanEstimateEffectV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanEstimateEffectGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanEstimateEffectGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanEstimateEffectGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanEstimateEffectGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanEstimateEffectGet Method for OpenApiV10QianchuanEstimateEffectGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanEstimateEffectGetRequest
*/
func (a *QianchuanEstimateEffectV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanEstimateEffectGetRequest {
	return &ApiOpenApiV10QianchuanEstimateEffectGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanEstimateEffectV10Response
func (a *QianchuanEstimateEffectV10ApiService) getExecute(r *ApiOpenApiV10QianchuanEstimateEffectGetRequest) (*QianchuanEstimateEffectV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *QianchuanEstimateEffectV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/estimate/effect/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.awemeId == nil {
		return localVarReturnValue, nil, ReportError("awemeId is required and must be specified")
	}
	if r.externalAction == nil {
		return localVarReturnValue, nil, ReportError("externalAction is required and must be specified")
	}
	if r.budgetMode == nil {
		return localVarReturnValue, nil, ReportError("budgetMode is required and must be specified")
	}
	if r.budget == nil {
		return localVarReturnValue, nil, ReportError("budget is required and must be specified")
	}
	if *r.budget < 300.0 {
		return localVarReturnValue, nil, ReportError("budget must be greater than 300.0")
	}
	if *r.budget > 9999999.99 {
		return localVarReturnValue, nil, ReportError("budget must be less than 9999999.99")
	}
	if r.liveScheduleType == nil {
		return localVarReturnValue, nil, ReportError("liveScheduleType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_id", r.awemeId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "external_action", r.externalAction)
	if r.deepExternalAction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "deep_external_action", r.deepExternalAction)
	}
	if r.deepBidType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "deep_bid_type", r.deepBidType)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "budget_mode", r.budgetMode)
	parameterAddToHeaderOrQuery(localVarQueryParams, "budget", r.budget)
	parameterAddToHeaderOrQuery(localVarQueryParams, "live_schedule_type", r.liveScheduleType)
	if r.startTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
	}
	if r.endTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
	}
	if r.scheduleTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "schedule_time", r.scheduleTime)
	}
	if r.scheduleFixedRange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "schedule_fixed_range", r.scheduleFixedRange)
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
