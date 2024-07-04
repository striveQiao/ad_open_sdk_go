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

// ToolsClueRefundReportGetV2ApiService ToolsClueRefundReportGetV2Api service
type ToolsClueRefundReportGetV2ApiService service

type ApiOpenApi2ToolsClueRefundReportGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsClueRefundReportGetV2ApiService
	advertiserId *int64
	month        *string
}

// 广告主id
func (r *ApiOpenApi2ToolsClueRefundReportGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsClueRefundReportGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 查询账单月份，格式YYYYMM，只支持查询202112月及之后的账单
func (r *ApiOpenApi2ToolsClueRefundReportGetGetRequest) Month(month string) *ApiOpenApi2ToolsClueRefundReportGetGetRequest {
	r.month = &month
	return r
}

func (r *ApiOpenApi2ToolsClueRefundReportGetGetRequest) Execute() (*ToolsClueRefundReportGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsClueRefundReportGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsClueRefundReportGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsClueRefundReportGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsClueRefundReportGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsClueRefundReportGetGet Method for OpenApi2ToolsClueRefundReportGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsClueRefundReportGetGetRequest
*/
func (a *ToolsClueRefundReportGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsClueRefundReportGetGetRequest {
	return &ApiOpenApi2ToolsClueRefundReportGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsClueRefundReportGetV2Response
func (a *ToolsClueRefundReportGetV2ApiService) getExecute(r *ApiOpenApi2ToolsClueRefundReportGetGetRequest) (*ToolsClueRefundReportGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsClueRefundReportGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/clue/refund_report/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 1 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 1")
	}
	if r.month == nil {
		return localVarReturnValue, nil, ReportError("month is required and must be specified")
	}
	if strlen(*r.month) < 6 {
		return localVarReturnValue, nil, ReportError("month must have at least 6 elements")
	}
	if strlen(*r.month) > 6 {
		return localVarReturnValue, nil, ReportError("month must have less than 6 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "month", r.month)
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
