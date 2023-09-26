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

// QianchuanTodayLiveRoomFlowPerformanceGetV10ApiService QianchuanTodayLiveRoomFlowPerformanceGetV10Api service
type QianchuanTodayLiveRoomFlowPerformanceGetV10ApiService service

type ApiOpenApiV10QianchuanTodayLiveRoomFlowPerformanceGetGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanTodayLiveRoomFlowPerformanceGetV10ApiService
	advertiserId *int64
	roomId       *int64
	flowSource   *QianchuanTodayLiveRoomFlowPerformanceGetV10FlowSource
}

func (r *ApiOpenApiV10QianchuanTodayLiveRoomFlowPerformanceGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanTodayLiveRoomFlowPerformanceGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanTodayLiveRoomFlowPerformanceGetGetRequest) RoomId(roomId int64) *ApiOpenApiV10QianchuanTodayLiveRoomFlowPerformanceGetGetRequest {
	r.roomId = &roomId
	return r
}

func (r *ApiOpenApiV10QianchuanTodayLiveRoomFlowPerformanceGetGetRequest) FlowSource(flowSource QianchuanTodayLiveRoomFlowPerformanceGetV10FlowSource) *ApiOpenApiV10QianchuanTodayLiveRoomFlowPerformanceGetGetRequest {
	r.flowSource = &flowSource
	return r
}

func (r *ApiOpenApiV10QianchuanTodayLiveRoomFlowPerformanceGetGetRequest) Execute() (*QianchuanTodayLiveRoomFlowPerformanceGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanTodayLiveRoomFlowPerformanceGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanTodayLiveRoomFlowPerformanceGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanTodayLiveRoomFlowPerformanceGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanTodayLiveRoomFlowPerformanceGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanTodayLiveRoomFlowPerformanceGetGet Method for OpenApiV10QianchuanTodayLiveRoomFlowPerformanceGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanTodayLiveRoomFlowPerformanceGetGetRequest
*/
func (a *QianchuanTodayLiveRoomFlowPerformanceGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanTodayLiveRoomFlowPerformanceGetGetRequest {
	return &ApiOpenApiV10QianchuanTodayLiveRoomFlowPerformanceGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanTodayLiveRoomFlowPerformanceGetV10Response
func (a *QianchuanTodayLiveRoomFlowPerformanceGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanTodayLiveRoomFlowPerformanceGetGetRequest) (*QianchuanTodayLiveRoomFlowPerformanceGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *QianchuanTodayLiveRoomFlowPerformanceGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/today_live/room/flow_performance/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.roomId == nil {
		return localVarReturnValue, nil, ReportError("roomId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "room_id", r.roomId)
	if r.flowSource != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "flow_source", r.flowSource)
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
