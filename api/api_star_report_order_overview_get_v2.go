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

// StarReportOrderOverviewGetV2ApiService StarReportOrderOverviewGetV2Api service
type StarReportOrderOverviewGetV2ApiService service

type ApiOpenApi2StarReportOrderOverviewGetGetRequest struct {
	ctx        context.Context
	ApiService *StarReportOrderOverviewGetV2ApiService
	orderId    *int64
	starId     *int64
}

func (r *ApiOpenApi2StarReportOrderOverviewGetGetRequest) OrderId(orderId int64) *ApiOpenApi2StarReportOrderOverviewGetGetRequest {
	r.orderId = &orderId
	return r
}

func (r *ApiOpenApi2StarReportOrderOverviewGetGetRequest) StarId(starId int64) *ApiOpenApi2StarReportOrderOverviewGetGetRequest {
	r.starId = &starId
	return r
}

func (r *ApiOpenApi2StarReportOrderOverviewGetGetRequest) Execute() (*StarReportOrderOverviewGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarReportOrderOverviewGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarReportOrderOverviewGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarReportOrderOverviewGetGetRequest) WithLog(enable bool) *ApiOpenApi2StarReportOrderOverviewGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarReportOrderOverviewGetGet Method for OpenApi2StarReportOrderOverviewGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarReportOrderOverviewGetGetRequest
*/
func (a *StarReportOrderOverviewGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarReportOrderOverviewGetGetRequest {
	return &ApiOpenApi2StarReportOrderOverviewGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarReportOrderOverviewGetV2Response
func (a *StarReportOrderOverviewGetV2ApiService) getExecute(r *ApiOpenApi2StarReportOrderOverviewGetGetRequest) (*StarReportOrderOverviewGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *StarReportOrderOverviewGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/report/order_overview/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.orderId == nil {
		return localVarReturnValue, nil, ReportError("orderId is required and must be specified")
	}
	if r.starId == nil {
		return localVarReturnValue, nil, ReportError("starId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "order_id", r.orderId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
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
