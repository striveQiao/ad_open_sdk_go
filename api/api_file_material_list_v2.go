/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
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

// FileMaterialListV2ApiService FileMaterialListV2Api service
type FileMaterialListV2ApiService service

type ApiOpenApi2FileMaterialListGetRequest struct {
	ctx              context.Context
	ApiService       *FileMaterialListV2ApiService
	advertiserId     *int64
	materialSource   *FileMaterialListV2MaterialSource
	propertiesFilter *[]*FileMaterialListV2PropertiesFilter
	startTime        *string
	endTime          *string
	page             *int32
	pageSize         *int32
}

// 广告主 id
func (r *ApiOpenApi2FileMaterialListGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2FileMaterialListGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 素材来源，允许值： QIANCHUAN 千川 AD 广告平台
func (r *ApiOpenApi2FileMaterialListGetRequest) MaterialSource(materialSource FileMaterialListV2MaterialSource) *ApiOpenApi2FileMaterialListGetRequest {
	r.materialSource = &materialSource
	return r
}

// 素材标签过滤项， 允许值： INEFFICIENT_MATERIAL低效素材； SIMILAR_MATERIAL 同质化挤压严重素材； SIMILAR_QUEUE_MATERIAL 同质化素材风险-排队投放素材 AD_HIGH_QUALITY_MATERIAL AD 优质素材 ECP_HIGH_QUALITY_MATERIAL 千川优质素材 FIRST_PUBLISH_MATERIAL  首发素材 如果不传，则默认为全部素材标签
func (r *ApiOpenApi2FileMaterialListGetRequest) PropertiesFilter(propertiesFilter []*FileMaterialListV2PropertiesFilter) *ApiOpenApi2FileMaterialListGetRequest {
	r.propertiesFilter = &propertiesFilter
	return r
}

// 素材创建时间，格式为yyyy-mm-dd HH:MM:SS， 最早允许传入时间：2022-01-01 00:00:00
func (r *ApiOpenApi2FileMaterialListGetRequest) StartTime(startTime string) *ApiOpenApi2FileMaterialListGetRequest {
	r.startTime = &startTime
	return r
}

// 素材结束时间，格式为yyyy-mm-dd HH:MM:SS
func (r *ApiOpenApi2FileMaterialListGetRequest) EndTime(endTime string) *ApiOpenApi2FileMaterialListGetRequest {
	r.endTime = &endTime
	return r
}

// 页数默认值: 1
func (r *ApiOpenApi2FileMaterialListGetRequest) Page(page int32) *ApiOpenApi2FileMaterialListGetRequest {
	r.page = &page
	return r
}

// 页面大小默认值:10, 最大值：3000
func (r *ApiOpenApi2FileMaterialListGetRequest) PageSize(pageSize int32) *ApiOpenApi2FileMaterialListGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2FileMaterialListGetRequest) Execute() (*FileMaterialListV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2FileMaterialListGetRequest) AccessToken(accessToken string) *ApiOpenApi2FileMaterialListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2FileMaterialListGetRequest) WithLog(enable bool) *ApiOpenApi2FileMaterialListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2FileMaterialListGet Method for OpenApi2FileMaterialListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2FileMaterialListGetRequest
*/
func (a *FileMaterialListV2ApiService) Get(ctx context.Context) *ApiOpenApi2FileMaterialListGetRequest {
	return &ApiOpenApi2FileMaterialListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FileMaterialListV2Response
func (a *FileMaterialListV2ApiService) getExecute(r *ApiOpenApi2FileMaterialListGetRequest) (*FileMaterialListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *FileMaterialListV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/file/material/list/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.materialSource == nil {
		return localVarReturnValue, nil, ReportError("materialSource is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "material_source", r.materialSource)
	if r.propertiesFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "properties_filter", r.propertiesFilter)
	}
	if r.startTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
	}
	if r.endTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
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
