/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
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

// QianchuanReportUniPromotionDimensionDataRoomGetV10ApiService QianchuanReportUniPromotionDimensionDataRoomGetV10Api service
type QianchuanReportUniPromotionDimensionDataRoomGetV10ApiService service

type ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanReportUniPromotionDimensionDataRoomGetV10ApiService
	advertiserId *int64
	roomId       *int64
	startTime    *string
	endTime      *string
	dimension    *QianchuanReportUniPromotionDimensionDataRoomGetV10Dimension
	metrics      *[]string
	orderField   *string
	orderType    *QianchuanReportUniPromotionDimensionDataRoomGetV10OrderType
	page         *int32
	pageSize     *int32
}

func (r *ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest) RoomId(roomId int64) *ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest {
	r.roomId = &roomId
	return r
}

func (r *ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest) StartTime(startTime string) *ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest {
	r.startTime = &startTime
	return r
}

func (r *ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest) EndTime(endTime string) *ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest) Dimension(dimension QianchuanReportUniPromotionDimensionDataRoomGetV10Dimension) *ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest {
	r.dimension = &dimension
	return r
}

func (r *ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest) Metrics(metrics []string) *ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest {
	r.metrics = &metrics
	return r
}

func (r *ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest) OrderField(orderField string) *ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest {
	r.orderField = &orderField
	return r
}

// 排序方式
func (r *ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest) OrderType(orderType QianchuanReportUniPromotionDimensionDataRoomGetV10OrderType) *ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest {
	r.orderType = &orderType
	return r
}

// 页码
func (r *ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest) Page(page int32) *ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest {
	r.page = &page
	return r
}

// 页面大小
func (r *ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest) PageSize(pageSize int32) *ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest) Execute() (*QianchuanReportUniPromotionDimensionDataRoomGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGet Method for OpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest
*/
func (a *QianchuanReportUniPromotionDimensionDataRoomGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest {
	return &ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanReportUniPromotionDimensionDataRoomGetV10Response
func (a *QianchuanReportUniPromotionDimensionDataRoomGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequest) (*QianchuanReportUniPromotionDimensionDataRoomGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanReportUniPromotionDimensionDataRoomGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/report/uni_promotion/dimension_data/room/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
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
	if r.startTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
	}
	if r.endTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
	}
	if r.dimension != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dimension", r.dimension)
	}
	if r.metrics != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "metrics", r.metrics)
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
