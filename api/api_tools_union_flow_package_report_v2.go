/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
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

// ToolsUnionFlowPackageReportV2ApiService ToolsUnionFlowPackageReportV2Api service
type ToolsUnionFlowPackageReportV2ApiService service

type ApiOpenApi2ToolsUnionFlowPackageReportGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsUnionFlowPackageReportV2ApiService
	advertiserId *int64
	filter       *ToolsUnionFlowPackageReportV2Filter
	orderField   *string
	orderType    *ToolsUnionFlowPackageReportV2OrderType
	page         *int64
	pageSize     *int64
}

func (r *ApiOpenApi2ToolsUnionFlowPackageReportGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsUnionFlowPackageReportGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsUnionFlowPackageReportGetRequest) Filter(filter ToolsUnionFlowPackageReportV2Filter) *ApiOpenApi2ToolsUnionFlowPackageReportGetRequest {
	r.filter = &filter
	return r
}

func (r *ApiOpenApi2ToolsUnionFlowPackageReportGetRequest) OrderField(orderField string) *ApiOpenApi2ToolsUnionFlowPackageReportGetRequest {
	r.orderField = &orderField
	return r
}

func (r *ApiOpenApi2ToolsUnionFlowPackageReportGetRequest) OrderType(orderType ToolsUnionFlowPackageReportV2OrderType) *ApiOpenApi2ToolsUnionFlowPackageReportGetRequest {
	r.orderType = &orderType
	return r
}

func (r *ApiOpenApi2ToolsUnionFlowPackageReportGetRequest) Page(page int64) *ApiOpenApi2ToolsUnionFlowPackageReportGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ToolsUnionFlowPackageReportGetRequest) PageSize(pageSize int64) *ApiOpenApi2ToolsUnionFlowPackageReportGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ToolsUnionFlowPackageReportGetRequest) Execute() (*ToolsUnionFlowPackageReportV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsUnionFlowPackageReportGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsUnionFlowPackageReportGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsUnionFlowPackageReportGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsUnionFlowPackageReportGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsUnionFlowPackageReportGet Method for OpenApi2ToolsUnionFlowPackageReportGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsUnionFlowPackageReportGetRequest
*/
func (a *ToolsUnionFlowPackageReportV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsUnionFlowPackageReportGetRequest {
	return &ApiOpenApi2ToolsUnionFlowPackageReportGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsUnionFlowPackageReportV2Response
func (a *ToolsUnionFlowPackageReportV2ApiService) getExecute(r *ApiOpenApi2ToolsUnionFlowPackageReportGetRequest) (*ToolsUnionFlowPackageReportV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsUnionFlowPackageReportV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/union/flow_package/report/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter)
	}
	if r.orderField != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_field", r.orderField)
	}
	if r.orderType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_type", r.orderType)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
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
