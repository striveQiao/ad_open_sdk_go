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

// StarDemandOrderListV2ApiService StarDemandOrderListV2Api service
type StarDemandOrderListV2ApiService service

type ApiOpenApi2StarDemandOrderListGetRequest struct {
	ctx        context.Context
	ApiService *StarDemandOrderListV2ApiService
	starId     *int64
	demandId   *int64
	filtering  *StarDemandOrderListV2Filtering
	page       *int32
	pageSize   *int32
}

// 星图id，星图客户授权后，通过“获取已授权账户”接口，查询到账号角色为”6-星图账号“的账户id，即为星图id
func (r *ApiOpenApi2StarDemandOrderListGetRequest) StarId(starId int64) *ApiOpenApi2StarDemandOrderListGetRequest {
	r.starId = &starId
	return r
}

// 任务id，可通过“获取星图客户任务列表”获取
func (r *ApiOpenApi2StarDemandOrderListGetRequest) DemandId(demandId int64) *ApiOpenApi2StarDemandOrderListGetRequest {
	r.demandId = &demandId
	return r
}

// 过滤条件，若此字段不传，或传空则视为无限制条件
func (r *ApiOpenApi2StarDemandOrderListGetRequest) Filtering(filtering StarDemandOrderListV2Filtering) *ApiOpenApi2StarDemandOrderListGetRequest {
	r.filtering = &filtering
	return r
}

// 页码，默认为1
func (r *ApiOpenApi2StarDemandOrderListGetRequest) Page(page int32) *ApiOpenApi2StarDemandOrderListGetRequest {
	r.page = &page
	return r
}

// 页面大小，默认10，最大值50
func (r *ApiOpenApi2StarDemandOrderListGetRequest) PageSize(pageSize int32) *ApiOpenApi2StarDemandOrderListGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2StarDemandOrderListGetRequest) Execute() (*StarDemandOrderListV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarDemandOrderListGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarDemandOrderListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarDemandOrderListGetRequest) WithLog(enable bool) *ApiOpenApi2StarDemandOrderListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarDemandOrderListGet Method for OpenApi2StarDemandOrderListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarDemandOrderListGetRequest
*/
func (a *StarDemandOrderListV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarDemandOrderListGetRequest {
	return &ApiOpenApi2StarDemandOrderListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarDemandOrderListV2Response
func (a *StarDemandOrderListV2ApiService) getExecute(r *ApiOpenApi2StarDemandOrderListGetRequest) (*StarDemandOrderListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarDemandOrderListV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/demand/order/list/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.starId == nil {
		return localVarReturnValue, nil, ReportError("starId is required and must be specified")
	}
	if r.demandId == nil {
		return localVarReturnValue, nil, ReportError("demandId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "demand_id", r.demandId)
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
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
